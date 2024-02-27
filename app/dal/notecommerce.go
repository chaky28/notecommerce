package dal

import (
	"reflect"
)

// -------------------- DB constants --------------------

const PgCredsFilePath = "/notecommerce_db_creds.txt"
const DbHost = "192.168.1.6"
const DbPort = "5432"
const SslMode = "disable"
const DbName = "notecommerce"
const NotEcommerceDbVersion = 2 //Change version HERE to update DB with versioning

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

}
