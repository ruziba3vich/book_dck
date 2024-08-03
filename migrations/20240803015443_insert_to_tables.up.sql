
CREATE TABLE books (
    book_id UUID PRIMARY KEY,
    author VARCHAR(255) NOT NULL,
    title VARCHAR(255) NOT NULL,
    published_year INT NOT NULL
);

