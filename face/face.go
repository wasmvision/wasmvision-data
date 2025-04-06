package face

import (
	"encoding/binary"
	"errors"

	"wasmcv.org/wasm/cv/types"
)

var (
	errInvalidSize = errors.New("face: invalid size")
)

type Data struct {
	ID               uint32
	Rect             types.Rect
	RightEye         types.Size
	LeftEye          types.Size
	NoseTip          types.Size
	RightMouthCorner types.Size
	LeftMouthCorner  types.Size
}

func (f *Data) Read(p []byte) (n int, err error) {
	if len(p) < 60 {
		return 0, errInvalidSize
	}

	f.ID = binary.LittleEndian.Uint32(p[0:4])
	f.Rect.Min.X = int32(binary.LittleEndian.Uint32(p[4:8]))
	f.Rect.Min.Y = int32(binary.LittleEndian.Uint32(p[8:12]))
	f.Rect.Max.X = int32(binary.LittleEndian.Uint32(p[12:16]))
	f.Rect.Max.Y = int32(binary.LittleEndian.Uint32(p[16:20]))
	f.RightEye.X = int32(binary.LittleEndian.Uint32(p[20:24]))
	f.RightEye.Y = int32(binary.LittleEndian.Uint32(p[24:28]))
	f.LeftEye.X = int32(binary.LittleEndian.Uint32(p[28:32]))
	f.LeftEye.Y = int32(binary.LittleEndian.Uint32(p[32:36]))
	f.NoseTip.X = int32(binary.LittleEndian.Uint32(p[36:40]))
	f.NoseTip.Y = int32(binary.LittleEndian.Uint32(p[40:44]))
	f.RightMouthCorner.X = int32(binary.LittleEndian.Uint32(p[44:48]))
	f.RightMouthCorner.Y = int32(binary.LittleEndian.Uint32(p[48:52]))
	f.LeftMouthCorner.X = int32(binary.LittleEndian.Uint32(p[52:56]))
	f.LeftMouthCorner.Y = int32(binary.LittleEndian.Uint32(p[56:60]))

	return 60, nil
}

func (f *Data) Write(b []byte) (n int, err error) {
	if len(b) < 60 {
		return 0, errInvalidSize
	}

	binary.LittleEndian.PutUint32(b[0:4], f.ID)
	binary.LittleEndian.PutUint32(b[4:8], uint32(f.Rect.Min.X))
	binary.LittleEndian.PutUint32(b[8:12], uint32(f.Rect.Min.Y))
	binary.LittleEndian.PutUint32(b[12:16], uint32(f.Rect.Max.X))
	binary.LittleEndian.PutUint32(b[16:20], uint32(f.Rect.Max.Y))
	binary.LittleEndian.PutUint32(b[20:24], uint32(f.RightEye.X))
	binary.LittleEndian.PutUint32(b[24:28], uint32(f.RightEye.Y))
	binary.LittleEndian.PutUint32(b[28:32], uint32(f.LeftEye.X))
	binary.LittleEndian.PutUint32(b[32:36], uint32(f.LeftEye.Y))
	binary.LittleEndian.PutUint32(b[36:40], uint32(f.NoseTip.X))
	binary.LittleEndian.PutUint32(b[40:44], uint32(f.NoseTip.Y))
	binary.LittleEndian.PutUint32(b[44:48], uint32(f.RightMouthCorner.X))
	binary.LittleEndian.PutUint32(b[48:52], uint32(f.RightMouthCorner.Y))
	binary.LittleEndian.PutUint32(b[52:56], uint32(f.LeftMouthCorner.X))
	binary.LittleEndian.PutUint32(b[56:60], uint32(f.LeftMouthCorner.Y))

	return 60, nil
}
