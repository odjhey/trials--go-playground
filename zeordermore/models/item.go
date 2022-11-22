package models

type warehouse struct {
	name string
}

type Item struct {
	Sku string
	Uom string
}

type itemInventory struct {
	sku       string
	uom       string
	warehouse warehouse
}

type Order struct {
	Items []struct {
		Item Item
		Qty  int
	}
	Number string
}
