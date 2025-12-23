const std = @import("std");
const fs = std.fs;

var stdout_buf: [1024]u8 = undefined;
var stdout_wr = fs.File.stdout().writer(&stdout_buf);
const stdout = &stdout_wr.interface;

pub fn main() !void {
    try stdout.print("foo\n", .{});
    try stdout.flush();
}

