package linux

import (
	"strings"
)

func PrettyName() string {
	prettyName := OsRelease()
	return strings.Trim(prettyName["PRETTY_NAME"], "\"")
}

