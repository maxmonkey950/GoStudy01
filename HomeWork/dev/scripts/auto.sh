#!/bin/sh
echo "build docker images..."
docker ps
if [ "$?" != 0 ];then
    echo "docker not install or command not found..."
    exit 1
fi
# build image
docker build -t dev:v2 .
if [ "$?" != 0 ];then
    echo "docker build err..."
    exit 1
fi
