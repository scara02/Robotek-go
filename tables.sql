CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    fullName VARCHAR(50) NOT NULL,
    email VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(50) NOT NULL,
    phoneNumber VARCHAR(12) UNIQUE,
    role VARCHAR(8),
    CHECK (role in ('admin', 'teacher', 'student'))
);

CREATE TABLE groups (
    id SERIAL PRIMARY KEY,
    groupName VARCHAR(50) NOT NULL UNIQUE
);

CREATE TABLE student_group (
    studentID int,
    groupID INT,
    PRIMARY KEY (studentID),
    FOREIGN KEY (studentID) REFERENCES users(id),
    FOREIGN KEY (groupID) REFERENCES groups(id)
);

CREATE TABLE teacher_groups (
    teacherID INT REFERENCES users(id),
    groupID INT REFERENCES groups(id),
    PRIMARY KEY (teacherID, groupID)
);
