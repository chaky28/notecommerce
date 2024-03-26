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

func (p Product) GetColumns() []string {
	return []string{
		"Id",
		"Name",
		"Salesprice",
		"Price",
		"CurrencyId",
		"OfferId",
		"Description",
		"InstalmentId",
		"BreadcrumbId",
		"ShippingId",
		"Stock",
		"Spec1Id",
		"Spec2Id",
		"Spec3Id",
		"Spec4Id",
		"Spec5Id",
		"Datetime",
	}
}

type Instalment struct {
	Id        string
	CardId    string
	Amount    int
	Surcharge float32
	Datetime  time.Time
}

func (ins Instalment) GetColumns() []string {
	return []string{
		"Id",
		"CardId",
		"Amount",
		"Surcharge",
		"Datetime",
	}
}

type ProductInstalments struct {
	Id           string
	ProductId    string
	InstalmentId string
}

func (pi ProductInstalments) GetColumns() []string {
	return []string{
		"Id",
		"ProductId",
		"InstalmentId",
	}
}

type Currency struct {
	Id       string
	Name     string
	Symbol   string
	Datetime time.Time
}

func (curr Currency) GetColumns() []string {
	return []string{
		"Id",
		"Name",
		"Symbol",
		"Datetime",
	}
}

type Offer struct {
	Id         string
	Name       string
	Multiplier string
	Datetime   time.Time
}

func (off Offer) GetColumns() []string {
	return []string{
		"Id",
		"Name",
		"Multiplier",
		"Datetime",
	}
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

func (br Breadcrumb) GetColumns() []string {
	return []string{
		"Id",
		"L1",
		"L2",
		"L3",
		"L4",
		"L5",
		"Datetime",
	}
}

type Shipping struct {
	Id       string
	Name     string
	Datetime time.Time
}

func (sh Shipping) GetColumns() []string {
	return []string{
		"Id",
		"Name",
		"Datetime",
	}
}

type Spec struct {
	Id       string
	Name     string
	Datetime time.Time
}

func (sp Spec) GetColumns() []string {
	return []string{
		"Id",
		"Name",
		"Datetime",
	}
}
