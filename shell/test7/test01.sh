#!/bin/bash

wc=`wc -m a.txt | tr 'a.txt' ' '`

if [ $wc -gt 1 ]
then
	echo 1
else
	echo 0
fi

exit 0
