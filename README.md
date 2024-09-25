# lf

## Setup

- Install [Taskfile](https://taskfile.dev/installation/)
- Install [PostgreSQL](https://www.postgresql.org/download/)
- run `task install`

## Development

### Create local dev database

```sql
CREATE USER root WITH SUPERUSER PASSWORD 'adminadmin';
CREATE DATABASE go_template WITH OWNER = root;
```

### Run dev

```sh
task dev
```

## Migrations

### Create migration script

```sh
task migration-create --<migration-name> #task migration-create -- init
```

## Build

```sh
task build 
```