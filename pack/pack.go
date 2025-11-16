package pack

import (
	"strings"
	"fmt"
)

const p1 = "]"
const p2 = "["
const N = 4
func Pack(s string) string {
	if len(s) % 2 != 0 {
		s = fmt.Sprintf(" %s", s)
	}
	return fmt.Sprintf("%s\n%s%s%s\n%s",strings.Repeat("[]", N+len(s)/2),  strings.Repeat(p1, N), s, strings.Repeat(p2, N), strings.Repeat("[]", N+len(s)/2))
}

