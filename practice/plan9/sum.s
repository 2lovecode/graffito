#include "textflag.h"

TEXT ·sum(SB),NOSPLIT, $0-32
    MOVQ $0, SI
    MOVQ s1+0(FP), BX
    MOVQ s1+8(FP), CX
    INCQ CX

start:
    DECQ CX
    JZ done
    ADDQ (BX), SI
    ADDQ $8, BX
    JMP start

done:
    MOVQ SI, ret+24(FP)
    RET
    