CREATE TABLE books (
                             id bigint NOT NULL,
                             title varchar(64) DEFAULT NULL,
                             price DECIMAL(10,2) DEFAULT(0),
                             author varchar(64) DEFAULT NULL,
                             PRIMARY KEY (id)
)