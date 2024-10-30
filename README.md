# `liquigen`

A lightweight `Liquibase` file generator `cmd` tool implemented in `Golang`.

## 1.`Usage`

### 1.1.`Install`

```shell
$ go install github.com/photowey/liquigen@latest
```

### 1.2.`Home`

```shell
# Home: ~/.liquibase
# Config file: ~/.liquibase/liquibase.json
```

## 2.`Commands`

- `usage`
- `version`
- `config`
- `changelog`

### 2.1.`Mode`

#### 2.1.1.`SQL file`

```shell
# Example:
$ liquigen[.exe] changelog -a changjun -D mysql -s ./testdata/sql/mysql/company.sql
```

#### 2.1.2.`Database`

- `Unsupported now`
