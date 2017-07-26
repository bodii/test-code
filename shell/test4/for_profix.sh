#!/bin/bash

# 在当前目录下创建测试目录
mkdir test &> /dev/null

# 如果创建目录不成功

if [ $? -ne 0 ]
then
	echo "Directory test has already existed."
	echo "Do Nothing, Abort!"
	exit 1
else
	echo "Ctreate directory test"
fi

echo "cp *.sh files into driectory test, and prefix \"test_\"
befor each filename."

# 遍历当前目录下的所有*.sh 文件
for FILE in `ls *.sh`
	# = for FILE in *.sh
	# = for FILE in `ls -1 *.sh` -1 表示每个文件显示一行
do
	# 把文件复制到刚创建的目录下，并重命名为test_*
	cp $FILE test/test_$FILE
	# 为每一个文件添加可执行权限
	chmod u+x test/test_$FILE
done

exit 0
