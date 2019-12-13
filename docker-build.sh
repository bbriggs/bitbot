#! /bin/sh

# Get latest commit, tag, and branch we build on
export GIT_COMMIT=$(git rev-parse --short HEAD) 
export GIT_TAG=$(git describe --abbrev=0 --tags)
export GIT_BRANCH=$(git rev-parse --abbrev-ref HEAD) 

echo $GIT_COMMIT
echo $GIT_TAG
echo $GIT_BRANCH

XFLAGS="-X github.com/bbriggs/bitbot/bitbot.GitCommit=$GIT_COMMIT -X github.com/bbriggs/bitbot/bitbot.GitBranch=$GIT_BRANCH -X github.com/bbriggs/bitbot/bitbot.GitTag=$GIT_TAG"

GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -ldflags "-s -w -extldflags '-static' ${XFLAGS}" -o ./opt/bitbot

if [ $? -eq 0 ]; then
	echo -e "Compiled bitbot:\n\tGit tag: $GIT_TAG\n\tGit commit: $GIT_COMMIT\n\tGit branch: $GIT_BRANCH\n"
else
	echo "The build failed"
fi
