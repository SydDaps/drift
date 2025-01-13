# Drift

Drift is a simple yet effective database migration tool written in Go. It allows you to easily manage your database schema changes, ensuring smooth migrations and rollbacks.

## 📁 Configuration

To get started, you need to create a configuration file to store your database credentials.
	1.	Create a directory named config in your application’s root.
	2.	Inside config, create a file named database.yml with the following structure:
```yaml
host: 
port: 
user: 
password: 
dbname: 
```
## 🚀 Installation

Add Drift to your project using go get:
```go
go get github.com/SydDaps/drift@v1.0.0
```

## 🛠️ Creating a Migration File

Generate a new migration file using the Create method:
```go
drift.Create("migration file name")
```
This command will create a migration file containing up and down sections for you to define your migration scripts.
Example Migration File:
```yaml
version: 20250108013715
name: Create User Table
description: Users table
up: |
  CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
  );
down: |
  DROP TABLE users;
```
- up: SQL commands to apply the migration (e.g., creating tables or columns).
- down: SQL commands to reverse the migration (e.g., dropping tables or columns).

## 🔼 Running Migrations

Run all pending migrations using:
```go
drift.MigrateUp()
```
This will apply all up scripts in your migration files.

## 🔽 Rolling Back Migrations

Rollback migrations using the MigrateDown method:
```go
drift.MigrateDown(step int)
```
- step: Number of migrations to roll back. For example, step = 1 rolls back the most recent migration, while step = 3 rolls back the last three migrations.

## 📝 Example Workflow

1. Create a Migration:
```go
drift.Create("create users table")
```

2. Define Migration Scripts:

Edit the generated file to add up and down SQL scripts.
```yaml
up: |
  CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
  );
down: |
  DROP TABLE users;
```
3. Apply Migrations:

Run all pending migrations:
```go
drift.MigrateUp()
```

4. Rollback if Needed:

Undo the last migration:
```go
drift.MigrateDown(1)
```

🛡️ Best Practices
	•	Use descriptive names for your migration files (e.g., add users table).
	•	Test migrations in a staging environment before running them in production.
	•	Always keep backups of your database before applying significant changes.

💬 Support

For questions or issues, please reach out via GitHub Issues.

⚖️ License

Drift is open-source software licensed under the MIT License.

Happy Migrating! 🚀
