#!/bin/bash
rm -rf .godeps
rm -rf pkg
source gvp
cp -rf /data/share/golang/* .godeps/src/
gpm git
echo "编译环境准备完毕"