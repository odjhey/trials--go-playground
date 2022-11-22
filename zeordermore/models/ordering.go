package models

func CreateOrder(items []Item) (Order, error) {

	return Order{
		Number: "OR001",
		Items: []struct {
			Item Item
			Qty  int
		}{
			{Item: Item{Sku: "chicken", Uom: "bag"}, Qty: 9},
		},
	}, nil

}
