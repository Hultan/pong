package pong

import (
	"image/color"
	"time"

	"github.com/gotk3/gotk3/gtk"
)

type Pong struct {
	window      *gtk.ApplicationWindow
	drawingArea *gtk.DrawingArea

	p1 *player
	p2 *player

	game
	position
	size
}

type game struct {
	speed    time.Duration
	isActive bool
	ticker   ticker
}

type player struct {
	paddle
	score int
}

type paddle struct {
	position
	size
	color color.Color
}

type ticker struct {
	tickerQuit chan struct{}
	ticker     *time.Ticker
}

type position struct {
	X float64
	Y float64
}

type size struct {
	width  float64
	height float64
}
