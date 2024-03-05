CREATE TABLE teachers (
    id SERIAL PRIMARY KEY,
    fullName VARCHAR(50) NOT NULL,
    email VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(50) NOT NULL,
    phoneNumber VARCHAR(12) UNIQUE
);

CREATE TABLE groups (
    id SERIAL PRIMARY KEY,
    groupName VARCHAR(50) NOT NULL UNIQUE
);

CREATE TABLE students (
    id SERIAL PRIMARY KEY,
    fullName VARCHAR(50) NOT NULL,
    email VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(50) NOT NULL,
    phoneNumber VARCHAR(12) UNIQUE,
    groupID INT,
    FOREIGN KEY (groupID) REFERENCES groups(id)
);

CREATE TABLE teacherGroup (
    teacherID INT REFERENCES teachers(id),
    groupID INT REFERENCES groups(id),
    PRIMARY KEY (teacherID, groupID)
);




