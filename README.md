# TaskManagementApp

## Getting started

### Prerequisites

- install [Go](https://go.dev/) v1.21
- install [postgreSQL](https://www.postgresql.org/)
- (optional) install [Make](https://www.gnu.org/software/make/)

### Config

1. copy `config-example.ini` and rename it to `config.ini`.
2. adjust each values in config respectively, such as App's Port, DB's Host, Port and Database name.

### Migrations

Run all sql in `migration` directory.

### Running the App

```bash
$ make run
# OR
$ go run main.go
```
