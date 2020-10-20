#!/bin/bash

path=""
path=`pwd`/${path}
file=${path}file.txt

LineNumber=0
while read Line || [[ -n ${Line} ]]
do
    LineNumber=`expr $LineNumber + 1`
    if [ $LineNumber == 10 ]
    then 
        echo $Line
    fi
done < ${file}