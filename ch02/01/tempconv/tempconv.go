// tempconv패키지는 섭씨와 화씨 변환을 수행한다.
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
	ZeroK Kelvin = 0
)

func (c Celsius) String() string    { return fmt.Sprintf("%g℃", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g℉", f) }
func (k Kelvin) String() string { return fmt.Sprintf("%gK", k) }