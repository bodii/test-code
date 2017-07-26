[ 表的导出和导入 ]
mysql数据库中的表可以导出成文本文件、xml文件或者html文件

1. 用select ... into outfile 导出文本文件
语法：
select [列名] from table [where 语句]
		into outfile '目标文件' [option];

option 参数有常用的5个选项：
fields terminated by '字符串':设置字符串为字段的分隔符，默认值是‘\t’;
fields enclosed by '字符': 设置字符来括上字段的值。默认情况下不使用任何符号;
fields optionally enclosed by '字符': 设置字符来括上char、varchar和text等字
符型字段。默认情况下不使用任何符号;
fields escaped by '字符': 设置转义字符，默认值为‘\’;
lines starting by '字符串': 设置每行开头的字符，默认情况下野无任何字符;
lines terminated by '字符串': 设置每行的结束符，默认值是'\n';


2.用mysqldump命令导出文本文件
语法：
mysqldump -u root -pPassword -T 目标目录 dbname table [option];

option选项如下：
--fields-termainated-by=字符串： 设置字符串为字段的分隔符，默认值是'\t';
--fields-enclosed-by=字符： 设置字符来括上字段的值;
--fields-optionally-enclosed-by=字符： 设置字符括上char、varchar和text等
--fields-escaped-by=字符：设置转义字符;
--lines-terminated-by=字符串：设置每行的结束符。

[例] mysqldump -u root -phuang -T ~/ test student 
		"--fields-terminated-by=," "--fields-optionally-enclosed-by=""


3.用mysql命令导出文本文件
语法：
文本：
mysql -u root -p -e 'select 语句' dbname > name.txt;
xml：
mysql -u root -p  --xml | -X -e 'select 语句' dbname > name.xml;
html：
mysql -u root -p  --html | -H -e 'select 语句' dbname > name.html;



4.用load data infile方式导入文本文件
语法：
load data [local] infile filename into table tablename [option];



5.用mysqlimport命令导入文本文件
mysqlimport -u root -p  [--Local] dbname file [option]





