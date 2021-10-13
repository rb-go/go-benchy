package sony

import (
	"time"

	sf "github.com/sony/sonyflake"
)

type SF struct {
	sf *sf.Sonyflake
}

func NewSF(StartTime time.Time) *SF {
	st := sf.Settings{
		StartTime:      StartTime,
		MachineID:      nil,
		CheckMachineID: nil,
	}
	// time.Date(2014, 9, 1, 0, 0, 0, 0, time.UTC)
	return &SF{
		sf: sf.NewSonyflake(st),
	}
}

func (s *SF) CreateID() (uint64, error) {
	return s.sf.NextID()
}

func (s *SF) CreateHex() ([]byte, error) {
	val, err := s.sf.NextID()
	if err != nil {
		return nil, err
	}
	return i64tob(val), nil
}

func (s *SF) ParseID(id uint64) error {
	return nil
}

func (s *SF) ParseHex(hexid []byte) error {
	return nil
}

func i64tob(val uint64) []byte {
	r := make([]byte, 8)
	for i := uint64(0); i < 8; i++ {
		r[i] = byte((val >> (i * 8)) & 0xff)
	}
	return r
}
