// +build !windows

package parser

import (
	"fmt"
	"github.com/gofrs/flock"
	"github.com/google/renameio"
)

func (p *configParser) save(data []byte, filename string) error {
	f := flock.New(filename)
	if err := f.Lock(); err != nil {
		return err
	}
	err := renameio.WriteFile(filename, data, 0644)
	if err != nil {
		f.Unlock() //nolint:errcheck
		return err
	}
	if err := f.Unlock(); err != nil {
		errMsg := err.Error()
		return fmt.Errorf("%w %s", UnlockError{}, errMsg)
	}
	return nil
}
