; message to print '0x0A' is newline character
section .data
  msg db "foo", 0x0A
  len equ $ - msg ; gets the length of the string

; linker entry
section .text
  global _start
  
_start:
  mov rax, 1   ; Linux syscall: write
  mov rdi, 1   ; Linux file descriptor: stdout
  mov rsi, msg ; string address (from 'section .data')
  mov rdx, len ; number of bytes that need to be written 
  syscall      ; make syscall to kernel

  mov rax, 60 ; Linux syscall: exit
  mov rdi, 0  ; status code (Unix '0' for success)
  syscall     ; make syscall to kernel
