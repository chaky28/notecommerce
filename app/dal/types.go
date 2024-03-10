package dal

import (
	"time"
)

type Product struct {
	Id           string
	Name         string
	Salesprice   float32
	Price        float32
	CurrencyId   string
	OfferId      string
	Description  string
	InstalmentId string
	BreadcrumbId string
	ShippingId   string
	Stock        int
	Spec1Id      string
	Spec2Id      string
	Spec3Id      string
	Spec4Id      string
	Spec5Id      string
	Datetime     time.Time
}

type Instalment struct {
	Id        string
	CardId    string
	Amount    int
	Surcharge float32
	Datetime  time.Time
}

type ProductInstalments struct {
	Id           string
	ProductId    string
	InstalmentId string
}

type Currency struct {
	Id       string
	Name     string
	Symbol   string
	Datetime time.Time
}

type Offer struct {
	Id         string
	Name       string
	Multiplier string
	Datetime   time.Time
}

type Breadcrumb struct {
	Id       string
	L1       string
	L2       string
	L3       string
	L4       string
	L5       string
	Datetime time.Time
}

type Shipping struct {
	Id       string
	Name     string
	Datetime time.Time
}

type Spec struct {
	Id       string
	Name     string
	Datetime time.Time
}
