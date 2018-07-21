#!/bin/bash

EXEC_BIN=fomo-proxy
SERVICE_FILE=fomo-proxy.service
SERVICE_PATH=/lib/systemd/system/
CONFIG_PATH=/etc/fomo-proxy/
CONFIG_FILE=/etc/fomo-proxy/config.json
LOCAL_BIN=/usr/local/bin/

CUSER=$(who am i | cut -d' ' -f1)
CGROUP=$(groups $CUSER | cut -d' ' -f1)

##############################################
# Check if run as root
##############################################
if [ xroot != x$(whoami) ]
then
    echo "You must run as root (Hint: Try prefix 'sudo' while execution the script)"
    exit
fi

echo "To install will stop the fomo-proxy service. (Y/N)?\c"
read answer
if [ $answer = 'y' ] || [ $answer = 'Y' ]; then
    echo "Contine installation..."
else
    echo "Abort install..."
    exit
fi

if [ ! -d $CONFIG_PATH ]; then
    mkdir -p $CONFIG_PATH
fi

if [ ! -f $CONFIG_FILE ]; then
    if [ ! -f config.json ]; then
	echo 'Not config.json file found, check it.'
	exit
    fi
    cp config.json $CONFIG_PATH
fi

if [ ! -x $EXEC_BIN ]; then
    echo "Not executable file found, check it."
    exit
fi

useradd fomo -U -M -s /sbin/nologin
cp $SERVICE_FILE $SERVICE_PATH
cp $EXEC_BIN $LOCAL_BIN
service $EXEC_BIN start

echo "fomo-proxy has successfully installed."
echo "run service fomo-proxy restart to make it work."
