[ 带exists 关键字的子查询 ]
exists 关键字表示存在。使用exists关键字时，内层查询语句不返回查询的记录。
而是返回一个真假值。内层查询满足条件，返回true, 否则返回false;

[例] select * from test02 where exists 
			( select id from test01 where id=1 );
如果 内查询的id值存在，则查询test02表

not exists 与 exists 相反。




[ 带 any 关键字的子查询 ]
any 关键字表示满足其中任一条件。使用 any 关键字时，只要满足内层查询语句返回
的结果中的任何一个，就可以通过该条件来执行外层查询语句。
any 关键字通常与比较运算符一起使用。例如，>ANY 表示大于任何一个值，=ANY 表示
等于任何一个值。

[例] select * from test02 where d_id >= any (select d_id from test01);

**  带 any 的子查询，会将子查询的结果与外层条件值与内层的最小值比较，满足条件则显示。 



[ 带 all 关键字的子查询 ]
all 关键字表示满足所有条件。使用 all 关键字时，只有满足内层查询语句返回的所有
结果，才可以执行外层查询语句。 all 关键字也经常与比较运算符一起使用。例如，>all
表示大于所有值， <all 表示小于所有值。

[例] select * from test02 where d_id > all (select d_id from test01);

**  带 all 的子查询，会将子查询的结果与外层条件值与内层的最大值比较，满足条件则显示。



 
