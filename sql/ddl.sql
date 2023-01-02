

CREATE TABLE `breed_info` (
  `BreedID` int NOT NULL AUTO_INCREMENT,
  `BreedName` varchar(128) NOT NULL,
  `GoodWithKids` int NOT NULL,
  `GoodWithDogs` int NOT NULL,
  `Shedding` int NOT NULL,
  `Grooming` int NOT NULL,
  `Energy` int NOT NULL,
  `BreedImg` varchar(256) NOT NULL,
  PRIMARY KEY (`BreedID`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

CREATE TABLE `dogs` (
  `KennelID` int NOT NULL,
  `BreedID` int NOT NULL,
  `DogID` int NOT NULL AUTO_INCREMENT,
  `DogName` varchar(52) NOT NULL,
  `Sex` varchar(6) NOT NULL,
  PRIMARY KEY (`DogID`),
  KEY `BreedID` (`BreedID`),
  KEY `KennelID` (`KennelID`),
  CONSTRAINT `dogs_ibfk_1` FOREIGN KEY (`BreedID`) REFERENCES `breed_info` (`BreedID`),
  CONSTRAINT `dogs_ibfk_2` FOREIGN KEY (`KennelID`) REFERENCES `kennels` (`KennelID`)
) ENGINE=InnoDB AUTO_INCREMENT=53 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

CREATE TABLE `kennel_addr` (
  `ID_Kennel` int DEFAULT NULL,
  `Numero` varchar(12) NOT NULL,
  `Rua` varchar(128) NOT NULL,
  `Bairro` varchar(28) NOT NULL,
  `CEP` varchar(9) NOT NULL,
  `Cidade` varchar(28) NOT NULL,
  KEY `ID_Kennel` (`ID_Kennel`),
  CONSTRAINT `kennel_addr_ibfk_1` FOREIGN KEY (`ID_Kennel`) REFERENCES `kennels` (`KennelID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

CREATE TABLE `kennels` (
  `KennelID` int NOT NULL AUTO_INCREMENT,
  `KennelName` varchar(128) NOT NULL,
  `ContactNumber` varchar(20) NOT NULL,
  PRIMARY KEY (`KennelID`)
) ENGINE=InnoDB AUTO_INCREMENT=45 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

CREATE TABLE `user_dog_prefs` (
  `UserID` int NOT NULL,
  `GoodWithKids` int NOT NULL,
  `GoodWithDogs` int NOT NULL,
  `Shedding` int NOT NULL,
  `Grooming` int NOT NULL,
  `Energy` int NOT NULL,
  KEY `UserID` (`UserID`),
  CONSTRAINT `user_dog_prefs_ibfk_1` FOREIGN KEY (`UserID`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `nome` varchar(128) NOT NULL,
  `email` varchar(128) NOT NULL,
  `passwd` varchar(128) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=33 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci