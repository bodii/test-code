存储过程和存储函数

MySQL的存储过程（stored procedure）和函数（stored function）统称为stored routines。

1. MySQL存储过程和函数的区别

函数只能通过return语句返回单个值或者表对象。而存储过程不允许执行return，但是通过out参数返回多个值。 函数是可以嵌入在sql中使用的，可以在select中调用，而存储过程不行。

函数限制比较多，比如不能用临时表，只能用表变量．还有一些函数都不可用等等．而存储过程的限制相对就比较少

一般来说，存储过程实现的功能要复杂一点，而函数的实现的功能针对性比较强。

当存储过程和函数被执行的时候，SQL Manager会到procedure cache中去取相应的查询语句，如果在procedure cache里没有相应的查询语句，SQL Manager就会对存储过程和函数进行编译。

Procedure cache中保存的是执行计划 (execution plan) ，当编译好之后就执行procedure cache中的execution plan，之后SQL SERVER会根据每个execution plan的实际情况来考虑是否要在cache中保存这个plan，评判的标准一个是这个execution plan可能被使用的频率；其次是生成这个plan的代价，也就是编译的耗时。保存在cache中的plan在下次执行时就不用再编译了。

存储过程应用实例

1. 创建没有参数的存储过程

创建存储过程取得最高考分

create PROCEDURE getMaxMark()

begin

select max(mark) 最高分 from `TScore`;

end

调用存储过程

call getMaxMark();

2. 创建带输入参数的存储过程

以下存储过程能够输入的学生号查出该学生的信息 参数Sid代表学号，IN代表输入参数。

CREATE PROCEDURE getStudentByID(IN sid varchar(15))

BEGIN

select * from `TStudent` where studentID=sid;

END

调用该参数

call getStudentByID('00009')

3. 创建带输入和输出参数的存储过程

创建存储过程，能够输出指定课程的最高分

create PROCEDURE getMaxMarkBySubject(IN subName varchar(30),OUT maxMark int)

begin

select MAX(mark) into maxMark from `TScore` a join `TSubject` b on a.`subJectID`=b.`subJectID`

where b.`subJectName`=subName;

end

调用存储过程，将取出的最大值放到变量@maxScore

CALL getMaxMarkBySubject('计算机网络',@maxScore);

select @maxScore 计算机网络最高分;

CALL getMaxMarkBySubject('数据结构',@maxScore);

select @maxScore 数据结构最高分;

4. 思考：创建存储过程，能够根据指定学号删除学生记录

5. 删除存储过程

drop PROCEDURE `getMaxMarkBySubject`

6. 查看创建的存储过程

创建存储函数

7. 根据学生成绩判断是否优秀

以下函数能够根据输入值范围输出成绩是否优良差。

create FUNCTION getGrad1(score int)

RETURNS varchar(50)

BEGIN

return IF(score>90,'成绩优秀',IF(score<90 and score>80,'成绩良好',IF(score<80 and score>70,'成绩一般',IF(score<70 and score>60,'成绩及格','不及格'))));

END

在查询中使用定义的函数

select b.sname 姓名,mark 分数,getGrad1(mark) 成绩级别 from `TScore` a join `TStudent` b on a.`StudentID`=b.`StudentID`

8. 取汉字拼音首字母的函数

先创建获取汉字拼音函数需要用到的表

DROP TABLE IF EXISTS `pinyin`;

CREATE TABLE `pinyin` (

`letter` char(1) NOT NULL,

`chinese` char(1) NOT NULL,

PRIMARY KEY (`letter`)

) ENGINE=MyISAM DEFAULT CHARSET=gbk;

INSERT INTO `pinyin` VALUES ('A','驁'),('B','簿'),('C','錯'),('D','鵽'),('E','樲'),('F','鰒'),('G','腂'),

('H','夻'),('J','攈'),('K','穒'),('L','鱳'),('M','旀'),('N','桛'),('O','漚'),('P','曝'),('Q','囕'),('R','鶸'),

('S','蜶'),('T','籜'),('W','鶩'),('X','鑂'),('Y','韻'),('Z','咗');

创建获取汉字拼音的函数
delimiter && (将mysql 原有的结束符;改为&&)

DROP FUNCTION IF EXISTS `PINYIN`

CREATE FUNCTION PINYIN(str CHAR(255))

RETURNS char(255)

BEGIN

DECLARE hexCode char(4);

DECLARE pinyin varchar(255);

DECLARE firstChar char(1);

DECLARE aChar char(1);

DECLARE pos int;

DECLARE strLength int;

SET pinyin = '';

SET strLength = CHAR_LENGTH(LTRIM(RTRIM(str)));

SET pos = 1;

SET @str = (CONVERT(str USING gbk));

WHILE pos <= strLength DO

SET @aChar = SUBSTRING(@str,pos,1);

SET hexCode = HEX(@aChar);

IF hexCode >= "8140" AND hexCode <= "FEA0" THEN

SELECT letter into firstChar

FROM pinyin

WHERE chinese >= @aChar

LIMIT 1;

ELSE

SET firstChar = @aChar;

END IF;

SET pinyin = CONCAT(pinyin,firstChar);

SET pos = pos + 1;

END WHILE;

RETURN UPPER(pinyin);

END

使用函数获取用户的姓名拼音首写字母

select sname 姓名,pinyin(sname) 拼音首字母 from `TStudent`

9. 思考：更改用户邮箱，将用户的邮箱地址设置姓名的拼音缩写

10. 数字转汉字

以下函数能够把阿拉伯数字转化成财务中用到的汉字

create FUNCTION tohanzi (n_LowerMoney DECIMAL)

RETURNS VARCHAR(120)

BEGIN

Declare v_LowerStr VARCHAR(200) ;

Declare v_UpperPart VARCHAR(200) ;

Declare v_UpperStr VARCHAR(200) ;

Declare i_I int ;

set v_LowerStr = LTRIM(RTRIM(ROUND(n_LowerMoney,2 ) ) ) ;

set i_I = 1 ;

set v_UpperStr = '' ;

while ( i_I <=char_length(v_LowerStr ) ) do

set v_UpperPart = CONCAT( case substring(v_LowerStr,char_length(v_LowerStr) - i_I + 1,1 )

WHEN '.' THEN '元'

WHEN '0' THEN '零'

WHEN '1' THEN '壹'

WHEN '2' THEN '贰'

WHEN '3' THEN '叁'

WHEN '4' THEN '肆'

WHEN '5' THEN '伍'

WHEN '6' THEN '陆'

WHEN '7' THEN '柒'

WHEN '8' THEN '捌'

WHEN '9' THEN '玖'

END,

case i_I

WHEN 1 THEN '分'

WHEN 2 THEN '角'

WHEN 3 THEN ''

WHEN 4 THEN ''

WHEN 5 THEN '拾'

WHEN 6 THEN '佰'

WHEN 7 THEN '仟'

WHEN 8 THEN '万'

WHEN 9 THEN '拾'

WHEN 10 THEN '佰'

WHEN 11 THEN '仟'

WHEN 12 THEN '亿'

WHEN 13 THEN '拾'

WHEN 14 THEN '佰'

WHEN 15 THEN '仟'

WHEN 16 THEN '万'

ELSE ''

END );

set v_UpperStr =CONCAT( v_UpperPart , v_UpperStr) ;

set i_I = i_I + 1 ;

end while;

set v_UpperStr = REPLACE(v_UpperStr,'零拾','零') ;

set v_UpperStr = REPLACE(v_UpperStr,'零佰','零') ;

set v_UpperStr = REPLACE(v_UpperStr,'零仟','零') ;

set v_UpperStr = REPLACE(v_UpperStr,'零零零','零') ;

set v_UpperStr = REPLACE(v_UpperStr,'零零','零') ;

set v_UpperStr = REPLACE(v_UpperStr,'零角零分','整') ;

set v_UpperStr = REPLACE(v_UpperStr,'零分','整') ;

set v_UpperStr = REPLACE(v_UpperStr,'零角','零') ;

set v_UpperStr = REPLACE(v_UpperStr,'零亿零万零元','亿元') ;

set v_UpperStr = REPLACE(v_UpperStr,'亿零万零元','亿元') ;

set v_UpperStr = REPLACE(v_UpperStr,'零亿零万','亿') ;

set v_UpperStr = REPLACE(v_UpperStr,'零万零元','万元') ;

set v_UpperStr = REPLACE(v_UpperStr,'万零元','万元') ;

set v_UpperStr = REPLACE(v_UpperStr,'零亿','亿') ;

set v_UpperStr = REPLACE(v_UpperStr,'零万','万') ;

set v_UpperStr = REPLACE(v_UpperStr,'零元','元') ;

set v_UpperStr = REPLACE(v_UpperStr,'零零','零') ;

if ( '元' = substring(v_UpperStr,1,1)) then

set v_UpperStr = substring(v_UpperStr,2,(char_length(v_UpperStr) - 1));

end if;

if ( '零' = substring(v_UpperStr,1,1)) then

set v_UpperStr = substring(v_UpperStr,2,(char_length(v_UpperStr) - 1)) ;

end if;

if ( '角' = substring(v_UpperStr,1,1)) then

set v_UpperStr = substring(v_UpperStr,2,(char_length(v_UpperStr) - 1)) ;

end if;

if ( '分' = substring(v_UpperStr,1,1)) then

set v_UpperStr = substring(v_UpperStr,2,(char_length(v_UpperStr) - 1)) ;

end if;

if ('整' = substring(v_UpperStr,1,1)) then

set v_UpperStr = '零元整' ;

end if;

return v_UpperStr;

END

select tohanzi(20321)

11. 随机产生姓名的函数

该函数，使用三个字符串，存放用户的姓名，使用随机函数从姓名中随机排列组合成人名。

create function create_name()

RETURNS varchar(3)

begin

DECLARE LN VARCHAR(300);

DECLARE MN VARCHAR(200);

DECLARE FN VARCHAR(200);

DECLARE LN_N INT;

DECLARE MN_N INT;

DECLARE FN_N INT;

SET LN='李王张刘陈杨黄赵周吴徐孙朱马胡郭林何高梁郑罗宋谢唐韩曹许邓萧冯曾程蔡彭潘袁于董余苏叶吕魏蒋田杜丁沈姜范江傅钟卢汪戴崔任陆廖姚方金邱夏谭韦贾邹石熊孟秦阎薛侯雷白龙段郝孔邵史毛常万顾赖武康贺严尹钱施牛洪龚';

SET MN='伟刚勇春菊毅俊峰强军平保东文辉力明永健世广志瑗琰韵融园艺咏卿聪澜纯毓悦昭冰爽琬茗羽希宁欣飘育滢馥新利筠柔竹霭凝晓欢霄枫芸菲寒伊亚宜可姬舒义兴良海山仁波宁贵福生龙元全国胜学祥亮政谦亨奇固之岚苑富顺信子杰涛昌成康星光天达安岩中茂进林有坚和彪博诚先敬震振壮会思群豪清飞彬娜静淑惠珠翠雅芝妍茜秋珊莎锦黛青倩婷姣婉娴瑾颖露瑶怡婵雁蓓纨仪荷丹蓉眉君琴蕊薇菁梦素伟刚勇毅俊峰强军平保东文辉力明永健世广志义兴良海山仁波宁贵福生龙元全国胜学祥才发武新利清飞彬富顺信子杰涛昌成康星光天达安岩中茂进林有坚和彪博诚先敬震振壮会思群豪心邦承乐绍功松善厚庆磊民友裕河哲江超浩亮政谦亨奇固之轮翰朗伯宏言若鸣朋斌梁栋维启克伦翔旭鹏泽晨辰士以建家致树炎德行时泰盛雄琛钧冠策腾楠榕风航弘';

SET FN='伟刚勇毅俊云莲真环雪荣爱妹霞香月莺媛艳瑞凡佳嘉琼勤珍贞莉桂娣叶璧才发武丽琳轮翰朗伯宏言若鸣朋斌梁栋维启克伦翔旭鹏泽晨辰士以建家致树炎德河哲江超浩璐娅琦晶裕华慧巧美婕馨影荔枝思心邦承乐绍功松善厚庆磊民友玉萍红娥玲芬芳燕彩兰凤洁梅秀娟英行时泰盛雄琛钧冠策腾楠榕风航弘峰强军平保东文辉力明永健世广志义兴良海山仁波宁贵福生龙元全国胜学祥才发武新利清飞彬富顺信子杰涛昌成康星光天达安岩中茂进林有坚和彪博诚先敬震振壮会思群豪心邦承乐绍功松善厚庆磊民友裕河哲江超浩亮政谦亨奇固之轮翰朗伯宏言若鸣朋斌梁栋维启克伦翔旭鹏泽晨辰士以建家致树炎德行时泰盛雄琛钧冠策腾楠榕风航弘';

SET LN_N=CHAR_LENGTH(LN);

SET MN_N=CHAR_LENGTH(MN);

SET FN_N=CHAR_LENGTH(FN);

return Concat(substring(LN,ceil(rand()*LN_N),1),substring(MN,ceil(rand()*MN_N),1),substring(FN,ceil(rand()*FN_N),1));

end

调用函数产生姓名

select create_name() 随机姓名 union

select create_name() 随机姓名 union

select create_name() 随机姓名 union

select create_name() 随机姓名 union

select create_name() 随机姓名 union

select create_name() 随机姓名 union

select create_name() 随机姓名 union

select create_name() 随机姓名 union

select create_name() 随机姓名 union

select create_name() 随机姓名 union

select create_name() 随机姓名 union

select create_name() 随机姓名 union

select create_name() 随机姓名

查看结果

12. 查看创建的函数

13. 删除存储的函数

drop FUNCTION `tohanzi`

在存储过程和函数中使用变量、判断和循环

MySQL中变量、判断和循环只能在存储过程和存储函数中使用。

14. 在存储过程中使用循环、变量

写一个存储过程，能够给TStudent表插入指定数量的学生记录。身份证号随机产生，姓名随机产生，性别随机，班级随机产生。这其中用到了随机函数，以及上面创建的产生姓名的函数。生日有随机函数产生，范围在1980-1989年，用户的邮箱由用户的姓名首写字母组合而成。

如果已有存储过程，必须先删除

drop procedure addStudent

创建的存储过程

create procedure addStudent(in num int)

begin

declare i int;

set i=1;

delete from TStudent;

while num>=i do

insert TStudent values (

LPAD(convert(i,char(5)),5,'0'),

create_name(),

if(ceil(rand()*10)%2=0,'男','女'),

RPAD(convert(ceil(rand()*1000000000000000000),char(18)),18,'0'),

Concat('198',convert(ceil(rand()*10),char(1)),'-',LPAD(convert(ceil(rand()*12),char(2)),2,'0'),'-',LPAD(convert(ceil(rand()*30),char(2)),2,'0')),

Concat(PINYIN(sname),'@hotmail.com'),

case ceil(rand()*3) when 1 then '网络与网站开发' when 2 then 'JAVA' ELSE 'NET' END,

NOW());

set i=i+1;

end while;

select * from TStudent;

end

调用存储过程

call addStudent(100)

15. 创建使用while的存储过程插入学生成绩

插入了100名学生后，执行以下命令。

创建存储过程，能够插入为学生插入分数。存储过程使用两个循环，分数在50-100分之间，使用随机数实现。

drop procedure fillSore

创建存储过程

create procedure fillSore()

begin

DECLARE St_Num INT;

DECLARE Sb_Num INT;

DECLARE i1 INT default 1;

DECLARE i2 INT default 1;

delete from TScore;

select count(*) into St_Num from TStudent;

select count(*) into Sb_Num from TSubject;

while St_Num>=i1 do

set i2=1;

while Sb_Num>=i2 do

insert TScore values

(LPAD(convert(i1,char(5)),5,'0'),LPAD(convert(i2,char(4)),4,'0'),ceil(50+rand()*50));

set i2=i2+1;

END WHILE;

set i1=i1+1;

END WHILE;

End

调用存储过程

call fillSore()

查询

select * from TScore

16. 创建使用if的函数

If函数支持多层嵌套

create FUNCTION getGrad2(score int)

RETURNS varchar(50)

BEGIN

declare grad varchar(50);

if score>90 then

set grad='成绩优秀';

else if score>80 then

set grad='成绩良好';

else if score>70 then

set grad='成绩一般';

else set grad='刚刚及格';

end if;

end if;

end if;

return grad;

END

使用函数查询数据库

select b.sname 姓名,mark 分数,getGrad2(mark) 成绩级别

from `TScore` a join `TStudent` b on a.`StudentID`=b.`StudentID` limit 5

17. 创建使用case的函数

该函数根据学生的分数，给出评价。

create FUNCTION getGrad3(score int)

RETURNS varchar(50)

BEGIN

declare grad varchar(50);

declare mark int;

set mark=ceil(score/10);

case mark

when 9 then set grad='成绩优秀';

when 8 then set grad='成绩良好';

when 7 then set grad='成绩一般';

else set grad='刚刚及格';

end case;

return grad;

END

测试函数

select b.sname 姓名,mark 分数,getGrad3(mark) 成绩级别

from `TScore` a join `TStudent` b on a.`StudentID`=b.`StudentID` limit 5

查看存储过程和存储函数的语句

运行一下命令可以查看创建fillSore存储过程语句

18. 使用show create查看存储过程内容

show create procedure fillSore

19. 使用管理工具产生查看创建存储过程的语句

作业：

20. 统计男生和女生人数

21. 统计各班人数

22. 把重姓的学生找出来

23. 找出身份证号末尾是偶数的学生。

24. 查询出生年月在1985-01-00到1988-01-00之间的学生。

25. 统计各班“计算机网络”平均分

26. 找出“计算机网络”不及格的男同学

来源： 51cto   作者：韩立刚  
