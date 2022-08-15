package version

import (
	"regexp"
	"runtime"
	"strings"
)

const versionCore = "0.1.0"

var (
	kRegex    = regexp.MustCompile(`^v?(0|[1-9][0-9]*)\.(0|[1-9][0-9]*)\.(0|[1-9][0-9]*)([-0-9a-zA-Z_\.+]*)?$`)
	gitCommit = ""
)

func Version() string {
	v := versionCore

	if gitCommit != "" {
		v += "+" + truncate(gitCommit, 14)
	}
	return v
}

func DisplayVersion() string {
	return "k v" + Version() + " " + runtime.Version() + " " + runtime.GOOS + "/" + runtime.GOARCH
}

func truncate(s string, maxLen int) string {
	if len(s) < maxLen {
		return s
	}
	return s[:maxLen]
}

// normalizedBuildVersion used to returns normalized build version (with "v" prefix if needed)
// If input doesn't match known version pattern, returns empty string.
func NormalizedBuildVersion(version string) string {
	if kRegex.MatchString(version) {
		if strings.HasPrefix(version, "v") {
			return version
		}
		return "v" + version
	}
	return ""
}
