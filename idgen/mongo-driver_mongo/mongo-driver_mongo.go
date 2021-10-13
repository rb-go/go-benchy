package mongo

import (
	"encoding/binary"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TODO

type SF struct {
}

func NewSF() *SF {

	// time.Date(2014, 9, 1, 0, 0, 0, 0, time.UTC)
	return &SF{}
}

func (j *SF) CreateID() (uint64, error) {
	mgoID := primitive.NewObjectID()
	num := binary.LittleEndian.Uint64(mgoID[:])
	return num, nil
}

func (j *SF) CreateHex() ([]byte, error) {
	mgoID := primitive.NewObjectID()
	return mgoID[:], nil
}

func (j *SF) ParseID(id uint64) error {
	return nil
}

func (j *SF) ParseHex(hexid []byte) error {
	return nil
}
