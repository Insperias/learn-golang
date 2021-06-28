package main

import (
	"fmt"
)

type celsius float64
type kelvin float64
type fahrenheit float64

func (k kelvin) celsius() celsius {
	k -= 273.15
	return celsius(k)
}

func (k kelvin) fahrenheit() fahrenheit {
	return k.celsius().fahrenheit()
}

func (c celsius) fahrenheit() fahrenheit{
	return fahrenheit(c * 9.0 / 5.0) + 32.0
}

func (c celsius) kelvin() kelvin {
	return kelvin(c + 273.15)
}

func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

func (f fahrenheit) kelvin() kelvin {
	return f.celsius().kelvin()
}

const (
	line         = "======================="
	rowFormat    = "| %8s | %8s |\n"
	numberFormat = "%.1f"
)

type getRowFn func(row int) (string, string)

// drawTable создает таблицу с двумя столбцами.
func drawTable(hdr1, hdr2 string, rows int, getRow getRowFn) {
	fmt.Println(line)
	fmt.Printf(rowFormat, hdr1, hdr2)
	fmt.Println(line)
	for row := 0; row < rows; row++ {
		cell1, cell2 := getRow(row)
		fmt.Printf(rowFormat, cell1, cell2)
	}
	fmt.Println(line)
}

func ctof(row int) (string, string) {
	c := celsius(row*5 - 40)
	f := c.fahrenheit()
	cell1 := fmt.Sprintf(numberFormat, c)
	cell2 := fmt.Sprintf(numberFormat, f)
	return cell1, cell2
}

func ftoc(row int) (string, string) {
	f := fahrenheit(row*5 - 40)
	c := f.celsius()
	cell1 := fmt.Sprintf(numberFormat, f)
	cell2 := fmt.Sprintf(numberFormat, c)
	return cell1, cell2
}

func main() {
	drawTable("°C", "°F", 29, ctof)
	fmt.Println()
	drawTable("°F", "°C", 29, ftoc)
}
