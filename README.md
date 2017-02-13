# sqlabble

Type supported SQL query builder.

## Supports

### Clauses

- [x] `SELECT {COLUMN}`
- [x] `SELECT DISTINCT {COLUMN}`
- [x] `FROM {TABLE}`
- [x] `WHERE {OPERATION}`
- [x] `GROUP BY {COLUMN}`
- [x] `HAVING`
- [x] `ORDER BY {ORDER}`
- [x] `LIMIT`
- [x] `INSERT INTO {TABLE} ({COLUMN})`
- [x] `VALUES ({VALUE})`
- [x] `UPDATE {TABLE}`
- [x] `SET ({ASSIGNMENT})`
- [x] `DELETE`
- [x] `CREATE TABLE {TABLE}`

### Aliases

- [x] `{TABLE} AS {ALIAS}`
- [x] `{COLUMN} AS {ALIAS}`

### Assignment

- [x] `{COLUMN} = {VALUE}`

### Logical Operators

- [x] `{OPERATION} AND {OPERATION}`
- [x] `{OPERATION} OR {OPERATION}`
- [x] `NOT ({OPERATION})`

### Comparison Operators

- [x] `{COLUMN} = {VALUE}`
- [x] `{COLUMN} != {VALUE}`
- [x] `{COLUMN} > {VALUE}`
- [x] `{COLUMN} >= {VALUE}`
- [x] `{COLUMN} < {VALUE}`
- [x] `{COLUMN} <= {VALUE}`
- [x] `{COLUMN} BETWEEN {VALUE}`
- [x] `{COLUMN} IN {VALUE}`
- [x] `{COLUMN} NOT IN {VALUE}`
- [x] `{COLUMN} LIKE {VALUE}`
- [x] `{COLUMN} BETWEEN {VALUE}`
- [x] `{COLUMN} IS NULL {VALUE}`
- [x] `{COLUMN} IS NOT NULL {VALUE}`

### Joins

- [x] `{TABLE} JOIN {TABLE}`
- [x] `{TABLE} INNER JOIN {TABLE}`
- [x] `{TABLE} LEFT JOIN {TABLE}`
- [x] `{TABLE} RIGHT JOIN {TABLE}`

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

### Subqueries

N/A
