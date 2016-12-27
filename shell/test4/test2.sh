#!/bin/bash

for file in ./*
do
	if [ -f $file ]
	then
		echo $file
	fi
done

exit 0
