package brianvoe

import (
	"encoding/binary"

	"github.com/rs/xid"
)

// Problems with uint64!!!

type SF struct {
}

func NewSF() *SF {
	return &SF{}
}

func (j *SF) CreateID() (uint64, error) {
	guid := xid.New()
	//var num uint64
	//err := binary.Read(bytes.NewBuffer(guid[:]), binary.LittleEndian, &num)
	num := binary.LittleEndian.Uint64(guid[:])
	//binary.LittleEndian.PutUint64(guid[:], num)
	return num, nil
}

func (j *SF) CreateHex() ([]byte, error) {
	guid := xid.New()
	xid.NilID()
	return guid[:], nil
}

func (j *SF) ParseID(id uint64) error {
	return nil
}

func (j *SF) ParseHex(hexid []byte) error {
	return nil
}
