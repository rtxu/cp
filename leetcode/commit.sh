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
mv solution*.* $PROBLEM
git add $PROBLEM

# $2 maybe unset variable
set +u
if [ -z $2 ]; then
  MSG="solution($PROBLEM)"
else
  MSG="solution($PROBLEM): $2"
fi
set -u

git commit -m "$MSG"
