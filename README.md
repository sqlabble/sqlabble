# sqlabble

Type supported SQL query builder.

## Supports

### Clauses

- [] `CREATE TABLE {TABLE}`
- [] `INSERT INTO {TABLE}`
- [x] `SELECT {COLUMN}`
- [] `SELECT DISTINCT {COLUMN}`
- [x] `FROM {TABLE}`
- [x] `WHERE {OPERATION}`
- [] `GROUP BY {COLUMN}`
- [] `HAVING`
- [x] `ORDER BY {ORDER}`
- [] `LIMIT`
- [] `UPDATE {TABLE}`
- [] `SET {OPERATION}`
- [] `DELETE`

### Aliases

- [x] `{TABLE} AS {ALIAS}`
- [x] `{COLUMN} AS {ALIAS}`

### Joins

- [x] `{TABLE} JOIN {TABLE}`
- [x] `{TABLE} INNER JOIN {TABLE}`
- [x] `{TABLE} LEFT JOIN {TABLE}`
- [x] `{TABLE} RIGHT JOIN {TABLE}`

### Join Conditions

- [x] `ON {COLUMN} = {COLUMN}`
- [] `USING {COLUMN}`

### Logical Operators

- [x] `{OPERATION} AND {OPERATION}`
- [x] `{OPERATION} OR {OPERATION}`
- [x] `NOT {OPERATION}`

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

### Orders

- [x] `{COLUMN} ASC`
- [x] `{COLUMN} DESC`

### Sets

- [] `UNION`
- [] `UNION ALL`
- [] `INTERSECT`
- [] `INTERSECT ALL`
- [] `EXCEPT`
- [] `EXCEPT ALL`

### Subqueries
