package main

import (
	"image/color"
	"math"
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func newLayer() *Layer {
	l := &Layer{}

	l.color = randColor()

	l.vx, l.vy = int32(rand.Intn(10))-5, int32(rand.Intn(10))-5

	l.size = rand.Float32()*30 + 1

	if l.vx == 0 {
		l.vx = 1
	}

	if l.vy == 0 {
		l.vy = 1
	}

	l.x, l.y = int32(rand.Intn(1024)), int32(rand.Intn(768))

	return l
}

type Layer struct {
	vx, vy int32
	x, y   int32
	size   float32
	color  color.RGBA
}

func (l *Layer) OnRender() {
	nextX := l.x + l.vx
	nextY := l.y + l.vy

	modX := modulo(nextX, 1024)
	modY := modulo(nextY, 768)

	if nextX != modX {
		l.vx = -l.vx // hit left or right screen edge
	}

	if nextY != modY {
		l.vy = -l.vy // hit top or bottom screen edge
	}

	l.x = nextX
	l.y = nextY

	rl.DrawCircle(l.x, l.y, l.size, l.color)
}

func modulo(n, m int32) int32 {
	if m == 0 {
		panic("Modulo divisor cannot be zero.")
	}

	result := n % m

	if result < 0 {
		result += int32(math.Abs(float64(m)))
	}

	return result
}

func randColor() color.RGBA {
	rand.Seed(time.Now().UnixNano())

	r := uint8(rand.Intn(256))
	g := uint8(rand.Intn(256))
	b := uint8(rand.Intn(256))

	return color.RGBA{r, g, b, 255}
}

func randomString(length int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	rand.Seed(time.Now().UnixNano())

	bytes := make([]byte, length)
	for i := 0; i < length; i++ {
		bytes[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	return string(bytes)
}
