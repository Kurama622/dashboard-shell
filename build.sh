#!/usr/bin/env bash

if [ ! -d /usr/local/bin  ]; then
  sudo mkdir /usr/local/bin
fi

sysOS=`uname -s`
if [ $sysOS == "Darwin" ];then
    sudo cp dashboard-shell-mac /usr/local/bin/dashboard-shell
    if [ ! -x "$(command -v figlet)" ]; then
        echo "please install figlet!"
        exit
    fi

elif [ $sysOS == "Linux" ];then
    sudo cp dashboard-shell /usr/local/bin/dashboard-shell
    if [ ! -x "$(command -v figlet)" ]; then
        os=$(cat /etc/os-release 2>/dev/null | grep ^ID= | awk -F= '{print $2}')
        case "$os" in
            ubuntu)
                sudo apt install figlet;;
            arch)
                sudo pacman -S figlet;;
        esac
    fi
fi


CONFIG=$HOME/.config/dashboard-shell
mkdir $CONFIG -p
cp config.ini dashboard-shell-run.sh add-recent-path.sh $CONFIG
pwd > $CONFIG/.AddRecFolder
cp $CONFIG/.AddRecFolder $CONFIG/RecFolder

CURSHELL=`echo $SHELL | awk -F "/" '{print $NF}'`

if [ $CURSHELL = bash ]; then
    if [ ! -f $CONFIG/.dashboard-shell.bash ]; then
        doWrite=true
    else
        doWrite=false
    fi
   cp .dashboard-shell.sh $CONFIG/.dashboard-shell.bash
   if $doWrite; then
       echo """
# dashboard-shell

[ -f $CONFIG/.dashboard-shell.bash ] && source $CONFIG/.dashboard-shell.bash""" >> $HOME/.bashrc
   fi

elif [ $CURSHELL = zsh ]; then
   if [ ! -f $CONFIG/.dashboard-shell.zsh ]; then
        doWrite=true
    else
        doWrite=false
   fi
   cp .dashboard-shell.sh $CONFIG/.dashboard-shell.zsh
   if $doWrite; then
       echo """
# dashboard-shell

[ -f $CONFIG/.dashboard-shell.zsh ] && source $CONFIG/.dashboard-shell.zsh""" >> $HOME/.zshrc
   fi

else
   if [ ! -f $CONFIG/.dashboard-shell.fish ]; then
        doWrite=true
    else
        doWrite=false
   fi
   cp .dashboard-shell.sh $CONFIG/.dashboard-shell.fish
   if $doWrite; then
       echo """
# dashboard-shell

[ -f $CONFIG/.dashboard-shell.fish ] && source $CONFIG/.dashboard-shell.fish""" >> $HOME/.fishrc
   fi
fi
