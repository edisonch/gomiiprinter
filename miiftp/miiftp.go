package miiftp

import (
	"errors"
	"gomiiprinter/model"
)

func Set(miiftpconfig *model.Miiftpconfig) error {
	if (model.Miiftpconfig{}) == *miiftpconfig {
		return errors.New("invalid empty param")
	}

	return nil
}

func Get() *model.Miiftpconfig {
	return nil
}
