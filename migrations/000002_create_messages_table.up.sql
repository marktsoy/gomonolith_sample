CREATE TABLE messages(
    id SERIAL PRIMARY KEY,
    content TEXT,
    status int DEFAULT 0 ,
    priority int DEFAULT 0,
    bundle_id INT NOT NULL,
    
    FOREIGN KEY (bundle_id)
        REFERENCES bundles(id)
)