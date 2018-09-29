#! /bin/sh

# Get latest commit, tag, and branch we build on
export GIT_COMMIT=$(git rev-list -1 HEAD) 
export GIT_TAG=$(git describe --abbrev=0 --tags)
export GIT_BRANCH=$(git rev-parse --abbrev-ref HEAD) 

# We don't want empty tags overriding the default value
if [ ! -z $GIT_TAG ]; then
	go install -ldflags "-X github.com/bbriggs/bitbot/bitbot.GitCommit=$GIT_COMMIT -X github.com/bbriggs/bitbot/bitbot.GitVersion=$GIT_TAG -X github.com/bbriggs/bitbot/bitbot.GitBranch=$GIT_BRANCH"
else
	go install -ldflags "-X github.com/bbriggs/bitbot/bitbot.GitCommit=$GIT_COMMIT -X github.com/bbriggs/bitbot/bitbot.GitBranch=$GIT_BRANCH" 
fi

echo "Compiled bitbot:\n\tGit tag: $GIT_TAG\n\tGit commit: $GIT_COMMIT\n\tGit branch: $GIT_BRANCH\n"
