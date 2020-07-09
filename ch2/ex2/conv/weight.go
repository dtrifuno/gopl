package conv

import "fmt"

type Kilograms float64
type Pounds float64

func (k Kilograms) ToPounds() Pounds    { return Pounds(k * 2.20462) }
func (p Pounds) ToKilograms() Kilograms { return Kilograms(p / 2.20462) }

func (k Kilograms) String() string { return fmt.Sprintf("%g kg", k) }
func (p Pounds) String() string    { return fmt.Sprintf("%g lbs", p) }
