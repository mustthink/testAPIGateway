
CREATE TABLE users
(
    Id SERIAL PRIMARY KEY,
    Username  CHARACTER VARYING(70),
    Email     CHARACTER VARYING(30),
    Dob       DATE,
    Age       INT,
    Number    CHARACTER VARYING(10)
);

INSERT INTO users (Username, Email, Dob, Age, Number)
VALUES
('d890bb8d82a2f0b3176f976916225786ecb6dd55237d0f1f9ca827cdf9aa6ff7', '1@email.com', '1999-01-04', 23, '111111111'),
('3fd8852a281e05215f8aae8d119fd341a7320bd3008500f391b34ad91745e9fb', '2@email.com', '2000-02-05', 22, '222222222'),
('27621deea941b70f6c4ff0446b3740f209881d2111825a5e9be2a09607b15acd', '3@email.com', '2001-03-06', 21, '333333333')