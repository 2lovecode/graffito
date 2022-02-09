#include "textflag.h"

TEXT ·output(SB), NOSPLIT, $8-48
    MOVQ 24(SP), DX
    MOVQ DX, ret3+24(FP)
    MOVQ perhapsArg1+16(SP), BX
    MOVQ BX, ret2+16(FP)
    MOVQ arg1+0(FP), AX
    MOVQ AX, ret1+8(FP)
    RET

TEXT ·output1(SB), NOSPLIT, $24-24
    MOVQ a+0(FP), DX
    MOVQ DX, 0(SP)
    MOVQ b+8(FP), CX
    MOVQ CX, 8(SP)
    CALL ·add(SB)
    MOVQ 16(SP), AX
    MOVQ AX, ret+16(FP)
    RET

