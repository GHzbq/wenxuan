#!/bin/bash

case $1 in 
	start)
		nohup ./wenxuan 2>&1 > /dev/null 2>&1 &
		echo "服务已启动..."
		sleep 1
	;;
	stop)
		killall wenxuan 
		echo "服务已停止..."
		sleep 1
	;;
	restart)
		killall wenxuan 
		sleep 1
		nohup ./wenxuan 2>&1 > /dev/null 2>&1 &
		echo "服务已重启..."
		sleep 1
	;;
	*) 
		echo "$0 {start|stop|restart}"
		exit 4
	;;
esac
