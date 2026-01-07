const std = @import("std");

var stdout_buf:[1024]u8 = undefined;
var stdout_wr = std.fs.File.stdout().writer(&stdout_buf);
var stdout = &stdout_wr.interface;

const log = struct {
    fn info(comptime msg: []const u8, args: anytype) void {
        const prefix = "\x1b[38;2;150;150;150m[\x1b[1;33m" ++
                        "info\x1b[38;2;150;150;150m]:\x1b[0m ";
        nosuspend stdout.print(prefix ++ msg ++ "\n", args) catch return;
        nosuspend stdout.flush() catch return;
    }
    fn req(comptime msg: []const u8, args: anytype) void {
        const prefix = "\x1b[38;2;150;150;150m[\x1b[1;34m" ++
                        "req\x1b[38;2;150;150;150m]:\x1b[0m ";
        nosuspend stdout.print(prefix ++ msg ++ "\n", args) catch return;
        nosuspend stdout.flush() catch return;
    }
};
