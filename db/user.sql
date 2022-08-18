
CREATE TABLE users
(
    Id SERIAL PRIMARY KEY,
    Username  CHARACTER VARYING(30),
    Email     CHARACTER VARYING(30),
    Dob       DATE,
    Age       INT,
    Number    CHARACTER VARYING(10)
);

INSERT INTO users (Username, Email, Dob, Age, Number)
VALUES
('Test1', '1@email.com', '1999-01-04', age('1999-01-04'), '111111111'),
('Test2', '2@email.com', '2000-02-05', age('2000-02-05'), '222222222'),
('Test3', '3@email.com', '2001-03-06', age('2001-03-06'), '333333333')