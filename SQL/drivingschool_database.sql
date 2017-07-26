-- MySQL dump 10.13  Distrib 5.7.17, for Linux (x86_64)
--
-- Host: localhost    Database: drivingschool
-- ------------------------------------------------------
-- Server version	5.7.17-0ubuntu0.16.10.1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `courseInfo`
--

DROP TABLE IF EXISTS `courseInfo`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `courseInfo` (
  `cno` int(4) NOT NULL COMMENT '科目号',
  `cname` varchar(20) NOT NULL COMMENT '科目名称',
  `before_cour` int(4) NOT NULL DEFAULT '0' COMMENT '先行考试科目',
  PRIMARY KEY (`cno`),
  UNIQUE KEY `cno` (`cno`),
  UNIQUE KEY `cname` (`cname`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='考试科目信息表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `courseInfo`
--

LOCK TABLES `courseInfo` WRITE;
/*!40000 ALTER TABLE `courseInfo` DISABLE KEYS */;
/*!40000 ALTER TABLE `courseInfo` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `gradeInfo`
--

DROP TABLE IF EXISTS `gradeInfo`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `gradeInfo` (
  `id` int(8) NOT NULL AUTO_INCREMENT COMMENT '编号',
  `sno` int(8) NOT NULL COMMENT '学号',
  `cno` int(4) NOT NULL COMMENT '科目号',
  `last_time` date DEFAULT NULL COMMENT '考试时间',
  `times` int(4) DEFAULT '1' COMMENT '考试次数',
  `grade` float DEFAULT '0' COMMENT '成绩',
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`),
  KEY `grade_sno_fk` (`sno`),
  KEY `grade_cno_fk` (`cno`),
  CONSTRAINT `grade_cno_fk` FOREIGN KEY (`cno`) REFERENCES `courseInfo` (`cno`),
  CONSTRAINT `grade_sno_fk` FOREIGN KEY (`sno`) REFERENCES `studentInfo` (`sno`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='学员的成绩信息表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `gradeInfo`
--

LOCK TABLES `gradeInfo` WRITE;
/*!40000 ALTER TABLE `gradeInfo` DISABLE KEYS */;
/*!40000 ALTER TABLE `gradeInfo` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Temporary view structure for view `grade_view`
--

DROP TABLE IF EXISTS `grade_view`;
/*!50001 DROP VIEW IF EXISTS `grade_view`*/;
SET @saved_cs_client     = @@character_set_client;
SET character_set_client = utf8;
/*!50001 CREATE VIEW `grade_view` AS SELECT 
 1 AS `id`,
 1 AS `sno`,
 1 AS `sname`,
 1 AS `cname`,
 1 AS `last_time`,
 1 AS `times`,
 1 AS `grade`*/;
SET character_set_client = @saved_cs_client;

--
-- Table structure for table `healthInfo`
--

DROP TABLE IF EXISTS `healthInfo`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `healthInfo` (
  `id` int(8) NOT NULL AUTO_INCREMENT COMMENT '编号',
  `sno` int(8) NOT NULL COMMENT '学号',
  `sname` varchar(20) NOT NULL COMMENT '姓名',
  `height` float DEFAULT NULL COMMENT '身高',
  `weight` float DEFAULT NULL COMMENT '体重',
  `differentiate` enum('正常','色弱','色盲') NOT NULL COMMENT '辨色',
  `left_sight` float DEFAULT NULL COMMENT '左眼视力',
  `right_sight` float DEFAULT NULL COMMENT '右眼视力',
  `left_ear` enum('正常','偏弱') DEFAULT NULL COMMENT '左耳听力',
  `right_ear` enum('正常','偏弱') DEFAULT NULL COMMENT '右耳听力',
  `legs` enum('正常','不相等') DEFAULT NULL COMMENT '腿长是否相等',
  `pressure` enum('正常','不相等') DEFAULT NULL COMMENT '血压',
  `history` varchar(50) DEFAULT NULL COMMENT '病史',
  `h_text` text COMMENT '备注',
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`),
  UNIQUE KEY `sno` (`sno`),
  KEY `index_health_name` (`sname`),
  CONSTRAINT `health_fk` FOREIGN KEY (`sno`) REFERENCES `studentInfo` (`sno`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='学员体检表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `healthInfo`
--

LOCK TABLES `healthInfo` WRITE;
/*!40000 ALTER TABLE `healthInfo` DISABLE KEYS */;
/*!40000 ALTER TABLE `healthInfo` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `licenseInfo`
--

DROP TABLE IF EXISTS `licenseInfo`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `licenseInfo` (
  `id` int(8) NOT NULL AUTO_INCREMENT COMMENT '编号',
  `sno` int(8) NOT NULL COMMENT '学号',
  `sname` varchar(20) NOT NULL COMMENT '姓名',
  `lno` varchar(18) NOT NULL COMMENT '驾驶证号',
  `receive_time` date DEFAULT NULL COMMENT '领证时间',
  `receive_name` varchar(20) DEFAULT NULL COMMENT '领证人',
  `l_text` text COMMENT '备注',
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`),
  UNIQUE KEY `lno` (`lno`),
  KEY `license_sno_fk` (`sno`),
  KEY `index_license_name` (`sname`),
  KEY `index_license_receive_name` (`receive_name`),
  CONSTRAINT `license_sno_fk` FOREIGN KEY (`sno`) REFERENCES `studentInfo` (`sno`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='学员领证表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `licenseInfo`
--

LOCK TABLES `licenseInfo` WRITE;
/*!40000 ALTER TABLE `licenseInfo` DISABLE KEYS */;
/*!40000 ALTER TABLE `licenseInfo` ENABLE KEYS */;
UNLOCK TABLES;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8 */ ;
/*!50003 SET character_set_results = utf8 */ ;
/*!50003 SET collation_connection  = utf8_general_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
/*!50003 CREATE*/ /*!50017 DEFINER=`root`@`localhost`*/ /*!50003 trigger license_stu after insert 
on licenseInfo for each row
begin
update stuentInfo set leave_time=new.receive_time,scondition='结业'

where sno=new.sno;
end */;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;

--
-- Table structure for table `studentInfo`
--

DROP TABLE IF EXISTS `studentInfo`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `studentInfo` (
  `sno` int(8) NOT NULL COMMENT '学号',
  `sname` varchar(20) NOT NULL COMMENT '姓名',
  `sex` enum('男','女') NOT NULL COMMENT '性别',
  `age` int(3) DEFAULT NULL COMMENT '年龄',
  `identify` varchar(18) NOT NULL COMMENT '身份证号',
  `tel` varchar(15) DEFAULT NULL COMMENT '电话',
  `car_type` varchar(4) NOT NULL COMMENT '报考车型',
  `enroll_time` date NOT NULL COMMENT '入学时间',
  `leave_time` date DEFAULT NULL COMMENT '毕业时间',
  `scondition` enum('学习','结业','退学') NOT NULL COMMENT '学业状态',
  `s_text` text COMMENT '备注',
  PRIMARY KEY (`sno`),
  UNIQUE KEY `sno` (`sno`),
  UNIQUE KEY `identify` (`identify`),
  KEY `index_stu_name` (`sname`),
  KEY `index_car` (`car_type`),
  KEY `index_con` (`scondition`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='学员学籍信息';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `studentInfo`
--

LOCK TABLES `studentInfo` WRITE;
/*!40000 ALTER TABLE `studentInfo` DISABLE KEYS */;
/*!40000 ALTER TABLE `studentInfo` ENABLE KEYS */;
UNLOCK TABLES;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8 */ ;
/*!50003 SET character_set_results = utf8 */ ;
/*!50003 SET collation_connection  = utf8_general_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
/*!50003 CREATE*/ /*!50017 DEFINER=`root`@`localhost`*/ /*!50003 trigger update_sname after update 
on studentInfo for each row
begin
update healthInfo set sname=new.sname where sno=new.sno;
update licenseInfo set sname=new.sname where sno=new.sno;
end */;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8 */ ;
/*!50003 SET character_set_results = utf8 */ ;
/*!50003 SET collation_connection  = utf8_general_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
/*!50003 CREATE*/ /*!50017 DEFINER=`root`@`localhost`*/ /*!50003 trigger delete_stu after delete
on studentInfo for each row
begin
delete from gradeInfo where sno=old.sno;
delete from healthInfo where sno=old.sno;
delete from licenseInfo where sno=old.sno;
end */;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user` (
  `username` varchar(20) NOT NULL,
  `password` varchar(20) NOT NULL,
  PRIMARY KEY (`username`),
  UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='user table';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Final view structure for view `grade_view`
--

/*!50001 DROP VIEW IF EXISTS `grade_view`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8 */;
/*!50001 SET character_set_results     = utf8 */;
/*!50001 SET collation_connection      = utf8_general_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`root`@`localhost` SQL SECURITY DEFINER */
/*!50001 VIEW `grade_view` AS select `g`.`id` AS `id`,`g`.`sno` AS `sno`,`s`.`sname` AS `sname`,`c`.`cname` AS `cname`,`g`.`last_time` AS `last_time`,`g`.`times` AS `times`,`g`.`grade` AS `grade` from ((`studentInfo` `s` join `courseInfo` `c`) join `gradeInfo` `g`) where ((`g`.`sno` = `s`.`sno`) and (`g`.`cno` = `c`.`cno`)) */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2017-02-15 16:33:15
