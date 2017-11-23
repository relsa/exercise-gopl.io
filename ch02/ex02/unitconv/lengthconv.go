package unitconv

import "fmt"

type Feet float64
type Metre float64

func (f Feet) String() string  { return fmt.Sprintf("%gft", f) }
func (m Metre) String() string { return fmt.Sprintf("%gm", m) }
