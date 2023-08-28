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

// Calculate the discount ratio as percentage.
func (self Item) calculateDiscountRatio() float64 {
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
func (self Item) CalculatePrice() float64 {
	return float64(self.calculatePrice()) / 100.0
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
				self.calculateDiscountRatio(),
				self.CalculatePrice(),
			)
		} else {
			fmt.Fprintf(f, "%s - %v TL", self.Name, self.CalculatePrice())
		}
	default:
		fmt.Fprint(f, self)
	}
}

// Sum total price of items with discount.
func TotalPrice(items []Item) float64 {
	var total uint
	for _, item := range items {
		total += item.calculatePrice()
	}
	return kurusToLira(total)
}

func main() {
	items := []Item{
		{Name: "Elma", RawPrice: 75, Discount: 7},
		{Name: "Portakal", RawPrice: 100, Discount: 0},
	}

	for _, item := range items {
		fmt.Printf("%Q\n", item)
	}

	total := TotalPrice(items)
	fmt.Printf("Toplam Fiyat: %v TL\n", total)
}
