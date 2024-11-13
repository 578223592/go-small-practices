main.adder STEXT size=80 args=0x0 locals=0x28 funcid=0x0 align=0x0
	0x0000 00000 (/Users/didi/projects/practices/test1/main.go:3)	TEXT	main.adder(SB), ABIInternal, $48-0
	0x0000 00000 (/Users/didi/projects/practices/test1/main.go:3)	MOVD	16(g), R16
	0x0004 00004 (/Users/didi/projects/practices/test1/main.go:3)	PCDATA	$0, $-2
	0x0004 00004 (/Users/didi/projects/practices/test1/main.go:3)	CMP	R16, RSP
	0x0008 00008 (/Users/didi/projects/practices/test1/main.go:3)	BLS	64
	0x000c 00012 (/Users/didi/projects/practices/test1/main.go:3)	PCDATA	$0, $-1
	0x000c 00012 (/Users/didi/projects/practices/test1/main.go:3)	MOVD.W	R30, -48(RSP)
	0x0010 00016 (/Users/didi/projects/practices/test1/main.go:3)	MOVD	R29, -8(RSP)
	0x0014 00020 (/Users/didi/projects/practices/test1/main.go:3)	SUB	$8, RSP, R29
	0x0018 00024 (/Users/didi/projects/practices/test1/main.go:3)	FUNCDATA	$0, gclocals·g2BeySu+wFnoycgXfElmcg==(SB)
	0x0018 00024 (/Users/didi/projects/practices/test1/main.go:3)	FUNCDATA	$1, gclocals·g2BeySu+wFnoycgXfElmcg==(SB)
	0x0018 00024 (/Users/didi/projects/practices/test1/main.go:5)	MOVD	$type:noalg.struct { F uintptr; X0 int }(SB), R0
	0x0020 00032 (/Users/didi/projects/practices/test1/main.go:5)	PCDATA	$1, $0
	0x0020 00032 (/Users/didi/projects/practices/test1/main.go:5)	CALL	runtime.newobject(SB)
	0x0024 00036 (/Users/didi/projects/practices/test1/main.go:5)	MOVD	$main.adder.func1(SB), R1
	0x002c 00044 (/Users/didi/projects/practices/test1/main.go:5)	MOVD	R1, (R0)
	0x0030 00048 (/Users/didi/projects/practices/test1/main.go:5)	MOVD	ZR, 8(R0)
	0x0034 00052 (/Users/didi/projects/practices/test1/main.go:5)	LDP	-8(RSP), (R29, R30)
	0x0038 00056 (/Users/didi/projects/practices/test1/main.go:5)	ADD	$48, RSP
	0x003c 00060 (/Users/didi/projects/practices/test1/main.go:5)	RET	(R30)
	0x0040 00064 (/Users/didi/projects/practices/test1/main.go:5)	NOP
	0x0040 00064 (/Users/didi/projects/practices/test1/main.go:3)	PCDATA	$1, $-1
	0x0040 00064 (/Users/didi/projects/practices/test1/main.go:3)	PCDATA	$0, $-2
	0x0040 00064 (/Users/didi/projects/practices/test1/main.go:3)	MOVD	R30, R3
	0x0044 00068 (/Users/didi/projects/practices/test1/main.go:3)	CALL	runtime.morestack_noctxt(SB)
	0x0048 00072 (/Users/didi/projects/practices/test1/main.go:3)	PCDATA	$0, $-1
	0x0048 00072 (/Users/didi/projects/practices/test1/main.go:3)	JMP	0
	0x0000 90 0b 40 f9 ff 63 30 eb c9 01 00 54 fe 0f 1d f8  ..@..c0....T....
	0x0010 fd 83 1f f8 fd 23 00 d1 00 00 00 90 00 00 00 91  .....#..........
	0x0020 00 00 00 94 01 00 00 90 21 00 00 91 01 00 00 f9  ........!.......
	0x0030 1f 04 00 f9 fd fb 7f a9 ff c3 00 91 c0 03 5f d6  .............._.
	0x0040 e3 03 1e aa 00 00 00 94 ee ff ff 17 00 00 00 00  ................
	rel 24+8 t=R_ADDRARM64 type:noalg.struct { F uintptr; X0 int }+0
	rel 32+4 t=R_CALLARM64 runtime.newobject+0
	rel 36+8 t=R_ADDRARM64 main.adder.func1+0
	rel 68+4 t=R_CALLARM64 runtime.morestack_noctxt+0
main.adder.func1 STEXT size=16 args=0x8 locals=0x0 funcid=0x0 align=0x0 leaf
	0x0000 00000 (/Users/didi/projects/practices/test1/main.go:5)	TEXT	main.adder.func1(SB), LEAF|NEEDCTXT|NOFRAME|ABIInternal, $0-8
	0x0000 00000 (/Users/didi/projects/practices/test1/main.go:5)	FUNCDATA	$0, gclocals·g2BeySu+wFnoycgXfElmcg==(SB)
	0x0000 00000 (/Users/didi/projects/practices/test1/main.go:5)	FUNCDATA	$1, gclocals·g2BeySu+wFnoycgXfElmcg==(SB)
	0x0000 00000 (/Users/didi/projects/practices/test1/main.go:5)	FUNCDATA	$5, main.adder.func1.arginfo1(SB)
	0x0000 00000 (/Users/didi/projects/practices/test1/main.go:5)	FUNCDATA	$6, main.adder.func1.argliveinfo(SB)
	0x0000 00000 (/Users/didi/projects/practices/test1/main.go:5)	PCDATA	$3, $1
	0x0000 00000 (/Users/didi/projects/practices/test1/main.go:5)	MOVD	8(R26), R1
	0x0004 00004 (/Users/didi/projects/practices/test1/main.go:6)	ADD	R1, R0, R0
	0x0008 00008 (/Users/didi/projects/practices/test1/main.go:7)	RET	(R30)
	0x0000 41 07 40 f9 00 00 01 8b c0 03 5f d6 00 00 00 00  A.@......._.....
main.main STEXT size=16 args=0x0 locals=0x0 funcid=0x0 align=0x0 leaf
	0x0000 00000 (/Users/didi/projects/practices/test1/main.go:10)	TEXT	main.main(SB), LEAF|NOFRAME|ABIInternal, $0-0
	0x0000 00000 (/Users/didi/projects/practices/test1/main.go:10)	FUNCDATA	$0, gclocals·g2BeySu+wFnoycgXfElmcg==(SB)
	0x0000 00000 (/Users/didi/projects/practices/test1/main.go:10)	FUNCDATA	$1, gclocals·g2BeySu+wFnoycgXfElmcg==(SB)
	0x0000 00000 (/Users/didi/projects/practices/test1/main.go:14)	RET	(R30)
	0x0000 c0 03 5f d6 00 00 00 00 00 00 00 00 00 00 00 00  .._.............
main.main.adder.func1 STEXT size=16 args=0x8 locals=0x0 funcid=0x0 align=0x0 leaf
	0x0000 00000 (/Users/didi/projects/practices/test1/main.go:5)	TEXT	main.main.adder.func1(SB), LEAF|NEEDCTXT|NOFRAME|ABIInternal, $0-8
	0x0000 00000 (/Users/didi/projects/practices/test1/main.go:5)	FUNCDATA	$0, gclocals·g2BeySu+wFnoycgXfElmcg==(SB)
	0x0000 00000 (/Users/didi/projects/practices/test1/main.go:5)	FUNCDATA	$1, gclocals·g2BeySu+wFnoycgXfElmcg==(SB)
	0x0000 00000 (/Users/didi/projects/practices/test1/main.go:5)	FUNCDATA	$5, main.main.adder.func1.arginfo1(SB)
	0x0000 00000 (/Users/didi/projects/practices/test1/main.go:5)	FUNCDATA	$6, main.main.adder.func1.argliveinfo(SB)
	0x0000 00000 (/Users/didi/projects/practices/test1/main.go:5)	PCDATA	$3, $1
	0x0000 00000 (/Users/didi/projects/practices/test1/main.go:5)	MOVD	8(R26), R1
	0x0004 00004 (/Users/didi/projects/practices/test1/main.go:6)	ADD	R1, R0, R0
	0x0008 00008 (/Users/didi/projects/practices/test1/main.go:7)	RET	(R30)
	0x0000 41 07 40 f9 00 00 01 8b c0 03 5f d6 00 00 00 00  A.@......._.....
go:cuinfo.producer.<unlinkable> SDWARFCUINFO dupok size=0
	0x0000 72 65 67 61 62 69                                regabi
go:cuinfo.packagename.main SDWARFCUINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
go:info.main.adder$abstract SDWARFABSFCN dupok size=26
	0x0000 05 6d 61 69 6e 2e 61 64 64 65 72 00 01 03 01 0e  .main.adder.....
	0x0010 73 75 6d 00 04 00 00 00 00 00                    sum.......
	rel 21+4 t=R_DWARFSECREF go:info.int+0
go:info.main.main.adder.func1$abstract SDWARFABSFCN dupok size=35
	0x0000 05 6d 61 69 6e 2e 6d 61 69 6e 2e 61 64 64 65 72  .main.main.adder
	0x0010 2e 66 75 6e 63 31 00 01 05 01 13 78 00 00 00 00  .func1.....x....
	0x0020 00 00 00                                         ...
	rel 30+4 t=R_DWARFSECREF go:info.int+0
main..inittask SNOPTRDATA size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
runtime.memequal64·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=R_ADDR runtime.memequal64+0
runtime.gcbits.0100000000000000 SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
type:.namedata.*func(int) int- SRODATA dupok size=16
	0x0000 00 0e 2a 66 75 6e 63 28 69 6e 74 29 20 69 6e 74  ..*func(int) int
type:*func(int) int SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 4f fe 94 82 08 08 08 36 00 00 00 00 00 00 00 00  O......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=R_ADDR runtime.memequal64·f+0
	rel 32+8 t=R_ADDR runtime.gcbits.0100000000000000+0
	rel 40+4 t=R_ADDROFF type:.namedata.*func(int) int-+0
	rel 48+8 t=R_ADDR type:func(int) int+0
type:func(int) int SRODATA dupok size=72
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 1d b2 35 81 02 08 08 33 00 00 00 00 00 00 00 00  ..5....3........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 01 00 01 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00                          ........
	rel 32+8 t=R_ADDR runtime.gcbits.0100000000000000+0
	rel 40+4 t=R_ADDROFF type:.namedata.*func(int) int-+0
	rel 44+4 t=RelocType(-32763) type:*func(int) int+0
	rel 56+8 t=R_ADDR type:int+0
	rel 64+8 t=R_ADDR type:int+0
type:.namedata.*struct { F uintptr; X0 int }- SRODATA dupok size=31
	0x0000 00 1d 2a 73 74 72 75 63 74 20 7b 20 46 20 75 69  ..*struct { F ui
	0x0010 6e 74 70 74 72 3b 20 58 30 20 69 6e 74 20 7d     ntptr; X0 int }
type:*struct { F uintptr; X0 int } SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 08 69 ea 8e 08 08 08 36 00 00 00 00 00 00 00 00  .i.....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=R_ADDR runtime.memequal64·f+0
	rel 32+8 t=R_ADDR runtime.gcbits.0100000000000000+0
	rel 40+4 t=R_ADDROFF type:.namedata.*struct { F uintptr; X0 int }-+0
	rel 48+8 t=R_ADDR type:noalg.struct { F uintptr; X0 int }+0
runtime.gcbits. SRODATA dupok size=0
type:.namedata.F. SRODATA dupok size=3
	0x0000 01 01 46                                         ..F
type:.namedata.X0. SRODATA dupok size=4
	0x0000 01 02 58 30                                      ..X0
type:noalg.struct { F uintptr; X0 int } SRODATA dupok size=128
	0x0000 10 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 0a b9 58 bb 02 08 08 19 00 00 00 00 00 00 00 00  ..X.............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 02 00 00 00 00 00 00 00 02 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0060 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0070 00 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	rel 32+8 t=R_ADDR runtime.gcbits.+0
	rel 40+4 t=R_ADDROFF type:.namedata.*struct { F uintptr; X0 int }-+0
	rel 44+4 t=RelocType(-32763) type:*struct { F uintptr; X0 int }+0
	rel 56+8 t=R_ADDR type:noalg.struct { F uintptr; X0 int }+80
	rel 80+8 t=R_ADDR type:.namedata.F.+0
	rel 88+8 t=R_ADDR type:uintptr+0
	rel 104+8 t=R_ADDR type:.namedata.X0.+0
	rel 112+8 t=R_ADDR type:int+0
gclocals·g2BeySu+wFnoycgXfElmcg== SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
main.adder.func1.arginfo1 SRODATA static dupok size=3
	0x0000 00 08 ff                                         ...
main.adder.func1.argliveinfo SRODATA static dupok size=2
	0x0000 00 00                                            ..
main.main.adder.func1.arginfo1 SRODATA static dupok size=3
	0x0000 00 08 ff                                         ...
main.main.adder.func1.argliveinfo SRODATA static dupok size=2
	0x0000 00 00                                            ..
