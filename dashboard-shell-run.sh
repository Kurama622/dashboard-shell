COMMAND=`dashboard-shell`
clear
eval $COMMAND

cur_sh=`echo $SHELL | awk -F "/" '{print $NF}'`
sh_num=`ps | grep $cur_sh | wc -l`

if [ $sh_num -le 1 ]; then
    eval $SHELL
fi

#echo $COMMAND
