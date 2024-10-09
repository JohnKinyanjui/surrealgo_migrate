# SurrealGo Migrate

SurrealGo Migrate is a powerful tool designed to streamline database migrations for SurrealDB, simplifying the management of structural changes and keeping your workflow organized. It now supports SurrealDB v2.0.1.

## Features

- Create new migration files with timestamps
- Execute "up" and "down" migrations
- Support for surrealql migrations
- Transactional migrations for data integrity
- SSL support for secure connections

## Installation

Install SurrealGo Migrate using the following command:

```bash
go install github.com/JohnKinyanjui/surrealgo_migrate@latest
```

## Usage

### Creating a New Migration

To create a new migration, use the following command:

```bash
surrealgo_migrate --path [MIGRATION_PATH] migrate new [FILE_NAME]
```

Example:

```bash
surrealgo_migrate migrate --path database/migrations new add_users
```

This command generates new "up" and "down" migration files. Always provide a descriptive name for your migration.

**Note:** The `event` and `migrate` commands serve different purposes:

- `event` generates files for writing clean event functions
- `migrate` is used to define tables, fields, and other database structures

### Applying Migrations

To apply pending migrations, use the "up" command:

```bash
surrealgo_migrate migrate up --path [MIGRATION_PATH] --host [HOST] --db [DATABASE] --ns [NAMESPACE] --user [USERNAME] --pass [PASSWORD]
```

Example:

```bash
surrealgo_migrate migrate up --path database/migrations --host localhost:8000 --db root --ns root --user root --pass root --path /internal/database/migrations
```

For secure connections, add the `--ssl true` flag:

```bash
surrealgo_migrate migrate up --path database/migrations --host localhost:8000 --db root --ns root --user root --pass root --ssl true
```

### Rolling Back Migrations

To roll back migrations, use the "down" command (syntax similar to "up" command):

```bash
surrealgo_migrate migrate down --host [HOST] --db [DATABASE] --ns [NAMESPACE] --user [USERNAME] --pass [PASSWORD] --path [MIGRATION_PATH]
```

## Transactional Migrations

SurrealGo Migrate executes each migration (both up and down) within a SurrealDB transaction. This approach offers several benefits:

1. **Atomicity**: If any part of a migration fails, the entire migration is rolled back. This prevents your database from being left in an inconsistent state.
2. **Data Integrity**: Transactions help ensure your database remains valid and consistent throughout the migration process.
3. **Reliability**: In case of unexpected errors or interruptions, your database won't be partially migrated.

## Contributing

Contributions to SurrealGo Migrate are welcome! Please feel free to submit pull requests, create issues, or suggest improvements.

## Support

For questions, issues, or feature requests, please [open an issue](https://github.com/JohnKinyanjui/surrealgo_migrate/issues) on our GitHub repository.
