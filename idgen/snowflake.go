package idgen

type Snowflaker interface {
	CreateID() (uint64, error)
	CreateHex() ([]byte, error)
	ParseID(id uint64) error
	ParseHex(hexid []byte) error
}
