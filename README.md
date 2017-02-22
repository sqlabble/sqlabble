# sqlabble [ ![Codeship Status for minodisk/sqlabble](https://img.shields.io/codeship/f3642650-d5ab-0134-3d76-0246ca48a45f/master.svg)](https://app.codeship.com/projects/202522) [![Go Report Card](https://goreportcard.com/badge/github.com/minodisk/sqlabble)](https://goreportcard.com/report/github.com/minodisk/sqlabble) [![Coverage Status](https://coveralls.io/repos/github/minodisk/sqlabble/badge.svg?branch=master)](https://coveralls.io/github/minodisk/sqlabble?branch=master) [![GoDoc](https://godoc.org/github.com/minodisk/sqlabble?status.png)](https://godoc.org/github.com/minodisk/sqlabble) [![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)

SQL query builder with type support.

## Features

- Quote the alias according to the specification of each SQL server.
  - Standard: `"`
  - MySQL: `` ` ``

## Supports

### Clauses

- [x] `SELECT {COLUMN|SUBQUERY}`
- [x] `SELECT DISTINCT {COLUMN|SUBQUERY}`
- [x] `FROM {TABLE|SUBQUERY}`
- [x] `WHERE {OPERATION}`
- [x] `GROUP BY {COLUMN}`
- [x] `HAVING`
- [x] `ORDER BY {ORDER}`
- [x] `LIMIT {COUNT}`
- [x] `OFFSET {COUNT}`
- [x] `INSERT INTO {TABLE} ({COLUMN})`
- [x] `VALUES ({VALUE})`
- [x] `DEFAULT VALUES`
- [x] `UPDATE {TABLE}`
- [x] `SET ({COLUMN} = {VALUE|SUBQUERY})`
- [x] `DELETE`
- [x] `CREATE TABLE {TABLE}`
- [x] `CREATE TABLE IF NOT EXISTS {TABLE}`

### Aliases

- [x] `{TABLE} AS {ALIAS}`
- [x] `{COLUMN} AS {ALIAS}`

### Assignment

- [x] `{COLUMN} = {VALUE}`

### Operators

#### Logical Operators

- [x] `{OPERATION} AND {OPERATION}`
- [x] `{OPERATION} OR {OPERATION}`
- [x] `NOT ({OPERATION})`

#### Comparison Operators

- [x] `{COLUMN|SUBQUERY} = {VALUE|SUBQUERY}`
- [x] `{COLUMN|SUBQUERY} != {VALUE|SUBQUERY}`
- [x] `{COLUMN|SUBQUERY} > {VALUE|SUBQUERY}`
- [x] `{COLUMN|SUBQUERY} >= {VALUE|SUBQUERY}`
- [x] `{COLUMN|SUBQUERY} < {VALUE|SUBQUERY}`
- [x] `{COLUMN|SUBQUERY} <= {VALUE|SUBQUERY}`
- [x] `{COLUMN|SUBQUERY} LIKE {VALUE|SUBQUERY}`
- [x] `{COLUMN|SUBQUERY} REGEXP {VALUE|SUBQUERY}`
- [x] `{COLUMN|SUBQUERY} BETWEEN {VALUE|SUBQUERY} AND {VALUE|SUBQUERY}`
- [x] `{COLUMN|SUBQUERY} IN {VALUE|SUBQUERY}`
- [x] `{COLUMN|SUBQUERY} NOT IN {VALUE|SUBQUERY}`
- [x] `{COLUMN|SUBQUERY} IS NULL`
- [x] `{COLUMN|SUBQUERY} IS NOT NULL`

### Joins

- [x] `JOIN {TABLE|SUBQUERY}`
- [x] `INNER JOIN {TABLE|SUBQUERY}`
- [x] `LEFT JOIN {TABLE|SUBQUERY}`
- [x] `RIGHT JOIN {TABLE|SUBQUERY}`

#### Join Conditions

- [x] `ON {COLUMN} = {COLUMN}`
- [x] `USING {COLUMN}`

### Orders

- [x] `{COLUMN} ASC`
- [x] `{COLUMN} DESC`

### Sets

- [x] `({STATEMENT}) UNION ({STATEMENT})`
- [x] `({STATEMENT}) UNION ALL ({STATEMENT})`
- [x] `({STATEMENT}) INTERSECT ({STATEMENT})`
- [x] `({STATEMENT}) INTERSECT ALL ({STATEMENT})`
- [x] `({STATEMENT}) EXCEPT ({STATEMENT})`
- [x] `({STATEMENT}) EXCEPT ALL ({STATEMENT})`

### Table Definitions

- [x] `({COLUMN} {DEFINITION})`

### Conditional Logics

- [ ] `CASE`
- [ ] `WHEN`
- [ ] `ELSE`
- [ ] `THEN`
- [ ] `END`

#### Nonscalar Operation

- [ ] `{COLUMN} = ALL ({SUBQUERY})`
- [ ] `{COLUMN} != ALL ({SUBQUERY})`
- [ ] `{COLUMN} > ALL ({SUBQUERY})`
- [ ] `{COLUMN} >= ALL ({SUBQUERY})`
- [ ] `{COLUMN} < ALL ({SUBQUERY})`
- [ ] `{COLUMN} <= ALL ({SUBQUERY})`
- [ ] `{COLUMN} = ANY ({SUBQUERY})`
- [ ] `{COLUMN} != ANY ({SUBQUERY})`
- [ ] `{COLUMN} > ANY ({SUBQUERY})`
- [ ] `{COLUMN} >= ANY ({SUBQUERY})`
- [ ] `{COLUMN} < ANY ({SUBQUERY})`
- [ ] `{COLUMN} <= ANY ({SUBQUERY})`
- [ ] `EXISTS ({SUBQUERY})`
- [ ] `NOT EXISTS ({SUBQUERY})`
