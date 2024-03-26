package dal

import (
	"database/sql"
	"log"
	"reflect"
	"strconv"

	"github.com/google/uuid"
	_ "github.com/lib/pq"

	"github.com/chaky28/notecommerce/app/app/files"
)

type DB struct {
	user     string
	host     string
	name     string
	password string
	ssl      string
	driver   string
	port     string
	version  int
}

func getDbCreds(filePath string) (string, string) {
	fileData := files.ReadFile(filePath)
	return files.GetUserAndPasswordFromFileData(fileData)
}

func (db DB) getDbConn() *sql.DB {
	connStr := "user=" + db.user + " dbname=" + db.name + " password=" + db.password + " host=" + db.host + " port=" + db.port + " sslmode=" + db.ssl
	conn, err := sql.Open(db.driver, connStr)
	if err != nil {
		log.Fatal("ERROR: Opening connection to DB -->", err.Error())
	}

	return conn
}

// ------------------------------- DB versioning -------------------------------

func checkVersioning(db DB, reflectValue reflect.Value) {
	conn := db.getDbConn()
	defer conn.Close()

	if !db.doesVersioningExists() {
		db.Db_v0()
	}

	currentDbVersion := db.getCurrentDbVersion()

	for db.version > currentDbVersion {
		currentDbVersion++
		f := reflectValue.MethodByName("Db_v" + strconv.Itoa(currentDbVersion))
		f.Call(nil)
		db.updateDbVersioning(currentDbVersion)
	}
}

func (db DB) updateDbVersioning(version int) {
	conn := db.getDbConn()
	defer conn.Close()

	_, err := conn.Exec("UPDATE versioning SET version = $1", version)
	if err != nil {
		log.Fatal("ERROR: Updating DB versioning --> ", err.Error())
	}
}

func (db DB) getCurrentDbVersion() int {
	conn := db.getDbConn()
	defer conn.Close()

	row, err := conn.Query("SELECT version FROM versioning")
	if err != nil {
		log.Fatal("ERROR: Getting DB version --> ", err.Error())
	}

	var version int
	row.Next()
	row.Scan(&version)
	return version
}

func (db DB) doesVersioningExists() bool {
	conn := db.getDbConn()
	defer conn.Close()

	row, err := conn.Query("SELECT EXISTS (SELECT 1 FROM pg_tables WHERE schemaname = 'public' AND tablename = 'versioning')")
	if err != nil {
		log.Fatal("ERROR: Checking if versioning table exists --> ", err.Error())
	}

	var exists bool
	row.Next()
	row.Scan(&exists)
	return exists
}

func (db DB) Db_v0() {
	conn := db.getDbConn()
	defer conn.Close()

	_, err := conn.Exec("CREATE TABLE versioning (id varchar(36) PRIMARY KEY, version int)")
	if err != nil {
		log.Fatal("ERROR: Executing v0 versioning function --> ", err.Error())
	}

	uuid := uuid.New()
	_, err = conn.Exec("INSERT INTO versioning VALUES ($1, $2)", uuid, 0)
	if err != nil {
		log.Fatal("ERROR: Executing v0 versioning function --> ", err.Error())
	}
}

// ------------------------------- Helper functions -------------------------------
