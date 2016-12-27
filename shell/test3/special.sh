#!/bin/bash

NEWDIR=./newdir
NEWFILE=newfile
LOGFILE=
# 指定log文件为special.pid.log
ls `basename $0.sh.*.log` > /dev/null 2>&1
if [ "$?" != "0" ]
then
	LOGFILE=`basename $0.sh`.$$.log
else
	LOGFILE=$(ls `basename $0.sh.*.log`)
fi

# 在主目录下创建一个目录
if [ ! -d $NEWDIR ]
then
	mkdir $NEWDIR
fi

if [ $? -eq 0 ]
then
	#成功地创建log目录
	echo "`date`:Creating Directory '$NEWDIR' succeed."
	# 在log目录中创建log文件
	echo "`date`:Creating Directory '$NEWDIR' succeed." >> $LOGFILE

	#改变当前工作目录为./newdir
	cd $NEWDIR

	if [ $? -eq 0 ]
	then
		echo "`date`:Changing Current Directory to '$NEWDIR' succeed."
		# 当前工作目录改变，因此使用OLDPWD
		echo "`date`:Changing current Directory to '$NEWDIR' succeed." >> $OLDPWD/$LOGFILE

		# 然后创建一个新的文件
		if [ ! -f "$NEWFILE" ]
		then
			touch $NEWFILE
		fi

		if [ $? -eq 0 ]
		then
			echo "`date`:Creating new file in directory '$NEWDIR' succeed."
			#当前工作目录改变，因此使用OLDPWD
			echo "`date`:Creating new file in directory '$NEWDIR' succeed." >> $OLDPWD/$LOGFILE
		else
			# $? != 0 意味着失败
			echo "`date`:Creating newfile in directory '$NEWDIR' failed."

			# 当前工作目录改变，因此作用OLDPWD
			echo "`date`:Creating newfile in directory '$NEWDIR' failed." >> $OLDPWD/$LOGFILE
		fi
	else
		# $? != 0 意味着失败
		echo "`date`:Changing Current Directory to '$NEWDIR' failed!"
		# 当前工作目录改变，因此使用OLDPWD
		echo "`date`:Changing Current Directory to '$NEWDIR' failed!" >> $OLDPWD/$LOGFILE
	fi
else
	# $? != 0 意味着失败
	echo "`date`:Creating Directory '$NEWDIR' failed."
	echo "`date`:Creating Directory '$NEWDIR' failed." >> $LOGFILE
fi

while :
do
	read -p "read $LOGFILE? (n or y):"
	
	case "$REPLY" in
		[Nn]|[Nn][Oo][Tt])

			echo "Bye."
			# 如果输入的是大\小写的not或n的话退出
			exit 0
			;;
			
		[Yy]|[Yy][Ee][Ss])
		
			cat -n $OLDPWD/$LOGFILE | less
			exit 0
			;;

		*)
	esac
done
echo 

exit 0
