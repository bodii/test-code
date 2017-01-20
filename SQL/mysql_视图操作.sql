 [ 视图的作用 ]

视图是在原有的表或者视图的基础上重新定义的虚拟表，这
可以从原有的表上选取对用户有用的信息。那些对用户没有
用，或者用户没有权限了解的信息，都可以直接屏蔽掉。这
样做既使应用简单化，也保证了系统的安全。视图起着类似
于筛选的作用。视图的作用归纳为如下几点：
1.使操作简单化
2.增加数据的安全性
3.提高表的逻辑独立性

[ 创建视图的语法 ]
create [algorithm = {undefined | merge | temptable }]
	view 视图名 [( 属性清单 )]
	as select 语句
	[ with [cascaded | local ] check option ];

algorithm [可选参数]，表示视图选择的算法;
“视图名” 表示要创建的视图的名称;
“属性清单” [可选参数]，其指定了视图中各个属性的名词，
默认情况下与select语句中查询的属性相同;
select语句参数是一个完整的查询语句，表示从某个表中查出某
些满足条件的记录，将这些记录导入视图中;
with check option [可选参数] 表示更新视图时要保证在该视
图的权限范围之内。



查看用户是否Select_priv和Create_view_priv权限 Y=yes   N=no
select Select_priv, Create_view_priv from mysql.user where user='root';

创建普通视图
create view department_view1
as select * from department;

创建视图
create view department_view2(name,function,location)
as select d_name,function,address
from department;


[多表上创建视图]

create algorithm=merge view
worker_view1(name,department,sex,age,address)
as select name,department.d_name,sex,2009-birthday,address
	from worker,department where worker.d_id=department.d_id
	with local check option;

worker_wiew1 的属性分别为：name、department、sex、age和 location
视图指定的属性列表对应着两个不同的表的属性列。视图的属性名与属性列表
中的属性名相同。
select 语句查询出了department表的d_name字段，还有worker表的name、sex
、birthday和address。其中，department表的d_name字段对应视图的department
字段; worker表的birthday字段进行减法操作后，对应视图的age字段。而且，视图
worker_view1的algorithm的值指定为merge。还增加了with local check option
约束。
视图可以将多个表上的操作简洁的表示出来。


[ 查看视图信息 ]
1.
show table status like '%view1'\G   （view1 为视图名称）
show table status 语句虽然也可以查看视图的基本信息，但是通常很少使用
因为，使用show table status语句查询视图信息时，各个属性显示的值都是
null。只有Comment属性显示值为VIEW。


2.
show create view 语句查看视图详细信息
语法： show create view 视图名

例如:show create view worker_view1\G
执行结果显示详细的信息。包括视图的各个属性\ with local check option条件
和字符编码（character_set_client)等信息。通过show create view 语句，可
以查看视图的所有信息。

3.在views表中查看视图详细信息
在MySQL中，所有视图的定义都存在information_schema数据库下的views表中。查
询views表，可以查看到数据库中所有视图的详细信息。查询的语句如下：
select * from information_schema.views\G
select * from information_schema.views where table_name like '%view1'\G



[ 修改视图 ]
当基本表的某些字段发生改变时，可以通过修改视图来保持视图和基本表之间一致。
mysql中通过 create or replace view 语句和 alter 语句来修改视图。

1. create or replace view 语句
该语句的使用非常灵活。在视图已经存在的情况下，对视图进行修改;视图不存在时，
可以创建视图。
语法： create or replace [algorithm= {undefined | merge | temptalbe}]
					view 视图名 [( 属性清单 )]
					as select 语句
					[with  [cascaded | local ] check option];
这里的所有参数都与创建视图的参数是一样的。


2. alter 语句修改视图
alter 语句可以修改表的定义，可以创建索引，还可以修改视图
语法： alter [ algorithm={undefined | merge | temptable}]
			view 视图名 [( 属性清单 )]
			as select 语句
			[ with  [cascaded | local ] check option ];




[ 视图的更新 ]
update 视图名 set 字段名称1=字段值1 [,字段名称=字段值] [where 条件表达式]

	// 不能被更新的视图类型 //
1.视图中包含 sum()\ count()\ max() 和 min()等 函数
2.视图中包含 union\ union all\ destinct\ group by\ 和 havig等 关键字
3.常量视图 如 create view worker_view as select 'Aric' as name;视图中的name
字段是个字符串常量'Aric',所以该视图也是不能更新的。使用update语句更新时，会出现
系统报错。
4.视图中的select中包含子查询，如create view worker_view(name) as select
(select name from worker);
5.由不可更新的视图导出的视图 如 create view worker_view2 as select * from
 worker_view1; 因为 worker_view1是不可更新的视图，所以worker_view2也是不
可以更新的视图。
6.创建视图时，algorithm为temptable类型 。
create algorithm=timetable view worker_view as select * from worker;
因为该视图的algorithm为temptable类型，所以worker_view为不可以更新的视图
。temptable类型就是临时表类型。系统默认临时表是不能更新的。
7.视图对应的表上存在没有默认值的列，而且该列没有包含在视图里。例如，表中
包含的name字段没有默认值，但是视图中不包括该字段。那么这个视图是不能更新的。
因为，在更新视图时，这个没有默认值的记录将没有值插入，也没有null值插入。数据
库系统是不会允许这样的情况出现的，其会阻止这个视图更新。

8. with [ cascaded | local ] check option也决定视图能否更新。
local参数表示更新视图时要满足该视图本身的定义的条件即可;
cascaded 参数表示更新视图时要满足所有相关视图和表的条件。没有指明，默认为
cascaded.

|××| 视图中虽然可以更新数据，但是有很多的限制。一般情况下，最好将视图作为查询
数据的虚拟表，而不要通过视图更新数据。因为，使用视图更新数据时，如果没有全
面考虑在视图中更新数据的限制，可能会造成数据更新失败。




[ 删除视图 ]
drop view [if exists] 视图名列表 [restrict | cascade]
*删除视图必须是具有删除视图权限的用户才能删除
select Drop_priv from mysql.user where user='用户名' (查看drop权限)

