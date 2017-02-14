mysql 用户主要包括普通用户和root用户。
root用户是超级管理员，拥有所有权限。root用户的权限包括创建用户、
删除用户和修改普通用户的密码等管理权限。而普通用户只拥有创建时
赋予它的权限。



[ 权限表 ]
名为 mysql 的数据库。mysql数据库下面存储的都是权限表。用户登录
以后，mysql 数据库系统会根据这些权限表的内容为每个用户赋予相应
的权限。这些权限表中最重要的是user表、db表和host表、tables_priv
表、columns_priv表和proc_priv表等。


user表是mysql中最重要的一个权限表。
有39个字段，可分为4类，分别是用户列、权限列、安全列和资源控制列

tables_priv 表可以对单个表进行权限设置。
columns_priv表可以对单个数据列进行权限设置。
procs_priv表可以对存储过程和存储函数进行权限设置




2.账户管理
登录mysql：
mysql -h hostname|hostIP -P prot -u username -p DatebaseName -e 'sql 语句'

-h 参数后面接主机名或者主机IP
-P mysql服务的端口号，默认为3306
-u 用户名
-p 密码
DatabaseName 参数指明登录到哪一个数据库中。
-e 参数后面可以直接加sql语句。登录mysql服务器以后即可执行这个sql语句，然后退出
mysql服务。




[ 新建普通用户 ]

1.create user
语法：
create user user[identified by [password] 'password']
			[,user [identified by [password] 'password']]...

user 由用户名(user) 和主机名(host)构成;
identified by 关键字用来设置用户的密码;如果是一个普通的字符串，
就不需要使用password关键字。
create user语句可以同时创建多个用户。新用户可以没有初始密码。

[例] create user 'test1'@'localhost' identified by 'test1';



2.insert 创建普通用户
语法：
insert into mysql.user(Host,User,Password) values('hostname','username',
password('password'));


[例 version :5.7.17]
 insert into mysql.user(host,user,authentication_string,ssl_cipher,x509_issuer,x509_subject) values
('localhost','test2',password('test2'),'','','');

[ 使用用户生效 ]
flush privileges;
使用用上一个命令创建成功后，使用这个命令使其生效。
使用这个命令可以从mysql数据库中的user表中重新装载权限。
** 但是执行flush命令需要reload权限。


3. 用grant语句创建普通用户
在创建用户时，可以为用户授权。但必须拥有grant。
语法：
grant priv_type on database.table
	to user [identified by [password] 'password']
	[,user [identified by [password] 'password']...

其中，priv_type参数表示新用户的权限;

[例]
grant select on *.* to 'test3'@'localhost' identified by 'test3';

** grant 语句不仅可以创建用户，也可以修改用户密码。而且，还可以设置
用户权限。因此，grant 语句是mysql中一个非常重要的语句。




[ 删除用户 ]
1.用drop user语句来删除普通用户
drop user user[,user]...;

[例] drop user test1@localhost;


2.用delete语句来删除普通用户
语法：
delete from mysql.user where host='hostname' and user='username';

[例] delete from mysql.user where host='localhost' and user='test2';

** 删除成功后执行 
flush privileges; 
使其生效。mysql数据库中的user表中重新装载权限。




[ root用户修改自己的密码 ]
1. 使用mysqladmin命令来修改root用户的密码
语法：
mysqladmin -u username -p password 'new_password';

2. 修改mysql数据库下的user表
语法：
update mysql.user set password=password('new_password')
	where user='root' and host='localhost';

3. 使用set语句来修改root用户的密码
语法：
set password=password('new_password');
新密码必须使用password()来加密。
修改后可使用 flush privileges;




[ root 用户修改普通用户密码 ]
1.作用set语句
语法：
set password for 'username'@'hostname'=password('new_password');


2.修改mysql数据库下的user表
语法：
update mysql.user set passwrod=password('new_password')
	where user='username' and host='localhost';


3.用grant语句
语法：
grant priv_type on database.table
	to user [identified by [password] 'password'];
