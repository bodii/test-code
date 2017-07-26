create [ UNIQUE | FULLTEXT | SPATIAL ]  INDEX 索引名
ON 表名 （属性名   [( 长度 )]   [ ASC | DESC ])

1.创建普通索引
CREATE INDEX index7_id on example0(id);

2.创建唯一性索引
CREATE UNIQUE INDEX index7_id on example0(id);

3.创建全文索引
CREATE FULLTEXT INDEX index7_info on example0(info);
其中，FULLTEXT用来设置索引为全文索引;表example0的存储引擎必须是MyISAM类型
info字段必须为（字符串）：CHAR、VARCHAR和TEXT等类型

4.创建单列索引
CREATE INDEX index7_add on index10(address(4));
查询时可以只查询address字段的前4个字符，而不需要全部查询。

5.创建多列索引
CREATE INDEX index7_na on index7(name,address);
该索引创建好了以后，查询条件中必须有name字段才能使用索引。

6.创建空间索引
CREATE SPATIAL INDEX index7_line on index7(line);
其中，SPATIAL用来设置索引为空间索引；表index7的存储引擎必须是MyISAM类型；line字段必须为空间数据类型，而且是非空。


【或】

建表时

1.创建普通索引
CREATE TABLE index1(
     id int,
     name varchar(20),
     sex char(4),
    INDEX (id)
);

2.创建唯一性索引
CREATE TABLE index1(
     id int,
     name varchar(20),
     sex char(4),
     UNIQUE  INDEX index1_id(id ASC)
);

3.创建全文索引
CREATE TABLE index1(
     id int,
     name varchar(20),
     sex char(4),
     info varchar(20),
     FULLTEXT  INDEX index1_info(info)
);

4.创建单列索引
CREATE TABLE index1(
     id int,
     name varchar(20),
     sex char(4),
     info varchar(20),
     INDEX index1_info(info(4))
);

5.创建多列索引
CREATE TABLE index1(
     id int,
     name varchar(20),
     sex char(4),
     info varchar(20),
     INDEX index1_ns(name,sex)
);

6.创建空间索引
CREATE TABLE index1(
     id int,
     name varchar(20),
     sex char(4),
     info varchar(20),
     space geometray not null,
     SPATIAL INDEX index1_sp(space)
);
空间类型包括：GEOMETRY、POINT、LINESTRING和POLYGON类型，这些空间数据类型平时很少用到。

【或】

ALTER TABLE 表名  ADD  [ UNIQUE | FULLTEXT | SPATIAL ] INDEX
索引名（属性名   [(长度)]  [ ASC | DESC ]   

其中，UNIQUE 是可选参数，表示索引为唯一性索引，且属性名(也可称为字段名) 必须具有唯一约束
FULLTEXT是可选参数，表示索引为全文索引
SPATIAL 也是可选参数，表示索引为空间索引
INDEX 参数用来指定字段为索引（和 KEY 的作用相同）
索引名 参数是给创建的索引取的新名称（一般为"表名_字段[或字段简称]"的形式）
表名 参数是指需要创建索引的表的名称，该表必须是已经 存在的， 如果不存在，需要先创建
属性名 参数指定索引对应的字段名称，该字段必须为前面定义好的字段
长度 是可选参数，其指索引的长度，（一般索引字符串（有且必须是字符串类型）的长度可能很长，索引时只取前几位，加外索引时的查询速度）
ASC 和 DESC 都是可选参数，ASC参数表示升序排序， DESC参数表示降序排列。


 【 删除索引 】

drop index 索引名 on 表名;
普通索引、唯一索引、全文索引、单列索引、多列索引、空间索引都可以使用这样删除



