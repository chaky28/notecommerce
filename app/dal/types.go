package dal

import (
	"time"
)

type Product struct {
	Id            string
	Name          string
	Salesprice    float32
	Price         float32
	Currency_id   string
	Offer_id      string
	Description   string
	Breadcrumb_id string
	Shipping_id   string
	Stock         int
	Spec1_id      string
	Spec2_id      string
	Spec3_id      string
	Spec4_id      string
	Spec5_id      string
}

type Instalment struct {
	Id        string
	Card_id   string
	Amount    int
	Surcharge float32
	Datetime  time.Time
}

type ProductInstalments struct {
	Id           string
	ProductId    string
	InstalmentId string
}
