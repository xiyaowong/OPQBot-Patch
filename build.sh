#! /bin/bash

exe="opq-patch"

go build -ldflags '-s -w' -o $exe && upx -9 $exe
