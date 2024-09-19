#!/bin/bash

pkg_prefix=obsgo
if [[ -d "build" ]]; then
    rm -r build
fi
archs=(amd64 arm64 arm ppc64le ppc64 s390x) 
for arch in ${archs[@]} 
do
    env CGO_ENABLED=0 GOOS=linux GOARCH=${arch} go build -o ./build/obsgo_${arch} -v *.go
done 
echo "Build done."
echo "$(file ./build/obsgo_*)"