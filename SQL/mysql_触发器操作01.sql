[ 什么是触发器 ]
触发器是由insert、update和delete等事件来触发某种特定操作。满足触发器的
触发条件时，数据库系统就会执行触发器中定义的程序语句。这样做可以保证某
些操作之间的一致性。例如，当学生表中增加了一个学生的信息时，学生的总数就
必须同时改变。可以在这里创建一个触发器，每次增加一个学生的记录，就执行一
次计算学生总数的操作。这样就可以保证每次增加学生的记录后，学生总数是与记
录的数是一致的。触发器触发的执行语句可能只有一个，也可能有多个。



[ 触发器的创建语句 ] 
create trigger 触发器名  before | after 触发事件 on 表名 
for each row 执行语句

触发器名：触发器名参数指要创建的触发器的名字；
before 和 after 参数 指字的解发器执行的时间, "before"指在触发事件之前执
行触发语句，"afeter"表示在触发事件之后执行触发语句；
"触发事件"参数指触发的条件，其中包括 insert\update\delete;
"表名"参数指触发事件操作的表的名称;
for each row 表示任何一条记录上的操作满足解发事件都会触发该触发器;
"执行语句"参数指触发器被触发后执行的程序。（执行语句涉及的表必须已经存在）


例如
create trigger dept_trig1 before insert on department for each row
insert into trigger_time values(new());


[ 创建多个执行语句的触发器 ]
create trigger 触发器名 before | after 触发事件
on 表名 for each row 
begin
执行语句列表
end 

begin 与 end 直接的"执行语句列表" 参数表示需要执行的多个执行语句的内容
不同的执行语句之间用分号隔开。
=注=：一般情况下，mysql默认是以";"作为结束执行语句。在创建触发器过程中需
要用到“;”。为了解决这个问题，可以用delimiter(定界符)语句。如"delimiter&&",
可以将结束符号变成"&&"。当触发器创建 完成后，可以用命令"delimiter;" 来将
结束符号变成";"

例如步骤如下：
1.delimiter &&
2.create trigger dept_trig2 after delete
     on department for each row
     begin
     insert into trigger_time values('21:01:01');
     insert into trigger_time values('22:01:01');
     end
     &&
3.delimiter ;



[ 查看触发器 ]
查看触发器是指查看数据库中已存在的触发器的定义\状态和语法等信息。
查看触发器的方法包括sho triggers 语句和查询information_schema数据
库下的triggers表
show triggers\G
或
select * from information_schema.triggers\G
select * from information_schema.triggers where trigger_name=触发器名\G

[ 删除触发器 ]

drop trigger 触发器名;
或
drop trigger 数据库名.触发器名;


