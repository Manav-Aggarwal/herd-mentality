package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// QuestionKeyPrefix is the prefix to retrieve all Question
	QuestionKeyPrefix = "Question/value/"
)

// QuestionKey returns the store key to retrieve a Question from the index fields
func QuestionKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
