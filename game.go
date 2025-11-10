package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	// 起始位置与方向
	X, Y  float32
	Angle float64
	// ik 节点
	Node *Node
}

func NewGame(x float32, y float32, angle float64, node *Node) *Game {
	return &Game{X: x, Y: y, Angle: angle, Node: node}
}

func (g *Game) Update() error {
	x, y := ebiten.CursorPosition()
	// 迭代 3 次进行调整
	for i := 0; i < 3; i++ {
		g.Node.UpdateBase(g.X, g.Y, g.Angle)
		g.Node.Calculate(float32(x), float32(y))
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Node.Draw(screen, g.X, g.Y, g.Angle)
}

func (g *Game) Layout(w, h int) (int, int) {
	return w, h
}
