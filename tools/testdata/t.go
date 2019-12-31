// This package and the interfaces defined in it are used to test the
// interface implementation generator.
package testdata

import (
	"bytes"

	"github.com/mattermost/mattermost-server/v5/model"
)

// A very simple interface.
type Simple interface {
	Run(a int, b string) error
}

// A more complex interface.
type Complex interface {
	A()                                          // Empty method.
	B(a int, b string) (C, error)                // A method returning a locally defined type.
	C(a *bytes.Buffer, b []byte) ([]byte, error) // Require core import, use slices.
	D(_ *model.CommandArgs) error                // 3rd party import.
}

// C is a locally defined type
type C int
