package godruoyi

import (
	"time"

	"github.com/godruoyi/go-snowflake"
)

type SF struct {
}

func NewSF() *SF {
	snowflake.SetMachineID(1) // change to your machineID
	snowflake.SetStartTime(time.Date(2014, 9, 1, 0, 0, 0, 0, time.UTC))
	return &SF{}
}

func (j *SF) CreateID() (uint64, error) {
	return snowflake.NextID()
}

func (j *SF) CreateHex() ([]byte, error) {
	val, err := snowflake.NextID()
	if err != nil {
		return nil, err
	}
	return i64tob(val), nil
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
