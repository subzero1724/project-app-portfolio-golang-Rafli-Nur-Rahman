CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    bio TEXT,
    title VARCHAR(100),
    avatar VARCHAR(500),
    github VARCHAR(255),
    linkedin VARCHAR(255),
    twitter VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Insert sample user
INSERT INTO users (name, email, bio, title, github, linkedin, twitter) 
VALUES (
    'Your Name',
    'your.email@example.com',
    'I am a passionate full-stack developer with expertise in building scalable web applications using modern technologies.',
    'Full Stack Developer',
    'https://github.com/yourusername',
    'https://linkedin.com/in/yourusername',
    'https://twitter.com/yourusername'
) ON CONFLICT (email) DO NOTHING;
