DROP TABLE IF EXISTS projects;

CREATE TABLE IF NOT EXISTS projects (
    id INTEGER  PRIMARY KEY AUTOINCREMENT,
    project_name VARCHAR(255) NOT NULL,
    category VARCHAR(100),
    project_type VARCHAR(100),
    release_year INTEGER,
    age_category VARCHAR(50),
    duration TIME,
    director VARCHAR(255),
    producer VARCHAR(255)
);

INSERT INTO projects (project_name, category, project_type, release_year, age_category, duration, director, producer)
VALUES 
    ('The Shawshank Redemption', 'Drama', 'Movie', 1994, 'Adult', '02:22:00', 'Frank Darabont', 'Niki Marvin'),
    ('Inception', 'Science Fiction', 'Movie', 2010, 'Adult', '02:28:00', 'Christopher Nolan', 'Emma Thomas'),
    ('Breaking Bad', 'Drama', 'TV Show', 2008, 'Adult', '00:49:00', 'Vince Gilligan', 'Vince Gilligan');
