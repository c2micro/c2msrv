// Code generated by ent, DO NOT EDIT.

package operator

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the operator type in the database.
	Label = "operator"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldToken holds the string denoting the token field in the database.
	FieldToken = "token"
	// FieldColor holds the string denoting the color field in the database.
	FieldColor = "color"
	// FieldLast holds the string denoting the last field in the database.
	FieldLast = "last"
	// EdgeChat holds the string denoting the chat edge name in mutations.
	EdgeChat = "chat"
	// EdgeGroup holds the string denoting the group edge name in mutations.
	EdgeGroup = "group"
	// Table holds the table name of the operator in the database.
	Table = "operator"
	// ChatTable is the table that holds the chat relation/edge.
	ChatTable = "chat"
	// ChatInverseTable is the table name for the Chat entity.
	// It exists in this package in order to avoid circular dependency with the "chat" package.
	ChatInverseTable = "chat"
	// ChatColumn is the table column denoting the chat relation/edge.
	ChatColumn = "from"
	// GroupTable is the table that holds the group relation/edge.
	GroupTable = "group"
	// GroupInverseTable is the table name for the Group entity.
	// It exists in this package in order to avoid circular dependency with the "group" package.
	GroupInverseTable = "group"
	// GroupColumn is the table column denoting the group relation/edge.
	GroupColumn = "author"
)

// Columns holds all SQL columns for operator fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldUsername,
	FieldToken,
	FieldColor,
	FieldLast,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/c2micro/c2msrv/internal/ent/runtime"
var (
	Hooks        [1]ent.Hook
	Interceptors [1]ent.Interceptor
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	UsernameValidator func(string) error
	// TokenValidator is a validator for the "token" field. It is called by the builders before save.
	TokenValidator func(string) error
	// DefaultColor holds the default value on creation for the "color" field.
	DefaultColor uint32
	// DefaultLast holds the default value on creation for the "last" field.
	DefaultLast func() time.Time
)

// OrderOption defines the ordering options for the Operator queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// ByUsername orders the results by the username field.
func ByUsername(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUsername, opts...).ToFunc()
}

// ByToken orders the results by the token field.
func ByToken(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldToken, opts...).ToFunc()
}

// ByColor orders the results by the color field.
func ByColor(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldColor, opts...).ToFunc()
}

// ByLast orders the results by the last field.
func ByLast(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLast, opts...).ToFunc()
}

// ByChatCount orders the results by chat count.
func ByChatCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newChatStep(), opts...)
	}
}

// ByChat orders the results by chat terms.
func ByChat(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newChatStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByGroupCount orders the results by group count.
func ByGroupCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newGroupStep(), opts...)
	}
}

// ByGroup orders the results by group terms.
func ByGroup(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newGroupStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newChatStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ChatInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ChatTable, ChatColumn),
	)
}
func newGroupStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(GroupInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, GroupTable, GroupColumn),
	)
}