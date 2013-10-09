#!/bin/bash -e
subprojects=bloom_filter

dependencies="github.com/spaolacci/murmur3 \
              github.com/willf/bitset"

for dep in $dependencies; do
  echo "Getting $dep..."
  go get $dep
done

for subproj in $subprojects; do
  pushd $subproj > /dev/null
  echo "Building and testing $subproj..."
  go build
  go test
  popd > /dev/null
done
