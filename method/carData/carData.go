/* carData.go */

package carData

type Car struct {
	Color string
	Body  string
}

func (Auto Car) Accelerate() string {
	return "accelerating-->"
}

func (Mobile Car) Zooming() string {
	return "zooming-->"
}
