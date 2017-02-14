1.授权
语法：
grant priv_type [(column_list)] on database.table
	to user [indentified by [password] 'password']
	[,user [indentified by [password] 'password']...
	[with with_option [with_option] ...]
其中，priv_type 参数表示权限的类型;
column_list表示权限用于哪些列上，没有该参数时作用于整个表上;
user 用户名和主机名构成，形式是“‘user’@‘localhost’”;
identified by 用来为用户设置密码;
password 用户的新密码
with 关键字后面带有一个或多个with_option参数。这个参数有5个选项，如下：
grant option: 被授权的用户可以将这些权限赋予给别的用户;
max_queries_per_hour count: 设置每个小时可以允许执行count次查询;
max_updates_per_hour count: 设置每个小时可以允许执行count次查询;
max_connections_pre_hour count: 设置每小时可以创建count连接;
max_user_connections count: 设置单个用户可以同时具有的 count个连接数

[ 例 ]
grant select, update on *.*
	to 'test1'@'localhost' indentified by 'test'
	with grant option;


3.收回权限
语法：
revoke priv_type [(column_list)]...
	on database.table
	from user [,user]...

收回全部权限的语法:
revoke all privileges,grant option from user [,user]...

[例] revoke update on *.* from test1@localhost;

select host,user,update_priv from mysql.user where user='test1';
+-----------+-------+-------------+
| host      | user  | update_priv |
+-----------+-------+-------------+
| localhost | test1 | N           |
+-----------+-------+-------------+


移出test1的所有权限：
revoke all privileges,grant option from 'test1'@'localhost';

select host,user,select_priv,update_priv,grant_priv from mysql.user where
user='test1'\G
*************************** 1. row ***************************
       host: localhost
       user: test1
select_priv: N
update_priv: N
 grant_priv: N


4.查看权限
select * from mysql.user;
或
show grants for 'username'@'hostname';

