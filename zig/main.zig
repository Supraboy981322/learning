const std = @import("std");
const fs = std.fs;
const net = std.net;

var stdout_buf:[1024]u8 = undefined;
var stdout_wr = fs.File.stdout().writer(&stdout_buf);
const stdout = &stdout_wr.interface;

pub fn main() !void {
    const addr = try net.Address.resolveIp("::", 9999);
    var server = try addr.listen(net.Address.ListenOptions{});
    defer server.deinit();
    log.info("listening on port {d}", .{addr.getPort()});
    log.req("foo", .{});
}
