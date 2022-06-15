package fdbutils

import (
	"bytes"
	"encoding/binary"
)

func EncodeInt64(v int64) []byte {
	var buf bytes.Buffer

	data := make([]byte, binary.MaxVarintLen64)
	binary.PutVarint(data, v)

	buf.Write(data)

	return buf.Bytes()
}

func DecodeInt64(value []byte, dest *int64) (int, error) {
	res, _ := binary.Varint(value)
	*dest = res
	return binary.MaxVarintLen64, nil
}

func EncodeUInt64(v uint64) []byte {
	var buf bytes.Buffer

	data := make([]byte, 8)
	binary.LittleEndian.PutUint64(data, v)
	buf.Write(data)

	return buf.Bytes()
}

func DecodeUInt64(value []byte) uint64 {
	return binary.LittleEndian.Uint64(value)
}
