package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// StoredBetKeyPrefix is the prefix to retrieve all StoredBet
	StoredBetKeyPrefix = "StoredBet/value/"
)

// StoredBetKey returns the store key to retrieve a StoredBet from the index fields
func StoredBetKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
