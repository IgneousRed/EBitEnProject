package main

import (
	"math"
	"time"

	et "github.com/IgneousRed/EduTen"
	m "github.com/IgneousRed/gomisc"
)

var windowSize = m.Vec2I(800, 600)
var font et.Font

type Game struct{}

func (g *Game) Update() {
}
func (g *Game) Draw() {
	et.DrawRectangleI(m.Vec2I(50, 100), m.Vec2I(150, 100), et.Red)
	point := windowSize.Div1(2)
	ang := float64(time.Now().UnixMilli()%10000) / 10000. * math.Pi * 2.
	other := m.Vec2F(m.Cos(ang), m.Sin(ang)).Mul1(200.).RoundI().Add(point)
	et.DrawLineI(point, other, 20, et.Green)
	et.DrawCircleI(m.Vec2I(600, 500), 50, 8, et.Blue)
	et.DrawTextI(font, 40, m.Vec2I(0, 10), "^_^", et.White)
}

func main() {
	f, err := et.FontNew("FiraCode-Medium.ttf")
	m.FatalErr("", err)
	font = f
	game := Game{}
	et.InitGame("Title", windowSize, &game)
}
