package main

import "fmt"

type Descirabable interface {
	Description() string
}

type Item struct {
	// Name of the product
	Name string
	// Price without the discount
	RawPrice uint
	// Discount amount
	Discount uint
}

type Items []Item

// Calculate the discount ratio as percentage.
func (item Item) DiscountRatio() float64 {
	return (float64(item.Discount) / float64(item.RawPrice)) * 100
}

// Convert kuruş to lira
func kurusToLira(kurus uint) float64 {
	return float64(kurus) / 100.0
}

// Calculate the discounted price in kuruş.
func (item Item) calculatePrice() uint {
	return item.RawPrice - item.Discount
}

// Calculate the discounted price in liras.
func (item Item) Price() float64 {
	return float64(item.calculatePrice()) / 100.0
}

// Sum total price of items with discount.
func (items Items) TotalPrice() float64 {
	var total uint
	for _, item := range items {
		total += item.calculatePrice()
	}
	return kurusToLira(total)
}

// Return a description string.
//
// Includes the name and the price information about the item.
func (item Item) Description() string {
	return fmt.Sprintf("%Q", item)
}

func (items Items) Description() string {
	return fmt.Sprintf("%d item(s) worth %.2f liras", len(items), items.TotalPrice())
}

// Custom formatter that prints the description of the item using the 'Q' verb.
func (item Item) Format(f fmt.State, verb rune) {
	switch verb {
	case rune('Q'):
		if item.Discount != 0 {
			fmt.Fprintf(
				f,
				"%s - %v (%%%.2f TL indirimle %v TL)",
				item.Name,
				kurusToLira(item.RawPrice),
				item.DiscountRatio(),
				item.Price(),
			)
		} else {
			fmt.Fprintf(f, "%s - %v TL", item.Name, item.Price())
		}
	default:
		fmt.Fprint(f, item)
	}
}

func main() {
	items := Items{
		{Name: "Elma", RawPrice: 75, Discount: 7},
		{Name: "Portakal", RawPrice: 100, Discount: 0},
	}

	for _, item := range items {
		fmt.Printf("%Q\n", item)
	}

	fmt.Println(items.Description())
}
