// Code generated by entc, DO NOT EDIT.

package instance

import (
	"time"

	"github.com/marcustut/fyp/backend/ent/schema/ulid"
)

const (
	// Label holds the string label denoting the instance type in the database.
	Label = "instance"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldInstanceID holds the string denoting the instance_id field in the database.
	FieldInstanceID = "instance_id"
	// FieldInstanceType holds the string denoting the instance_type field in the database.
	FieldInstanceType = "instance_type"
	// FieldPrivateDNSName holds the string denoting the private_dns_name field in the database.
	FieldPrivateDNSName = "private_dns_name"
	// FieldPrivateIPAddress holds the string denoting the private_ip_address field in the database.
	FieldPrivateIPAddress = "private_ip_address"
	// FieldPublicDNSName holds the string denoting the public_dns_name field in the database.
	FieldPublicDNSName = "public_dns_name"
	// FieldPublicIPAddress holds the string denoting the public_ip_address field in the database.
	FieldPublicIPAddress = "public_ip_address"
	// FieldImageID holds the string denoting the image_id field in the database.
	FieldImageID = "image_id"
	// FieldArchitecture holds the string denoting the architecture field in the database.
	FieldArchitecture = "architecture"
	// FieldAvailabilityZone holds the string denoting the availability_zone field in the database.
	FieldAvailabilityZone = "availability_zone"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeSlide holds the string denoting the slide edge name in mutations.
	EdgeSlide = "slide"
	// Table holds the table name of the instance in the database.
	Table = "instances"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "instances"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_instances"
	// SlideTable is the table that holds the slide relation/edge.
	SlideTable = "instances"
	// SlideInverseTable is the table name for the Slide entity.
	// It exists in this package in order to avoid circular dependency with the "slide" package.
	SlideInverseTable = "slides"
	// SlideColumn is the table column denoting the slide relation/edge.
	SlideColumn = "slide_instance"
)

// Columns holds all SQL columns for instance fields.
var Columns = []string{
	FieldID,
	FieldInstanceID,
	FieldInstanceType,
	FieldPrivateDNSName,
	FieldPrivateIPAddress,
	FieldPublicDNSName,
	FieldPublicIPAddress,
	FieldImageID,
	FieldArchitecture,
	FieldAvailabilityZone,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "instances"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"slide_instance",
	"user_instances",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// ArchitectureValidator is a validator for the "architecture" field. It is called by the builders before save.
	ArchitectureValidator func(string) error
	// AvailabilityZoneValidator is a validator for the "availability_zone" field. It is called by the builders before save.
	AvailabilityZoneValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() ulid.ID
)