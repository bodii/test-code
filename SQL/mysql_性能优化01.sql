性能优化是通过某些有效的方法提高mysql数据库的性能。
性能优化的目的是为了使用mysql数据库运行速度更快、占用的磁盘空间更小
性能优化包括很多方面，例如优化查询速度、优化更新速度和优化mysql服务器等;


可以使用show status 语句查询mysql数据库的性能。
语法：
show status like 'value';
其中，value参数是常用的几个统计参数，如下：
Connections: 连接mysql服务器的次数;
Uptiem: mysql服务器的上线时间;
Slow_queries:慢查询的次数;
Com_slect: 查询操作的次数;
Com_insert: 插入操作的次数;
Com_update: 更新操作的次数;
Com_delete: 删除操作的次数;



2.分析查询语句
语法：
explain select 语句;
[例] explain select * from student\G

describe select 语句;(describe 可以缩写成desc)
查询结果和explain语句是一样的。



索引对查询速度的影响
如果查询时不使用索引，查询语句将查询表中的所有字段。这样查询的速度会很慢。
如果使用索引进行查询，查询语句只查询索引字段。这样可以减少查询的记录数，达
到提高查询速度的目的。


3.1.将字段很多的表分解成多个表
有些表在设计时设置了很多的字段。这个表中有些字段的使用频率很低。当这个表的
数据量很大时，查询数据的速度就会很慢。
对于这种字段特别多且有些字段的使用频率很低的表，可以将其分解成多个表。
可通过id关联


3.2. 增加中间表
有时需要经常查询某两个表中的几个字段。如果经常进行联表查询，会降低mysql数据库
的查询速度。对于这种情况，可以建立中间表来提高查询速度。
如学生表 经常要查学生的学号、姓名和成绩。根据这种情况可以创建一个temp_score表
字段分别是id、name、grade。
create table temp_socre(
	id int not null,
	name varchar(20) not null,
	grade float
	);
然后从student表和score表中将记录导入到temp_socre表中。
insert into temp_score select student.id,student.name,score.grade
	from student,score where student.id=socre.stu_id;


3.3 增加冗余字段
设计数据库表时尽量让表达到三范式。但是，有时为了提高查询速度，可以有意识
地在表中增加冗余字段。
表的规范化程度越高，表与表之间的关系就越多;查询时可能经常需要在多个表之间
进行连接查询;而进行连接操作会降低查询速度。

如，学生表student存学生信息，department表存院系信息。通过student表中的dept_id
字段与department表建立关联。如果要查询一个学生所在系的名称，必须从student表
中查找学生所在院系的编号（dept_id),然后根据这个编号去department查找系的名称。
如果经常需要进行这个操作时，连接查询会浪费很多的时间。因此可以在student表中增
加一个冗余字段dept_name,该字段用来存储学生所在院系的名称。这样就不用每次进行
连接操作了。

** 分解表、增加中间表和增加冗余字段都浪费了一定的磁盘空间。从数据库性能来看，
增加少量的冗余来提高数据库的查询速度是可以接受的。是否通过增加冗余来提高数据
库性能，这要根据mysql服务器的具体要求来定。


3.4 优化插入记录的速度
插入记录时，索引、唯一性校验都会影响到插入记录的速度。而且，一次插入多条记录和
多次插入记录所耗费的时间是不一样的。根据这些情况，分别进行不同的优化。
	1. 禁用索引
	插入记录时，mysql会根据表的索引对插入的记录进行排序。如果插入大量数据时，这
	些排序会降低插入记录的速度。为了解决这种情况，在插入记录之前先禁用索引。等到
	记录都插入完毕后再开启索引。
	禁用索引的语句：
	alter table 表名 disable keys;
	重新开启索引的语句：
	alter table 表名 enable keys;
	
	对于新创建的表，可以先不创建索引。等到记录都导入以后再创建索引。这样可以提高
	导入数据的速度。

	
	2. 禁用唯一性检查
	禁用唯一性检查的语句：
	set unique_checks=0;
	重新开启唯一性检查的语句：
	set unique_checks=1;

	
	3. 优化 insert 语句
	当插入数据时，insert插入多表比多次插入一条要快。

	** 当插入大量数据时，建议使用一个insert插入多条记录的方式。而且，如
	果能用load data infile语句，就尽量用load data infile 语句。因为load
	data infile 语句导入数据的速度比insert语句的速度快。



5. 分析表、检查表和优化表
分析表主要作用是分析关键字的分布。检查表主要作用是检查表是否存在错误。优化表主
要是消除删除或者更新造成的空间浪费。
	
	1.分析表
	mysql中使用analyze table 语句来分析表：
	analhyze table 表名1 [,表名2...];
	使用analyze table 分析表的过程中，数据库系统会对表加一个只读锁。在分析期间，
	只能读取表中的记录，不能更新和插入记录。analyze table 语句能够分析innodb和
	myisam类型的表。
	如： analyze table test.t1;
	+---------+---------+----------+----------+
	| Table   | Op      | Msg_type | Msg_text |
	+---------+---------+----------+----------+
	| test.t1 | analyze | status   | OK       |
	+---------+---------+----------+----------+
	
	Op:表示执行的操作。analyze表示进行分析操作。check表示进行检查查找。optimize
	表示进行优化操作;
	Msg_type: 表示信息类型，其显示的值通常是状态、警告、错误和信息这四者之一;
	Msg_text: 显示信息。
	

	2.检查表
	mysql中使用check table语句来检查表。check table 语句能够检查innodb和myisam类型
	的表是否存在错误。而且，该语句还可以检查视图是否存在错误。
	语法：
	check table 表名1 [,表名2...] [option];
	其中，option 参数有5个参数，分别是quick,fast,changed,medium 和extended这5个参数
	的执行效率依次降低。option选项只对myisam类型的表有效，对InnoDB类型的表无效。check
	table 语句在执行过程中也会给表加上只读锁。


	3.优化表
	mysql 中使用optimize table 语句来优化表。该语句对innodb和myisam类型的表都有效。但
	是，optilmize table 语句只能优化表中的varchar、blob 和 text类型的字段。
	语法：
	optimize table 表名1[,表名2...];
	通过optimize table 语句可以消除删除和更新造成的磁盘碎片，从而减少空间的浪费。optimize
	table 语句在执行过程中也会给表加上只读锁。

	**如果一个表使用了text或者blob这样的数据类型，那么更新、删除等操作就会造成磁盘空间的
	浪费。因为，更新和删除操作后，以前分配的磁盘空间不会自动收回。使用optimize table 语句
	就可以将这些磁盘碎片整理出来，以便以后再利用。




