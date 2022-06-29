#! /bin/sh

# set -o xtrace


ACTION=$1

ROOT_PATH=$(cd "$(dirname "$0")" && pwd)
APP_HOME=$ROOT_PATH/$APP_NAME

START_TIME=$(date -u '+%Y-%m-%dT%H:%M:%SZ')

STATUS_PORT=8004
DF_BIN=$APP_HOME/target/manager
SERVICE_PID=$APP_HOME/${APP_NAME}.pid
mkdir -p $APP_HOME/logs
SERVICE_OUT=$APP_HOME/logs/service_stdout.log
NGINXCTL=/home/admin/cai/bin/nginxctl



usage() {
    echo "Usage: $PROG_NAME {stop|start|restart|status}"
    exit 1;
}

start() {
    echo "INFO:app star at ${START_TIME}"
    $DF_BIN > $SERVICE_OUT 2>&1 &
    echo $! > $SERVICE_PID

    "$NGINXCTL" start
    start_waf

    echo "INFO:app start end, pid: $SERVICE_PID"
}


start_waf(){
    # 修改 xagent 应用名
    sed -i "s/^-application=.*/-application=$APP_NAME/g" /home/admin/xagent/conf/xagent.flags
    # OXS 区需要修改 xagent 连接的 redis 地址（弹内不需要）
    if [ -n "$ENV_OXS" ]; then
       sed -i "s/^-redis_sentinel=.*/-redis_sentinel=oxs.security-faraday-sentinel-0.alibaba-inc.com:50000/g" /home/admin/xagent/conf/xagent.flags
    fi

    #启动安全的agent waf
    startpath=`pwd`
    cd /home/admin/xagent/;
    /home/tops/bin/supervisord -c /home/admin/xagent/supervisord.conf ;
    /home/tops/bin/supervisorctl -c /home/admin/xagent/supervisord.conf restart xagent;
    cd $startpath
}

stop() {
    # gracefully shutdown
    pid=`cat $SERVICE_PID`
    pid=${pid:-$(ps -eo pid,ppid,command |grep "$DF_BIN" |awk '$3 !~ /grep/ {print $1}')}

    kill -2 $pid
    WAIT=10
    while status > /dev/null && [ $WAIT -gt 0 ] ; do
        sleep 1;
        WAIT=$(( WAIT - 1 ))
        echo "waiting for terminating, ${WAIT}s...";
    done

    # force shutdown
    if [ $WAIT -eq 0 ]; then
        echo "force terminating"
        kill -9 $pid
        sleep 1
    fi

    $NGINXCTL stop
    /home/tops/bin/supervisorctl -c /home/admin/xagent/supervisord.conf stop xagent

    echo "status: $(status)"
}

status() {
   cat /dev/null | nc -w 5  localhost $STATUS_PORT > /dev/null 2>&1
   STATUS=$?
   if [ $STATUS -eq 0 ]; then
        echo "up"
   else
        echo "down"
   fi
   return $STATUS
}

case "$ACTION" in
    start)
        start
        ;;
    stop)
        stop
        ;;
    restart)
        stop
        start
        ;;
    status)
        status
        ;;
    upgrade)
        stop
        start
        ;;
   *)
        usage
        ;;
esac