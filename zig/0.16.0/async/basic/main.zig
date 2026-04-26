const std = @import("std");

pub const Msg = struct {
    name:[]u8,
    msg:[]u8,
    pub fn init(name:[]u8, msg:[]u8) Msg {
        return .{
            .name = name,
            .msg = msg,
        };
    }
};

pub const Chat = struct {
    messages:std.ArrayList(Msg) = .empty,
    alloc:std.mem.Allocator,
    mutex:std.Io.Mutex = .init,

    pub fn init(alloc:std.mem.Allocator) !Chat {
        return .{ .alloc = alloc };
    }

    pub fn deinit(self:*Chat, io:std.Io) !void {
        try self.mutex.lock(io);
        defer self.mutex.unlock(io);
        for (self.messages.items) |msg| {
            self.alloc.free(msg.name);
            self.alloc.free(msg.msg);
        }
        self.messages.deinit(self.alloc);
    }

    pub fn append(self:*Chat, io:std.Io, name:[]u8, msg:[]u8) !void {
        try self.mutex.lock(io);
        defer self.mutex.unlock(io);
        try self.messages.append(self.alloc, .init(
            try self.alloc.dupe(u8, name),
            try self.alloc.dupe(u8, msg)
        ));
        std.debug.print("|{s}| just said {s}\n", .{name, msg});
    }

    pub fn poll(self:*Chat, io:std.Io, have:usize) !?[]Msg {
        try self.mutex.lock(io);
        defer self.mutex.unlock(io);
        return
            if (have != self.messages.items.len)
                self.messages.items[have..]
            else
                null;
    }
};

pub fn main(init:std.process.Init) !void {

    var threaded:std.Io.Threaded = .init(init.gpa, .{});
    const io = threaded.io();

    const ip = try std.Io.net.IpAddress.parse("::", 8934);
    var server = try ip.listen(io, .{ .reuse_address = true });
    defer server.deinit(io);

    var group:std.Io.Group = .init;
    defer group.cancel(io);

    var chat:Chat = try .init(threaded.allocator);
    defer chat.deinit(io) catch chat.messages.deinit(chat.alloc);

    while (true) {
        const stream = try server.accept(io);
        group.async(io, connection_shim, .{ stream, io, threaded.allocator, &chat});
    }
}

pub fn connection_shim(
    stream:std.Io.net.Stream,
    io:std.Io,
    alloc:std.mem.Allocator,
    chat:*Chat
) !void {
    defer stream.close(io);

    var arena = std.heap.ArenaAllocator.init(alloc);
    defer _ = arena.deinit();

    std.debug.print("connection\n", .{});
    new_connection(stream, io, arena.allocator(), chat) catch |e| {
        std.debug.print("{t}\n", .{e});
        return error.Canceled;
    };
}

pub fn new_connection(
    stream:std.Io.net.Stream,
    io_outer:std.Io,
    life_time_alloc:std.mem.Allocator,
    chat_outer:*Chat
) !void {
    var arena_alloc:std.heap.ArenaAllocator = .init(life_time_alloc);
    defer _ = arena_alloc.deinit();
    const allocator = arena_alloc.allocator();

    var wr_buf:[1_024]u8 = undefined;
    var useless_writer = stream.writer(io_outer, &wr_buf);
    var writer = &useless_writer.interface;

    var re_buf:[1_024]u8 = undefined;
    var useless_reader = stream.reader(io_outer, &re_buf);
    var reader = &useless_reader.interface;

    try writer.print("welcome\r\n", .{});
    try writer.flush();

    const name_outer:[]u8 = b: {
        defer _ = arena_alloc.reset(.free_all);
        var ok:bool = false;
        var bad:bool = false;
        while (!ok) {
            defer {
                writer.flush() catch {};
                _ = arena_alloc.reset(.free_all);
            }
            if (bad) try writer.print(
                "\r\nsorry, but I want a 'y' (yes) or 'n' (no) response\r\n",
            .{});
            bad = false;
            try writer.print("might I ask your name? (y/n): ", .{});
            try writer.flush();
            var buf:std.Io.Writer.Allocating = .init(allocator);
            defer buf.deinit();
            const n = try reader.streamDelimiterEnding(&buf.writer, '\n');
            if (reader.peekByte() catch null) |b|
                if (std.ascii.isWhitespace(b))
                    reader.toss(1);
            if (n != 1) {
                std.debug.print("length not 1\n", .{});
                bad = true;
                continue;
            }
            var input:[]u8 = @constCast(
                std.mem.trim(u8, try buf.toOwnedSlice(), &std.ascii.whitespace)
            );
            if (input.len > 1) {
                bad = true;
                continue;
            }
            defer allocator.free(input);
            for (0..input.len) |i|
                input[i] = std.ascii.toLower(input[i]);
            const response = std.meta.stringToEnum(
                enum{ y, n }, input,
            ) orelse {
                std.debug.print("no match\n", .{});
                bad = true;
                continue;
            };
            switch (response) {
                .y => ok = true,
                .n => {
                    try writer.print("\nsorry, but I need a name; goodbye\n", .{});
                    try writer.flush();
                    return;
                },
            }
        }
        ok = false;
        bad = false;
        while (!ok) {
            defer {
                writer.flush() catch {};
                _ = arena_alloc.reset(.free_all);
            }
            if (bad)
                try writer.print("\r\ncome on, I just need a name\r\n", .{});
            try writer.print("\r\nplease enter your name: ", .{});
            try writer.flush();
            var buf:std.Io.Writer.Allocating = .init(allocator);
            defer buf.deinit();
            const n = try reader.streamDelimiterEnding(&buf.writer, '\n');
            if (reader.peekByte() catch null) |b|
                if (std.ascii.isWhitespace(b))
                    reader.toss(1);
            if (n < 1) {
                bad = true;
                continue;
            }
            break :b @constCast(std.mem.trim(u8,
                try life_time_alloc.dupe(u8, try buf.toOwnedSlice()),
                &std.ascii.whitespace
            ));
        }
        return;
    };
    defer {
        std.debug.print("{s} just disconnected", .{name_outer});
        life_time_alloc.free(name_outer);
        writer.print("goodbye...\r\n\r\n", .{}) catch {};
        writer.flush() catch {};
    }
    std.debug.print("|{s}| ({x}) just logged in\n", .{name_outer, name_outer});

    const chat_loop = struct {
        pub fn poller_shim(
            io:std.Io,
            wr:*std.Io.Writer,
            chat:*Chat,
        ) !void {
            poller(io, wr, chat) catch
                return error.Canceled;
        }
        pub fn poller(
            io:std.Io,
            wr:*std.Io.Writer,
            chat:*Chat,
        ) !void {
            var last_seen:usize = 0;
            while (true) {
                if (chat.poll(io, last_seen) catch null) |new| for (new) |msg| {
                    wr.print("\r\x1b[2K|{s}|: {s}\r\n", .{msg.name, msg.msg}) catch {};
                    wr.flush() catch {};
                    last_seen += 1;
                };
            }
        }
        pub fn main_shim(
            io:std.Io,
            wr:*std.Io.Writer,
            re:*std.Io.Reader,
            arena:*std.heap.ArenaAllocator,
            chat:*Chat,
            name:[]u8,
        ) !void {
            @This().main(io, wr, re, arena, chat, name) catch
                return error.Canceled;
        }
        pub fn main(
            io:std.Io,
            wr:*std.Io.Writer,
            re:*std.Io.Reader,
            arena:*std.heap.ArenaAllocator,
            chat:*Chat,
            name:[]u8,
        ) !void {
            const alloc = arena.allocator();
            var quit:bool = false;
            while (!quit) {
                defer _ = arena.reset(.free_all);
                const buf = re.takeDelimiterExclusive('\n') catch |e| switch (e) {
                    error.EndOfStream => continue,
                    else => return e,
                };
                if (buf.len < 1) {
                    wr.flush() catch return;
                    continue; // TODO: proper disconnection
                }
                re.toss(1);
                var msg = @constCast(std.mem.trim(u8, buf, &std.ascii.whitespace));
                blk: {
                    const cmd = std.meta.stringToEnum(
                        enum{ EXIT }, msg 
                    ) orelse break :blk;
                    switch (cmd) {
                        .EXIT => quit = true,
                    }
                    continue;
                }
                {
                    var i:usize = 0;
                    var mem:std.ArrayList(u8) = .empty;
                    var esc:bool = false;
                    while (i < msg.len) : (i += 1) {
                        if (esc) {
                            defer esc = false;
                            try mem.append(alloc, switch (msg[i]) {
                                'x' => {
                                    if (msg[i..].len > 3) {
                                        if (std.mem.eql(u8, "x1b", msg[i..i+3])) {
                                            try mem.append(alloc, '\x1b');
                                            i += 2;
                                            continue;
                                        } else std.debug.print(
                                            "not match: |{s}| ({x})",
                                            .{msg[i..i+3], msg[i..i+3]}
                                        );
                                    }
                                    try mem.appendSlice(alloc, "\\x");
                                    continue;
                                },
                                'r' => '\r',
                                't' => '\t',
                                'b' => std.ascii.control_code.bs,
                                'v' => std.ascii.control_code.vt,
                                'a' => std.ascii.control_code.bel,
                                else => {
                                    try mem.appendSlice(alloc, &[_]u8{ '\\', msg[i] });
                                    continue;
                                },
                            });
                            continue;
                        }
                        switch (msg[i]) {
                            '\\' => esc = true,
                            else => try mem.append(alloc, msg[i]),
                        }
                    }
                    alloc.free(msg);
                    msg = try mem.toOwnedSlice(alloc);
                }
                try chat.append(io, name, msg);
            }
        }
    };
    var group:std.Io.Group = .init;
    group.async(io_outer, chat_loop.poller_shim, .{
        io_outer, writer, chat_outer
    });
    group.async(io_outer, chat_loop.main_shim, .{
        io_outer, writer, reader, &arena_alloc, chat_outer, name_outer
    });
    try group.await(io_outer);
}
