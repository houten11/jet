// Modeling of tables.  This is where query preparation starts

package sqlbuilder

import (
	"errors"
)

type readableTable interface {
	// Generates a select query on the current tableName.
	SELECT(projections ...projection) SelectStatement

	// Creates a inner join tableName Expression using onCondition.
	INNER_JOIN(table ReadableTable, onCondition BoolExpression) ReadableTable

	// Creates a left join tableName Expression using onCondition.
	LEFT_JOIN(table ReadableTable, onCondition BoolExpression) ReadableTable

	// Creates a right join tableName Expression using onCondition.
	RIGHT_JOIN(table ReadableTable, onCondition BoolExpression) ReadableTable

	FULL_JOIN(table ReadableTable, onCondition BoolExpression) ReadableTable

	CROSS_JOIN(table ReadableTable) ReadableTable
}

// The sql tableName write interface.
type writableTable interface {
	INSERT(columns ...Column) InsertStatement
	UPDATE(columns ...Column) UpdateStatement
	DELETE() DeleteStatement

	LOCK() LockStatement
}

type ReadableTable interface {
	readableTable
	clause
}

type WritableTable interface {
	writableTable
	clause
}

type Table interface {
	readableTable
	writableTable
	clause
	SchemaName() string
	TableName() string
	AS(alias string)
}

type readableTableInterfaceImpl struct {
	parent ReadableTable
}

// Generates a select query on the current tableName.
func (r *readableTableInterfaceImpl) SELECT(projections ...projection) SelectStatement {
	return newSelectStatement(r.parent, projections)
}

// Creates a inner join tableName Expression using onCondition.
func (r *readableTableInterfaceImpl) INNER_JOIN(table ReadableTable, onCondition BoolExpression) ReadableTable {
	return newJoinTable(r.parent, table, innerJoin, onCondition)
}

// Creates a left join tableName Expression using onCondition.
func (r *readableTableInterfaceImpl) LEFT_JOIN(table ReadableTable, onCondition BoolExpression) ReadableTable {
	return newJoinTable(r.parent, table, leftJoin, onCondition)
}

// Creates a right join tableName Expression using onCondition.
func (r *readableTableInterfaceImpl) RIGHT_JOIN(table ReadableTable, onCondition BoolExpression) ReadableTable {
	return newJoinTable(r.parent, table, rightJoin, onCondition)
}

func (r *readableTableInterfaceImpl) FULL_JOIN(table ReadableTable, onCondition BoolExpression) ReadableTable {
	return newJoinTable(r.parent, table, fullJoin, onCondition)
}

func (r *readableTableInterfaceImpl) CROSS_JOIN(table ReadableTable) ReadableTable {
	return newJoinTable(r.parent, table, crossJoin, nil)
}

type writableTableInterfaceImpl struct {
	parent WritableTable
}

func (w *writableTableInterfaceImpl) INSERT(columns ...Column) InsertStatement {
	return newInsertStatement(w.parent, columns...)
}

func (w *writableTableInterfaceImpl) UPDATE(columns ...Column) UpdateStatement {
	return newUpdateStatement(w.parent, columns)
}

func (w *writableTableInterfaceImpl) DELETE() DeleteStatement {
	return newDeleteStatement(w.parent)
}

func (w *writableTableInterfaceImpl) LOCK() LockStatement {
	return LOCK(w.parent)
}

func NewTable(schemaName, name string, columns ...Column) Table {

	t := &tableImpl{
		schemaName: schemaName,
		name:       name,
		columns:    columns,
	}
	for _, c := range columns {
		c.setTableName(name)
	}

	t.readableTableInterfaceImpl.parent = t
	t.writableTableInterfaceImpl.parent = t

	return t
}

type tableImpl struct {
	readableTableInterfaceImpl
	writableTableInterfaceImpl

	schemaName string
	name       string
	alias      string
	columns    []Column
}

func (t *tableImpl) AS(alias string) {
	t.alias = alias

	for _, c := range t.columns {
		c.setTableName(alias)
	}
}

// Returns the tableName's name in the database
func (t *tableImpl) SchemaName() string {
	return t.schemaName
}

// Returns the tableName's name in the database
func (t *tableImpl) TableName() string {
	return t.name
}

func (t *tableImpl) SchemaTableName() string {
	return t.schemaName
}

func (t *tableImpl) serialize(statement statementType, out *queryData, options ...serializeOption) error {
	if t == nil {
		return errors.New("tableImpl is nil. ")
	}

	out.writeString(t.schemaName)
	out.writeString(".")
	out.writeString(t.TableName())

	if len(t.alias) > 0 {
		out.writeString(" AS ")
		out.writeString(t.alias)
	}

	return nil
}

type joinType int

const (
	innerJoin joinType = iota
	leftJoin
	rightJoin
	fullJoin
	crossJoin
)

// Join expressions are pseudo readable tables.
type joinTable struct {
	readableTableInterfaceImpl

	lhs         ReadableTable
	rhs         ReadableTable
	join_type   joinType
	onCondition BoolExpression
}

func newJoinTable(
	lhs ReadableTable,
	rhs ReadableTable,
	join_type joinType,
	onCondition BoolExpression) ReadableTable {

	joinTable := &joinTable{
		lhs:         lhs,
		rhs:         rhs,
		join_type:   join_type,
		onCondition: onCondition,
	}

	joinTable.readableTableInterfaceImpl.parent = joinTable

	return joinTable
}

func (t *joinTable) SchemaName() string {
	return ""
}

func (t *joinTable) TableName() string {
	return ""
}

func (t *joinTable) serialize(statement statementType, out *queryData, options ...serializeOption) (err error) {
	if t == nil {
		return errors.New("Join table is nil. ")
	}

	if isNil(t.lhs) {
		return errors.New("left hand side of join operation is nil table")
	}

	if err = t.lhs.serialize(statement, out); err != nil {
		return
	}

	out.nextLine()

	switch t.join_type {
	case innerJoin:
		out.writeString("INNER JOIN")
	case leftJoin:
		out.writeString("LEFT JOIN")
	case rightJoin:
		out.writeString("RIGHT JOIN")
	case fullJoin:
		out.writeString("FULL JOIN")
	case crossJoin:
		out.writeString("CROSS JOIN")
	}

	if isNil(t.rhs) {
		return errors.New("right hand side of join operation is nil table")
	}

	if err = t.rhs.serialize(statement, out); err != nil {
		return
	}

	if t.onCondition == nil && t.join_type != crossJoin {
		return errors.New("join condition is nil")
	}

	if t.onCondition != nil {
		out.writeString("ON")
		if err = t.onCondition.serialize(statement, out); err != nil {
			return
		}
	}

	return nil
}
