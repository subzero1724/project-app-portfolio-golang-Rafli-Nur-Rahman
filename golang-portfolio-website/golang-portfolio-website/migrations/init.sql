-- Create extensions
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create users table
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    bio TEXT,
    avatar_url VARCHAR(500),
    github_url VARCHAR(255),
    linkedin_url VARCHAR(255),
    twitter_url VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create projects table
CREATE TABLE IF NOT EXISTS projects (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    image_url VARCHAR(500),
    tech_stack TEXT[] NOT NULL,
    github_url VARCHAR(255),
    demo_url VARCHAR(255),
    is_featured BOOLEAN DEFAULT false,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create contacts table
CREATE TABLE IF NOT EXISTS contacts (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    subject VARCHAR(255) NOT NULL,
    message TEXT NOT NULL,
    is_read BOOLEAN DEFAULT false,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create skills table
CREATE TABLE IF NOT EXISTS skills (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    level INTEGER NOT NULL CHECK (level >= 0 AND level <= 100),
    category VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create experiences table
CREATE TABLE IF NOT EXISTS experiences (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    company VARCHAR(255) NOT NULL,
    position VARCHAR(255) NOT NULL,
    description TEXT,
    start_date DATE NOT NULL,
    end_date DATE,
    is_current BOOLEAN DEFAULT false,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes for better query performance
CREATE INDEX IF NOT EXISTS idx_projects_featured ON projects(is_featured);
CREATE INDEX IF NOT EXISTS idx_contacts_read ON contacts(is_read);
CREATE INDEX IF NOT EXISTS idx_contacts_created ON contacts(created_at DESC);
CREATE INDEX IF NOT EXISTS idx_skills_category ON skills(category);
CREATE INDEX IF NOT EXISTS idx_experiences_dates ON experiences(start_date DESC, end_date DESC);

-- Insert default admin user (password: admin123)
INSERT INTO users (name, email, password, bio) 
VALUES (
    'Admin User',
    'admin@portfolio.com',
    '$2a$10$rZJ8X.YHvHqVJK4xF5K9WOqKvGZGJ9mH7K2xJ5K9WOqKvGZGJ9mH7',
    'Full Stack Developer'
) ON CONFLICT (email) DO NOTHING;

-- Insert sample projects
INSERT INTO projects (title, description, tech_stack, github_url, demo_url, is_featured) 
VALUES 
(
    'E-Commerce Platform',
    'A full-featured e-commerce platform with shopping cart, payment integration, and admin dashboard.',
    ARRAY['Go', 'PostgreSQL', 'React', 'Tailwind CSS'],
    'https://github.com/example/ecommerce',
    'https://demo.example.com',
    true
),
(
    'Task Management API',
    'RESTful API for task management with authentication, authorization, and real-time updates.',
    ARRAY['Go', 'Chi Router', 'PostgreSQL', 'JWT'],
    'https://github.com/example/task-api',
    'https://api.example.com',
    true
),
(
    'Weather Dashboard',
    'Real-time weather dashboard with forecasts, maps, and historical data visualization.',
    ARRAY['Go', 'HTML/CSS', 'Chart.js', 'OpenWeather API'],
    'https://github.com/example/weather',
    'https://weather.example.com',
    false
) ON CONFLICT DO NOTHING;

-- Insert sample skills
INSERT INTO skills (name, level, category) 
VALUES 
('Go', 90, 'Backend'),
('PostgreSQL', 85, 'Database'),
('JavaScript', 80, 'Frontend'),
('React', 75, 'Frontend'),
('Docker', 70, 'DevOps'),
('Git', 85, 'Tools') ON CONFLICT DO NOTHING;

-- Insert sample experience
INSERT INTO experiences (company, position, description, start_date, is_current) 
VALUES 
(
    'Tech Company Inc',
    'Senior Backend Developer',
    'Leading the development of microservices architecture using Go and PostgreSQL. Managing a team of 5 developers.',
    '2022-01-01',
    true
),
(
    'Startup XYZ',
    'Full Stack Developer',
    'Built and maintained web applications using various technologies. Worked on both frontend and backend features.',
    '2020-06-01',
    false
) ON CONFLICT DO NOTHING;
