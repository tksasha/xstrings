package xstrings

import "strings"

func Transliterate(input string) string {
	mapping := map[rune][]rune{
		'а':  {'a'},
		'б':  {'b'},
		'в':  {'v'},
		'г':  {'h'},
		'ґ':  {'g'},
		'д':  {'d'},
		'е':  {'e'},
		'є':  {'i', 'e'},
		'ж':  {'z', 'h'},
		'з':  {'z'},
		'и':  {'y'},
		'і':  {'i'},
		'ї':  {'i'},
		'й':  {'i'},
		'к':  {'k'},
		'л':  {'l'},
		'м':  {'m'},
		'н':  {'n'},
		'о':  {'o'},
		'п':  {'p'},
		'р':  {'r'},
		'с':  {'s'},
		'т':  {'t'},
		'у':  {'u'},
		'ф':  {'f'},
		'х':  {'k', 'h'},
		'ц':  {'t', 's'},
		'ч':  {'c', 'h'},
		'ш':  {'s', 'h'},
		'щ':  {'s', 'h', 'c', 'h'},
		'ю':  {'i', 'u'},
		'я':  {'i', 'a'},
		'ь':  {},
		'\'': {},
	}

	output := []rune{}

	for _, i := range strings.ToLower(input) {
		rs, ok := mapping[i]
		if !ok {
			output = append(output, '_')

			continue
		}

		output = append(output, rs...)
	}

	return string(output)
}
