#!/bin/bash
set -x

rm -rf .git
git init
git config user.name daodao97
git config user.email daodao97@foxmail.com
git add .
git commit -m 'init'
git remote add origin git@github.com:daodao97/EasyClash.git
git push origin master -f
