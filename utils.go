package contentscraper

import (
	"strings"
)

func sliceFormatter(slice []string) string {
	slj := strings.Join(slice, " ")
	r := strings.NewReplacer("     ", " ", "    ", " ", "   ", " ", "  ", " ")
	s := strings.TrimSpace(r.Replace(slj))
	return s
}

func replTxt(s string) string {
	rep := strings.NewReplacer(`\u00a0`, ` `, `"`, ``, `.""`, `. `)
	r := rep.Replace(s)
	return r
}
