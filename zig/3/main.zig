const std = @import("std");
const heap = std.heap;
const mem = std.mem;
const fmt = std.fmt;
const io = std.io;
const fs = std.fs;

//makes printing to stdout easier
//  `try stdout.print("...", .{});`
var stdout_buf: [1024]u8 = undefined;
var stdout_wr = fs.File.stdout().writer(&stdout_buf);
const stdout = &stdout_wr.interface;

pub fn main() !void {
    //create an allocator
    var gpa = heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const alloc = gpa.allocator();

    var fi = try fs.openFileAbsolute("/proc/partitions", .{ .mode = .read_only });
    defer fi.close();
    var fi_buf: [1024]u8 = undefined;
    var fi_R = fi.reader(&fi_buf);

    //create string array to hold json
    var out = std.ArrayList(u8).empty;
    defer out.deinit(alloc);
    
    //open json array
    try out.appendSlice(alloc, "[\n");


    var li_N:usize = 0;
    const fi_I = &fi_R.interface;
    while (try fi_I.takeDelimiter('\n')) |li| {
        li_N += 1;

        //skip first two lines
        //  (table head and empty line)
        if (li_N <= 2) continue;
        
        //open a new opject
        try out.appendSlice(alloc, "  {\n");

        //create backwards iterator, so the name is read first
        var fields = mem.splitBackwardsScalar(u8, li, ' ');

        var i:i8 = 0;
        var used:bool = false;
        var previous: [1024]u8 = undefined;
        while (fields.next()) |itm| {
            //skip item if empty
            //  (mem.splitScalar doesn't count for repeated delimiters)
            if (itm.len == 0) continue;

            //what the current field is
            const w = switch(i) {
                0 => "name",
                1 => "#blocks",
                2 => "minor",
                3 =>  "major",
                else => "unknown field",
            };

            if (i == 0) {
                if (mem.containsAtLeast(u8, &previous, 1, itm)) {
                    continue;
                } else {
                    previous = undefined;
                    mem.copyForwards(u8, previous[0..itm.len], itm);
                }
            }

            //create line 
            const nLi = try fmt.allocPrint(alloc, "    \"{s}\": \"{s}\",\n", .{w, itm});
            try out.appendSlice(alloc, nLi); //add to slice
            alloc.free(nLi); //free mem

            used = true;
            i += 1;
        }
       
        //cleanup if unused
        if (used) {
            //remove comma from last item in object
            for (0..2) |_| _ = out.pop();
        } else {
            //remove opened object
            for (0..10) |_| _ = out.pop();
        }

        try out.appendSlice(alloc, "\n  },\n");
    }

    //remove the comma and newline
    // from last object's closing brace
    for (0..2) |_| _ = out.pop();

    //close json array
    try out.appendSlice(alloc, "\n]\n");

    //print the constructed json
    try stdout.print("{s}\n", .{out.items});
    try stdout.flush();
}
