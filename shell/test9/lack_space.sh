#!/bin/bash 

ANSWER="yes"

# [ 是一个命令，所有后面的元素都是它的参数
if [$ANSWER = "yes" ]
	echo $?
then
	echo "The answer is yes, so do operation"
	echo "...."
fi

echo "Finished."

exit 0
