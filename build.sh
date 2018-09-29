#! /bin/sh

export GIT_COMMIT=$(git rev-list -1 HEAD) && go install -ldflags "-X github.com/bbriggs/bitbot/bitbot.GitCommit=$GIT_COMMIT"
