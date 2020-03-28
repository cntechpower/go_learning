### build
turn of code optimize and func inline to support one step debug and breakpoint
```
[root@10-0-0-2 objdump]# go build -gcflags "-N -l" -o test test.go
```

### debug
```
[root@10-0-0-2 objdump]# gdb test
GNU gdb (GDB) Red Hat Enterprise Linux 7.6.1-115.el7
Copyright (C) 2013 Free Software Foundation, Inc.
License GPLv3+: GNU GPL version 3 or later <http://gnu.org/licenses/gpl.html>
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.  Type "show copying"
and "show warranty" for details.
This GDB was configured as "x86_64-redhat-linux-gnu".
For bug reporting instructions, please see:
<http://www.gnu.org/software/gdb/bugs/>...
Reading symbols from /ssd/code/go_learning/src/go_learning/goNote/objdump/test...done.
Loading Go Runtime support.
(gdb) info files
Symbols from "/ssd/code/go_learning/src/go_learning/goNote/objdump/test".
Local exec file:
        `/ssd/code/go_learning/src/go_learning/goNote/objdump/test', file type elf64-x86-64.
        Entry point: 0x455340
        0x0000000000401000 - 0x00000000004583d2 is .text
        0x0000000000459000 - 0x00000000004881f4 is .rodata
        0x00000000004883c0 - 0x0000000000488a08 is .typelink
        0x0000000000488a08 - 0x0000000000488a10 is .itablink
        0x0000000000488a10 - 0x0000000000488a10 is .gosymtab
        0x0000000000488a20 - 0x00000000004cc2d3 is .gopclntab
        0x00000000004cd000 - 0x00000000004cd020 is .go.buildinfo
        0x00000000004cd020 - 0x00000000004ce620 is .noptrdata
        0x00000000004ce620 - 0x00000000004d0430 is .data
        0x00000000004d0440 - 0x00000000004f9990 is .bss
        0x00000000004f99a0 - 0x00000000004fc1a8 is .noptrbss
        0x0000000000400f9c - 0x0000000000401000 is .note.go.buildid
(gdb) b *0x455340
Breakpoint 1 at 0x455340
(gdb) r
Starting program: /ssd/code/go_learning/src/go_learning/goNote/objdump/test 

Breakpoint 1, 0x0000000000455340 in _rt0_amd64_linux ()
(gdb) n
Single stepping until exit from function _rt0_amd64_linux,
which has no line number information.
_rt0_amd64 () at /usr/local/go/src/runtime/asm_amd64.s:15
15              MOVQ    0(SP), DI       // argc
```