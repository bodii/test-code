#!/bin/bash

recursive ()
{
	if [ $1 -eq 1 ]
	then
		echo 1
		return 0
	fi

	local TEMP=`expr $1 - 1`

	local REVIOUS=`recursive $TEMP`

	RESULT=`expr $1 \* $REVIOUS`

	echo $RESULT

	return 0
}

echo 
echo -n " Enter a number(<=20) to do n! operation(type 'quit' to exit):"
read INPUT

until [ "$INPUT" = "quit" -o "$INPUT" = "q" ]
do
	case $INPUT in
		[1-9]|[1-9][0-9]|20)
			
			OUTCOME=`recursive "$INPUT"`
			echo " After computing, $INUPT!=$OUTCOME"
			echo
			;;
		*)
			echo
			echo " Not a valid number, Enter again!"
			;;
	esac

	echo -n " Enter a number(<=20)to do n! operation(type \"quit\" or 'q' to exit):"
	read INPUT
done

echo
echo " Bye."

exit 0
