package conv

import "fmt"

type Meters float64
type Feet float64

func (m Meters) ToFeet() Feet   { return Feet(m * 3.28084) }
func (f Feet) ToMeters() Meters { return Meters(f / 3.28084) }

func (m Meters) String() string { return fmt.Sprintf("%g m", m) }
func (f Feet) String() string   { return fmt.Sprintf("%g ft", f) }
