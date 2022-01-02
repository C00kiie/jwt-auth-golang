-- MySQL dump 10.13  Distrib 8.0.27, for Linux (x86_64)
--
-- Host: localhost    Database: goauth
-- ------------------------------------------------------
-- Server version	8.0.27-0ubuntu0.20.04.1

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
-- Table structure for table `authentication`
--

DROP TABLE IF EXISTS `authentication`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `authentication` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(256) DEFAULT NULL,
  `password` varchar(256) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username_UNIQUE` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `authentication`
--

LOCK TABLES `authentication` WRITE;
/*!40000 ALTER TABLE `authentication` DISABLE KEYS */;
INSERT INTO `authentication` VALUES (1,'nigga','$2a$14$KLhlm2QW7Spel2Xj0nKnq.fb3dysX/6/YRuFam7EzGPKwnU95vFRW'),(2,'salom','$2a$14$JKg3YBuvD0cS6/W8cHKapOCsbbTAlXo2QJB3E/AV255h6iurque6q'),(8,'hasan','$2a$14$AHuZ4bMEAWJ2XRH0Y0tvau755WWnxDUvmdHObtsmvqQz/R383dRfK'),(10,'unique','$2a$14$7oIKRMLNj4H/r6oetT9OiOM3xQMHwMK1L8Tn48ubMaJdh2wf21Yku'),(11,'cookie','$2a$14$uHUF7Um9gnds2dTAlxzkcuOD34WP8uPBpUy5FDLci3H84t3HH9.Nu'),(12,'lolnoyou','$2a$14$G07PPD.vU0vrVo/0Bk8cCOHm37ZehyTWFMGLPSsa3hiuC1dLuXqvO'),(15,'test123','$2a$14$pf8eUhq/2g2Hlct7HY8hGu4JNHy3u40REdiL85SakEzJjK3ZIRKtK');
/*!40000 ALTER TABLE `authentication` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `jwt_sessions`
--

DROP TABLE IF EXISTS `jwt_sessions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `jwt_sessions` (
  `id` int NOT NULL AUTO_INCREMENT,
  `userId` int DEFAULT NULL,
  `jwt_token` text,
  `created_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_jwt_sessions_1_idx` (`userId`),
  CONSTRAINT `fk_jwt_sessions_1` FOREIGN KEY (`userId`) REFERENCES `authentication` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `jwt_sessions`
--

LOCK TABLES `jwt_sessions` WRITE;
/*!40000 ALTER TABLE `jwt_sessions` DISABLE KEYS */;
INSERT INTO `jwt_sessions` VALUES (1,11,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NDExMTYzNjAsInVzZXJfaWQiOjExfQ.YPUsesra_sKkvJSjncLrkLWJYEkXwqbJ9LQOdeeXmoM','2022-01-02 06:24:21'),(2,12,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NDExMjcxMTgsInVzZXJfaWQiOjEyfQ.92_zMGq4Tgmq3-QEgSPML5nEsjDDAtHbjMXKN8D2mpo','2022-01-02 08:38:39'),(3,15,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NDExMjcxNjUsInVzZXJfaWQiOjE1fQ.PWcVqToNVLbfL3jrbSDQ04MEhntJHLKmUpGt3MvbnY4','2022-01-02 08:39:25'),(4,15,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NDExMjczNjIsInVzZXJfaWQiOjE1fQ.D4wsfa8CUHLvO063Qi1ISdUtZxDvL92kKFOqgn1uXAc','2022-01-02 08:42:42'),(5,15,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NDExMjg2OTYsInVzZXJfaWQiOjE1fQ.i-jNQpdB1PFLDT6ge7hN19sLl623Hr5x1MTL6zaiml4','2022-01-02 09:04:56'),(6,12,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NDExMjkwODYsInVzZXJfaWQiOjEyfQ.AGnZPneSUVvynkDJqtusGXtal15ZdKtI2ywm6AycNwM','2022-01-02 09:11:27'),(7,12,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NDExMzMwMzIsInVzZXJfaWQiOjEyfQ.4rF5545Wl4I0lkeCLD3a4DF0AnNt4xc4P85YV6Os0Xg','2022-01-02 10:17:12');
/*!40000 ALTER TABLE `jwt_sessions` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-01-02 16:20:41
