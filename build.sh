#!/bin/bash
echo "Initialing to build system"
mkdir build
cd build
echo "building..."
echo "$(tput setaf 3)Red text $(tput setab 7)and white background$(tput sgr 0)"
go build ../app/smartid/smartid.go
echo "Build success!"
