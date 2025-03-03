package xstrings_test

import (
	"testing"

	"github.com/tksasha/xstrings"
	"gotest.tools/v3/assert"
)

func TestTransliterate(t *testing.T) {
	data := map[string]string{
		"м'ясо":   "miaso",
		"їжа":     "izha",
		"під'їзд": "pidizd",
		"львів":   "lviv",
		"ЇЖАК":    "izhak",
		"Їжак":    "izhak",
		"щука":    "shchuka",
		"борщ":    "borshch",
		"об'єкт":  "obiekt",
		"україна": "ukraina",
		"дощ$":    "doshch_",
		"":        "",
	}

	for uk, tr := range data {
		actual := xstrings.Transliterate(uk)

		assert.Equal(t, actual, tr)
	}
}
