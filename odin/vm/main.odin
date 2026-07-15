package vm

import "core:fmt"

OpCode :: enum u8 {
  push = 0,
  pop = 1,

  add = 2,
  sub = 3,
  mul = 4,
  div = 5,

  sadd = 6,
  ssub = 7,
  smul = 8,
  sdiv = 9,

  mov = 10,

  puts = 11,
  putc = 12,
  print = 13,

  jam = 14,
}

main :: proc() {
  bin :[]u8= { 0, 1, 0, 2, 6, 13, 14 }
  stack_buf:[256]u8;
  state := mkState(stack_buf[:], bin)

  loop: for {
    op := transmute(OpCode)state->next();

    switch (op) {
      case .jam: { break loop }

      case .push: { state->push(state->next()) }; break

      case .sadd: {
        v := state->pop()
        state.stack->peek()^ += v;
      }; break

      case .ssub: {
        v := state->pop()
        state.stack->peek()^ -= v;
      }; break

      case .smul: {
        v := state->pop()
        state.stack->peek()^ *= v;
      }; break

      case .sdiv: {
        v := state->pop()
        state.stack->peek()^ /= v;
      }; break

      case .print: { fmt.printf("%d", state->pop()) }; break

      case .pop: { panic("todo") }; break
      case .mov: { panic("todo") }; break
      case .add: { panic("todo: registers") }; break
      case .sub: { panic("todo: registers") }; break
      case .mul: { panic("todo: registers") }; break
      case .div: { panic("todo: registers") }; break

      case .puts: { panic("todo: pointers") }; break
      case .putc: { panic("todo: pointers") }; break

      case: fmt.eprintf("invalid opcode: %d", );
    }
  }
}
