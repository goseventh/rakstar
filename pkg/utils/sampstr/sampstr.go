package rakstar

import "golang.org/x/text/encoding/charmap"

func Decode(input string) string {
	str, _ := charmap.Windows1252.NewDecoder().String(input)

	return str
}

func Encode(input string) string {
	str, _ := charmap.Windows1252.NewEncoder().String(input)

	return str
}
