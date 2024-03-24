package dal

import (
	"log"
	"reflect"

	"github.com/chaky28/notecommerce/app/app/common"
)

// -------------------- DB constants --------------------

const PgCredsFilePath = common.CredentialsDirectory + "/notecommerce_db.txt"
const DbHost = "192.168.1.17"
const DbPort = "5432"
const SslMode = "disable"
const DbName = "notecommerce"
const NotEcommerceDbVersion = 3 //Change version HERE to update DB with versioning

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

func (ndb NotEcommerceDB) InsertNewProduct(pr Product) error {
	conn := ndb.db.getDbConn()
	defer conn.Close()

	sql := `INSERT INTO products VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17
	)`
	_, err := conn.Exec(sql, pr.Id, pr.Name, pr.Salesprice, pr.Price, pr.CurrencyId, pr.OfferId,
		pr.Description, pr.InstalmentId, pr.BreadcrumbId, pr.ShippingId,
		pr.Stock, pr.Spec1Id, pr.Spec2Id, pr.Spec3Id, pr.Spec4Id, pr.Spec5Id, pr.Datetime)

	return err
}

// ------------------------------- Updates -------------------------------

func (ndb NotEcommerceDB) UpdateStock(productId string, newStock int) error {
	conn := ndb.db.getDbConn()
	defer conn.Close()

	sql := "UPDATE products SET stock = $1 WHERE id = $2"
	_, err := conn.Exec(sql, productId, newStock)

	return err
}

// ------------------------------- Selections -------------------------------

func (ndb NotEcommerceDB) GetProductsByCatId(catId string) ([]Product, error) {
	conn := ndb.db.getDbConn()
	defer conn.Close()

	sql := `SELECT pr.* FROM products 
			pr JOIN breadcrumbs br 
			ON pr.breadcrumb_id = br.id WHERE 
			br.lev1_category_id = $1 OR
			br.lev2_category_id = $2 OR
			br.lev3_category_id = $3 OR
			br.lev4_category_id = $4 OR
			br.lev5_category_id = $5
			GROUP BY pr.id`

	rows, err := conn.Query(sql, catId, catId, catId, catId, catId)
	if err != nil {
		return nil, err
	}

	prods := []Product{}
	for rows.Next() {
		var pr Product
		err = rows.Scan(&pr.Id, &pr.Name, &pr.Salesprice, &pr.Price, &pr.CurrencyId, &pr.OfferId,
			&pr.Description, &pr.InstalmentId, &pr.BreadcrumbId, &pr.ShippingId,
			&pr.Stock, &pr.Spec1Id, &pr.Spec2Id, &pr.Spec3Id, &pr.Spec4Id, &pr.Spec5Id, &pr.Datetime)

		if err != nil {
			return prods, err
		}

		prods = append(prods, pr)
	}
	return prods, nil
}

func (ndb NotEcommerceDB) GetInstalmentById(id string) (Instalment, error) {
	conn := ndb.db.getDbConn()
	defer conn.Close()

	sql := `SELECT * FROM currency WHERE id = ?`
	row := conn.QueryRow(sql, id)

	var inst Instalment
	err := row.Scan(&inst.Id, &inst.CardId, &inst.Amount, &inst.Surcharge, &inst.Datetime)

	return inst, err
}

func (ndb NotEcommerceDB) GetCurrencyById(id string) (Currency, error) {
	conn := ndb.db.getDbConn()
	defer conn.Close()

	sql := `SELECT * FROM currency WHERE id = ?`
	row := conn.QueryRow(sql, id)

	var curr Currency
	err := row.Scan(&curr.Id, &curr.Name, &curr.Symbol, &curr.Datetime)

	return curr, err
}

func (ndb NotEcommerceDB) GetOfferById(id string) (Offer, error) {
	conn := ndb.db.getDbConn()
	defer conn.Close()

	sql := `SELECT * FROM offers WHERE id = ?`
	row := conn.QueryRow(sql, id)

	var offer Offer
	err := row.Scan(&offer.Id, &offer.Name, &offer.Multiplier, &offer.Datetime)

	return offer, err
}

func (ndb NotEcommerceDB) GetProductInstalments(productId string) ([]Instalment, error) {
	conn := ndb.db.getDbConn()
	defer conn.Close()

	sql := `SELECT instalment_id FROM product_instalments WHERE product_id = $1`
	piRows, err := conn.Query(sql, productId)

	result := []Instalment{}
	for piRows.Next() {
		var instalmentId string
		piRows.Scan(&instalmentId)

		sql = "SELECT * FROM instalments WHERE id = $1"
		iRows := conn.QueryRow(sql, instalmentId)

		var ins Instalment
		err = iRows.Scan(&ins.Id, &ins.CardId, &ins.Amount, &ins.Surcharge, &ins.Datetime)
		result = append(result, ins)
	}

	return result, err
}

func (ndb NotEcommerceDB) GetBreadcrumbById(id string) (Breadcrumb, error) {
	conn := ndb.db.getDbConn()
	defer conn.Close()

	sql := `SELECT * FROM breadcrumbs WHERE id = ?`
	row := conn.QueryRow(sql, id)

	var bc Breadcrumb
	err := row.Scan(&bc.Id, &bc.L1, &bc.L2, &bc.L3, &bc.L4, &bc.L5, &bc.Datetime)

	return bc, err
}

func (ndb NotEcommerceDB) GetShippingById(id string) (Shipping, error) {
	conn := ndb.db.getDbConn()
	defer conn.Close()

	sql := `SELECT * FROM shipping WHERE id = ?`
	row := conn.QueryRow(sql, id)

	var sh Shipping
	err := row.Scan(&sh.Id, &sh.Name, &sh.Datetime)

	return sh, err
}

func (ndb NotEcommerceDB) GetProductSpecById(id string) (Spec, error) {
	conn := ndb.db.getDbConn()
	defer conn.Close()

	sql := `SELECT * FROM shipping WHERE id = ?`
	row := conn.QueryRow(sql, id)

	var spec Spec
	err := row.Scan(&spec.Id, &spec.Name, &spec.Datetime)

	return spec, err
}

// ------------------------------- Insertions -------------------------------

func (ndb NotEcommerceDB) InsertInstalment(inst Instalment) error {
	conn := ndb.db.getDbConn()
	defer conn.Close()

	sql := "INSERT INTO instalments VALUES ($1, $2, $3, $4, $5)"
	_, err := conn.Exec(sql, inst.Id, inst.CardId, inst.Amount, inst.Surcharge, inst.Datetime)

	return err
}

func (ndb NotEcommerceDB) InsertCurrency(curr Currency) error {
	conn := ndb.db.getDbConn()
	defer conn.Close()

	sql := "INSERT INTO currency VALUES ($1, $2, $3, $4)"
	_, err := conn.Exec(sql, curr.Id, curr.Name, curr.Symbol, curr.Datetime)

	return err
}

func (ndb NotEcommerceDB) InsertOffer(offer Offer) error {
	conn := ndb.db.getDbConn()
	defer conn.Close()

	sql := "INSERT INTO offers VALUES ($1, $2, $3, $4)"
	_, err := conn.Exec(sql, offer.Id, offer.Name, offer.Multiplier, offer.Datetime)

	return err
}

func (ndb NotEcommerceDB) InsertBreadcrumb(bc Breadcrumb) error {
	conn := ndb.db.getDbConn()
	defer conn.Close()

	sql := "INSERT INTO offers VALUES ($1, $2, $3, $4, $5, $6, $7)"
	_, err := conn.Exec(sql, bc.Id, bc.L1, bc.L2, bc.L3, bc.L4, bc.L5, bc.Datetime)

	return err
}

func (ndb NotEcommerceDB) InsertShipping(sh Shipping) error {
	conn := ndb.db.getDbConn()
	defer conn.Close()

	sql := "INSERT INTO shipping VALUES ($1, $2, $3)"
	_, err := conn.Exec(sql, sh.Id, sh.Name, sh.Datetime)

	return err
}

func (ndb NotEcommerceDB) InsertSpec(spec Spec) error {
	conn := ndb.db.getDbConn()
	defer conn.Close()

	sql := "INSERT INTO specs VALUES ($1, $2, $3)"
	_, err := conn.Exec(sql, spec.Id, spec.Name, spec.Datetime)

	return err
}

// ------------------------------- Versioning -------------------------------

func (ndb NotEcommerceDB) Db_v3() {
	conn := ndb.db.getDbConn()
	defer conn.Close()

	sql := `CREATE TABLE product_instalments (
			id varchar(36) PRIMARY KEY,
			product_id varchar(36),
			instalment_id varchar(36)
			)`
	conn.Exec(sql)
}

// Create user related tables
func (ndb NotEcommerceDB) Db_v2() {
	conn := ndb.db.getDbConn()
	defer conn.Close()

	sql := `CREATE TABLE users 
		    (id varchar(36) PRIMARY KEY,
			user_name varchar(128),
			name_id varchar(36),
			last_name_id varchar(36),
			country_id varchar(36),
			pass_hash varchar(60),
			datetime timestamp without time zone NOT NULL DEFAULT (current_timestamp AT TIME ZONE 'ADT')
			)`
	if _, err := conn.Exec(sql); err != nil {
		log.Fatal("ERROR running query ", sql, " --> ", err.Error())
	}

	sql = `CREATE TABLE names 
		    (id varchar(36) PRIMARY KEY,
			name varchar(128),
			datetime timestamp without time zone NOT NULL DEFAULT (current_timestamp AT TIME ZONE 'ADT')
			)`
	if _, err := conn.Exec(sql); err != nil {
		log.Fatal("ERROR running query ", sql, " --> ", err.Error())
	}

	sql = `CREATE TABLE last_names 
		    (id varchar(36) PRIMARY KEY,
			last_name varchar(128),
			datetime timestamp without time zone NOT NULL DEFAULT (current_timestamp AT TIME ZONE 'ADT')
			)`
	if _, err := conn.Exec(sql); err != nil {
		log.Fatal("ERROR running query ", sql, " --> ", err.Error())
	}

	sql = `CREATE TABLE favs 
		   (id varchar(36) PRIMARY KEY,
		   user_id varchar(36),
		   product_id varchar(36),
		   datetime timestamp without time zone NOT NULL DEFAULT (current_timestamp AT TIME ZONE 'ADT')
		   )`
	if _, err := conn.Exec(sql); err != nil {
		log.Fatal("ERROR running query ", sql, " --> ", err.Error())
	}

	sql = `CREATE TABLE notifications 
		   (id varchar(36) PRIMARY KEY,
		   user_id varchar(36),
		   from_user_id varchar(36),
		   type_id varchar(36),
		   datetime timestamp without time zone NOT NULL DEFAULT (current_timestamp AT TIME ZONE 'ADT')
		   )`
	if _, err := conn.Exec(sql); err != nil {
		log.Fatal("ERROR running query ", sql, " --> ", err.Error())
	}

	sql = `CREATE TABLE notification_types 
		   (id varchar(36) PRIMARY KEY,
		   message varchar(128),
		   datetime timestamp without time zone NOT NULL DEFAULT (current_timestamp AT TIME ZONE 'ADT')
		   )`
	if _, err := conn.Exec(sql); err != nil {
		log.Fatal("ERROR running query ", sql, " --> ", err.Error())
	}

	sql = `CREATE TABLE purchases 
		   (id varchar(36),
		   user_id varchar(36),
		   product_id varchar(36),
		   amount int,
		   card_id varchar(36),
		   instalments int,
		   datetime timestamp without time zone NOT NULL DEFAULT (current_timestamp AT TIME ZONE 'ADT')
		   )`
	if _, err := conn.Exec(sql); err != nil {
		log.Fatal("ERROR running query ", sql, " --> ", err.Error())
	}
}

// Product related tables creation
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
			spec5_id varchar(36) NULL,
			datetime timestamp without time zone NOT NULL DEFAULT (current_timestamp AT TIME ZONE 'ADT')
			)`
	if _, err := conn.Exec(sql); err != nil {
		log.Fatal("ERROR running query ", sql, " --> ", err.Error())
	}

	sql = `CREATE TABLE currency
		   (id varchar(36) PRIMARY KEY,
		   name varchar(128),
		   symbol varchar(16),
		   datetime timestamp without time zone NOT NULL DEFAULT (current_timestamp AT TIME ZONE 'ADT')
		   )`
	if _, err := conn.Exec(sql); err != nil {
		log.Fatal("ERROR running query ", sql, " --> ", err.Error())
	}

	sql = `CREATE TABLE offers
		   (id varchar(36) PRIMARY KEY,
		   name varchar(128),
		   multiplier decimal(128, 64),
		   datetime timestamp without time zone NOT NULL DEFAULT (current_timestamp AT TIME ZONE 'ADT')
		   )`
	if _, err := conn.Exec(sql); err != nil {
		log.Fatal("ERROR running query ", sql, " --> ", err.Error())
	}

	sql = `CREATE TABLE instalments
		   (id varchar(36) PRIMARY KEY,
			card_id varchar(36),
			amount int,
			surcharge decimal(128, 64),
			datetime timestamp without time zone NOT NULL DEFAULT (current_timestamp AT TIME ZONE 'ADT')
		    )`
	if _, err := conn.Exec(sql); err != nil {
		log.Fatal("ERROR running query ", sql, " --> ", err.Error())
	}
	sql = `CREATE TABLE cards
		   (id varchar(36) PRIMARY KEY,
		   name varchar(128),
		   bank_id varchar(36),
		   institution_id varchar(36),
		   datetime timestamp without time zone NOT NULL DEFAULT (current_timestamp AT TIME ZONE 'ADT')
		   )`
	if _, err := conn.Exec(sql); err != nil {
		log.Fatal("ERROR running query ", sql, " --> ", err.Error())
	}
	sql = `CREATE TABLE banks
		   (id varchar(36) PRIMARY KEY,
		 	name varchar(128),
			country_id varchar(36),
			datetime timestamp without time zone NOT NULL DEFAULT (current_timestamp AT TIME ZONE 'ADT')
		    )`
	if _, err := conn.Exec(sql); err != nil {
		log.Fatal("ERROR running query ", sql, " --> ", err.Error())
	}

	sql = `CREATE TABLE countries
		   (id varchar(36) PRIMARY KEY,
			name varchar(128),
			datetime timestamp without time zone NOT NULL DEFAULT (current_timestamp AT TIME ZONE 'ADT')
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
			lev5_category_id varchar(36) DEFAULT NULL,
			datetime timestamp without time zone NOT NULL DEFAULT (current_timestamp AT TIME ZONE 'ADT')
		    )`
	if _, err := conn.Exec(sql); err != nil {
		log.Fatal("ERROR running query ", sql, " --> ", err.Error())
	}

	sql = `CREATE TABLE categories 
		   (id varchar(36) PRIMARY KEY,
			name varchar(128),
			datetime timestamp without time zone NOT NULL DEFAULT (current_timestamp AT TIME ZONE 'ADT')
		    )`
	if _, err := conn.Exec(sql); err != nil {
		log.Fatal("ERROR running query ", sql, " --> ", err.Error())
	}

	sql = `CREATE TABLE shipping 
		   (id varchar(36) PRIMARY KEY,
			name varchar(128),
			datetime timestamp without time zone NOT NULL DEFAULT (current_timestamp AT TIME ZONE 'ADT')
		    )`
	if _, err := conn.Exec(sql); err != nil {
		log.Fatal("ERROR running query ", sql, " --> ", err.Error())
	}

	sql = `CREATE TABLE specs 
		   (id varchar(36) PRIMARY KEY,
			name varchar(128),
			datetime timestamp without time zone NOT NULL DEFAULT (current_timestamp AT TIME ZONE 'ADT')
			)`
	if _, err := conn.Exec(sql); err != nil {
		log.Fatal("ERROR running query ", sql, " --> ", err.Error())
	}
}
