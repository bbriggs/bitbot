#! /bin/sh

# Get latest commit, tag, and branch we build on
export GIT_COMMIT=$(git rev-list -1 HEAD) 
export GIT_TAG=$(git describe --abbrev=0 --tags)
export GIT_BRANCH=$(git rev-parse --abbrev-ref HEAD) 

# We don't want empty tags overriding the default value
if [ ! -z $GIT_TAG ]; then
	XFLAGS="-X github.com/bbriggs/bitbot/bitbot.VERSION=$GIT_TAG"
else
	XFLAGS="-X github.com/bbriggs/bitbot/bitbot.VERSION=$GIT_COMMIT"
fi

# Passing variables in the build for the !info trigger
sed "s/.*GitCommit.*/	GitCommit  string = \"$GIT_COMMIT\"/" -i bitbot/version.go
sed "s/.*GitBranch.*/	GitBranch  string = \"$GIT_BRANCH\"/" -i bitbot/version.go

GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -ldflags "-s -w -extldflags '-static' ${XFLAGS}" -o ./opt/bitbot

if [ $? -eq 0 ]; then
	echo -e "Compiled bitbot:\n\tGit tag: $GIT_TAG\n\tGit commit: $GIT_COMMIT\n\tGit branch: $GIT_BRANCH\n"
else
	echo "The build failed"
fi
