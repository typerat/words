package words

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math"
)

func rand(seed interface{}) int {
	// convert data to []byte
	var buf []byte
	switch v := seed.(type) {
	case []byte:
		buf = v
	case string:
		buf = []byte(v)
	case int64:
		buf = make([]byte, 8)
		binary.BigEndian.PutUint64(buf, uint64(v))
	default:
		buf = []byte(fmt.Sprint(v))
	}

	// hash input data
	hash := sha256.Sum256(buf)

	// convert hash to int32 (to be compatible with 32bit int)
	i := binary.BigEndian.Uint32(hash[:4])
	i &= math.MaxInt32

	return int(i)
}
