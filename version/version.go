package version

var (
	GitCommit string
	Version string = "0.1.0"
)
func init() {
	if GitCommit != "" {
		Version += "-" + GitCommit
	}
}

