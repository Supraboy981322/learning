package vm

import "base:runtime"

Stack :: struct {
  top   : [^]u8,
  slice : []u8,

  push : proc(self:^Stack, value:u8),
  pop  : proc(self:^Stack) -> u8,
  peek : proc(self:^Stack) -> ^u8,
}
mkStack :: proc(buf:[]u8) -> Stack {
  return {
    top = &buf[0],
    slice = buf,

    push = proc(self:^Stack, value:u8) {
      self.top[0] = value
      self.top = self.top[1:]
    },
    pop = proc(self:^Stack) -> u8 {
      self.top = &self.top[-1]
      return self.top[0]
    },
    peek = proc(self:^Stack) -> ^u8 {
      return &self.top[-1];
    }
  }
}

State :: struct {
  stack : Stack,
  ip    : [^]u8,

  push : proc(self:^State, value:u8),
  pop  : proc(self:^State) -> u8,
  next : proc(self:^State) -> u8,
}
mkState :: proc(buf:[]u8, bin:[]u8) -> State {
  return {
    stack = mkStack(buf),
    ip = &bin[0],

    push = proc(self:^State, value:u8) {
      self.stack->push(value)
    },
    pop = proc(self:^State) -> u8 {
      return self.stack->pop()
    },
    next = proc(self:^State) -> u8 {
      defer self.ip = self.ip[1:]
      return self.ip[0];
    },
  }
}
