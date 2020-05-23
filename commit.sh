#!/bin/bash

# enable when debug
#set -x

# quit when any error happens
set -e
# quit when use undefined variable
set -u
# quit when error happens in pipe
set -o pipefail
trap "echo 'error: Script failed: see failed command above'" ERR

readonly PROBLEM=$1
mkdir $PROBLEM
mv solution.go $PROBLEM
git add $PROBLEM
git commit -m "solution($PROBLEM)"
