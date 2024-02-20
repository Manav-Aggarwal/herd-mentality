package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// AnswerKeyPrefix is the prefix to retrieve all Answer
	AnswerKeyPrefix = "Answer/value/"
)

// AnswerKey returns the store key to retrieve a Answer from the index fields
func AnswerKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
