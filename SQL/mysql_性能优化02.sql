4. 优化mysql服务器

优化mysql服务器可以从两个	方面来理解。一个是从硬件方面来进行优化;另一方面是从mysql
服务的参数进行优化。通过这些优化方式，可以提供mysql的运行速度。但是这部分的内容很
难理解，一般只有专业的数据库管理员才能进行这一类的优化。

	1.优化服务器硬件
	服务器的硬件性能直接决定着mysql数据库的性能。例如，增加内存和提高硬盘的读写速
	度，可以提高mysql数据库的查询、更新的速度。
	如果条件允许，可以将内存提高到4GB。并且选择my-innodb-heavy-4G.ini作为mysql
	数据库的配置文件。但是，这个配置文件主要支持innodb存储引擎的表。如果使用2GB
	内存，可以选择my-huge.ini作为配置文件。而且，mysql所在的计算机最好是专用数据
	库服务器。这样数据库可以完全利用该机器的资源。

	** 服务器类型分为Developer machine、Server Machine 和 Dedicate Mysql Server
	Machine。其中 Developer Machine用来做软件开发的时候使用，数据库占用的资源比
	较少。后面两者占用的资源比较多，尤其是Dedicate MySQL Server Machine,其几乎要占
	用所有的资源。
	
	另一种方式是提高mysql性能的方式是使用多块磁盘来存储数据。可以从多个磁盘并行读取
	数据，这样可以提高读取数据的速度。通过镜像机制可以将不同计算机上的mysql服务器进
	行同步,这些mysql服务器中的数据都是一样的。通过不同的mysql服务器来提供数据库服务
	，这样可以降低单个mysql服务器的压力，从而提高mysql的性能。


	2. 优化mysql的参数
	内存中会为mysql保留部分的缓存区。这些缓存区可以提高mysql数据库的处理速度。缓存区
	的大小都是在mysql的配置文件中进行设置的。

	mysql中比较重要的配置参数都在my.cnf或者my.ini文件的[mysqld]组中。
	几个很重要的参数详细：
	
	& key_beffer_size: 表示索引缓存的大小。这个值越大，使用索引进行查询的速度越快。
	& table_cache: 表示同时打开的表的个数。这个值越大，能够同时打开的表的个数越多。
		这个值不是越大越好，因为同时打开的表太多会影响操作系统的性能。
	& query_cache_size:表示查询缓存区的大小。使用查询缓存区可以提高查询的速度，这种
	方式只适用于修改操作少且经常执行相同的查询操作的情况; 其默认值为0。当取值为2时，
	只有select 语句中使用了sql_cache关键字，查询缓存区才会使用。例如，select sql_cache
	* from score。
	& query_cache_type:表示查询缓冲区的开启状态。其聚会为0时表示关闭，取值为1时表示开启，
	取值为2时表示按要求使用查询缓存区。
	& max_connections: 表示数据库的最大连接数。这个连接数不是越大越好，因为这些连接会浪费
	内存的资源。
	& sort_buffer_size:表示排序缓存区的大小。这个值越大，进行排序的速度赶快。
	& read_buffer_size:表示为每个线程保留的缓冲区的大小。当线程需要从表中连续读取记录时
	需要用到这个缓冲区。set session read_buffer_size=n 可以临时设置该参数的值。
	& read_rnd_buffer_size: 表示为每个线程保留的缓冲区的大小，与read_bueffer_size相似。
	但主要用于存储按特定顺序读取出来的记录。也可以用set session read_rnd_buffer_size=n
	业临时设置该参数的值。
	& innodb_buffer_pool_size:表示innodb类型的表和索引的最大缓存。这个值越大，查询的速度
	就会越快。但是这个值太大了会影响操作系统的性能。
	& innodb_flush_log_at_trx_commit:表示何时将缓冲区的数据写入日志文件，并且将日志文件
	写入日志文件并将日志文件写入磁盘;该参数有3个值，分别是0、1 和 2 。值为0时表示每隔1秒
	将数据写入日志文件并将日志文件写入磁盘；值为1时表示每次提交事务时将数据写入日志文件并
	将日志文件写入磁盘;值为2时，表示每次提交事务时将数据写入日志文件，每隔1秒将日志文件写
	入磁盘。该参数的默认值为1.这个默认值是最安全最合理的。
	
	合理的配置这些参数可以提高mysql服务器的性能。除上述参数外，还有	innodb_log_buffer_size
	、innodb_log_file_size等参数。配置完参数以后，需要重新启动mysql服务才会生效。



[例] 
1. 查看innodb表的查询次数和更新次数
	innodb_rows_read参数表示innodb表查询的记录数。
	innodb_rows_updated表示innodb表更新的记录数。
	语句：
	show status like 'innodb_rows_read'\G
	show status like 'innodb_rows_updated'\G

2.分析查询语句的性能
	语句：
	explain select * from score where stu_id=902\G
	或
	desc select * from score where stu_id=902\G

3.分析score表
	语句：
	analyze table score;
	Msg_test 为 OK 表是表的状态正常。

4. 查看mysql服务器的连接数、查询次数和慢查询的次数
Connection、Com_select 和 Slow_queries 3个参数分别表示mysql服务器的连接数、
查询次数和慢查询次数。
show status like 'Connections';

show status like 'Com_select';

show status like 'Slow_queries';


5.检查表
check table score;

6.优化表
optimize table score;
		
