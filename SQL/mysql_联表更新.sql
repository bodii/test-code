d = 表2的id 时，修改表1的值
update 表名1 as A，表名2 as B  set A.字段名=值1  where A.id = B.id;

2.当表2的where条件满足时， （表2作为临时表） 将表2 select 的结果 作为 更新表1的条件
update 表名1 as A，(select * from 表名2 where 表名2=值) as B  set A.字段名=值1  where A.id = B.id;

3.将两个表合并,如果A表的字段与B表的字段满足条件，修改A表字段的值等于B表字段的值
UPDATE A INNER JOIN B ON A.friendid=B.userid  SET A.friendname=B.username;

4.将两个表合并,如果A表的字段与B表的字段满足表件，修改A表字段的值
UPDATE A INNER JOIN B ON A.friendid=B.userid  SET A.表字段=值  where A.friendname=B.username;

5.如果A表的字段与B表的字段满足条件，修改A表某些的字段值 = B表某些字段的值
update table1,table2 set table1.f1=table2.f1,table1.f2=table2.f2 where table1.ID=table2.ID


left join(左联接) 返回包括左表中的所有记录和右表中联结字段相等的记录
right join(右联接) 返回包括右表中的所有记录和左表中联结字段相等的记录
inner join(等值连接) 只返回两个表中联结字段相等的行
