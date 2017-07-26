[ 调用存储过程 ]
mysql中使用call语句来调用存储过程。调用存储过程后，数据库系统将执行
存储过程中的语句。然后，将结果返回给输出值。
call sp_name([parameter[,...]);

其中，sp_name是存储过程的名称;
parameter是存储过程的参数。




[ show status 查看存储过程和函数的状态 ]
语法：
show {procedure | function} status [ like 'pattern' ];

其中，procedure参数表示查询存储过程;
function 参数表示查询存储函数;
like 'pattern'参数用来匹配存储过程或函数的名称。

[例] show procedure status like 'num_from_employee'\G




[ show create 语句查看存储过程和函数的定义 ]
语法：
show create {procedure | function} sp_name;


** show status 语句只能查看存储过程或函数是操作哪一个数据库、
存储过程或函数的名称、类型、谁定义的、创建和修改时间、字符编码等
信息。但是，这个语句不能查询存储过程或函数的具体定义。如果需要查
看详细定义，需要使用show create语句。




[ infornmation_schema.Routines 表中查看存储过程和函数的信息 ]
语法：
select * from information_shema.Routines where routine_name='sp_name';

** information_schema数据库下的Routiones表中，存储着所有存储过程和函数
的定义。如果使用select语句查询Routines表中的存储过程和函数的定义时，一定
要用 routine_name 字段指定存储过程和函数的名称。否则，将查询出所有的存储过
程或函数的定义。




[ 修改存储过程和函数 ]
语法：
alter {procedure | function} sp_name [characteristic ...]
characteristic:
	{contains sql | no sql | read sql data | modifies sql data}
	| sql security { definer | invoker}
	| comment 'string'

其中，sp_name 表示存储过程或函数的名称;
characteristic 表示指定存储函数的特性。
contains sql 表示子程序包含sql语句，但不包含读或写数据的语句;
no sql 表示子程序中不包含sql语句;
reads sql data 表示子程序中包含读数据的语句;
modifies sql data 表示子程序中包含写数据的语句。
sql security { definer | invoker } 指明谁有权限来执行。
definer 表示只有定义者自己才能够执行;
invoker 表示调用者可以执行;
comment 'string' 是注释信息。




[ 删除存储过程和函数 ]
语法：
drop {procedure | function} sp_name;
其中，sp_name 表示存储过程和函数的名称




[ 常见问题及解答 ]

1.一个存储过程中可以调用其他的存储过程吗？
存储过程是用户定义的sql语句的集合。用户通过call语句调用已经
定义好的存储过程来执行其中的sql语句。同时，存储过程中可以通
过call语句来调用其他的存储过程。

2.存储过程和存储函数的区别是什么？
存储过程的参数有3类，分别是in、out 和 inout。通过out、inout
将存储过程的执行结果输出。而且存储过程中可以有多个out、inout
类型的变量，可以输出多个值。
存储函数中的参数都是输入参数。函数中的运算结果通过 return 语句
来返回。return 语句只能返回一个结果。

3.存储函数和mysql内部函数有什么区别
存储函数是用户自己定义的函数。并且通过调用来执行函数中的sql语句。
函数执行完成后，通过return语句来返回执行结果。从原理上讲，存储函数
和mysql内部函数是一样的。只是内部函数比较常用，因此，数据库的设计者
将些函数集成到了数据库中。
