package abstract_factory

import (
	// inside imports
	"fmt"

	// outside imports.
	"github.com/pkg/errors"

type ISportsFactory interface {
	makeShoe() IShoe
	makeShirt() IShirt
}

func GetSportsFactory(brand string) (ISportsFactory, error) {
	if brand == "adidas" {
		return &Adidas{}, error.New("......")
	}

	if brand == "nike" {
		return &Nike{}, error.New("")
	}

	return nil, fmt.Errorf("Wrong brand type passed")
}
