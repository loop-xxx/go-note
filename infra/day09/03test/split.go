package splitutils

import (
	"strings"
)

// SplitStr ...
func SplitStr(str, step string) (substrs []string) {
	var start int
	for index := strings.Index(str, step); index != -1; index = strings.Index(str[start:], step) {
		substrs = append(substrs, str[start:start+index])
		start += index + len(step)
	}
	substrs = append(substrs, str[start:])
	return substrs
}
