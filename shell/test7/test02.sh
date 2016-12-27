#!/bin/bash

one_file='01.txt'

two_file='02.txt'
if [ -f "$1" -a -n "`cat $1`" ]
then
	echo 1
else
	echo 2
fi

exit 0
