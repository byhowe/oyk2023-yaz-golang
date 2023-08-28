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
func (self Item) DiscountRatio() float64 {
	return (float64(self.Discount) / float64(self.RawPrice)) * 100
}

// Convert kuruş to lira
func kurusToLira(kurus uint) float64 {
	return float64(kurus) / 100.0
}

// Calculate the discounted price in kuruş.
func (self Item) calculatePrice() uint {
	return self.RawPrice - self.Discount
}

// Calculate the discounted price in liras.
func (self Item) Price() float64 {
	return float64(self.calculatePrice()) / 100.0
}

// Sum total price of items with discount.
func (self Items) TotalPrice() float64 {
	var total uint
	for _, item := range self {
		total += item.calculatePrice()
	}
	return kurusToLira(total)
}

// Return a description string.
//
// Includes the name and the price information about the item.
func (self Item) Description() string {
	return fmt.Sprintf("%Q", self)
}

// Custom formatter that prints the description of the item using the 'Q' verb.
func (self Item) Format(f fmt.State, verb rune) {
	switch verb {
	case rune('Q'):
		if self.Discount != 0 {
			fmt.Fprintf(
				f,
				"%s - %v (%%%.1f TL indirimle %v TL)",
				self.Name,
				kurusToLira(self.RawPrice),
				self.DiscountRatio(),
				self.Price(),
			)
		} else {
			fmt.Fprintf(f, "%s - %v TL", self.Name, self.Price())
		}
	default:
		fmt.Fprint(f, self)
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

	total := items.TotalPrice()
	fmt.Printf("Toplam Fiyat: %v TL\n", total)
}
