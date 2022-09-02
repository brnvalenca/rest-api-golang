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
