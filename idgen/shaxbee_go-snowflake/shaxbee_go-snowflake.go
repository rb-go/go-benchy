package shaxbee

import (
	"gopkg.in/shaxbee/go-snowflake.v1"
)

type SF struct {
	sf snowflake.Snowflake
}

func NewSF(wrkID uint64) *SF {
	sf, _ := snowflake.New(wrkID)
	return &SF{sf: sf}
}

func (j *SF) CreateID() (uint64, error) {
	return uint64(<-j.sf), nil
}

func (j *SF) CreateHex() ([]byte, error) {

	return i64tob(uint64(<-j.sf)), nil
}

func (j *SF) ParseID(id uint64) error {
	return nil
}

func (j *SF) ParseHex(hexid []byte) error {
	return nil
}

func i64tob(val uint64) []byte {
	r := make([]byte, 8)
	for i := uint64(0); i < 8; i++ {
		r[i] = byte((val >> (i * 8)) & 0xff)
	}
	return r
}
