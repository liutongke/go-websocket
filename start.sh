#!/bin/sh
ps -ef | grep main | grep -v grep | awk '{print $2}' | xargs kill -9
nohup ./main &
