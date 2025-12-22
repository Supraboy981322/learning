const std = @import("std");
const http = std.http;
const net = std.net;
const fs = std.fs;
var stdout_buf: [1024]u8 = undefined;
var stdout_wr = fs.File.stdout().writer(&stdout_buf);
const stdout = &stdout_wr.interface();

pub fn main() !void {
}
