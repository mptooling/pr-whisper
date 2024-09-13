#!/bin/sh

set -eu

export GITHUB_TOKEN=${GITHUB_TOKEN}


sh -c "/app/prwhisper $*"