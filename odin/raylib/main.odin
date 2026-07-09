package foo

import rl "vendor:raylib"

screen_width : i32 : 640
screen_height : i32 : 480

//no 'juicy main'? (as Zig's community calls it)
main :: proc() {

  rl.InitWindow(screen_width, screen_height, "foo")
  defer rl.CloseWindow()

  //no 'while' loops? he really did copy Go (along (possibly) with Jai)
  for !rl.WindowShouldClose() {
    rl.BeginDrawing()
    defer rl.EndDrawing()

    rl.ClearBackground(rl.BLACK)

    rl.DrawText(
      text = "foo",
      posX = rl.GetScreenWidth() / 2,
      posY = rl.GetScreenHeight() / 2,
      color = rl.WHITE,
      fontSize = 20,
    )
  }
}
