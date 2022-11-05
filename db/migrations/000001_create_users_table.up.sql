CREATE TABLE IF NOT EXISTS users(
    id serial PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(300) NOT NULL,
    email VARCHAR(50) UNIQUE NOT NULL,
    age INT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create function to update timestamp on field updated_at
CREATE  FUNCTION update_updated_at_users()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ language 'plpgsql';

-- Create trigger to trigger the function
CREATE TRIGGER update_users_updated_at
    BEFORE UPDATE
    ON
        users
    FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_users();

INSERT INTO users(username, password, email, age) VALUES
('rifqoi', '$2y$10$54/1U4uOkA3pDQF/PDLnAuTWK7J4/ItaOoYWdGKKdfbVw2QFpyFxC', 'rifqoi@gmail.com', 20),
('romiari', '$2y$10$54/1U4uOkA3pDQF/PDLnAuTWK7J4/ItaOoYWdGKKdfbVw2QFpyFxC', 'romi@gmail.com', 21),
('anugaming', '$2y$10$54/1U4uOkA3pDQF/PDLnAuTWK7J4/ItaOoYWdGKKdfbVw2QFpyFxC', 'anu@gmail.com', 22);
