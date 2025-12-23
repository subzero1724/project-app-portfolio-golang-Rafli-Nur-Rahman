# Database Migrations

This folder contains database migration files for the portfolio application.

## Files

- `init.sql` - Initial database setup with all tables, indexes, and sample data

## Tables Created

1. **users** - User authentication and profile information
2. **projects** - Portfolio projects showcase
3. **contacts** - Contact form submissions
4. **skills** - Technical skills with proficiency levels
5. **experiences** - Work experience timeline

## Running Migrations

To run the migration:

```bash
psql -U your_username -d your_database -f migrations/init.sql
```

Or using the Go application (if migration runner is implemented):

```bash
go run cmd/server/main.go migrate
```

## Sample Data

The migration includes sample data for testing:
- 1 admin user (email: admin@portfolio.com, password: admin123)
- 3 sample projects
- 6 sample skills
- 2 sample work experiences

**Note:** Remember to change the default admin password in production!
