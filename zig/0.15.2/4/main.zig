const std = @import("std");
const mem = std.mem;
const heap = std.heap;

var stdout_buf:[1024]u8 = undefined;
var stdout_wr = std.fs.File.stdout().writer(&stdout_buf);
const stdout = &stdout_wr.interface;

pub fn main() !void {
    var gpa = heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const alloc = gpa.allocator();

    const fi = @embedFile("test.html");
    const t:[]const u8 = "<!-- split here -->";
    const new:[]const u8 = "/bT5M[>.n|'N@I\"_UYBHVA&oXryFkV&jq-b3#'5Gbv3t<d\"y-i.s+!BGhU[LKT|->DhlJ5@)WMZ4k'w.>o]b)<?}9hjJ!8[e\"-";
    
    const newSi = std.mem.replacementSize(u8, fi, t, new);

    const newStr = try alloc.alloc(u8, newSi);
    defer alloc.free(newStr);

    _ = std.mem.replace(u8, fi, t, new, newStr);

    try stdout.print("{s}", .{newStr});
    try stdout.flush();
}
