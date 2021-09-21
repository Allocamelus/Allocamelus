package errtools_test

import (
	"testing"

	"github.com/allocamelus/allocamelus/internal/pkg/errtools"
)

func TestInvalidLen(t *testing.T) {
	e := errtools.InvalidLen(1, 64)
	if e.Error() != "invalid-length-min1-max64" {
		t.Error("Failed errtools InvalidLen")
	}
}
