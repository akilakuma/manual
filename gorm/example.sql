-- phpMyAdmin SQL Dump
-- version 5.0.2
-- https://www.phpmyadmin.net/
--
-- 主機： mysql
-- 產生時間： 2020 年 08 月 17 日 06:38
-- 伺服器版本： 5.7.31
-- PHP 版本： 7.4.6

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- 資料庫： `example`
--
CREATE DATABASE IF NOT EXISTS `example` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;
USE `example`;

-- --------------------------------------------------------

--
-- 資料表結構 `trans_record`
--

CREATE TABLE `trans_record` (
  `id` bigint(20) NOT NULL,
  `user_id` bigint(20) NOT NULL,
  `account` varchar(100) NOT NULL,
  `point` decimal(20,4) NOT NULL,
  `opcode` int(11) NOT NULL,
  `created_at` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- 傾印資料表的資料 `trans_record`
--

INSERT INTO `trans_record` (`id`, `user_id`, `account`, `point`, `opcode`, `created_at`) VALUES
(1, 6653, 'bananaKing', '10.5000', 1001, 1597645920),
(2, 6653, 'bananaKing', '10.5000', 1001, 1597646085),
(3, 6653, 'bananaKing', '10.5000', 1001, 1597646204);

--
-- 已傾印資料表的索引
--

--
-- 資料表索引 `trans_record`
--
ALTER TABLE `trans_record`
  ADD PRIMARY KEY (`id`);

--
-- 在傾印的資料表使用自動遞增(AUTO_INCREMENT)
--

--
-- 使用資料表自動遞增(AUTO_INCREMENT) `trans_record`
--
ALTER TABLE `trans_record`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
