package main

import (
	"time"

	et "github.com/IgneousRed/EduTen"
	m "github.com/IgneousRed/gomisc"
)

type rad = m.Rad
type v2 = m.Vector2

var V2 = m.Vec2

type Image = et.Image

var font et.Font

type Game struct{}

func (g *Game) Update() {
}
func (g *Game) Draw(scr *Image) {
	et.DrawRectangle(scr, V2(50, 100), V2(150, 100), et.Red)
	ang := rad(time.Now().UnixMilli()%10000) / 10000 * m.Tau
	other := ang.Vec2().Mul1(200.).Add(et.WindowHalf())
	et.DrawLine(scr, et.WindowHalf(), other, 20, et.Green)
	et.DrawCircle(scr, V2(600, 500), 50, 8, et.Blue)
	et.DrawText(scr, font, 40, V2(0, 10), "^_^", et.White)
}

func main() {
	f, err := et.FontNew("FiraCode-Medium.ttf")
	m.FatalErr("", err)
	font = f
	et.WindowSizeSet(V2(800, 600))
	game := Game{}
	et.Run(&game)
}
