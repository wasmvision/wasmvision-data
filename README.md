# wasmVision Data

Experimental repo for binary data formats shared by various wasmVision processors.

All data is stored in little-endian format.

## Face

```go
type Data struct {
	ID               uint32
	Rect             types.Rect
	RightEye         types.Size
	LeftEye          types.Size
	NoseTip          types.Size
	RightMouthCorner types.Size
	LeftMouthCorner  types.Size
}
```

```mermaid
---
title: "Face Data (in bytes)"
---
packet-beta
0-3: "ID"
4-7: "Rect Mix X"
8-11: "Rect Min Y"
12-15: "Rect Max X"
16-19: "Rect Max Y"
20-23: "Right Eye X"
24-27: "Right Eye Y"
28-31: "Left Eye X"
32-35: "Left Eye Y"
36-39: "Nose Tip X"
40-43: "Nose Tip Y"
44-47: "Right Mouth Corner X"
48-51: "Right Mouth Corner Y"
52-55: "Left Mouth Corner X"
56-59: "Left Mouth Corner Y"
```
