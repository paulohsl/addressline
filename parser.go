package addressline

import (
	"regexp"
	"strings"
)

func Parse(address string) string {

	reg := regexp.MustCompile(`^([\p{L}\p{M}\s]+),* ([\w\s]+)$`)

	matches := reg.FindStringSubmatch(address)

	if len(matches) > 2 {
		values := []string{matches[1], matches[2]}
		return strings.Join(values, ",")
	}

	return ""
}
