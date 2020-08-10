package main

import (
	"fmt"
	"math"
)

func main() {
	redius := -3.0
	area, err := circleArea(redius)
	if err != nil {
		fmt.Println(err)
		if err,ok := err.(*areaError); ok {
			fmt.Printf("redius: %.2f,msg: %s\n",redius,err.msg)
		}
		return
	}
	fmt.Println("面积是",area)
}

type areaError struct {
	msg    string
	redius float64
}

func (e *areaError) Error() string {
	return fmt.Sprintf("error :半径,%.2f,msg: %s", e.redius, e.msg)
}

func circleArea(redius float64) (float64, error) {
	if redius < 0 {
		return 0, &areaError{"半径是非法的", redius}
	}
	return math.Pi * redius * redius, nil
}
