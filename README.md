# sqlabble

[![Codeship Status for sqlabble/sqlabble](https://img.shields.io/codeship/f3642650-d5ab-0134-3d76-0246ca48a45f/master.svg?style=flat)](https://app.codeship.com/projects/202522)
[![Build Status](https://travis-ci.org/sqlabble/sqlabble.svg?branch=master)](https://travis-ci.org/sqlabble/sqlabble)
[![Go Report Card](https://goreportcard.com/badge/github.com/sqlabble/sqlabble)](https://goreportcard.com/report/github.com/sqlabble/sqlabble)
[![codecov](https://codecov.io/gh/sqlabble/sqlabble/branch/master/graph/badge.svg)](https://codecov.io/gh/sqlabble/sqlabble)
[![GoDoc](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat)](https://godoc.org/github.com/sqlabble/sqlabble)
[![MIT License](https://img.shields.io/badge/License-MIT-blue.svg?style=flat)](LICENSE)



SQL query builder with type support.

## Features

- Type support:
  - Restrict method chain order.
  - Restrict expression that can be specified by  the interface type.
- Flexible formatter:
  - Prefix and Indent.
  - Quote the alias according to the specification of each SQL server.
    - Standard: `"`
    - MySQL: `` ` ``

## Installation

```
go get -u github.com/sqlabble/sqlabble/...
```

## Usage

### Intro

```go
import (
	"fmt"

	q "github.com/sqlabble/sqlabble"
	"github.com/sqlabble/sqlabble/builder"
)

func main() {
	stmt := q.Select(
		q.Column("person_id"),
		q.Column("fname"),
		q.Column("lname"),
		q.Column("birth_date"),
	).From(
		q.Table("person"),
	).Where(
		q.Column("lname").Eq(q.Param("Turner")),
	)

	query, values := builder.StandardIndented.Build(stmt)

	fmt.Println(query)
	// -> SELECT
	//      person_id
	//      , fname
	//      , lname
	//      , birth_date
	//    FROM
	//      person
	//    WHERE
	//      lname = ?

	fmt.Println(values)
	// -> [Turner]
}
```

If it is slightly redundant, there are short hands.

```go
q.Select(
  q.C("person_id"),
  q.C("fname"),
  q.C("lname"),
  q.C("birth_date"),
).From(
  q.T("person"),
).Where(
  q.C("lname").Eq(q.P("Turner")),
)
```

If you do not want to write table names or column names many times with strings, try [the code generation tool](#code-generation-tool).

### Insert

### Select

```go
q.Select(
  q.C("person_id").As("persion_id"),
  q.C("fname").As("persion_fname"),
  q.C("lname").As("persion_lname"),
  q.C("birth_date").As("persion_birth_date"),
).From(
  q.T("user"),
).Where(
  q.C("id").Eq(q.Param(3)),
),
```

### Update

### Delete

### Sets

### Subqueries

## Code Generation Tool

If you write table names and column names many times with strings, you will mistype someday. It would be nonsense to spend time finding mistypes. There is a code generation tool that implements a method that returns a table or column to a struct. Is declarative coding is fun, right?

First, create a file named `tables.go`:

```go
package tables

// +db:"persons"
type Person struct {
	PersonID             int
	FamilyName           string `db:"fname"`
	LastName             string `db:"lname"`
	BirthDate            time.Time
	SocialSecurityNumber string `db:"-"`
	password             string
}
```

And, call the following command at the terminal:

```sh
sqlabble tables.go
```

Then, a file named `tables_sqlabble.go` will be generated:

```go
N/A
```

Finally, you will be able to construct queries using the added methods:

```go
p := Person{}
q.Select(
    p.Columns()...,
  ).From(
    p.Table(),
  ).Where(
    p.ColumnLastName().Eq(q.P("Turner")),
  )
```

It's simple, and you never mistype table names or column names.

## Processing Layers

```
                                   Format
                                     |
         Nodeizer   Tokenizer    Generator
            |           |            |
Statement --+-> Nodes --+-> Tokens --+-> Query
                        |
                        +--------------> Values
```

## Supports

### Clauses

- [x] `CREATE TABLE {TABLE}`
- [x] `CREATE TABLE IF NOT EXISTS {TABLE}`
- [x] `SELECT {COLUMN|FUNCTION|SUBQUERY}`
- [x] `SELECT DISTINCT {COLUMN|FUNCTION|SUBQUERY}`
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
- [x] `SET ({ASSIGNMENT})`
- [x] `DELETE`

### Column Definition

- [x] `({COLUMN} {DEFINITION})`

### Joins

- [x] `JOIN {TABLE|SUBQUERY}`
- [x] `INNER JOIN {TABLE|SUBQUERY}`
- [x] `LEFT JOIN {TABLE|SUBQUERY}`
- [x] `RIGHT JOIN {TABLE|SUBQUERY}`

#### Conditions

- [x] `ON {COLUMN} = {COLUMN}`
- [x] `USING {COLUMN}`

### Orders

- [x] `{COLUMN} ASC`
- [x] `{COLUMN} DESC`

### Aliases

- [x] `{TABLE} AS {ALIAS}`
- [x] `{COLUMN} AS {ALIAS}`

### Assignment

- [x] `{COLUMN} = {VALUE|FUNCTION|SUBQUERY}`

### Sets

- [x] `({STATEMENT}) UNION ({STATEMENT})`
- [x] `({STATEMENT}) UNION ALL ({STATEMENT})`
- [x] `({STATEMENT}) INTERSECT ({STATEMENT})`
- [x] `({STATEMENT}) INTERSECT ALL ({STATEMENT})`
- [x] `({STATEMENT}) EXCEPT ({STATEMENT})`
- [x] `({STATEMENT}) EXCEPT ALL ({STATEMENT})`

### Conditional Logics

- [x] `CASE {VALUE|COLUMN|FUNCTION|SUBQUERY} WHEN {VALUE} THEN {VALUE|COLUMN|FUNCTION|SUBQUERY} ELSE {VALUE|COLUMN|FUNCTION|SUBQUERY} END`
- [x] `CASE WHEN {OPERATION} THEN {VALUE|COLUMN|FUNCTION|SUBQUERY} ELSE {VALUE|COLUMN|FUNCTION|SUBQUERY} END`

### Operators

#### Logical

- [x] `{OPERATION} AND {OPERATION}`
- [x] `{OPERATION} OR {OPERATION}`
- [x] `NOT ({OPERATION})`

#### Comparison

##### Scalar

- [x] `{VALUE|COLUMN|FUNCTION|SUBQUERY} = {VALUE|COLUMN|FUNCTION|SUBQUERY}`
- [x] `{VALUE|COLUMN|FUNCTION|SUBQUERY} != {VALUE|COLUMN|FUNCTION|SUBQUERY}`
- [x] `{VALUE|COLUMN|FUNCTION|SUBQUERY} > {VALUE|COLUMN|FUNCTION|SUBQUERY}`
- [x] `{VALUE|COLUMN|FUNCTION|SUBQUERY} >= {VALUE|COLUMN|FUNCTION|SUBQUERY}`
- [x] `{VALUE|COLUMN|FUNCTION|SUBQUERY} < {VALUE|COLUMN|FUNCTION|SUBQUERY}`
- [x] `{VALUE|COLUMN|FUNCTION|SUBQUERY} <= {VALUE|COLUMN|FUNCTION|SUBQUERY}`
- [x] `{VALUE|COLUMN|FUNCTION|SUBQUERY} LIKE {VALUE|COLUMN|FUNCTION|SUBQUERY}`
- [x] `{VALUE|COLUMN|FUNCTION|SUBQUERY} REGEXP {VALUE|COLUMN|FUNCTION|SUBQUERY}`
- [x] `{VALUE|COLUMN|FUNCTION|SUBQUERY} BETWEEN {VALUE|COLUMN|FUNCTION|SUBQUERY} AND {VALUE|COLUMN|FUNCTION|SUBQUERY}`
- [x] `{VALUE|COLUMN|FUNCTION|SUBQUERY} IN {VALUES|SUBQUERY}`
- [x] `{VALUE|COLUMN|FUNCTION|SUBQUERY} NOT IN {VALUES|SUBQUERY}`
- [x] `{COLUMN|SUBQUERY} IS NULL`
- [x] `{COLUMN|SUBQUERY} IS NOT NULL`

##### Nonscalar

- [x] `{VALUE|COLUMN|FUNCTION|SUBQUERY} = ALL {SUBQUERY}`
- [x] `{VALUE|COLUMN|FUNCTION|SUBQUERY} != ALL {SUBQUERY}`
- [x] `{VALUE|COLUMN|FUNCTION|SUBQUERY} > ALL {SUBQUERY}`
- [x] `{VALUE|COLUMN|FUNCTION|SUBQUERY} >= ALL {SUBQUERY}`
- [x] `{VALUE|COLUMN|FUNCTION|SUBQUERY} < ALL {SUBQUERY}`
- [x] `{VALUE|COLUMN|FUNCTION|SUBQUERY} <= ALL {SUBQUERY}`
- [x] `{VALUE|COLUMN|FUNCTION|SUBQUERY} = ANY {SUBQUERY}`
- [x] `{VALUE|COLUMN|FUNCTION|SUBQUERY} != ANY {SUBQUERY}`
- [x] `{VALUE|COLUMN|FUNCTION|SUBQUERY} > ANY {SUBQUERY}`
- [x] `{VALUE|COLUMN|FUNCTION|SUBQUERY} >= ANY {SUBQUERY}`
- [x] `{VALUE|COLUMN|FUNCTION|SUBQUERY} < ANY {SUBQUERY}`
- [x] `{VALUE|COLUMN|FUNCTION|SUBQUERY} <= ANY {SUBQUERY}`
- [x] `EXISTS {SUBQUERY}`
- [x] `NOT EXISTS {SUBQUERY}`

### Functions

#### Control Flow

N/A

#### String

N/A

#### Numeric

N/A

#### Date and Time

- [x] `ADDDATE({VALUE|COLUMN|FUNCTION|SUBQUERY})`
- [x] `ADDTIME({VALUE|COLUMN|FUNCTION|SUBQUERY})`
- [x] `CONVERT_TZ({VALUE|COLUMN|FUNCTION|SUBQUERY})`
- [x] `CURDATE({VALUE|COLUMN|FUNCTION|SUBQUERY})`
- [x] `CURRENT_DATE({VALUE|COLUMN|FUNCTION|SUBQUERY})`
- [x] `CURRENT_TIME({VALUE|COLUMN|FUNCTION|SUBQUERY})`
- [x] `CURRENT_TIMESTAMP({VALUE|COLUMN|FUNCTION|SUBQUERY})`
- [x] `CURTIME({VALUE|COLUMN|FUNCTION|SUBQUERY})`
- [x] `DATE({VALUE|COLUMN|FUNCTION|SUBQUERY})`
- [x] `ATE_AD({VALUE|COLUMN|FUNCTION|SUBQUERY})`
- [x] `DATE_FORMAT({VALUE|COLUMN|FUNCTION|SUBQUERY})`
- [x] `DATE_SUB({VALUE|COLUMN|FUNCTION|SUBQUERY})`
- [x] `DATEDIFF({VALUE|COLUMN|FUNCTION|SUBQUERY})`
- [x] `DAY({VALUE|COLUMN|FUNCTION|SUBQUERY})`
- [x] `DAYNAME({VALUE|COLUMN|FUNCTION|SUBQUERY})`
- [x] `DAYOFMONTH({VALUE|COLUMN|FUNCTION|SUBQUERY})`
- [x] `DAYOFWEEK({VALUE|COLUMN|FUNCTION|SUBQUERY})`
- [x] `DAYOFYEAR({VALUE|COLUMN|FUNCTION|SUBQUERY})`
- [x] `EXTRACT({VALUE|COLUMN|FUNCTION|SUBQUERY})`
- [x] `FROM_DAYS({VALUE|COLUMN|FUNCTION|SUBQUERY})`
- [x] `FROM_UNIXTIME({VALUE|COLUMN|FUNCTION|SUBQUERY})`
- [x] `GET_FORMAT({VALUE|COLUMN|FUNCTION|SUBQUERY})`
- [x] `HOUR({VALUE|COLUMN|FUNCTION|SUBQUERY})`
- [x] `LAST_DAY({VALUE|COLUMN|FUNCTION|SUBQUERY})`
- [x] `LOCALTIME({VALUE|COLUMN|FUNCTION|SUBQUERY})`
- [x] `LOCALTIMESTAMP({VALUE|COLUMN|FUNCTION|SUBQUERY})`
- [x] `MAKEDATE({VALUE|COLUMN|FUNCTION|SUBQUERY})`
- [x] `MAKETIME({VALUE|COLUMN|FUNCTION|SUBQUERY})`
- [x] `MICROSECOND({VALUE|COLUMN|FUNCTION|SUBQUERY})`
- [x] `MINUTE({VALUE|COLUMN|FUNCTION|SUBQUERY})`
- [x] `MONTH({VALUE|COLUMN|FUNCTION|SUBQUERY})`
- [x] `MONTHNAME({VALUE|COLUMN|FUNCTION|SUBQUERY})`
- [x] `NOW({VALUE|COLUMN|FUNCTION|SUBQUERY})`
- [x] `PERIOD_ADD({VALUE|COLUMN|FUNCTION|SUBQUERY})`
- [x] `PERIOD_DIFF({VALUE|COLUMN|FUNCTION|SUBQUERY})`
- [x] `QUARTER({VALUE|COLUMN|FUNCTION|SUBQUERY})`
- [x] `SEC_TO_TIME({VALUE|COLUMN|FUNCTION|SUBQUERY})`
- [x] `SECOND({VALUE|COLUMN|FUNCTION|SUBQUERY})`
- [x] `STR_TO_DATE({VALUE|COLUMN|FUNCTION|SUBQUERY})`
- [x] `SUBDATE({VALUE|COLUMN|FUNCTION|SUBQUERY})`
- [x] `SUBTIME({VALUE|COLUMN|FUNCTION|SUBQUERY})`
- [x] `SYSDATE({VALUE|COLUMN|FUNCTION|SUBQUERY})`
- [x] `TIME({VALUE|COLUMN|FUNCTION|SUBQUERY})`
- [x] `IME_FORMA({VALUE|COLUMN|FUNCTION|SUBQUERY})`
- [x] `TIME_TO_SEC({VALUE|COLUMN|FUNCTION|SUBQUERY})`
- [x] `TIMEDIFF({VALUE|COLUMN|FUNCTION|SUBQUERY})`
- [x] `TIMESTAMP({VALUE|COLUMN|FUNCTION|SUBQUERY})`
- [x] `TIMESTAMPADD({VALUE|COLUMN|FUNCTION|SUBQUERY})`
- [x] `TIMESTAMPDIFF({VALUE|COLUMN|FUNCTION|SUBQUERY})`
- [x] `TO_DAYS({VALUE|COLUMN|FUNCTION|SUBQUERY})`
- [x] `TO_SECONDS({VALUE|COLUMN|FUNCTION|SUBQUERY})`
- [x] `UNIX_TIMESTAMP({VALUE|COLUMN|FUNCTION|SUBQUERY})`
- [x] `UTC_DATE({VALUE|COLUMN|FUNCTION|SUBQUERY})`
- [x] `UTC_TIME({VALUE|COLUMN|FUNCTION|SUBQUERY})`
- [x] `UTC_TIMESTAMP({VALUE|COLUMN|FUNCTION|SUBQUERY})`
- [x] `WEEK({VALUE|COLUMN|FUNCTION|SUBQUERY})`
- [x] `WEEKDAY({VALUE|COLUMN|FUNCTION|SUBQUERY})`
- [x] `WEEKOFYEAR({VALUE|COLUMN|FUNCTION|SUBQUERY})`
- [x] `YEAR({VALUE|COLUMN|FUNCTION|SUBQUERY})`
- [x] `YEARWEEK({VALUE|COLUMN|FUNCTION|SUBQUERY})`
