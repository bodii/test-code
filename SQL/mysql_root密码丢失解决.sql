步骤：

1.使用 skip-grant-tables选项启动mysql服务
skip-grant-tables选项将使mysql服务器停止权限判断，任何用户都有访问数据
库的权力。这个选项是跟在mysql服务的命令的后面。
windows操作系统中， 使用mysqld或者mysqld-nt来启动mysql服务。也可以使用
net start mysql命令，来启动mysql服务。

mysqld命令如下：
mysqld --skip-grant-tables

mysqld-nt命令如下：
mysqld-nt --skip-grant-tables

net start mysql 命令如下:
net start mysql --skip-grant-tables


linux 操作系统中，使用mysqld_safe来启动mysql,也可以使用/etc/init.d/mysql
来启动mysql服务。

mysqld_safe命令如下：
mysqld_safe --skip-grant-tables user=mysql

使用/etc/init.d/mysql执行如下：
/etc/init.d/mysql start --mysqld --skip-grant-tables

启动mysql服务后，就可以使用root用户登录了。


2.登录root用户，并且设置新的密码
通过上述方式启动mysql服务以后，可以不输入密码就登录root用户。登录以后，
可以使用update语句来修改密码。
mysql -u root

update mysql.user set password=password('root') where user='root' and host='localhost';

**这里必须使用update语句来更新mysql.user,而不能使用set语句。


3.加载权限表
修改完密码以后，必须用flush privileges 语句来加载权限表。加载权限表后，新密码才能有效。
而且，mysql服务器开始进行权限认证。用户必须输入用户名和密码才能登录mysql数据库
flush pirvileges


