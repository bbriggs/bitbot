package bitbot

import (
	"encoding/binary"
	"time"
	"fmt"
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

func fmtDuration(d time.Duration) string {
    day := d / time.Hour * 24
    d -= day * time.Hour * 24

    h := d / time.Hour
    d -= h * time.Hour

    m := d / time.Minute
    d -= m * time.Minute

    return fmt.Sprintf("%02d days, %02d hours, %02d minutes", day, h, m)
}
