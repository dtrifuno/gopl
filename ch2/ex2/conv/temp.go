package conv

import "fmt"

type Celsius float64
type Fahrenheit float64

func (c Celsius) ToF() Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func (f Fahrenheit) ToC() Celsius { return Celsius((f - 32) * 5 / 9) }

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
