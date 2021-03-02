#!/usr/bin/env bash

CONFIG=$HOME/.config/dashboard-shell
mkdir $CONFIG -p
cp config.ini dashboard-shell-run.sh add-recent-path.sh $CONFIG

if [ ! -d /usr/local/bin  ]; then
  sudo mkdir /usr/local/bin
fi

sysOS=`uname -s`
if [ $sysOS == "Darwin" ];then
    sudo cp dashboard-shell-mac /usr/local/bin/dashboard-shell
elif [ $sysOS == "Linux" ];then
    sudo cp dashboard-shell /usr/local/bin/dashboard-shell
fi


CURSHELL=`echo $SHELL | awk -F "/" '{print $NF}'`
[ $CURSHELL==zsh ]

if [ $CURSHELL = bash ]; then
   cp .dashboard-shell.sh $CONFIG/.dashboard-shell.bash
   echo """
# dashboard-shell

[ -f $CONFIG/.dashboard-shell.bash ] && source $CONFIG/.dashboard-shell.bash""" >> $HOME/.bashrc

elif [ $CURSHELL = zsh ]; then
   cp .dashboard-shell.sh $CONFIG/.dashboard-shell.zsh
   echo """
# dashboard-shell

[ -f $CONFIG/.dashboard-shell.zsh ] && source $CONFIG/.dashboard-shell.zsh""" >> $HOME/.zshrc

else
   cp .dashboard-shell.sh $CONFIG/.dashboard-shell.fish
   echo """
# dashboard-shell

[ -f $CONFIG/.dashboard-shell.fish ] && source $CONFIG/.dashboard-shell.fish""" >> $HOME/.fishrc
fi
