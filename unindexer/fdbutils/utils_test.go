package fdbutils

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncodingUint64(t *testing.T) {
	result := EncodeUInt64(20)
	assert.Equal(t, "1400000000000000", hex.EncodeToString(result), "failed to marshal %s", "uint32")
}

func TestDecodingUint64(t *testing.T) {
	decodedBytes, err := hex.DecodeString("1400000000000000")
	if err != nil {
		return
	}
	result := DecodeUInt64(decodedBytes)
	assert.Equal(t, uint64(20), result, "failed to unmarshal %s", "uint32")
}
