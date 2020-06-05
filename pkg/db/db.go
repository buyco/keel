package db

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	"strings"

	// Needed by GORM
	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/sirupsen/logrus"
)

// Database GORM wrapper
type Database struct {
	DbHandler *gorm.DB
}

// NewDatabase is Database constructor
func NewDatabase(db *gorm.DB) *Database {
	return &Database{DbHandler: db}
}

// MigrateDB migrates DB from table structs
func (d *Database) MigrateDB(tableStruct ...interface{}) error {
	db := d.DbHandler.AutoMigrate(tableStruct...)
	if db.Error != nil {
		log.Error("Unable to migrate db")
		d.DbHandler.Rollback()
		return db.Error
	}
	return nil
}

// AddForeignKey adds constraints to table
func (d *Database) AddForeignKey(tableStruct interface{}, field, dest, onDelete, onUpdate string) error {
	db := d.DbHandler.Model(tableStruct).AddForeignKey(field, dest, onDelete, onUpdate)
	if db.Error != nil {
		log.Error("Unable to create constraint")
		d.DbHandler.Rollback()
		return db.Error
	}
	return nil
}

// AddUniqueIndex adds unique index to table
func (d *Database) AddUniqueIndex(tableStruct interface{}, indexName string, fields ...string) error {
	db := d.DbHandler.Model(tableStruct).AddUniqueIndex(indexName, fields...)
	if db.Error != nil {
		log.Error("Unable to create unique index constraint")
		d.DbHandler.Rollback()
		return db.Error
	}
	return nil
}

// CreateDB creates database
func (d *Database) CreateDB(dbName string) error {
	db := d.DbHandler.Exec(fmt.Sprintf("CREATE DATABASE %s;", dbName))
	if db.Error != nil {
		log.Errorf("Unable to create db [%s]", dbName)
		return db.Error
	}
	return nil
}

// DropDB drops database
func (d *Database) DropDB(dbName string) error {
	db := d.DbHandler.Exec(fmt.Sprintf("DROP DATABASE %s;", dbName))
	if db.Error != nil {
		log.Errorf("Unable to drop db [%s]", dbName)
		return db.Error
	}
	return nil
}

// CreateEnum creates enum in database
func (d *Database) CreateEnum(enumName string, values []string) error {
	db := d.DbHandler.Exec(fmt.Sprintf("CREATE TYPE %s AS ENUM('%s');", enumName, strings.Join(values, "','")))
	if db.Error != nil {
		log.Errorf("Unable to create enum in db [%s]", enumName)
		return db.Error
	}
	return nil
}

// UpdateEnum updates enum in database
func (d *Database) UpdateEnum(enumName, value string) error {
	db := d.DbHandler.Exec(fmt.Sprintf("ALTER TYPE %s ADD VALUE '%s';", enumName, value))
	if db.Error != nil {
		log.Errorf("Unable to update enum in db [%s] with value [%s]", enumName, value)
		return db.Error
	}
	return nil
}

// DropEnum deletes enum in database
func (d *Database) DropEnum(enumName string) error {
	db := d.DbHandler.Exec(fmt.Sprintf("DROP TYPE %s;", enumName))
	if db.Error != nil {
		log.Errorf("Unable to drop enum in db [%s]", enumName)
		return db.Error
	}
	return nil
}

// CreateTables creates DB tables from table structs
func (d *Database) CreateTables(tablesStruct ...interface{}) error {
	db := d.DbHandler.CreateTable(tablesStruct...)
	if db.Error != nil {
		log.Errorf("Unable to create tables [%v]", db.Error)
		d.DbHandler.Rollback()
		return db.Error
	}
	return nil
}

// WithLimitAndOffset is a query scope function for pagination
func (d *Database) WithLimitAndOffset(limit uint, offset uint) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Limit(limit).Offset(offset)
	}
}

// WithIds is a query scope function to include given IDs (integer IDs)
func (d *Database) WithIds(ids []uint, tableName string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(tableName+".id in (?)", ids)
	}
}

// WithNotIds is a query scope function to exclude given ids
func (d *Database) WithNotIds(ids []uint, tableName string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(tableName+".id not in (?)", ids)
	}
}

// WithIDs is a query scope to include the given IDs (string IDs)
func (d *Database) WithIDs(tableName string, ids []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(tableName+".id in (?)", ids)
	}
}

// WithStartingAfter is a query scope for organization ID
func (d *Database) WithOrgID(tableName string, id uint) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(tableName+".org_id = ?", id)
	}
}

// WithStartingAfter is a query scope for external IDs
func (d *Database) WithExternalIDs(tableName string, externalIDs []sql.NullString) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(tableName+".external_id in (?)", externalIDs)
	}
}

// WithStartingAfter is a query scope for cursors
func (d *Database) WithStartingAfter(tableName string, limit uint, after string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(tableName+".id > ?", after).Limit(limit)
	}
}

// WithStartingAfter is a query scope for cursors
func (d *Database) WithEndingBefore(tableName string, limit uint, before string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(tableName+".id < ?", before).Limit(limit)
	}
}
