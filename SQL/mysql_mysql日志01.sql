[ mysql日志 ]
mysql日志是用来记录mysql数据库的客户端连接情况、sql语句的执行情况和
错误信息等。
日志可以分为4种，分别是二进制日志、错误日志、通用查询日志和慢查询日志。
二进制日志也叫变更日志（update log），主要用于记录数据库的变化情况。
慢查询日志：记录执行时间超过指定时间的操作。



1. 启动和设置二进制日志
默认情况下，二进制日志功能是关闭的。通过my.cnf或my.ini文件的log-bin选项可以
开启二进制日志。
如下：
在my.cnf (linux) 或 my.ini(windows)
文件中：
[mysqld]
log-bin [=DIR \ [filename] ]
其中，DIR参数指定二进制文件的存储路径; filename参数指定二进制文件的文件名，
其形式为filename.number,number的形式为000001、000002等。每次重启mysql服务
后，都会生成一个新的二进制日志文件，这些日志文件的‘number’会不断递增。除了
生成上述文件外，还会生成一个名为filename.index的文件。这个文件中存储所有二
进制日志文件的清单。



2.查看二进制日志
语法：
mysqlbinlog filename.number


3.1删除所有二进制日志
语法：
reset master;

3.2根据编号来删除二进制日志
语法：
purge master logs to 'filename.number';

3.3根据创建时间来删除二进制日志
purge master logs to 'yyyy-mm-dd hh:MM:ss';


4 使用二进制日志还原数据库
命令：
mysqlbinlog filename.number | mysql -u root -p



5.暂时停止二进制日志功能
语句：
set sql_log_bin=0;
重新开启：
set sql_log_bin=1;




[ 错误日志 ]

1.启动和设置错误日志
设置：
my.cnf(linux) 或 my.ini(windows)
[mysqld]
log-error=Dir / [filename]


2.查看错误日志
可使用vi或gedit工具来查看

3.删除错误日志
命令：
mysqlamdin -u root -p flush-logs
执行该命令后，数据库系统会自动创建一个新的错误日志。旧的错误日志仍然保留着，只是
已经更名为 filename.err-old。
除了mysqladmin命令外，也可以使用flush logs语句来开启新的错误日志。使用该语句之前
必须登录到mysql数据库中。



[ 通用查询日志 ]
1.开启通用查询日志:
my.cnf(linux) 或 my.ini(windows)下
[mysqld]
log [=dir \ [filename]]


2. 查看通用查询日志
可使用vi或gedit。


3.删除通用查询日志
mysqladmin -u root -p flush-logs




[ 慢查询日志 ]
1.启动慢查询日志
my.cnf(linux) 或 my.ini(windows) 下：
[mysqld]
log-slow-queries [=dir\ [filename] ]
long_query_time=n
如果不指定文件名，默认文件名为hostname-slow.log,
hostname是mysql服务器的主机名。
"n"参数是设定的时间值，该值的单位是秒。如果不设置long_query_time选项，默认
为10秒。

2. 查看慢查询日志
可以使用普通文本文件查看。


3.删除慢查询日志
语法：
mysqladmin -u root -p flush-logs



