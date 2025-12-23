CREATE TABLE IF NOT EXISTS projects (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    description TEXT NOT NULL,
    image_url VARCHAR(500),
    demo_url VARCHAR(500),
    github_url VARCHAR(500),
    tags TEXT[],
    featured BOOLEAN DEFAULT false,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create index for featured projects
CREATE INDEX IF NOT EXISTS idx_projects_featured ON projects(featured);

-- Insert sample projects
INSERT INTO projects (title, description, image_url, demo_url, github_url, tags, featured) VALUES
    ('E-Commerce Platform', 'A full-stack e-commerce platform with payment integration and admin dashboard', '/placeholder.svg?height=300&width=400', 'https://demo.example.com', 'https://github.com/yourusername/ecommerce', ARRAY['React', 'Node.js', 'PostgreSQL', 'Stripe'], true),
    ('Task Management App', 'Real-time collaborative task management application with team features', '/placeholder.svg?height=300&width=400', 'https://demo.example.com', 'https://github.com/yourusername/taskapp', ARRAY['Next.js', 'TypeScript', 'Prisma', 'WebSocket'], true),
    ('Portfolio CMS', 'Content management system for portfolio websites with dynamic templates', '/placeholder.svg?height=300&width=400', 'https://demo.example.com', 'https://github.com/yourusername/portfolio-cms', ARRAY['Go', 'PostgreSQL', 'HTML/CSS'], false)
ON CONFLICT DO NOTHING;
