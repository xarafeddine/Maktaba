-- GRANT CREATE ON DATABASE maktaba TO maktaba;
CREATE EXTENSION IF NOT EXISTS citext;
ALTER DATABASE maktaba OWNER TO maktaba;
CREATE TABLE IF NOT EXISTS books (
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    title text NOT NULL,
    year integer NOT NULL,
    page_count integer NOT NULL,
    genres text [] NOT NULL,
    version integer NOT NULL DEFAULT 1
);