package piaohua

import (
	sf "github.com/piaohua/snowflake-golang"
)

type SF struct {
	sf *sf.Node
}

func NewSF(id uint32) *SF {
	return &SF{sf: sf.NewNode(id)}
}

func (j *SF) CreateID() (uint64, error) {
	id := j.sf.GenerateID()
	return id.Uint64(), nil
}

func (j *SF) CreateHex() ([]byte, error) {
	id := j.sf.GenerateID()
	return id.Bytes(), nil
}

func (j *SF) ParseID(id uint64) error {
	return nil
}

func (j *SF) ParseHex(hexid []byte) error {
	return nil
}
