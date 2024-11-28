ALTER TABLE books
ADD CONSTRAINT books_page_count_check CHECK (page_count >= 0);
ALTER TABLE books
ADD CONSTRAINT books_year_check CHECK (
        year BETWEEN 0 AND date_part('year', now())
    );
ALTER TABLE books
ADD CONSTRAINT genres_length_check CHECK (
        array_length(genres, 1) BETWEEN 1 AND 5
    );