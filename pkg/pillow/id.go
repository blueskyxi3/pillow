package pillow

import (
	"regexp"
	"strings"
)

func ID(id string) string {
	reg := regexp.MustCompile("[^A-Za-z0-9]+")

	id = reg.ReplaceAllString(id, "-")
	id = strings.ReplaceAll(id, " ", "-")
	id = strings.ReplaceAll(id, "--", "-")

	return id
}
