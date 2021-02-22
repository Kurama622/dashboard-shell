#!/bin/sh

CONFIG=$HOME/.config/dashboard-shell
mkdir $CONFIG -p
cp folders.ini files.ini config.ini dashboard-shell-run.sh $CONFIG
sudo cp dashboard-shell /usr/local/bin/dashboard-shell



