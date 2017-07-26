CREATE TABLE pruview
(
   id            INTEGER,
   module_name   VARCHAR (50),
   father_id     INTEGER,
   module_code   VARCHAR (10),
   pruview       VARCHAR (2),
   primary key (id),
   foreign key (father_id) references pruview(id)
);

foreign key (father_id) references pruview(id)
constraint pruview_fk foreign key (father_id) references pruview(id)


constraint 外键别名
foreign key(要添加外键的字段名)
references 表名（字段名）

CREATE TABLE Ordersei
(
Id_O int NOT NULL PRIMARY KEY,
OrderNo int NOT NULL,
Id_P int FOREIGN KEY REFERENCES Persons(Id_P)
)




MYSQL外键(Foreign Key)的使用

在MySQL 3.23.44版本后，InnoDB引擎类型的表支持了外键约束。
外键的使用条件：
1.两个表必须是InnoDB表，MyISAM表暂时不支持外键（据说以后的版本有可能支持，但至少目前不支持）；
2.外键列必须建立了索引，MySQL 4.1.2以后的版本在建立外键时会自动创建索引，但如果在较早的版本则需要显示建立；
3.外键关系的两个表的列必须是数据类型相似，也就是可以相互转换类型的列，比如int和tinyint可以，而int和char则不可以；

外键的好处：可以使得两张表关联，保证数据的一致性和实现一些级联操作；

外键的定义语法：
[CONSTRAINT symbol] FOREIGN KEY [id] (index_col_name, ...)
    REFERENCES tbl_name (index_col_name, ...)
    [ON DELETE {RESTRICT | CASCADE | SET NULL | NO ACTION | SET DEFAULT}]
    [ON UPDATE {RESTRICT | CASCADE | SET NULL | NO ACTION | SET DEFAULT}]
该语法可以在 CREATE TABLE 和 ALTER TABLE 时使用，如果不指定CONSTRAINT symbol，MYSQL会自动生成一个名字。
ON DELETE、ON UPDATE表示事件触发限制，可设参数：
RESTRICT（限制外表中的外键改动）
CASCADE（跟随外键改动）
SET NULL（设空值）
SET DEFAULT（设默认值）
NO ACTION（无动作，默认的）

搞个例子，简单演示一下使用，做dage和xiaodi两个表，大哥表是主键，小弟表是外键：
建表：
 1CREATE TABLE `dage` (
 2  `id` int(11) NOT NULL auto_increment,
 3  `name` varchar(32) default '',
 4  PRIMARY KEY  (`id`)
 5) ENGINE=InnoDB DEFAULT CHARSET=latin1；
 6
 7CREATE TABLE `xiaodi` (
 8  `id` int(11) NOT NULL auto_increment,
 9  `dage_id` int(11) default NULL,
10  `name` varchar(32) default '',
11  PRIMARY KEY  (`id`),
12  KEY `dage_id` (`dage_id`),
13  CONSTRAINT `xiaodi_ibfk_1` FOREIGN KEY (`dage_id`) REFERENCES `dage` (`id`)
14) ENGINE=InnoDB DEFAULT CHARSET=latin1；

插入个大哥：
1mysql> insert into dage(name) values('铜锣湾');
2Query OK, 1 row affected (0.01 sec)
3mysql> select * from dage;
4+----+--------+
5| id | name   |
6+----+--------+
7|  1 | 铜锣湾 |
8+----+--------+
91 row in set (0.00 sec)

插入个小弟：
1mysql> insert into xiaodi(dage_id,name) values(1,'铜锣湾_小弟A');
2Query OK, 1 row affected (0.02 sec)
3
4mysql> select * from xiaodi;
5+----+---------+--------------+
6| id | dage_id | name         |
7+----+---------+--------------+
8|  1 |       1 | 铜锣湾_小弟A |
9+----+---------+--------------+

把大哥删除：
1mysql> delete from dage where id=1;
2ERROR 1451 (23000): Cannot delete or update a parent row: a foreign key constraint fails (`bstar/xiaodi`, CONSTRAINT `xiaodi_ibfk_1` FOREIGN KEY (`dage_id`) REFERENCES `dage` (`id`))


提示：不行呀，有约束的，大哥下面还有小弟，可不能扔下我们不管呀！

插入一个新的小弟：
1mysql> insert into xiaodi(dage_id,name) values(2,'旺角_小弟A');              
2ERROR 1452 (23000): Cannot add or update a child row: a foreign key constraint fails (`bstar/xiaodi`, CONSTRAINT `xiaodi_ibfk_1` FOREIGN KEY (`dage_id`) REFERENCES `dage` (`id`))
3


提示：小子，想造反呀！你还没大哥呢！

把外键约束增加事件触发限制：
 1mysql> show create table xiaodi;
 2
 3  CONSTRAINT `xiaodi_ibfk_1` FOREIGN KEY (`dage_id`) REFERENCES `dage` (`id`)
 4
 5mysql> alter table xiaodi drop foreign key xiaodi_ibfk_1; 
 6Query OK, 1 row affected (0.04 sec)
 7Records: 1  Duplicates: 0  Warnings: 
 8mysql> alter table xiaodi add foreign key(dage_id) references dage(id) on delete cascade on update cascade;
 9Query OK, 1 row affected (0.04 sec)
10Records: 1  Duplicates: 0  Warnings: 0

再次试着把大哥删了：
1mysql> delete from dage where id=1;
2Query OK, 1 row affected (0.01 sec)
3
4mysql> select * from dage;
5Empty set (0.01 sec)
6
7mysql> select * from xiaodi;
8Empty set (0.00 sec)

1. 建表时创建外键： 
CREATE TABLE`xh` (
 `id` int(100) unsigned NOT NULL AUTO_INCREMENT COMMENT ,
 `cl_id` smallint(3) unsigned NOT NULL COMMENT,
 `title` varchar(100) COLLATE utf8_unicode_ci NOT NULL COMMENT ,
 `details` text COLLATE utf8_unicode_ci NOT NULL COMMENT ,
 `date` datetime NOT NULL COMMENT ,
 `au_id` smallint(6) unsigned NOT NULL COMMENT ,
 `click` int(100) unsigned NOT NULL DEFAULT '0' COMMENT ,
 `reco` int(100) unsigned NOT NULL DEFAULT '0' COMMENT,
 PRIMARY KEY (`id`),
 KEY `fk_class` (`cl_id`),
 CONSTRAINT `fk_class`FOREIGN KEY  (`cl_id`)  REFERENCES  `fl` (`id`),
 KEY `fk_author` (`au_id`),
 CONSTRAINT `fk_author` FOREIGN KEY (`au_id`) REFERENCES `author` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci  

2. 建表后创建外键： 
Alter Table `ym` Add Constraint `fk_author` Foreign Key (`au_id`) References `author` (`id`);   

3. 查看表结构： 
Show Create Table `ym`; 

4.级联： 在末尾可加上（可单独添加，也可全部添加）：
ON UPDATE CASCADE（级联更新）
ON DELETE CASCADE（级联删除） 

比如：

ALTER TABLE `ym`  ADD  
CONSTRAINT  `fk_author`  
FOREIGN  KEY  (`au_id`)  
REFERENCES  `author`  (`id`) ON UPDATE CASCADE
