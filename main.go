package main

import (
	"math"
	"time"

	eb "github.com/IgneousRed/EBitEn"
	m "github.com/IgneousRed/gomisc"
)

var windowSize = m.Vec2I(800, 600)
var font eb.Font

type Game struct{}

func (g *Game) Update() {
}
func (g *Game) Draw() {
	eb.DrawRectangleI(m.Vec2I(50, 100), m.Vec2I(150, 100), eb.Red)
	point := windowSize.Div1(2)
	ang := float64(time.Now().UnixMilli()%10000) / 10000. * math.Pi * 2.
	other := m.Vec2F(m.Cos(ang), m.Sin(ang)).Mul1(200.).RoundI().Add(point)
	eb.DrawLineI(point, other, 20, eb.Green)
	eb.DrawCircleI(m.Vec2I(600, 500), 50, 8, eb.Blue)
	eb.DrawTextI(font, 40, m.Vec2I(0, 10), "^_^", eb.White)
}

func main() {
	f, err := eb.FontNew("FiraCode-Medium.ttf")
	m.FatalErr("", err)
	font = f
	game := Game{}
	eb.InitGame("Title", windowSize, &game)
}
