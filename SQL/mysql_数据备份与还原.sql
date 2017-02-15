[ 数据备份 ]

1.使用mysqldump命令备份
mysqldump命令可以将数据库中的数据备份成一个文本文件。表的结构和表中
的数据将存储在生成的文本文件中。
语法：
mysqldump -u username -p dbname table1 table2 ... > backupName.sql

备份某个数据库下的所有表：
mysqldump -u username -p dbname > backupName.sql

** mysqldump命令备份的文件并非一定要求后缀名为.sql,备份成其他格式的文件
也是可以的。

[例] mysqldump -u root -p test student > /home/user/student.sql

备份文件的开头记录了mysql的版本、备份的主机名和数据库名。文件中，以'--'开头的
都是sql语言的注释。以'/*!40101'等形式开头的是与mysql有关的注释。40101是mysql
数据库的版本号，这里就表示mysql 4.1.1。如果还原数据时，mysql的版本比4.1.1高，
'/*!40101' 和 '*/' 之间的内容被作为sql命令来执行。如果比这个版本低，'/*!40101'
和 '*/'之间的内容被当作注释。


2. 备份所有数据库
mysqldump -u username -p --all-databases > backupName.sql

[例] mysqldump -u root -p --all-databases > 01.sql



3. 直接复制整个数据库目录
mysql有一种最简单的备份办法，就是将mysql中的数据库文件直接复制出来。这种
方法最简单，速度也是快。使用这种方法时，最好将服务器先停止。这样，可以保
证在复制期间数据库中的数据不会发生变化。如果在复制数据库的过程中还有数据写
入，就会造成数据不一致。
这种方法虽然简单快速，但不是最好的备份方法。因为，实际情况可能不允许停止
mysql服务器。而且，这种方法对innodb存储引擎的表不适用。对于myisam存储引擎
的表，这备份和还原很方便。但是还原时最好是相同版本的mysql数据库，否则可能
存在文件类型不同的情况。



4. 使用mysqlhotcope工具快速备份
如果备份时不能停止mysql服务器，可以采用mysqlhotcopy工具。mysqlhotcopy工具
的备份方式比mysqldump命令快。
mysqlhotcopy工具是一个perl脚本，主要在linux操作系统下使用。mysqlhotcopy工具
使用 lock tables, flush tables 和 cp 进行快速备份。工作原理是，先将需要备份
的数据库加上一个读操作锁，然后，用flush tables将内存中的数据写回到硬盘上的数据
库中，最后，将需要备份的数据库文件复制到目标目录。使用mysqlhotcopy的命令如下：

mysqlhotcopy [option] dbname1 dbname2 ...  backupDir/

mysqlhotcopy工具有一些常用的选项：
--help: 用来查看mysqlhotcopy的帮助;
--allowold: 如果备份目录下存在相同的备份文件，将旧的备份文件名加上_old;
--keepold: 如果备份目录下存在相同的备份文件，不删除旧的备份文件，而是将旧文件更名;
--flushlog: 本次备份之后，将对数据库的更新记录到目志中;
--noindices: 只备份数据文件，不备份索引文件;
--user=用户名：用来指定用户名，可以用-u代替;
--password=密码：用来指定密码，可以用-p代替。使用-p时，密码与-p紧挨着。或者只使用-p,
然后用交换的方式输入密码。这与登录数据库时的情况是一样的;
--port=商品号： 用来指定访问端口，可以用-P代替;
--socket=socket文件： 用来指定socket文件，可以用-S代替。




[ 数据还原 ]

1.使用mysql命令还原
语法：
mysql -u root -p [dbname] < backup.sql

** 如果是使用--all-databases 参数备份了所有的数据库，那么还原时不需要指定数据
库。因为，其对应的sql文件包含有create database语句。


2. 直接复制到数据库目录
linux数据库目录通常在/var/lib/mysql/ /usr/local/mysql/data /usr/local/mysql/var 
这3个目录下。
使用mysqlhotcopy命令备份的数据也是通过这种方式来还原的。在linux操作系统下，
复制到数据库目录后，一定要将数据库的用户和组变成mysql。如下：
chown -R myslq.mysql dataDir
‘-R’可以改变文件夹下的所有子文件用户和组。




[ 数据库的迁移 ]

1.相同版本下
最常用和最安全的方式是使用mysqldump命令来备份数据库。然后使用mysql命令将备份
文件还原到新的mysql数据库中。这里可以将备份和迁移同时进行。假设从一个名为host1
的机器中备份出所有数据库，然后，将这些数据库迁移到名为host2的机器上。命令如下：
mysqldump -h host1 -u root --password=password1 --all-databases |
mysql -h host2 -u root --password=passowrd2

其中，‘|’ 符号表示管道，其作用是将mysqldump备份的文件送给mysql命令



2.不同版本的mysql数据库之间的迁移
高版本的mysql数据库通常都会兼容低版本，因此可以从低版本的mysql数据库迁移到高版本的
mysql数据库。对于myisam类型的表可以直接复制，也可以使用mysqlhotcopy工具。但是innodb
类型的表不可以使用这两种方法。最常用的办法是使用mysqldump命令。

