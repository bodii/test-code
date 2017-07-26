[ 条件判断函数 ]



[ if(expr,v1,v2) 函数 ]
[例] select id, grade, if(grade>=60,'PASS','FAIL') from t1;
+------+-------+-----------------------------+
| id   | grade | if(grade>=60,'PASS','FAIL') |
+------+-------+-----------------------------+
| 1001 |    90 | PASS                        |
| 1002 |    50 | FAIL                        |
| 1003 |    60 | PASS                        |
| 1004 |  NULL | FAIL                        |
+------+-------+-----------------------------+



[ ifnull(v1,v2) 函数 ]
[例] select id,grade,ifnull(grade,'NO GRADE') from t1;
+------+-------+--------------------------+
| id   | grade | ifnull(grade,'NO GRADE') |
+------+-------+--------------------------+
| 1001 |    90 | 90                       |
| 1002 |    50 | 50                       |
| 1003 |    60 | 60                       |
| 1004 |  NULL | NO GRADE                 |
+------+-------+--------------------------+




[ case 函数 ]

1. case when expr1 then v1 [when expr2 then v2...] [else vn] end
case 表示函数开始，end 表示函数结束。如果表达式expr1成立时，返回v1
的值。如果表达式expr2成立时，返回v2的值。依次类推，最后遇到else时，返
回vn的值。

[例] select id,grade,case when grade>60 then 'GOOD' when grade=60 then 'PASS' else 'FAIL' end level from t1;
+------+-------+-------+
| id   | grade | level |
+------+-------+-------+
| 1001 |    90 | GOOD  |
| 1002 |    50 | FAIL  |
| 1003 |    60 | PASS  |
| 1004 |  NULL | FAIL  |
+------+-------+-------+


2. case expr when e1 then v1 [when e2 then v2...] [else vn] end
如果表过式 expr 取值等于e1时，返回v1的值，如果表达式expr取值等于e2
时，返回v2的值。依次类推，最后遇到else时，返回vn的值。case 表示函数
开始，end 表示函数结束。

[例] select id,grade,
	case grade when 90 then 'GOOD' 
			when 60 then 'PASS' 
			when 50 then 'FAIL' 
		else 'NO GRADE' 
	end level from t1;
+------+-------+----------+
| id   | grade | level    |
+------+-------+----------+
| 1001 |    90 | GOOD     |
| 1002 |    50 | FAIL     |
| 1003 |    60 | PASS     |
| 1004 |  NULL | NO GRADE |
+------+-------+----------+




[ 系统信息函数 ]
version()							返回数据库的版本号
connection_id()						返回服务器的连接数
database(),schema()					返回当前数据库名
user(),system_user(),session_user()	返回当前用户
currnet_user(),current_user			返回当前用户
charset(str)						返回字符串str的字符集
collation(str)						返回字符串str的字符排列方式
last_insert_id()					返回最近生成的auto_increment值




[ 加密函数 ]
password(str)			主要用来给用户的密码加密
md5(str)
encode(str,pswd_str)	可使用pswd_str加密str。加密结果是一个二进制数，
						必须使用blob类型的字段来保存它。
decode(crypt_str,pswd_str) 使用pswd_str解密crypt_str。




[ 其他函数 ]
format(x,n)			用来格式化数字x
ascll(s) 			返回字符串s的第一个字符的ASCII码
bin(x)				返回x的二进制编码
hex(x)				返回x的十六进制编码
oct(x)				返回x的八进制编码
conv(x,f1,f2)		将x从f1进制数变成f2进制数
inet_aton(IP)		将IP转换为数字
inet_ntoa(n)		将数字转换成IP
get_loct(name,time)	定义一个名为name、持续时间长度为time秒的锁
					如果锁定成功，返回 1;如果尝试超时，返回0;如果错误，返回null。
release_lock(name)	解除名称为name的锁。成功：1,超时：0,失败：null
is_free_lock(name)	判断是否使用名为name的锁。使用:0,否则：1。
benchmark(count,expr) 表达式expr重复执行count次，然后返回执行时间。该函数可以判断
						mysql处理表达式的速度。		

[例] select format(235.3456,3),format(234.3454,3);
+--------------------+--------------------+
| format(235.3456,3) | format(234.3454,3) |
+--------------------+--------------------+
| 235.346            | 234.345            |
+--------------------+--------------------+


[加锁 例] select get_lock('mysql',10);
+----------------------+
| get_lock('mysql',10) |
+----------------------+
|                    1 |
+----------------------+
1：成功
[查看锁 例]	select is_free_lock('mysql');
+-----------------------+
| is_free_lock('mysql') |
+-----------------------+
|                     0 |
+-----------------------+
0：存在

[解锁 例]  select release_lock('mysql');
+-----------------------+
| release_lock('mysql') |
+-----------------------+
|                     1 |
+-----------------------+
1：解锁成功


[改变字段类型 例]  select now(),cast(now() as date),convert(now(),time);
+---------------------+---------------------+---------------------+
| now()               | cast(now() as date) | convert(now(),time) |
+---------------------+---------------------+---------------------+
| 2017-02-07 18:19:54 | 2017-02-07          | 18:19:54            |
+---------------------+---------------------+---------------------+



