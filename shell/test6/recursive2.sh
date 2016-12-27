#!/bin/bash

recursive ()
{
	if [ $1 -eq 1 ]
	then
		echo 1
		return 0
	fi

	local TEMP=`expr $1 - 1`

	local PREVIOUS=`recursive $TEMP`

	RESULT=`expr $1 \* $PREVIOUS`

	echo $RESULT

	return 0
}

echo 
echo -n "Enter a number (<20) to do n! operation(type \"quit\" to exit):"
read INPUT

until [ "$INPUT" = "quit" ]
do
	case $INPUT in 
		[1-9]|[1][0-9]|20)
			OUTCOME=`recursive "$INPUT"`
			echo -e "\tAfter computing,$INPUT!=$OUTCOME"
			echo
			;;
		*)
			echo
			echo -e "\tNot a valid number,Enter again!"
			;;
	esac

	echo -n "Enetr a number (<20) to do n! operation (type \"quit\" to exit):"
	read INPUT

done

echo
echo -e "\tBye."
exit 0

