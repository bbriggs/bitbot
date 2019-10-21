package bitbot

import (
	"fmt"
	"testing"
	"unicode"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
)

func TestChoose(t *testing.T) {
	properties := gopter.NewProperties(nil)

	properties.Property("returns x when given x or x", prop.ForAll(
		func(x string) bool {
			content := fmt.Sprintf("%s or %s", x, x)
			m := choose(content)
			return m == x
		}, gen.AnyString(),
	))

	properties.Property("returns \"\" when given any junk string", prop.ForAll(
		func(x string) bool {
			content := fmt.Sprintf(x)
			m := choose(content)
			return m == ""
		}, gen.AnyString(),
	))

	for script, rt := range unicode.Scripts {
		properties.Property(fmt.Sprintf("returns \"\" when given any junk %s unicode string", script), prop.ForAll(
			func(x string) bool {
				content := fmt.Sprintf(x)
				m := choose(content)
				return m == ""
			}, gen.UnicodeString(rt),
		))
	}

	properties.TestingRun(t)
}
