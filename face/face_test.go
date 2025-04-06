package face

import (
	"encoding/binary"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"wasmcv.org/wasm/cv/types"
)

func TestReadFaceData(t *testing.T) {
	// Create a byte slice with the expected data
	data := make([]byte, 60)
	binary.LittleEndian.PutUint32(data[0:4], 1)     // ID
	binary.LittleEndian.PutUint32(data[4:8], 10)    // Rect.Min.X
	binary.LittleEndian.PutUint32(data[8:12], 20)   // Rect.Min.Y
	binary.LittleEndian.PutUint32(data[12:16], 30)  // Rect.Max.X
	binary.LittleEndian.PutUint32(data[16:20], 40)  // Rect.Max.Y
	binary.LittleEndian.PutUint32(data[20:24], 50)  // RightEye.X
	binary.LittleEndian.PutUint32(data[24:28], 60)  // RightEye.Y
	binary.LittleEndian.PutUint32(data[28:32], 70)  // LeftEye.X
	binary.LittleEndian.PutUint32(data[32:36], 80)  // LeftEye.Y
	binary.LittleEndian.PutUint32(data[36:40], 90)  // NoseTip.X
	binary.LittleEndian.PutUint32(data[40:44], 100) // NoseTip.Y
	binary.LittleEndian.PutUint32(data[44:48], 110) // RightMouthCorner.X
	binary.LittleEndian.PutUint32(data[48:52], 120) // RightMouthCorner.Y
	binary.LittleEndian.PutUint32(data[52:56], 130) // LeftMouthCorner.X
	binary.LittleEndian.PutUint32(data[56:60], 140) // LeftMouthCorner.Y

	// Create a FaceData instance and read from the byte slice
	var faceData Data
	n, err := faceData.Read(data)

	// Check for errors and the number of bytes read
	require.NoError(t, err)
	assert.Equal(t, len(data), n)

	// Check the values in the FaceData instance
	assert.Equal(t, uint32(1), faceData.ID)
	assert.Equal(t, int32(10), faceData.Rect.Min.X)
	assert.Equal(t, int32(20), faceData.Rect.Min.Y)
	assert.Equal(t, int32(30), faceData.Rect.Max.X)
	assert.Equal(t, int32(40), faceData.Rect.Max.Y)
	assert.Equal(t, int32(50), faceData.RightEye.X)
	assert.Equal(t, int32(60), faceData.RightEye.Y)
	assert.Equal(t, int32(70), faceData.LeftEye.X)
	assert.Equal(t, int32(80), faceData.LeftEye.Y)
	assert.Equal(t, int32(90), faceData.NoseTip.X)
	assert.Equal(t, int32(100), faceData.NoseTip.Y)
	assert.Equal(t, int32(110), faceData.RightMouthCorner.X)
	assert.Equal(t, int32(120), faceData.RightMouthCorner.Y)
	assert.Equal(t, int32(130), faceData.LeftMouthCorner.X)
	assert.Equal(t, int32(140), faceData.LeftMouthCorner.Y)
}

func TestWriteFaceData(t *testing.T) {
	// Create a FaceData instance with known values
	faceData := Data{
		ID:               1,
		Rect:             types.Rect{Min: types.Size{X: 10, Y: 20}, Max: types.Size{X: 30, Y: 40}},
		RightEye:         types.Size{X: 50, Y: 60},
		LeftEye:          types.Size{X: 70, Y: 80},
		NoseTip:          types.Size{X: 90, Y: 100},
		RightMouthCorner: types.Size{X: 110, Y: 120},
		LeftMouthCorner:  types.Size{X: 130, Y: 140},
	}

	// Create a byte slice to hold the written data
	data := make([]byte, 60)

	// Write the FaceData instance into the byte slice
	n, err := faceData.Write(data)

	// Check for errors and the number of bytes written
	require.NoError(t, err)
	assert.Equal(t, len(data), n)

	// Check the values in the byte slice
	assert.Equal(t, uint32(1), binary.LittleEndian.Uint32(data[0:4]))
	assert.Equal(t, uint32(10), binary.LittleEndian.Uint32(data[4:8]))
	assert.Equal(t, uint32(20), binary.LittleEndian.Uint32(data[8:12]))
	assert.Equal(t, uint32(30), binary.LittleEndian.Uint32(data[12:16]))
	assert.Equal(t, uint32(40), binary.LittleEndian.Uint32(data[16:20]))
	assert.Equal(t, uint32(50), binary.LittleEndian.Uint32(data[20:24]))
	assert.Equal(t, uint32(60), binary.LittleEndian.Uint32(data[24:28]))
	assert.Equal(t, uint32(70), binary.LittleEndian.Uint32(data[28:32]))
	assert.Equal(t, uint32(80), binary.LittleEndian.Uint32(data[32:36]))
	assert.Equal(t, uint32(90), binary.LittleEndian.Uint32(data[36:40]))
	assert.Equal(t, uint32(100), binary.LittleEndian.Uint32(data[40:44]))
	assert.Equal(t, uint32(110), binary.LittleEndian.Uint32(data[44:48]))
	assert.Equal(t, uint32(120), binary.LittleEndian.Uint32(data[48:52]))
	assert.Equal(t, uint32(130), binary.LittleEndian.Uint32(data[52:56]))
	assert.Equal(t, uint32(140), binary.LittleEndian.Uint32(data[56:60]))
}

func TestWriteFaceDataInvalidSize(t *testing.T) {
	// Create a FaceData instance with known values
	faceData := Data{
		ID:               1,
		Rect:             types.Rect{Min: types.Size{X: 10, Y: 20}, Max: types.Size{X: 30, Y: 40}},
		RightEye:         types.Size{X: 50, Y: 60},
		LeftEye:          types.Size{X: 70, Y: 80},
		NoseTip:          types.Size{X: 90, Y: 100},
		RightMouthCorner: types.Size{X: 110, Y: 120},
		LeftMouthCorner:  types.Size{X: 130, Y: 140},
	}

	// Create a byte slice with an invalid size
	data := make([]byte, 50)

	// Attempt to write the FaceData instance into the byte slice
	n, err := faceData.Write(data)

	// Check for errors and the number of bytes written
	assert.Equal(t, errors.New("face: invalid size"), err)
	assert.Equal(t, 0, n)
}
