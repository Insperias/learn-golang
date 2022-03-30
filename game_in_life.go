package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width = 80
	height = 15
)
type Universe [][]bool

func NewUniverse() Universe{
	universe := make(Universe, height)
	for i := range universe{
		universe[i] = make([]bool, width)
	}
	return universe
}
func (u Universe) String() string{
	var b byte
	buf := make([]byte, 0, (width+1)*height)

	for y :=0; y < height; y++{
		for x := 0; x< width; x++{
			b = ' '
			if u[y][x]{
				b = '*'
			}
			buf = append(buf, b)
		}
		buf = append(buf, '\n')
	}
	return string(buf)
}
func (u Universe) Show(){
	fmt.Print("\x1b[2J",u.String())
	fmt.Println("-------------------------------------------------------------------------")
}
func (u Universe) Set(x, y int, state bool){
	u[y][x] = state
}
func (u Universe) Seed(){
	for row := range u{
		for column := range u[row]{
			if cell:=rand.Intn(4); cell % 4 == 0{
				u.Set(column, row, true)
			}
		}
	}
}
func (u Universe) Alive(x, y int) bool{
	return u[(y + height) % height][(x + width) % width]
}
func (u Universe) Neighbors(x, y int) int{
	step := []int{0, -1 ,1}
	count := 0
	for i := range step{
		for j := range step{
			if !(i == 0 && j ==0) && u.Alive(x+j, y+i){
				count++
			}
		}
	}
	return count
}
func (u Universe) Next(x, y int) bool{
	count := u.Neighbors(x, y)
	return count == 3 || count == 2 && u.Alive(x, y)
}
func Step(a, b Universe){
	for y :=0; y < height; y++{
		for x := 0; x < width; x++{
			b.Set(x, y, a.Next(x, y))
		}
	}
}

func main(){
	a, b := NewUniverse(), NewUniverse()
	a.Seed()

	for i := 0; i < 300; i++ {
		Step(a, b)
		a.Show()
		time.Sleep(time.Second)
		a, b = b, a
	}
}
