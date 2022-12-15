CREATE TABLE Address (
	  Addr_ID INT NOT NULL AUTO_INCREMENT,
    ID_Kennel INT,
    ID_User INT,
    Numero VARCHAR(12) NOT NULL,
	  Rua VARCHAR(128) NOT NULL,
    Bairro VARCHAR(28) NOT NULL,
    CEP VARCHAR(9) NOT NULL,
    Cidade VARCHAR(28) NOT NULL,
    PRIMARY KEY (Addr_ID),
    FOREIGN KEY (ID_Kennel) REFERENCES kennels(KennelID),
    FOREIGN KEY (ID_User) REFERENCES users(id)
)

INSERT INTO Address
VALUES 
(
	2,
    "123",
    "Rua 3",
    "Graças",
    "52050-200",
    "Recife"
)

CREATE TABLE dogs (
    KennelID INT NOT NULL,
    BreedID INT NOT NULL,
    DogID INT NOT NULL AUTO_INCREMENT,
    DogName VARCHAR(52) NOT NULL,
    Sex VARCHAR(6) NOT NULL,
    PRIMARY KEY (DogID)
    FOREIGN KEY (BreedID) REFERENCES breed_info(BreedID)
    FOREIGN KEY (KennelID) REFERENCES kennels(KennelID)
)

INSERT INTO dogs
VALUES 
(
	2,
    1,
    1,
    "Banzé",
    "Male",
)

CREATE TABLE users (
  `id` int NOT NULL AUTO_INCREMENT,
  `nome` varchar(128) NOT NULL,
  `email` varchar(128) NOT NULL,
  `passwd` varchar(128) NOT NULL,
  PRIMARY KEY (`id`)
)

CREATE TABLE `user_dog_prefs` (
  `UserID` int NOT NULL,
  `GoodWithKids` int NOT NULL,
  `GoodWithDogs` int NOT NULL,
  `Shedding` int NOT NULL,
  `Grooming` int NOT NULL,
  `Energy` int NOT NULL,
  KEY `UserID` (`UserID`),
  CONSTRAINT `user_dog_prefs_ibfk_1` FOREIGN KEY (`UserID`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


CREATE TABLE `kennels` (
  `KennelID` int NOT NULL AUTO_INCREMENT,
  `KennelName` varchar(128) NOT NULL,
  `ContactNumber` varchar(20) NOT NULL,
  PRIMARY KEY (`KennelID`)
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

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
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


CREATE TABLE `kennel_addr` (
  `ID_Kennel` int DEFAULT NULL,
  `Numero` varchar(12) NOT NULL,
  `Rua` varchar(128) NOT NULL,
  `Bairro` varchar(28) NOT NULL,
  `CEP` varchar(9) NOT NULL,
  `Cidade` varchar(28) NOT NULL,
  KEY `ID_Kennel` (`ID_Kennel`),
  CONSTRAINT `kennel_addr_ibfk_1` FOREIGN KEY (`ID_Kennel`) REFERENCES `kennels` (`KennelID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;