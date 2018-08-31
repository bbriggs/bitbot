package bitbot

import (
	"encoding/binary"
)

func int64ToByte(i int64) []byte{
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(i))
	return b
}

func byteToInt64(b []byte) int64{
	i := int64(binary.LittleEndian.Uint64(b))
	return i
}
