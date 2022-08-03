#!/bin/sh

gox -os=linux -arch=amd64
mv easy-agent-server_linux_amd64 easy-agent-server
docker build -t 172.16.8.120:5443/dtstack-dev/easy-agent-server:4.2.1-fixregister .
docker push 172.16.8.120:5443/dtstack-dev/easy-agent-server:4.2.1-fixregister