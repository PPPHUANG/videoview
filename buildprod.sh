#! /bin/bash

cd ~/go/src/video_server/api
 go build -o ../bin/api

cd ~/go/src/video_server/scheduler
 go build -o ../bin/scheduler

cd ~/go/src/video_server/streamserver
 go build -o ../bin/streamserver

cd ~/go/src/video_server/web
 go build -o ../bin/web
