// Code generated by ent, DO NOT EDIT.

package task

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/c2micro/c2mshr/defaults"
)

const (
	// Label holds the string label denoting the task type in the database.
	Label = "task"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldGid holds the string denoting the gid field in the database.
	FieldGid = "gid"
	// FieldBid holds the string denoting the bid field in the database.
	FieldBid = "bid"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldPushedAt holds the string denoting the pushed_at field in the database.
	FieldPushedAt = "pushed_at"
	// FieldDoneAt holds the string denoting the done_at field in the database.
	FieldDoneAt = "done_at"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldCap holds the string denoting the cap field in the database.
	FieldCap = "cap"
	// FieldArgs holds the string denoting the args field in the database.
	FieldArgs = "args"
	// FieldOutput holds the string denoting the output field in the database.
	FieldOutput = "output"
	// FieldOutputBig holds the string denoting the output_big field in the database.
	FieldOutputBig = "output_big"
	// EdgeGroup holds the string denoting the group edge name in mutations.
	EdgeGroup = "group"
	// EdgeBeacon holds the string denoting the beacon edge name in mutations.
	EdgeBeacon = "beacon"
	// EdgeBlobberArgs holds the string denoting the blobber_args edge name in mutations.
	EdgeBlobberArgs = "blobber_args"
	// EdgeBlobberOutput holds the string denoting the blobber_output edge name in mutations.
	EdgeBlobberOutput = "blobber_output"
	// Table holds the table name of the task in the database.
	Table = "task"
	// GroupTable is the table that holds the group relation/edge.
	GroupTable = "task"
	// GroupInverseTable is the table name for the Group entity.
	// It exists in this package in order to avoid circular dependency with the "group" package.
	GroupInverseTable = "group"
	// GroupColumn is the table column denoting the group relation/edge.
	GroupColumn = "gid"
	// BeaconTable is the table that holds the beacon relation/edge.
	BeaconTable = "task"
	// BeaconInverseTable is the table name for the Beacon entity.
	// It exists in this package in order to avoid circular dependency with the "beacon" package.
	BeaconInverseTable = "beacon"
	// BeaconColumn is the table column denoting the beacon relation/edge.
	BeaconColumn = "bid"
	// BlobberArgsTable is the table that holds the blobber_args relation/edge.
	BlobberArgsTable = "task"
	// BlobberArgsInverseTable is the table name for the Blobber entity.
	// It exists in this package in order to avoid circular dependency with the "blobber" package.
	BlobberArgsInverseTable = "blobber"
	// BlobberArgsColumn is the table column denoting the blobber_args relation/edge.
	BlobberArgsColumn = "args"
	// BlobberOutputTable is the table that holds the blobber_output relation/edge.
	BlobberOutputTable = "task"
	// BlobberOutputInverseTable is the table name for the Blobber entity.
	// It exists in this package in order to avoid circular dependency with the "blobber" package.
	BlobberOutputInverseTable = "blobber"
	// BlobberOutputColumn is the table column denoting the blobber_output relation/edge.
	BlobberOutputColumn = "output"
)

// Columns holds all SQL columns for task fields.
var Columns = []string{
	FieldID,
	FieldGid,
	FieldBid,
	FieldCreatedAt,
	FieldPushedAt,
	FieldDoneAt,
	FieldStatus,
	FieldCap,
	FieldArgs,
	FieldOutput,
	FieldOutputBig,
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

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s defaults.TaskStatus) error {
	switch s.String() {
	case "new", "in-progress", "cancelled", "success", "error":
		return nil
	default:
		return fmt.Errorf("task: invalid enum value for status field: %q", s)
	}
}

// CapValidator is a validator for the "cap" field enum values. It is called by the builders before save.
func CapValidator(c defaults.Capability) error {
	switch c.String() {
	case "c_sleep", "c_ls", "c_pwd", "c_cd", "c_whoami", "c_ps", "c_cat", "c_exec", "c_cp", "c_jobs", "c_jobkill", "c_kill", "c_mv", "c_mkdir", "c_rm", "c_exec_assembly", "c_shell", "c_ppid", "c_exec_detach", "c_shellcode_injection", "c_download", "c_upload", "c_pause", "c_destruct", "c_exit", "c_socks5":
		return nil
	default:
		return fmt.Errorf("task: invalid enum value for cap field: %q", c)
	}
}

// OrderOption defines the ordering options for the Task queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByGid orders the results by the gid field.
func ByGid(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGid, opts...).ToFunc()
}

// ByBid orders the results by the bid field.
func ByBid(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBid, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByPushedAt orders the results by the pushed_at field.
func ByPushedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPushedAt, opts...).ToFunc()
}

// ByDoneAt orders the results by the done_at field.
func ByDoneAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDoneAt, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByCap orders the results by the cap field.
func ByCap(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCap, opts...).ToFunc()
}

// ByArgs orders the results by the args field.
func ByArgs(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldArgs, opts...).ToFunc()
}

// ByOutput orders the results by the output field.
func ByOutput(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOutput, opts...).ToFunc()
}

// ByOutputBig orders the results by the output_big field.
func ByOutputBig(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOutputBig, opts...).ToFunc()
}

// ByGroupField orders the results by group field.
func ByGroupField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newGroupStep(), sql.OrderByField(field, opts...))
	}
}

// ByBeaconField orders the results by beacon field.
func ByBeaconField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBeaconStep(), sql.OrderByField(field, opts...))
	}
}

// ByBlobberArgsField orders the results by blobber_args field.
func ByBlobberArgsField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBlobberArgsStep(), sql.OrderByField(field, opts...))
	}
}

// ByBlobberOutputField orders the results by blobber_output field.
func ByBlobberOutputField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBlobberOutputStep(), sql.OrderByField(field, opts...))
	}
}
func newGroupStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(GroupInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, GroupTable, GroupColumn),
	)
}
func newBeaconStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BeaconInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, BeaconTable, BeaconColumn),
	)
}
func newBlobberArgsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BlobberArgsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, BlobberArgsTable, BlobberArgsColumn),
	)
}
func newBlobberOutputStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BlobberOutputInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, BlobberOutputTable, BlobberOutputColumn),
	)
}
