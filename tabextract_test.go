package tabextract_test

import (
	"testing"

	"github.com/collatzc/tabextract"
)

func TestColAtoi(t *testing.T) {
	n, err := tabextract.ColAtoi("A")
	if err != nil {
		t.Fatal(err)
	}
	t.Error(n)
}
