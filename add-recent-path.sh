#! /bin/sh
RECDIR=`dirs | head -20`
DIR_LIST=(${RECDIR// /\\n})
AddRecFolder=$HOME/.config/dashboard-shell/.AddRecFolder
RecFolderPath=$HOME/.config/dashboard-shell/RecFolder
if [ ! -f $AddRecFolder ]; then
    touch $AddRecFolder
    existRFP=false
else
    sed -i '1!G;h;$!d' $AddRecFolder
fi
echo ${DIR_LIST[@]} >> $AddRecFolder
awk '!a[$0]++' $AddRecFolder > $RecFolderPath
if $existRFP ; then
    sed -i '1!G;h;$!d' $RecFolderPath
fi
sed -i '21,$d' $RecFolderPath
cp $RecFolderPath $AddRecFolder

