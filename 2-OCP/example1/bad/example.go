package bad

import "math"

type Calculator struct{}

func (c *Calculator) CalculateArea(shapeType string, params ...float64) float64 {
	switch shapeType {
	case "rectangle":
		return params[0] * params[1]
	case "circle":
		return math.Pi * params[0] * params[0]
		// Need to modify this function for each new shape
	}
	return 0
}
