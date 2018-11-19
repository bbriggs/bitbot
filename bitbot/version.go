package bitbot

const SourceRepo string = "https://github.com/bbriggs/bitbot"

var (
	GitVersion string = "1.0.0" // Dockerhub appears to use shallow clones which drop tag info. Set this as a default."
	GitCommit  string
	GitBranch  string
)
