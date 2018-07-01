#! /bin/bash

cp -R ./templates ./bin/

mkdir ./bin/videos

cd bin

nphup ./api &
nohup ./scheduler &
nohup ./streamserver &
nohup ./web &

echo "deploy finished"