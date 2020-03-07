// THIS FILE IS AUTOGENERATED BY THE MAKEFILE. DO NOT EDIT.

package version

import "fmt"

var (
	// Version is the application version (injected via Makefile)
	Version string
	// BuildDate is the date which the application was built (injected via Makefile)
	BuildDate string
	// CommitHash is the git Commit hash of the application (injected via Makefile)
	CommitHash string
	// Author is the author of the application (injected via Makefile)
	Author string
)

// PrintVersion returns the current version information
func PrintVersion() {
	fmt.Printf("Version %s, Date: %s (Commit: %s), Author: %s\n", Version, BuildDate, CommitHash, Author)
}
