package main

import (
	"time"

	et "github.com/IgneousRed/EduTen"
	m "github.com/IgneousRed/gomisc"
)

type Rad = et.Rad
type Vec2 = et.Vec2
type Image = et.Image

var font et.Font

type Game struct{}

func (g *Game) Update() {
}
func (g *Game) Draw(scr *Image) {
	et.DrawRectangle(scr, Vec2{50, 100}, Vec2{150, 100}, et.Red)
	ang := Rad(time.Now().UnixMilli()%10000) / 10000 * m.Tau
	other := ang.Vec2().Mul1(200.).Add(et.WindowHalf())
	et.DrawLine(scr, et.WindowHalf(), other, 20, et.Green)
	et.DrawCircle(scr, Vec2{600, 500}, 50, 8, et.Blue)
	et.DrawText(scr, font, 40, Vec2{0, 10}, "^_^", et.White)
}

func main() {
	f, err := et.FontNew("FiraCode-Medium.ttf")
	m.FatalErr(err, "")
	font = f
	et.WindowSizeSet(Vec2{800, 600})
	game := Game{}
	et.InitGame("Title", &game)
}
