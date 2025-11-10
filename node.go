package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/colornames"
)

type Node struct {
	// 当前长度与角度
	Len   float32
	Angle float64 // 最开始初始化必须是符合角度限制的
	// 最小与最大角度限制，相对于父节点的      0 0 相当于没有限制
	MinAngle, MaxAngle float64
	Next               *Node
	// 当前位置信息用于缓存更新
	BaseX, BaseY float32
	BaseAngle    float64
}

func (n *Node) UpdateBase(x float32, y float32, angle float64) {
	n.BaseX = x
	n.BaseY = y
	n.BaseAngle = angle
	angle += n.Angle
	x += float32(math.Cos(angle)) * n.Len
	y += float32(math.Sin(angle)) * n.Len
	if n.Next != nil {
		n.Next.UpdateBase(x, y, angle)
	}
}

func (n *Node) GetEndPos() (float32, float32) {
	if n.Next != nil {
		return n.Next.GetEndPos()
	}
	return n.BaseX + float32(math.Cos(n.BaseAngle+n.Angle))*n.Len, n.BaseY + float32(math.Sin(n.BaseAngle+n.Angle))*n.Len
}

// Calculate 调用前一定已经 UpdateBase 过位置了
func (n *Node) Calculate(x float32, y float32) {
	angle := math.Atan2(float64(y-n.BaseY), float64(x-n.BaseX))
	if n.Next == nil { // 最后一个
		n.AdjustAngle(angle - n.BaseAngle)
	} else { // 非最后一个
		n.Next.Calculate(x, y) // 先计算孩子节点
		ex, ey := n.GetEndPos()
		offset := angle - math.Atan2(float64(ey-n.BaseY), float64(ex-n.BaseX))
		n.AdjustAngle(n.Angle + offset)
	}
	n.UpdateBase(n.BaseX, n.BaseY, n.BaseAngle) // 触发位置的重新计算
}

func (n *Node) AdjustAngle(angle float64) {
	if n.MinAngle != 0 || n.MaxAngle != 0 {
		// 调整到  -Pi ~ Pi
		for angle > math.Pi {
			angle -= math.Pi * 2
		}
		for angle < -math.Pi {
			angle += math.Pi * 2
		}
		// 限制角度
		if angle < n.MinAngle {
			angle = n.MinAngle
		}
		if angle > n.MaxAngle {
			angle = n.MaxAngle
		}
	}
	n.Angle = angle
}

func (n *Node) Draw(screen *ebiten.Image, x float32, y float32, angle float64) {
	// 绘制可动范围
	if n.MinAngle != 0 || n.MaxAngle != 0 {
		DrawFan(screen, x, y, angle+n.MinAngle, angle+n.MaxAngle)
	}
	// 绘制水平参考线
	x0 := x + float32(math.Cos(angle+math.Pi/2))*100
	y0 := y + float32(math.Sin(angle+math.Pi/2))*100
	x1 := x + float32(math.Cos(angle-math.Pi/2))*100
	y1 := y + float32(math.Sin(angle-math.Pi/2))*100
	vector.StrokeLine(screen, x0, y0, x1, y1, 3, colornames.Red, false)
	// 计算信息
	angle += n.Angle
	tx := x + float32(math.Cos(angle))*n.Len
	ty := y + float32(math.Sin(angle))*n.Len
	// 绘制箭头
	vector.StrokeLine(screen, x, y, tx, ty, 3, colornames.Blue, false)
	ax := tx + float32(math.Cos(angle+math.Pi*5/6))*30
	ay := ty + float32(math.Sin(angle+math.Pi*5/6))*30
	vector.StrokeLine(screen, tx, ty, ax, ay, 3, colornames.Blue, false)
	ax = tx + float32(math.Cos(angle-math.Pi*5/6))*30
	ay = ty + float32(math.Sin(angle-math.Pi*5/6))*30
	vector.StrokeLine(screen, tx, ty, ax, ay, 3, colornames.Blue, false)
	if n.Next != nil {
		n.Next.Draw(screen, tx, ty, angle)
	}
}

var (
	whiteImage = ebiten.NewImage(1, 1)
	drawOption = &ebiten.DrawTrianglesOptions{}
)

func init() {
	whiteImage.Fill(color.White)
}

func DrawFan(screen *ebiten.Image, x, y float32, start, end float64) {
	vs := make([]ebiten.Vertex, 0)
	is := make([]uint16, 0)
	vs = append(vs, ebiten.Vertex{DstX: x, DstY: y, ColorG: 1, ColorA: 1})
	idx := uint16(1)
	for angle := start; angle < end; angle += math.Pi / 36 { // 每 5 度一个
		vs = append(vs, ebiten.Vertex{DstX: x + float32(math.Cos(angle))*100, DstY: y + float32(math.Sin(angle))*100, ColorG: 1, ColorA: 1})
		is = append(is, 0, idx, idx+1)
		idx++
	}
	vs = append(vs, ebiten.Vertex{DstX: x + float32(math.Cos(end))*100, DstY: y + float32(math.Sin(end))*100, ColorG: 1, ColorA: 1})
	screen.DrawTriangles(vs, is, whiteImage, drawOption)
}
