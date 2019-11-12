package bitbot

const SourceRepo string = "https://github.com/bbriggs/bitbot"

var (
	GitVersion string = "1.1.0" // Dockerhub appears to use shallow clones which drop tag info. Set this as a default."
	GitCommit  string = "7d1a5990f8eaa380367406ef1f154e729556fb1c"
	GitBranch  string = "fix-info-trigger"
)
