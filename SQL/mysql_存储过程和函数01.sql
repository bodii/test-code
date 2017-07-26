[ 存储过程和函数 ]
存储过程和函数是在数据库中定义一些sql语句的集合，然后直接调用这些存储过程
和函数来执行已经定义好的sql语句。存储过程和函数可以避免开发人员重复的编写
相同sql语句。而且，存储过程和函数是在mysql服务器中存储和执行的，可以减少
客户端和服务器的数据传输。



存储过程的基本形式：
create procedure sp_name([proc_parameter[,...]])
		[characteristic ...] routine_body

sp_name 参数是存储过程的名称
proc_parameter 存储过程的参数列表
characteristic 存储过程的特性
routine_body sql代码的内容，可以用begin ... end 来标志sql代码的开始和结束。
proc_parameter 中每个参数由3部分组成。这3部分分别是输入输出类型、参数名称和
参数类型。
如： [ in | out | inout ] param_name type
in 表示输入参数; out 表示输出参数；inout 既可以是输入，也可以是输出;
param_name 参数是存储过程的参数名称;
type参数指定存储过程的参数类型，该类型可以是mysql数据库的任意数据类型。

characteristic 参数有多个取值。其取值说明如下：
language sql:说明routine_body 部分是由sql语言的语句组成，这也是数据库
系统默认的语言

[not] deterministic: 指明存储过程的执行结果是否是确定的。
deterministic 表示结果是确定的。每次执行存储过程时，相同的输入会得到相
同的输出。not deterministic 表示结果是非确定的，相同的输入可能得到不同
的输出。默认情况下，结果是非确定的。

{contains sql |no sql |reads sql data | modifies sql data}：指明子程序
使用sql语句的限制。contains sql表示子程序包含sql语句，但不包含读或写数据
的语句；no sql表示子程序中不包含sql语句;reads sql data表示子程序中包含读
数据的语句;modifies sql data表示子程序中包含写数据的语句。默认情况下，系统
会指定为contains sql。

sql security { defined | invoker}:指明谁有权限来执行。defined表示只有定义
者自己才能执行;invoker表示调用者可以执行。默认情况下，系统指定的权限是defined

comment 'string' ：注释信息。

××创建存储过程时，系统默认指定contains sql，表示存储过程中使用了sql语句。但是，
如果存储过程中没有使用sql语句，最好设置为no sql。而且，存储过程中最好在comment
部分对存储过程进行简单的注释，以便以后在阅读存储过程的代码时更加方便。



[例] create procedure num_form_employee(in emp_id int, out count_num int)
			reads sql data 
			begin
				select count(*) into count_num from employee where d_id=emp_id;
			end
上述代码中，存储过程名称为num_from_employee;输入变量为emp_id;输出变量为count_num.
select 语句从employee表查询d_id值等于emp_id的记录，并用count(*)计算d_id值相同的记录
的条数，最后将计算结果存入count_num中。




[ 创建存储函数 ]
基本形式如下：
create function sp_name([func_parameter[,...]])
		returns type
		[characteristic ...] routine_body

其中，sp_name 参数是存储函数的名称;
func_parameter 存储函数的参数列表;
returns type 指定返回值的类型;
characteristic 参数指定存储函数的特性，该参数的取值与存储过程中的取值是一样的;
routine_body 参数是sql代码的内容，可以用begin...end来标志sql代码的开始与结束。

func_parameter 可以由多个参数组成，其中每个参数由参数名称和参数类型组成

[例] create function name_from_employee(emp_id int)
			returns varchar(20)
			begin
				return (select name from employee where num=emp_id);
			end

存储函数的名称为 name_from_employee;该函数的参数为emp_id;返回值是varchar类型。
select语句从employee表查询num值等于emp_id的记录，并将该记录的name字段的值返回。




[ 变量的使用 ]
用户可以使用declare关键字来定义变量。然后可以为变量赋值。这些变量的作用范围是
begin...end程序段中。

定义变量 基本语法
declare var_name[,...] type [default value]
其中，declare关键字是用来声明变量的;var_name参数是变量的名称，这里可以同时定义
多个变量;type参数用来指定变量的类型;default value子句将变量默认值设置为value,
没有使用default子句时，默认值为null。
[例] declare my_sql int default 10;

为变量赋值 基本语法
set var_name=expr[,var_name=expr]...
[例] set my_sql=30;


mysql中还可以使用select ... into语句为变量赋值。
select col_name[,...] into var_name[,...]
	from table_name where condition
其中，col_name参数表示查询的字段名称;
var_name 参数是变量的名称;
table_name参数指表的名称;
condition 参数指查询条件。
[例] select d_id into my_sql from employee where id=2;




[ 定义条件和处理程序 ]
定义条件
declare condition_name condition for condition_value
condition_value:
	sqlstate[value] sqlstate_value | mysql_error_code

其中，condition_name 参数表示条件的名称;
condition_value参数表示条件的类型;
sqlstate_value参数和mysql_error_code参数都可以表示myslq
的错误。例如 error 1146(42s02)中，sqlstate_value值是42s02,
mysql_error_code值是1146。

[例] 下面定义'error 114(42s02)'这个错误，名称为can_not_find.
可以用两种不同的方法来定义
# 方法一： 使用sqlstate_value
declare can_not_find condition for sqlstate '42s02';
# 方法二：使用mysql_error_code
declare can_not_find condition for 1146;


2.定义处理程序
mysql 中可以使用declare关键字来定义处理程序。
declare handler_type handler for condition_value[,...] sp_statement
handler_type:
	continue | exit | undo
condition_value:
	sqlstate [value] sqlstate_value | condition_name | sqlwarning
		| not found | sqlexception | mysql_error_code


