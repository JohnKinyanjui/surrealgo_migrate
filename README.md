# surrealgo

**surrealgo_migrate** is a tool designed to streamline migrations for SurrealDB databases, simplifying the management of structural changes and keeping your workflow organized.

## Installation
```bash
go install github.com/JohnKinyanjui/surrealgo_migrate@latest
```

Key Commands

1. Init: Initializes a project to use surrealgo_migrate.This create a config file which you can add your database url, credentials and also the folder stucture:

```bash
surrealgo_migrate init
```

Example of config (surrealgo.yaml)
```yaml
endpoint: ws://localhost:8000/rpc

database:
    name: root
    namespace: root
    password: root
    user: root

folders:
    events: internal/database/events
    migrations: internal/database/migrations

```

idk am feeling this config is abit confusing i hope i can make it easy for you guys if you have a better format let me know issues.

2. New Migration:  Creates new up and down migration files and make sure you always the name of your migration

```Bash
surrealgo_migrate migrate new add_users
```

3. Up: Applies pending migrations in order based on their timestamps. Ensures database updates progress systematically.

```Bash
surrealgo_migrate migrate up
```

4. Down: Reverts the most recent migration. This helps with undoing changes or testing in development.

```Bash
surrealgo_migrate migrate down
```

### Important: Transactional Migrations

Each migration (both up and down) is executed within a SurrealDB transaction. This provides the following benefits:

1. Atomicity: If any part of a migration fails, the entire migration is rolled back, preventing your database from being left in an inconsistent state.
2. Data Integrity: Transactions help ensure your database remains valid and consistent throughout the migration process.
