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

### Joins

- [x] `JOIN {TABLE}`
- [x] `INNER JOIN {TABLE}`
- [x] `LEFT JOIN {TABLE}`
- [x] `RIGHT JOIN {TABLE}`

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

### Subqueries

- [ ] `SELECT ({SUBQUERY})`
- [ ] `FROM ({SUBQUERY})`

#### Scalar Operation

- [x] `{COLUMN} = ({SUBQUERY})`
- [x] `{COLUMN} != ({SUBQUERY})`
- [x] `{COLUMN} > ({SUBQUERY})`
- [x] `{COLUMN} >= ({SUBQUERY})`
- [x] `{COLUMN} < ({SUBQUERY})`
- [x] `{COLUMN} <= ({SUBQUERY})`
- [x] `{COLUMN} LIKE ({SUBQUERY})`
- [x] `{COLUMN} REGEXP ({SUBQUERY})`

- [ ] `({SUBQUERY}) = {VALUE}`
- [ ] `({SUBQUERY}) != {VALUE}`
- [ ] `({SUBQUERY}) > {VALUE}`
- [ ] `({SUBQUERY}) >= {VALUE}`
- [ ] `({SUBQUERY}) < {VALUE}`
- [ ] `({SUBQUERY}) <= {VALUE}`

#### Nonscalar Operation

- [ ] `{COLUMN} IN ({SUBQUERY})`
- [ ] `{COLUMN} NOT IN ({SUBQUERY})`
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

- [ ] `({SUBQUERY}) LIKE {VALUE}`
- [ ] `({SUBQUERY}) REGEXP {VALUE}`
- [ ] `({SUBQUERY}) IN {VALUE}`
- [ ] `({SUBQUERY}) NOT IN {VALUE}`
- [ ] `({SUBQUERY}) BETWEEN {VALUE}`
- [ ] `({SUBQUERY}) IS NULL {VALUE}`
- [ ] `({SUBQUERY}) IS NOT NULL {VALUE}`

- [ ] `EXISTS ({SUBQUERY})`
- [ ] `NOT EXISTS ({SUBQUERY})`

#### Set

- [ ] `SET {COLUMN} = ({SUBQUERY})`

#### Joins

- [ ] `JOIN ({SUBQUERY})`
- [ ] `INNER JOIN ({SUBQUERY})`
- [ ] `LEFT JOIN ({SUBQUERY})`
- [ ] `RIGHT JOIN ({SUBQUERY})`
