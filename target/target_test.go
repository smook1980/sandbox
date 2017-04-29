package target_test

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestTarget(t *testing.T) {
	RegisterTestingT(t)

	t.Run("blah", func(t *testing.T) {
		Expect(true).To(BeTrue())
	})
}
