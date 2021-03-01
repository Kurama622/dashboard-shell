#! /bin/sh

RECDIR=`dirs | head -20`
DIR_LIST=(${RECDIR// /\\n})
AddRecFolder=$HOME/.config/dashboard-shell/.AddRecFolder
RecFolderPath=$HOME/.config/dashboard-shell/RecFolder

if [ ! -f $AddRecFolder ]; then
    touch $AddRecFolder
    existRFP=false
else
    existRFP=true
fi

tmp=`cat $AddRecFolder`
echo ${DIR_LIST[@]} > $AddRecFolder
if $existRFP; then
    echo $tmp >> $AddRecFolder
fi
awk '!a[$0]++' $AddRecFolder > $RecFolderPath
sed -i '21,$d' $RecFolderPath
cp $RecFolderPath $AddRecFolder

