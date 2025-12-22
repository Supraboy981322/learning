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
    //temporary, will replace with reading file into memory
    const in = @embedFile("in.txt");

    //split file into array of lines 
    var liS = mem.splitAny(u8, in, "\r\n");

    //create an allocator
    var gpa = heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const alloc = gpa.allocator();
    
    //create string array to hold json
    var out = std.ArrayList(u8).empty;
    defer out.deinit(alloc);

    //open json array
    try out.appendSlice(alloc, "[\n");

    //iterate over each line
    var liN: usize = 0;

    //skip first two lines
    //  (table head and empty line)
    for (0..2) |_| _ = liS.next();

    while (liS.next()) |li| {
        liN += 1; //moves to next line (?), not sure how it knows that's what this's for

        //if empty line (likely EOF), end loop
        if (li.len == 0) break;

        //open a new opject
        try out.appendSlice(alloc, "  {\n");

        //create backwards iterator, so the name is read first
        var fields = mem.splitBackwardsScalar(u8, li, ' ');

        var i:i8 = 0;
        var used:bool = false;
        while (fields.next()) |itm| {
            //skip item if empty
            //  (mem.splitScalar doesn't count for repeated delimiters)
            if (itm.len == 0) continue;

            //stop if name is longer than 3 chars
            if (i == 0 and itm.len > 3) break;

            //what the current field is
            const w = switch(i) {
                0 => "name",
                1 => "#blocks",
                2 => "minor",
                3 =>  "major",
                else => "unknown field",
            };

            //create line 
            const nLi = try fmt.allocPrint(alloc, "    \"{s}\": \"{s}\",\n", .{w, itm});
            try out.appendSlice(alloc, nLi); //add to slice
            alloc.free(nLi); //free mem

            used = true;
            i += 1;
        }

        if (used) {
            //remove comma from last item in object
            for (0..2) |_| _ = out.pop();
        } else {
            //remove opened object
            for (0..10) |_| _ = out.pop();
        }

        try out.appendSlice(alloc, "\n  },\n");
    }

    //remove the comma and newline last object closing brace
    for (0..2) |_| _ = out.pop();

    //close json array
    try out.appendSlice(alloc, "\n]\n");

    //print the constructed json
    try stdout.print("{s}\n", .{out.items});
    try stdout.flush();
}
