const foo = @embedFile("foo.html");
const std = @import("std");
const net = std.net;
const http = std.http;
var stdout_buff: [1024]u8 = undefined;
var stdout_writer = std.fs.File.stdout().writer(&stdout_buff);
const stdout = &stdout_writer.interface;

pub fn main() !void {
    const addr = try net.Address.parseIp4("127.0.0.1", 8080);

    var server = try addr.listen(net.Address.ListenOptions{});
    defer server.deinit();

    while (true) {
        try hanConn(try server.accept());
    }
}

fn hanConn(conn: net.Server.Connection) !void {
    defer conn.stream.close();
    var buff: [1024]u8 = undefined;
    var http_server = http.Server.init(conn, &buff);
    var req = try http_server.receiveHead();
    try req.respond("foo\n", http.Server.Request.RespondOptions{});
}
