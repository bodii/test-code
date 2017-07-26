#!/bin/bash


ls `basename $0.*.log` &> /dev/null

if [ "$?" != "0" ]
then
	echo "create file"
	touch "$0.$$.log"
else
	echo "exists file:"
	echo "$(ls `basename $0.*.log`)"
fi

exit 0

