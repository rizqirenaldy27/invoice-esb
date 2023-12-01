-- MySQL dump 10.13  Distrib 5.7.44, for Linux (x86_64)
--
-- Host: localhost    Database: invoice
-- ------------------------------------------------------
-- Server version	5.7.44

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
-- Table structure for table `customers`
--

DROP TABLE IF EXISTS `customers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `customers` (
  `customer_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `customer_name` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `detail_address` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`customer_id`),
  UNIQUE KEY `unique_constraint_customer_name` (`customer_name`)
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `customers`
--

LOCK TABLES `customers` WRITE;
/*!40000 ALTER TABLE `customers` DISABLE KEYS */;
INSERT INTO `customers` VALUES (13,'PT Maju Jaya Sejahtera','Jl. Merdeka No. 123, Jakarta Barat','2023-11-29 06:45:44','2023-11-30 06:33:16'),(21,'CV Berkah Makmur',' Jl. Pahlawan No. 456, Surabaya','2023-11-29 07:01:32','2023-11-30 06:33:16'),(22,'PT Sentosa Abadi','Jl. Sudirman No. 789, Bandung','2023-11-29 07:01:34','2023-11-30 06:33:16'),(23,'PT Cemerlang Mandiri','Jl. Asia Tenggara No. 202, Semarang','2023-11-29 07:01:40','2023-11-30 06:33:16');
/*!40000 ALTER TABLE `customers` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `invoice_items`
--

DROP TABLE IF EXISTS `invoice_items`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `invoice_items` (
  `invoice_item_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `invoice_id` bigint(20) DEFAULT NULL,
  `item_id` bigint(20) DEFAULT NULL,
  `quantity` int(11) DEFAULT NULL,
  `amount` decimal(10,2) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`invoice_item_id`),
  KEY `invoice_id` (`invoice_id`),
  KEY `item_id` (`item_id`),
  CONSTRAINT `invoice_items_ibfk_1` FOREIGN KEY (`invoice_id`) REFERENCES `invoices` (`invoice_id`),
  CONSTRAINT `invoice_items_ibfk_2` FOREIGN KEY (`item_id`) REFERENCES `items` (`item_id`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `invoice_items`
--

LOCK TABLES `invoice_items` WRITE;
/*!40000 ALTER TABLE `invoice_items` DISABLE KEYS */;
INSERT INTO `invoice_items` VALUES (1,2,1,2,200.00,'2023-12-01 01:57:48.421','2023-12-01 03:33:08.261'),(2,2,5,3,600.00,'2023-12-01 01:57:48.472','2023-12-01 03:33:08.362'),(3,3,1,2,200.00,'2023-12-01 02:05:23.582','2023-12-01 02:05:23.582'),(4,3,5,3,600.00,'2023-12-01 02:05:23.634','2023-12-01 02:05:23.634'),(5,4,1,2,200.00,'2023-12-01 02:06:29.936','2023-12-01 02:06:29.936'),(6,4,5,3,600.00,'2023-12-01 02:06:29.982','2023-12-01 02:06:29.982'),(7,5,1,2,200.00,'2023-12-01 02:17:40.361','2023-12-01 02:17:40.361'),(8,5,5,3,600.00,'2023-12-01 02:17:40.412','2023-12-01 02:17:40.412'),(11,7,1,2,200.00,'2023-12-01 03:27:32.570','2023-12-01 03:27:32.570'),(12,7,5,3,600.00,'2023-12-01 03:27:32.618','2023-12-01 03:27:32.618'),(13,8,1,2,200.00,'2023-12-01 03:35:18.359','2023-12-01 03:35:18.359'),(14,8,5,3,600.00,'2023-12-01 03:35:18.411','2023-12-01 03:35:18.411'),(15,9,1,2,200.00,'2023-12-01 03:35:26.579','2023-12-01 03:35:26.579'),(16,9,5,3,600.00,'2023-12-01 03:35:26.629','2023-12-01 03:35:26.629'),(17,10,1,2,200.00,'2023-12-01 03:35:33.591','2023-12-01 03:35:33.591'),(18,10,5,3,600.00,'2023-12-01 03:35:33.640','2023-12-01 03:35:33.640');
/*!40000 ALTER TABLE `invoice_items` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `invoices`
--

DROP TABLE IF EXISTS `invoices`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `invoices` (
  `invoice_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `issue_date` date DEFAULT NULL,
  `customer_id` bigint(20) DEFAULT NULL,
  `subject` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `due_date` date DEFAULT NULL,
  `status` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `total_items` int(11) DEFAULT NULL,
  `sub_total` decimal(10,2) DEFAULT NULL,
  `tax` decimal(10,2) DEFAULT NULL,
  `grand_total` decimal(10,2) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`invoice_id`),
  KEY `customer_id` (`customer_id`),
  CONSTRAINT `invoices_ibfk_1` FOREIGN KEY (`customer_id`) REFERENCES `customers` (`customer_id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `invoices`
--

LOCK TABLES `invoices` WRITE;
/*!40000 ALTER TABLE `invoices` DISABLE KEYS */;
INSERT INTO `invoices` VALUES (2,'2018-05-06',13,'Spring Marketing Campaign','2018-05-06','UNPAID',2,800.00,88.00,888.00,'2023-12-01 01:57:48.419','2023-12-01 03:33:08.208'),(3,'2017-05-06',13,'Spring Marketing Campaign','2017-05-06','UNPAID',2,800.00,88.00,888.00,'2023-12-01 02:05:23.580','2023-12-01 02:05:23.580'),(4,'2017-05-06',13,'Spring Marketing Campaign','2017-05-06','UNPAID',2,800.00,88.00,888.00,'2023-12-01 02:06:29.934','2023-12-01 02:06:29.934'),(5,'2017-05-06',13,'Spring Marketing Campaign','2017-05-06','UNPAID',2,800.00,88.00,888.00,'2023-12-01 02:17:40.359','2023-12-01 02:17:40.359'),(7,'2017-05-06',13,'Spring Marketing Campaign','2017-05-06','UNPAID',2,800.00,88.00,888.00,'2023-12-01 03:27:32.569','2023-12-01 03:27:32.569'),(8,'2017-05-06',13,'Spring IT Campaign','2017-05-06','UNPAID',2,800.00,88.00,888.00,'2023-12-01 03:35:18.358','2023-12-01 03:35:18.358'),(9,'2017-05-06',13,'Spring ERP Campaign','2017-05-06','UNPAID',2,800.00,88.00,888.00,'2023-12-01 03:35:26.578','2023-12-01 03:35:26.578'),(10,'2017-05-06',13,'Spring Test Campaign','2017-05-06','UNPAID',2,800.00,88.00,888.00,'2023-12-01 03:35:33.590','2023-12-01 03:35:33.590');
/*!40000 ALTER TABLE `invoices` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `item_types`
--

DROP TABLE IF EXISTS `item_types`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `item_types` (
  `item_type_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `item_type_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`item_type_id`),
  UNIQUE KEY `idx_item_types_item_type_name` (`item_type_name`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `item_types`
--

LOCK TABLES `item_types` WRITE;
/*!40000 ALTER TABLE `item_types` DISABLE KEYS */;
INSERT INTO `item_types` VALUES (1,'Hardware','2023-11-29 15:31:38.841','2023-11-29 15:31:38.841'),(2,'Service','2023-11-29 15:32:04.040','2023-11-29 15:32:04.040');
/*!40000 ALTER TABLE `item_types` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `items`
--

DROP TABLE IF EXISTS `items`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `items` (
  `item_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `item_type_id` bigint(20) NOT NULL,
  `item_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `unit_price` double DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`item_id`),
  UNIQUE KEY `unique_item_name` (`item_name`),
  KEY `fk_items_item_types` (`item_type_id`),
  CONSTRAINT `fk_items_item_types` FOREIGN KEY (`item_type_id`) REFERENCES `item_types` (`item_type_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `items`
--

LOCK TABLES `items` WRITE;
/*!40000 ALTER TABLE `items` DISABLE KEYS */;
INSERT INTO `items` VALUES (1,2,'Design',100,'2023-11-30 05:14:45.299','2023-11-30 05:14:45.299'),(4,2,'Development',200,'2023-11-30 05:19:49.507','2023-11-30 05:19:49.507'),(5,2,'Meetings',150,'2023-11-30 05:19:55.859','2023-11-30 05:19:55.859'),(6,1,'Printer',90,'2023-11-30 05:20:04.288','2023-11-30 05:31:23.013'),(7,1,'Monitor',70,'2023-11-30 05:20:10.011','2023-11-30 05:20:10.011'),(8,1,'Meeting',60,'2023-12-01 03:27:21.085','2023-12-01 03:27:21.085');
/*!40000 ALTER TABLE `items` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-12-01  3:54:26
