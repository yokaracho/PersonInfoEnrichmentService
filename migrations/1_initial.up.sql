CREATE TABLE people (
                        id SERIAL PRIMARY KEY,
                        name VARCHAR(255) NOT NULL,
                        surname VARCHAR(255) NOT NULL,
                        patronymic VARCHAR(255) NOT NULL,
                        age INT NOT NULL,
                        gender VARCHAR(10) NOT NULL,
                        nationality VARCHAR(255) NOT NULL
);