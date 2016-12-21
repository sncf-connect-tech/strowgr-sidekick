#!/bin/bash 

git checkout master
echo $1 > VERSION
git add VERSION
git commit -m "[release] $1"
git push 
git tag $1
git push --tags
