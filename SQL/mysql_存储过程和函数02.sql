[ 流程控制的使用 ]

1.if语句
if search_condition then statement_list
[elseif search_condition then statement_list]...
[else statement_list]
end if

其中：
search_condition 表示条件判断语句
statement_list  表示不同条件的执行语句

[例] 
if age>20 then set @count1=@count1+1;
	elseif age=20 then @count2=@count2+1;
	else @count3=@count3+1;
end if;




2. case 语句
形式：
case case_value
	when when_value then statement_list
	[when when_value then statement_list]...
	[else statement_list]
end case
另外一种形式：
case 
	when search_condition then statement_list
	[when search_condition then statement_list]...
	[else statement_list]
end case

[例] 
case age
	when 20 then set @count1=@count1+1;
	else set @count2=@count2+1;
end case;
或
case 
	when age=20 then set @count1=@count1+1;
	else set @count2=@count2+1;
end case;




[ loop 语句 ]
loop语句可以使某些特定的语句重复执行，实现一个简单的循环。但是
loop语句本身没有停止循环的语句，必须是遇到leave语句等才能停止循
环。
语法：
[begin_label:] loop
	statement_list
end loop [end_label]

[例] 
add_num:loop
	set @count1=@count1+1;
end loop add_num;
××该示例循环执行count1+1的操作。因为没有跳出循环的语句，这个循环成了
一个死循环。loop循环都以end loop结束。




4. leave 语句
leave 语句用于跳出循环控制。
leave label

[例] 
add_num: loop
	set @count1=@count1+1;
	if @count1=100 then
		leave add_num;
end loop add_num;




5. iterate 语句
iterate 语句也是用来跳出循环的语句。但是，iterate语句是跳出本次循环，
然后直接进入下一次循环。
语法： iterate label

[例] 
add_num: loop
	set @count1=@count1+1;
	if @count1=100 then
		leave add_num;
	elseif mod(@count1,3)=0 then
		iterate add_num;
	select * from employee;
end loop add_num;




6. repeat 语句
当满足特定条件时，就会跳出循环语句。
语法：
[begin_label:] repeat
	statement_list
	until search_condition
end repeat [end_label]

[例]
repeat
	set @count1=@count1=1;
	until @count1=100
end repeat;
该示例循环执行count1 +1的操作，count1值为100时结束循环。




7. while 语句
语法：
[begin_label:] while search_condition do
	statement_list
end while [end_label]

search_condition 表示循环执行的条件，满足该条件时循环执行
statement_list 表示循环的执行语句

[例] 
while @count1<100 do
	set @count1=@count1+1;
end while;
该循环执行count1+1的操作，count1的值小于100时执行循环。如果
count1的值等于100了，则跳出循环。while循环需要使用end while
来结束。

