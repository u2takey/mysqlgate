package version

import (
	"bytes"
	"fmt"
	"strings"
)

// pls refer to https://semver.org/
var (
	VersionMain string
	// VersionPre indicates prerelease
	VersionPre string
	// VersionDev indicates development branch. Releases will be empty string.
	VersionDev string
)

// Version is the specification version that the package types support.
func Version() string {
	var buffer bytes.Buffer
	buffer.WriteString(VersionMain)

	if VersionPre != "" {
		buffer.WriteString(fmt.Sprintf("-%s", VersionPre))
	}
	if VersionDev != "" {
		buffer.WriteString(fmt.Sprintf("+%s", VersionDev))
	}
	return buffer.String()
}

// ParseVersionMain ...
func ParseVersionMain(in string) (main string) {
	vs := strings.Split(in, "-")
	if len(vs) > 0 {
		main = vs[0]
	}
	return
}

