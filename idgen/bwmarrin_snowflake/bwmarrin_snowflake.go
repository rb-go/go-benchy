package bwmarrin

import (
	"github.com/bwmarrin/snowflake"
)

type SF struct {
	sf *snowflake.Node
}

func NewSF(id int64) *SF {
	sf, _ := snowflake.NewNode(id)
	return &SF{sf: sf}
}

func (j *SF) CreateID() (uint64, error) {
	id := j.sf.Generate()
	return uint64(id.Int64()), nil
}

func (j *SF) CreateHex() ([]byte, error) {
	id := j.sf.Generate()
	return id.Bytes(), nil
}

func (j *SF) ParseID(id uint64) error {
	return nil
}

func (j *SF) ParseHex(hexid []byte) error {
	return nil
}
