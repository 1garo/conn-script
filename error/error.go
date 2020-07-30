package error

import (
	"errors"
	"fmt"
)

func CheckNil(err error, msg string) error {
	if err != nil {
		return errors.New(fmt.Sprintf("An error ocurred: %s", msg))
	}
	return nil
}

