package bmps

import "github.com/rocco-gossmann/GoWas/core"

var mem_CursorBMP = []uint32 {
0x0de06f8b, 0x00000000, 0x014d2326, 0x09be2633, 0x02000000, 0x014d2326, 0x05be2633, 0x014d2326, 0x04000000, 0x014d2326, 0x07be2633, 0x014d2326, 0x02000000, 0x014d2326, 0x01be2633, 0x014d2326, 0x05be2633, 0x014d2326, 0x00000000, 0x014d2326, 0x01be2633, 0x00000000, 0x014d2326, 0x05be2633, 0x034d2326, 0x04000000, 0x014d2326, 0x05be2633, 0x08000000, 0x014d2326, 0x01be2633, 0x00000000, 
}

var CursorBMP = core.CreateBitmapFromCompressed(8, 64, &mem_CursorBMP)