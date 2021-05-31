package config

import (
	"errors"

	"github.com/elastic/go-ucfg"
	"github.com/satori/go.uuid"
)

var _ ucfg.Unpacker = &WrapperUUID{}

type WrapperUUID struct {
	uuid.UUID
}

func (w *WrapperUUID) Unpack(v interface{}) error {
	input, ok := v.(string)
	if !ok {
		return errors.New("need string")
	}
	u, err := uuid.FromString(input)
	if err != nil {
		return err
	}
	copy(w.UUID[:], u[:])
	return nil
}
