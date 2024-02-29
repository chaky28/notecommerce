package dal

import (
	"log"
	"reflect"
)

// -------------------- DB constants --------------------

const PgCredsFilePath = "/notecommerce_db_creds.txt"
const DbHost = "192.168.1.6"
const DbPort = "5432"
const SslMode = "disable"
const DbName = "notecommerce"
const NotEcommerceDbVersion = 1 //Change version HERE to update DB with versioning

type NotEcommerceDB struct {
	db DB
}

func GetNotEcommerceDB() NotEcommerceDB {
	user, passw := getDbCreds(PgCredsFilePath)
	ndb := NotEcommerceDB{
		db: DB{
			name:     DbName,
			port:     DbPort,
			ssl:      SslMode,
			driver:   "postgres",
			version:  NotEcommerceDbVersion,
			user:     user,
			password: passw,
			host:     DbHost,
		},
	}

	//Update DB before returning the struct
	checkVersioning(ndb.db, reflect.ValueOf(ndb))

	return ndb

}

func (ndb NotEcommerceDB) Db_v1() {
	conn := ndb.db.getDbConn()
	defer conn.Close()

	sql := `CREATE TABLE products 
			(id varchar(36) PRIMARY KEY,
			name varchar(128),
			saleprice decimal(128, 64),
			price decimal(128, 64),
			currency_id varchar(36),
			offer_id varchar(36) DEFAULT NULL,
			description varchar(512) DEFAULT NULL,
			instalment_id varchar(36) DEFAULT NULL,
			breadcrumb_id varchar(36),
			shipping_id varchar(36) NULL,
			stock int DEFAULT 0,
			spec1_id varchar(36) NULL,
			spec2_id varchar(36) NULL,
			spec3_id varchar(36) NULL,
			spec4_id varchar(36) NULL,
			spec5_id varchar(36) NULL
			)`
	if _, err := conn.Exec(sql); err != nil {
		log.Fatal("ERROR running query ", sql, " --> ", err.Error())
	}

	sql = `CREATE TABLE currency
		   (id varchar(36) PRIMARY KEY,
		   name varchar(128),
		   symbol varchar(16)
		   )`
	if _, err := conn.Exec(sql); err != nil {
		log.Fatal("ERROR running query ", sql, " --> ", err.Error())
	}

	sql = `CREATE TABLE offers
		   (id varchar(36) PRIMARY KEY,
		   name varchar(128),
		   multiplier decimal(128, 64)
		   )`
	if _, err := conn.Exec(sql); err != nil {
		log.Fatal("ERROR running query ", sql, " --> ", err.Error())
	}

	sql = `CREATE TABLE instalments
		   (id varchar(36) PRIMARY KEY,
			card_id varchar(36),
			amount int,
			surcharge decimal(128, 64)
		    )`
	if _, err := conn.Exec(sql); err != nil {
		log.Fatal("ERROR running query ", sql, " --> ", err.Error())
	}
	sql = `CREATE TABLE cards
		   (id varchar(36) PRIMARY KEY,
		   name varchar(128),
		   bank_id varchar(36),
		   institution_id varchar(36)
		   )`
	if _, err := conn.Exec(sql); err != nil {
		log.Fatal("ERROR running query ", sql, " --> ", err.Error())
	}
	sql = `CREATE TABLE banks
		   (id varchar(36) PRIMARY KEY,
		 	name varchar(128),
			country_id varchar(36)
		    )`
	if _, err := conn.Exec(sql); err != nil {
		log.Fatal("ERROR running query ", sql, " --> ", err.Error())
	}

	sql = `CREATE TABLE countries
		   (id varchar(36) PRIMARY KEY,
			name varchar(128)
		    )`
	if _, err := conn.Exec(sql); err != nil {
		log.Fatal("ERROR running query ", sql, " --> ", err.Error())
	}

	sql = `CREATE TABLE breadcrumbs
		   (id varchar(36) PRIMARY KEY,
			lev1_category_id varchar(36),
			lev2_category_id varchar(36),
			lev3_category_id varchar(36) DEFAULT NULL,
			lev4_category_id varchar(36) DEFAULT NULL,
			lev5_category_id varchar(36) DEFAULT NULL
		    )`
	if _, err := conn.Exec(sql); err != nil {
		log.Fatal("ERROR running query ", sql, " --> ", err.Error())
	}

	sql = `CREATE TABLE categories 
		   (id varchar(36) PRIMARY KEY,
			name varchar(128)
		    )`
	if _, err := conn.Exec(sql); err != nil {
		log.Fatal("ERROR running query ", sql, " --> ", err.Error())
	}

	sql = `CREATE TABLE shipping 
		   (id varchar(36) PRIMARY KEY,
			name varchar(128)
		    )`
	if _, err := conn.Exec(sql); err != nil {
		log.Fatal("ERROR running query ", sql, " --> ", err.Error())
	}

	sql = `CREATE TABLE specs 
		   (id varchar(36) PRIMARY KEY,
			name varchar(128)
			)`
	if _, err := conn.Exec(sql); err != nil {
		log.Fatal("ERROR running query ", sql, " --> ", err.Error())
	}
}
