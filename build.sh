#!/usr/bin/env bash

CONFIG=$HOME/.config/dashboard-shell
mkdir $CONFIG -p
cp config.ini dashboard-shell-run.sh $CONFIG

if [ ! -d /usr/local/bin  ]; then
  sudo mkdir /usr/local/bin
fi

sysOS=`uname -s`
if [ $sysOS == "Darwin" ];then
    sudo cp dashboard-shell-mac /usr/local/bin/dashboard-shell
elif [ $sysOS == "Linux" ];then
    sudo cp dashboard-shell /usr/local/bin/dashboard-shell
fi
