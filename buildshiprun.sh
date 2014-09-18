#!/bin/bash

GOARCH=arm GOARM=7 GOOS=linux go build
scp gobot-demo root@192.168.7.2:~
ssh -t root@192.168.7.2 "./gobot-demo"
