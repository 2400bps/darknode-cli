#!/bin/sh
mkdir -p ../target/darknode/bin

cd ../cmd
xgo -go 1.10 --targets=darwin/amd64 .
xgo -go 1.10 --targets=linux/amd64 .
mv cmd-linux-amd64 ../target/darknode/bin/darknode_linux_amd64
mv cmd-darwin-10.6-amd64 ../target/darknode/bin/darknode_darwin_amd64
