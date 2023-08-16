#! /bin/bash

set -xe
go build main.go
sudo setcap CAP_NET_BIND_SERVICE=pei ./main
./main $@

