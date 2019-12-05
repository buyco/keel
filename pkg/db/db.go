package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
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

// WithIds is a query scope function to include given ids
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
