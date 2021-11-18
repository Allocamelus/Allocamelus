SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';
USE `allocamelus`;
SET NAMES utf8mb4;
CREATE TABLE `PostCommentClosures` (
  `parent` bigint(20) NOT NULL,
  `child` bigint(20) NOT NULL,
  `depth` bigint(20) NOT NULL,
  PRIMARY KEY (`parent`, `child`),
  UNIQUE KEY `parent_depth_child` (`parent`, `depth`, `child`),
  UNIQUE KEY `child_parent_depth` (`child`, `parent`, `depth`),
  KEY `parent` (`parent`),
  KEY `child` (`child`),
  KEY `depth` (`depth`),
  CONSTRAINT `PostCommentClosures_ibfk_7` FOREIGN KEY (`parent`) REFERENCES `PostComments` (`postCommentId`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `PostCommentClosures_ibfk_8` FOREIGN KEY (`child`) REFERENCES `PostComments` (`postCommentId`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
CREATE TABLE `PostComments` (
  `postCommentId` bigint(20) NOT NULL AUTO_INCREMENT,
  `postId` bigint(20) NOT NULL,
  `userId` bigint(20) NOT NULL,
  `parent` bigint(20) NOT NULL DEFAULT 0 COMMENT '0 or postCommentId',
  `created` int(11) NOT NULL COMMENT 'Time',
  `updated` int(11) NOT NULL DEFAULT 0 COMMENT 'Time',
  `content` varchar(4096) NOT NULL,
  PRIMARY KEY (`postCommentId`),
  KEY `parent` (`parent`),
  KEY `userId` (`userId`),
  KEY `postId` (`postId`),
  CONSTRAINT `PostComments_ibfk_8` FOREIGN KEY (`userId`) REFERENCES `Users` (`userId`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `PostComments_ibfk_9` FOREIGN KEY (`postId`) REFERENCES `Posts` (`postId`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
CREATE TABLE `PostMedia` (
  `postMediaId` bigint(20) NOT NULL AUTO_INCREMENT,
  `postId` bigint(20) NOT NULL,
  `created` int(11) NOT NULL COMMENT 'Time',
  `active` tinyint(1) NOT NULL DEFAULT 1 COMMENT 'Bool',
  `fileType` int(11) NOT NULL,
  `meta` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL,
  `hash` varchar(128) NOT NULL,
  PRIMARY KEY (`postMediaId`),
  KEY `postId` (`postId`),
  KEY `active` (`active`),
  KEY `hash` (`hash`),
  KEY `fileType` (`fileType`),
  CONSTRAINT `PostMedia_ibfk_7` FOREIGN KEY (`postId`) REFERENCES `Posts` (`postId`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
CREATE TABLE `Posts` (
  `postId` bigint(20) NOT NULL AUTO_INCREMENT,
  `userId` bigint(20) NOT NULL,
  `created` int(11) NOT NULL COMMENT 'Time',
  `published` int(11) NOT NULL COMMENT 'Time',
  `updated` int(11) NOT NULL DEFAULT 0 COMMENT 'Time',
  `content` text NOT NULL,
  PRIMARY KEY (`postId`),
  KEY `userId` (`userId`),
  CONSTRAINT `Posts_ibfk_1` FOREIGN KEY (`userId`) REFERENCES `Users` (`userId`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
CREATE TABLE `UserAvatars` (
  `userAvatarId` bigint(20) NOT NULL AUTO_INCREMENT,
  `userId` bigint(20) NOT NULL,
  `active` tinyint(1) NOT NULL DEFAULT 1 COMMENT 'Bool',
  `created` int(11) NOT NULL COMMENT 'Time',
  `fileType` int(11) NOT NULL,
  `hash` varchar(128) NOT NULL,
  PRIMARY KEY (`userAvatarId`),
  KEY `userId` (`userId`),
  KEY `active` (`active`),
  KEY `fileType` (`fileType`),
  KEY `hash` (`hash`),
  CONSTRAINT `UserAvatars_ibfk_1` FOREIGN KEY (`userId`) REFERENCES `Users` (`userId`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
CREATE TABLE `UserEventKeys` (
  `userEventKeyId` bigint(20) NOT NULL AUTO_INCREMENT,
  `userEventId` int(11) NOT NULL,
  `userKeyId` bigint(20) NOT NULL,
  `infoKey` varchar(1024) NOT NULL COMMENT 'Encrypted AES Key w/ Public key',
  PRIMARY KEY (`userEventKeyId`),
  KEY `userEventId` (`userEventId`),
  KEY `userKeyId` (`userKeyId`),
  CONSTRAINT `UserEventKeys_ibfk_1` FOREIGN KEY (`userEventId`) REFERENCES `UserEvents` (`userEventId`) ON DELETE CASCADE,
  CONSTRAINT `UserEventKeys_ibfk_2` FOREIGN KEY (`userKeyId`) REFERENCES `UserKeys` (`userKeyId`) ON DELETE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
CREATE TABLE `UserEvents` (
  `userEventId` int(11) NOT NULL AUTO_INCREMENT,
  `userId` bigint(20) NOT NULL,
  `eventType` tinyint(4) NOT NULL,
  `created` int(11) NOT NULL COMMENT 'Time',
  `info` varchar(4096) NOT NULL COMMENT 'Encrypted Bytes',
  PRIMARY KEY (`userEventId`),
  KEY `userId` (`userId`),
  KEY `created` (`created`),
  CONSTRAINT `UserEvents_ibfk_1` FOREIGN KEY (`userId`) REFERENCES `Users` (`userId`) ON DELETE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
CREATE TABLE `UserFollows` (
  `userFollowId` bigint(20) NOT NULL AUTO_INCREMENT,
  `userId` bigint(20) NOT NULL COMMENT 'Follower',
  `followUserId` bigint(20) NOT NULL COMMENT 'Following',
  `accepted` tinyint(1) NOT NULL COMMENT 'Bool',
  `created` int(11) NOT NULL COMMENT 'Time',
  PRIMARY KEY (`userFollowId`),
  KEY `userId` (`userId`),
  KEY `followUserId` (`followUserId`),
  CONSTRAINT `UserFollows_ibfk_3` FOREIGN KEY (`userId`) REFERENCES `Users` (`userId`) ON DELETE CASCADE,
  CONSTRAINT `UserFollows_ibfk_4` FOREIGN KEY (`followUserId`) REFERENCES `Users` (`userId`) ON DELETE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
CREATE TABLE `UserKeys` (
  `userKeyId` bigint(20) NOT NULL AUTO_INCREMENT,
  `userId` bigint(20) NOT NULL,
  `created` int(11) NOT NULL COMMENT 'Time',
  `replaced` int(11) NOT NULL DEFAULT 0 COMMENT 'Time',
  `publicKey` varchar(2048) NOT NULL,
  `keySalt` varchar(256) NOT NULL COMMENT 'Salt for encrypting privateKey',
  `privateKey` varchar(4096) NOT NULL COMMENT 'Encrypted',
  `recoveryKeyHash` varchar(128) NOT NULL,
  `backupKey` varchar(4096) NOT NULL COMMENT 'Encrypted w/ Recovery Key',
  PRIMARY KEY (`userKeyId`),
  KEY `userId` (`userId`),
  KEY `created` (`created`),
  KEY `replaced` (`replaced`),
  KEY `recoveryKeyHash` (`recoveryKeyHash`),
  CONSTRAINT `UserKeys_ibfk_1` FOREIGN KEY (`userId`) REFERENCES `Users` (`userId`) ON DELETE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
CREATE TABLE `Users` (
  `userId` bigint(20) NOT NULL AUTO_INCREMENT,
  `userName` varchar(64) NOT NULL,
  `name` varchar(128) NOT NULL,
  `email` varchar(254) NOT NULL,
  `bio` varchar(255) NOT NULL,
  `type` tinyint(4) NOT NULL DEFAULT 0,
  `permissions` bigint(20) NOT NULL DEFAULT 0,
  `created` int(11) NOT NULL COMMENT 'Time',
  PRIMARY KEY (`userId`),
  UNIQUE KEY `uniqueName` (`userName`),
  UNIQUE KEY `email` (`email`),
  KEY `created` (`created`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
CREATE TABLE `UserTokens` (
  `userTokenId` bigint(20) NOT NULL AUTO_INCREMENT,
  `userId` bigint(20) NOT NULL,
  `tokenType` tinyint(4) NOT NULL,
  `selector` varchar(12) NOT NULL,
  `token` varchar(128) NOT NULL COMMENT 'Hashed',
  `created` int(11) NOT NULL COMMENT 'Time',
  `expiration` int(11) NOT NULL COMMENT 'Time',
  PRIMARY KEY (`userTokenId`),
  KEY `userId` (`userId`),
  KEY `selector` (`selector`),
  KEY `expiration` (`expiration`),
  CONSTRAINT `UserTokens_ibfk_1` FOREIGN KEY (`userId`) REFERENCES `Users` (`userId`) ON DELETE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;