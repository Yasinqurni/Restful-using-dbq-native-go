-- MySQL dump 10.13  Distrib 8.0.33, for macos12.6 (arm64)
--
-- Host: localhost    Database: sa_academy_practice
-- ------------------------------------------------------
-- Server version	8.0.33

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `m_class`
--

DROP TABLE IF EXISTS `m_class`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `m_class` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `code` varchar(10) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `m_class_id_uindex` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `m_class`
--

LOCK TABLES `m_class` WRITE;
/*!40000 ALTER TABLE `m_class` DISABLE KEYS */;
INSERT INTO `m_class` VALUES (1,'Information Technology 2011','2011_TI'),(2,'Information Technology 2009','2009_TI'),(3,'Information Technology 2012','2012_TI');
/*!40000 ALTER TABLE `m_class` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `m_lesson`
--

DROP TABLE IF EXISTS `m_lesson`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `m_lesson` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `code` varchar(255) NOT NULL,
  `isExternal` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `m_lesson_code_uindex` (`code`),
  UNIQUE KEY `m_lesson_id_uindex` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `m_lesson`
--

LOCK TABLES `m_lesson` WRITE;
/*!40000 ALTER TABLE `m_lesson` DISABLE KEYS */;
INSERT INTO `m_lesson` VALUES (1,'Object Oriented Programming OOP','C_OOP',0),(2,'Structure Data','C_STRUCTDAT',0),(3,'Logic Algorithm','C_LOG_ALGO',0),(4,'Data Statistic','C_STATISTIC',0),(5,'Computer Science C++','C_COMP_C++',0),(6,'Toefl English','C_ENGLISH',0);
/*!40000 ALTER TABLE `m_lesson` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `m_mahasiswa`
--

DROP TABLE IF EXISTS `m_mahasiswa`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `m_mahasiswa` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `nama` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `nim` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `tempat_lahir` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `no_hp` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `jenis_kelamin` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `alamat` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `tanggal_lahir` date DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `id_class` int NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `m_mahasiswa`
--

LOCK TABLES `m_mahasiswa` WRITE;
/*!40000 ALTER TABLE `m_mahasiswa` DISABLE KEYS */;
INSERT INTO `m_mahasiswa` VALUES (1,'Maleakhi Rifandy','2100921231','Jakarta','081373100015','Pria','Jl Budi Rahayu III No 29 RT 03 RW 09','1993-05-12','2023-08-01 18:04:05',1),(2,'Defandry Gunawan','1776351233','Jakarta','085521372558','Pria','Jl Kawan Lama','1989-10-09','2023-08-01 18:04:05',2);
/*!40000 ALTER TABLE `m_mahasiswa` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `p_lesson_mahasiswa`
--

DROP TABLE IF EXISTS `p_lesson_mahasiswa`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `p_lesson_mahasiswa` (
  `id` int NOT NULL AUTO_INCREMENT,
  `lesson_id` int NOT NULL,
  `mahasiswa_id` int NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `p_lesson_mahasiswa_id_uindex` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `p_lesson_mahasiswa`
--

LOCK TABLES `p_lesson_mahasiswa` WRITE;
/*!40000 ALTER TABLE `p_lesson_mahasiswa` DISABLE KEYS */;
/*!40000 ALTER TABLE `p_lesson_mahasiswa` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-08-01 18:10:18
