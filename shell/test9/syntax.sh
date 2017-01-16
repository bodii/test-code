#!/bin/bash -nvx

ANSWER="yes"
if [ $ANSWER = "yes" ]
then
	echo "The answer is yes, so do operation."
	echo "..."
# 故意注释掉fi关键字，产生一个语法错误
#fi

echo "Finished."
exit 0
