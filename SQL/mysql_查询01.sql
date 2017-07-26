[ select的基本语法 ]
select 属性列表
		from  表名和视图列表
		[ where 条件表达式1 ]
		[ group by 属性名1 [ having 条件表达式2 ] ]
		[ order by 属性名2 [ asc | desc ] ]

如果group by 子句后带着 having 关键字，那么只有满足“条件表达式2"中指定的条件的才能够输出。
group by 子句通常 和 count() \ sun() 等聚合函数一起使用。

[ 单表查询 ]
select * from 表名;

[ where ]

查询条件            符号或关键字
--------------------------------------------------
比较                =  <  <=  >  >=  !=  <>  !>  !<
指定范围            between and \ not between and 
指定集合            in \ not in
匹配字符            like \ not like
是否为空值          is null  \ is not null
多个查询条件         and \ or
---------------------------------------------------
"<>" 等价于 "!=" 
"!>" 等价于 "<="
"!<" 等价于 ">="



[ between and ]  
[not] between  取值1  and  取值2

例如：select * from test where age between 15 and 25;



[ like ]
[not] like '字符串'
not 是可选参数 
"字符串"，必须加单引号或双引号。
"字符串"参数的值可以是一个完整的字符串，也可以是包含百分号( % )
或( _ )的通配符。但 % 和 _ 有很大的差别：
"%" 可以代表任意长度的字符串，长度可以为0
"_" 只能表示单个字符.
 


[ in() and or ]
例如： id in(1,3,4) and age=25 or sex='女'
会首先查找id 是1\3\4的行并且要满足 age=25 或 sex='女'的数据



[ 查询结果不重复 ]
select distinct 属性名
[例] select distinct d_id from test;



[ order by 字段1 asc,字段2 desc ]
select * from test order by d_id asc, age desc;
排序过程中，先按照d_id字段进行排序。遇到 d_id字段的值相等的情况时，
再把d_id值相等的记录按照age字段时行排序。


[ group by ]
group by 属性名 [having 条件表达式][ with rollup ]
with rollup 关键字将会在所有记录的最后加上一条记录。该记录是上面所有
记录的总和。
如果使用 with rollup 则 属性名必须使用集合函数，且会在最后加一条结果。
集合函数包括：count() \ sum() \ avg() \ max() \ min()   (agv()取平均数)
如：
select min(id) from test group by id with rollup ;
+---------+
| min(id) |
+---------+
|       1 |
|       2 |
|       1 |
+---------+

select max(id) from test group by id with rollup ;
+---------+
| max(id) |
+---------+
|       1 |
|       2 |
|       2 |
+---------+
select sum(id) from test group by id with rollup ;
+---------+
| sum(id) |
+---------+
|       1 |
|       2 |
|       3 |
+---------+

[ group by ... having ... ]
如 select sex,count(sex) from test group by sex having
count(sex) >=3;
having 返回的是boolean;


[ group_concat() 函数 ]
如 select group_concat(name) from test group by name with rollup;
+--------------------+
| group_concat(name) |
+--------------------+
| aa                 |
| bb                 |
| aa,bb              |
+--------------------+
查询结果显示，group_concat(name)显示了每个分组的name字段的值。同时，
最的一条记录的group_concat(name)列的值则好是上面分组name取值的总各。

** 注： group by 中 having 和 with rollup 不能同时出现。




