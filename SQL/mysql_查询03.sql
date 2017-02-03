[ 合并查询 ]

合并查询结果是将多个 select 语句的查询结果合并到一起。
语法：
	select 语句1  union | union all  select 语句2
union | union all select 语句n

** union 关键字和 union all 关键字都可以合并查询结果，但是两者有一点区别
union 关键字合并查询结果时，需要将相同的记录消除掉。而 union all 关键字则
相反，它不会消除掉相同的记录，而是将所有的记录合并到一起。
**！ 合并查询，只能合并字段相同的表内容。 

[例1] select id,name from a union select id,name from b ;

[例2] (select id from a) union (select id from b) 
order by id desc limit 0,10;




[ 别名 ]
属性名 [as] 别名
在给字段起别名时，应注意字段的的别只是用来显示，需不能出现在表名后的关键语句中。
[错例] select name as n from user where n='Tom';
[正例] select name as n from user where name='Tom';

** 表的别名不能与该数据库的其他表同名。字段的别名不能与该表的其他字段同名。
在条件表达式中不能使用字段的别名，否则将会出现"ERROR"




[ 正则表达式查询 ]
正则表达式查询，一般用于非常复杂的查询。
语法： 属性名 regexp '匹配方式'
mysql 的正则表达式与其他的编程语言中的正则表达式基本一致。




[ like ]
like 一般与通配符 % 和 _ 两个使用
其中 % 可以匹配任意长度任意字符，也可长度为0,即没有字符
其中 _ 只能匹配长度为一个的单一字符，若原代码为'%on__'即可以匹配长度为二的
任意字符

若要匹配字符串中有 % 和 _ 的字符串，则需要用到关键字escape;
使用escape,转义字符后的%或_就不作为通配符了。
[例] select name from user where name like '%TOM/_%' escape '/';
其中字符'/'为转义字符，可以将其后面的一个通配符失效，变成被 匹配的字符。
即匹配名字中含'TOM_'的学生数据。


 
