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
- [x] `LIMIT {COUNT}`
- [x] `OFFSET {COUNT}`
- [x] `INSERT INTO {TABLE} ({COLUMN})`
- [x] `VALUES ({VALUE})`
- [x] `DEFAULT VALUES`
- [x] `UPDATE {TABLE}`
- [x] `SET ({ASSIGNMENT})`
- [x] `DELETE`
- [x] `CREATE TABLE {TABLE}`
- [x] `CREATE TABLE IF NOT EXISTS {TABLE}`

#### With Sub Queries

- [ ] `SELECT ({SUBQUERY})`
- [ ] `FROM ({SUBQUERY})`

### Aliases

- [x] `{TABLE} AS {ALIAS}`
- [x] `{COLUMN} AS {ALIAS}`

### Assignment

- [x] `{COLUMN} = {VALUE}`

### Conditional Logics

- [ ] `CASE`
- [ ] `WHEN`
- [ ] `ELSE`
- [ ] `THEN`
- [ ] `END`

### Operators

#### Logical Operators

- [x] `{OPERATION} AND {OPERATION}`
- [x] `{OPERATION} OR {OPERATION}`
- [x] `NOT ({OPERATION})`

#### Comparison Operators

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
- [x] `{COLUMN} REGEXP {VALUE}`
- [x] `{COLUMN} IS NULL {VALUE}`
- [x] `{COLUMN} IS NOT NULL {VALUE}`

##### With Sub Queries

- [ ] `{COLUMN} = ({SUBQUERY})`
- [ ] `{COLUMN} != ({SUBQUERY})`
- [ ] `{COLUMN} > ({SUBQUERY})`
- [ ] `{COLUMN} >= ({SUBQUERY})`
- [ ] `{COLUMN} < ({SUBQUERY})`
- [ ] `{COLUMN} <= ({SUBQUERY})`
- [ ] `{COLUMN} IN ({SUBQUERY})`
- [ ] `{COLUMN} NOT IN ({SUBQUERY})`
- [ ] `({SUBQUERY}) = {VALUE}`
- [ ] `({SUBQUERY}) != {VALUE}`
- [ ] `({SUBQUERY}) > {VALUE}`
- [ ] `({SUBQUERY}) >= {VALUE}`
- [ ] `({SUBQUERY}) < {VALUE}`
- [ ] `({SUBQUERY}) <= {VALUE}`
- [ ] `({SUBQUERY}) BETWEEN {VALUE}`
- [ ] `({SUBQUERY}) IN {VALUE}`
- [ ] `({SUBQUERY}) NOT IN {VALUE}`
- [ ] `({SUBQUERY}) LIKE {VALUE}`
- [ ] `({SUBQUERY}) REGEXP {VALUE}`
- [ ] `({SUBQUERY}) IS NULL {VALUE}`
- [ ] `({SUBQUERY}) IS NOT NULL {VALUE}`

### Joins

- [x] `{TABLE} JOIN {TABLE}`
- [x] `{TABLE} INNER JOIN {TABLE}`
- [x] `{TABLE} LEFT JOIN {TABLE}`
- [x] `{TABLE} RIGHT JOIN {TABLE}`

#### Join Conditions

- [x] `ON {COLUMN} = {COLUMN}`
- [x] `USING {COLUMN}`

#### With Sub Queries

- [ ] `{TABLE} JOIN ({SUBQUERY})`
- [ ] `{TABLE} INNER JOIN ({SUBQUERY})`
- [ ] `{TABLE} LEFT JOIN ({SUBQUERY})`
- [ ] `{TABLE} RIGHT JOIN ({SUBQUERY})`

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
