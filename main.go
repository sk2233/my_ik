package main

import (
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

// https://zhuanlan.zhihu.com/p/546514925

func main() {
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("MyIK")
	node := &Node{Len: 100 + rand.Float32()*100, MinAngle: -rand.Float64() * math.Pi, MaxAngle: rand.Float64() * math.Pi}
	tmp := node
	for i := 0; i < 4; i++ {
		tmp.Next = &Node{Len: 150 + rand.Float32()*300, MinAngle: -rand.Float64() * math.Pi, MaxAngle: rand.Float64() * math.Pi}
		tmp = tmp.Next
	}
	err := ebiten.RunGame(NewGame(120, 360, 0, node))
	HandleErr(err)
}
