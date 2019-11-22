package version

import "fmt"

// GitCommit indicates which git commit the binary was built from
var GitCommit string

// String returns a pretty string concatenation of GitCommit
func String() string {
	return fmt.Sprintf("Marketplace source git commit: %s\n", GitCommit)
}
