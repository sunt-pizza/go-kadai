package main

import (
	"fmt"
	"math"
)

type ValidationError struct {
	Message string
	Field   string
	Value   float64
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// 図形のバリデーション
func validateShape(s Shape) *ValidationError {
	switch v := s.(type) {
	case Rectangle:
		if v.Width <= 0 {
			return &ValidationError{Message: "横幅は0より大きくしてください", Field: "Width", Value: v.Width}
		}
		if v.Height <= 0 {
			return &ValidationError{Message: "高さは0より大きくしてください", Field: "Height", Value: v.Height}
		}
	case Circle:
		if v.Radius <= 0 {
			return &ValidationError{Message: "半径は0より大きくしてください", Field: "Radius", Value: v.Radius}
		}
	}
	return nil
}

func main() {
	shapes := []Shape{
		Rectangle{Width: 10.0, Height: 5.0},
		Circle{Radius: 3.0},
		Rectangle{Width: -2.0, Height: 5.0}, // 不正（幅が不正）
		Circle{Radius: -1.0},                // 不正（半径が不正）
		Rectangle{Width: 1.0, Height: 0.0},  // 不正（高さが不正）
	}

	// 正しい図形を入れる箱（最初はからっぽ）
	calcShapes := []Shape{}

	for _, shape := range shapes {
		if err := validateShape(shape); err != nil {
			fmt.Println("エラー:", err.Message, err.Field, err.Value)
			continue
		}
		calcShapes = append(calcShapes, shape)
		fmt.Printf("面積:%v\n", shape.Area())
		fmt.Printf("周長:%v\n", shape.Perimeter())
	}
	fmt.Printf("型情報:%T\n", calcShapes)
}
