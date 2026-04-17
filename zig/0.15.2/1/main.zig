const std = @import("std");
const fs = std.fs;
const io = std.io;
var stdout_buff: [1024]u8 = undefined;
var stdout_writer = fs.File.stdout().writer(&stdout_buff);
const stdout = &stdout_writer.interface;

pub fn main() !void {
    var fi = try fs.cwd().openFile("foo.txt", .{});
    defer fi.close();
    var fi_buf: [1024]u8 = undefined;
    var fi_R = fi.reader(&fi_buf);

    var li_N:usize = 0;
    const fi_I = &fi_R.interface;
    while (try fi_I.takeDelimiter('\n')) |li| {
        li_N += 1;
        try stdout.print("{s}\n", .{li});
        try stdout.flush();
    }
}

