// Code generated by entc, DO NOT EDIT.

package user

import (
	"time"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldToken holds the string denoting the token field in the database.
	FieldToken = "token"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updatedat field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeApps holds the string denoting the apps edge name in mutations.
	EdgeApps = "apps"
	// EdgeVersions holds the string denoting the versions edge name in mutations.
	EdgeVersions = "versions"
	// Table holds the table name of the user in the database.
	Table = "users"
	// AppsTable is the table that holds the apps relation/edge. The primary key declared below.
	AppsTable = "user_apps"
	// AppsInverseTable is the table name for the App entity.
	// It exists in this package in order to avoid circular dependency with the "app" package.
	AppsInverseTable = "apps"
	// VersionsTable is the table that holds the versions relation/edge. The primary key declared below.
	VersionsTable = "user_versions"
	// VersionsInverseTable is the table name for the Version entity.
	// It exists in this package in order to avoid circular dependency with the "version" package.
	VersionsInverseTable = "versions"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldEmail,
	FieldPassword,
	FieldToken,
	FieldCreatedAt,
	FieldUpdatedAt,
}

var (
	// AppsPrimaryKey and AppsColumn2 are the table columns denoting the
	// primary key for the apps relation (M2M).
	AppsPrimaryKey = []string{"user_id", "app_id"}
	// VersionsPrimaryKey and VersionsColumn2 are the table columns denoting the
	// primary key for the versions relation (M2M).
	VersionsPrimaryKey = []string{"user_id", "version_id"}
)

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
	// DefaultCreatedAt holds the default value on creation for the "createdAt" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updatedAt" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updatedAt" field.
	UpdateDefaultUpdatedAt func() time.Time
)