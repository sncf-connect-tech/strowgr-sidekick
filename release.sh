#!/bin/bash 

git checkout master
echo $1 > VERSION
git add VERSION
git commit -m "[release] v$1"
git push 
git tag "v$1"
git push --tags
