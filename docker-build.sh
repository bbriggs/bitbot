#! /bin/sh

# Get latest commit, tag, and branch we build on
export GIT_COMMIT=$(git rev-list -1 HEAD) 
export GIT_TAG=$(git describe --abbrev=0 --tags)
export GIT_BRANCH=$(git rev-parse --abbrev-ref HEAD) 

# We don't want empty tags overriding the default value
if [ ! -z $GIT_TAG ]; then
	XFLAGS="-X github.com/bbriggs/bitbot/bitbot.GitCommit=$GIT_COMMIT -X github.com/bbriggs/bitbot/bitbot.GitVersion=$GIT_TAG -X github.com/bbriggs/bitbot/bitbot.GitBranch=$GIT_BRANCH"
else
	XFLAGS="-X github.com/bbriggs/bitbot/bitbot.GitCommit=$GIT_COMMIT -X github.com/bbriggs/bitbot/bitbot.GitBranch=$GIT_BRANCH"
fi

GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -ldflags "-s -w -extldflags '-static' ${XFLAGS}" -o /go/bin/bitbot

echo "Compiled bitbot:\n\tGit tag: $GIT_TAG\n\tGit commit: $GIT_COMMIT\n\tGit branch: $GIT_BRANCH\n"
