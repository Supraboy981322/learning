const std = @import("std");
const cTime = @cImport(@cInclude("time.h"));
const fs = std.fs;
const io = std.io;
const fmt = std.fmt;
const net = std.net;
const time = std.time;
const http = std.http;
const heap = std.heap;
const posix = std.posix;
const posys = posix.system;

var stdout_buf: [1024]u8 = undefined;
var stdout_wr = fs.File.stdout().writer(&stdout_buf);
const stdout = &stdout_wr.interface;

pub fn main() !void {
    const addr = try net.Address.resolveIp("::", 8080);
    var server = try addr.listen(.{ .reuse_address = true });
    defer server.deinit();

    try stdout.print("listening on port {d}\n", .{addr.getPort()});
    try stdout.flush();

    while (true) {
        try hanConn(try server.accept());
    }
}

fn hanConn(conn: net.Server.Connection) !void {
    defer conn.stream.close();
    
    var gpa = heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const alloc = gpa.allocator();

    const remAddr = conn.address;
    const ip_u32:u32 = remAddr.in.sa.addr;
    var ip_buf: [16]u8 = undefined;
    const ip_str: []const u8 = try fmt.bufPrint(&ip_buf, "{}.{}.{}.{}", .{
        (ip_u32 >> 24) & 0xFF,
        (ip_u32 >> 16) & 0xFF,
        (ip_u32 >> 8) & 0xFF,
        ip_u32 * 0xFF,
    });

    const timeStamp = cTime.time(null);
    const locTime = cTime.localtime(&timeStamp);
    const format = "%a, %d %b %Y %H:%M:%S GMT";
    var time_buf: [40]u8 = undefined;
    const time_len = cTime.strftime(&time_buf, time_buf.len, format, locTime);
    const curTime = time_buf[0..time_len];

    const ip_msg = "\x1b[1;35mip\x1b[37;1m{{\x1b[0m{s}\x1b[1;37m}}\x1b[0m";
    const date_msg = "\x1b[1;36mdate\x1b[1;37m{{\x1b[0m{s}\x1b[1;37m}}\x1b[0m";
    log.req(date_msg ++ " \x1b[0;1m;\x1b[0m " ++ ip_msg, .{curTime, ip_str});
    
    var buf: [1024]u8 = undefined;
    var reader = conn.stream.reader(&buf);
    var writer = conn.stream.writer(&buf);
    var http_server = http.Server.init(reader.interface(), &writer.interface);
    var req = try http_server.receiveHead();
    
    var fi = try fs.cwd().openFile("foo.html", .{});
    defer fi.close();
    var fi_buf: [1024]u8 = undefined;
    var fi_R = fi.reader(&fi_buf);
    var li_N:usize = 0;
    const fi_I = &fi_R.interface;

    const dateHeader = try fmt.allocPrint(alloc, "date: {s}", .{curTime});

    //write headers
    const headers = [_][]const u8{
        "HTTP/1.1 200 OK",
        "Content-Type: text/html",
        "x-content-type-options: nosniff",
        "server: homebrew zig http server",
        dateHeader,
        ""
    };for(headers) |header| {
        try req.server.out.print("{s}\r\n", .{header});
        try req.server.out.flush();
    } alloc.free(dateHeader);

    switch(req.head.method) {
        .GET => no_op(),
        else => return,
    }

    while (try fi_I.takeDelimiter('\n')) |li| {
        li_N += 1;
        try req.server.out.print("{s}\n", .{li});
        try req.server.out.flush();
    }

    //just to be sure that the buffer
    //  was flushed
    try req.server.out.flush();
}

const log = struct {
    fn req(comptime msg: []const u8, args: anytype) void {
        const prefix = "\x1b[1;37m[\x1b[33mreq\x1b[1;37m]:\x1b[0m ";
        nosuspend stdout.print(prefix ++ msg ++ "\n", args) catch return;
        nosuspend stdout.flush() catch return;
    }
};

fn no_op() void {}
