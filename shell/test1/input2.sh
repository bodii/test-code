#!/bin/bash

# 行号开始
linenumber=1

while read oneline
do
	echo -en "$linenumber : $oneline\n"

	linenumber=`expr $linenumber + 1`

done < /etc/passwd 

exit 0
