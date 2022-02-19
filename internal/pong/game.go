package pong

import (
	"fmt"
	"image/color"
	"time"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

func NewPong(w *gtk.ApplicationWindow, da *gtk.DrawingArea) *Pong {
	t := &Pong{window: w, drawingArea: da}
	t.window.Connect("key-press-event", t.onKeyPressed)
	t.drawingArea.Connect("size-allocate", t.onResize)
	// t.window.Connect("button-press-event", t.onKeyPressed)
	// t.window.Connect("event", t.onKeyPressed)

	return t
}

func (p *Pong) StartGame() {
	p.isActive = true
	p.speed = gameSpeed
	p.width = float64(p.drawingArea.GetAllocatedWidth())
	p.height = float64(p.drawingArea.GetAllocatedHeight())
	p.p1 = p.setupPlayer(distanceFromBorder, player1PaddleColor)
	p.p2 = p.setupPlayer(p.width-distanceFromBorder, player2PaddleColor)
	p.newBallPosition()
	p.drawingArea.Connect("draw", p.onDraw)

	p.ticker.ticker = time.NewTicker(p.speed * time.Millisecond)
	p.ticker.tickerQuit = make(chan struct{})

	go p.mainLoop()
}

func (p *Pong) mainLoop() {
	for {
		select {
		case <-p.ticker.ticker.C:
			p.X += ballSpeedX
			p.Y += ballSpeedY
			p.checkScore()
			p.checkBounce()
			p.checkPlayerBounce()
			p.drawingArea.QueueDraw()
		case <-p.ticker.tickerQuit:
			p.isActive = false
			p.ticker.ticker.Stop()
			return
		}
	}
}

// onKeyPressed : The onKeyPressed signal handler
func (p *Pong) onKeyPressed(_ *gtk.ApplicationWindow, e *gdk.Event) {
	key := gdk.EventKeyNewFromEvent(e)

	fmt.Println(key.KeyVal())

	switch key.KeyVal() {
	case 113: // Button "Q" => Quit game
		p.quit()
		p.window.Close() // Close window
	case 97: // A - P1 paddle up
		if p.p1.Y > p.p1.height/2 {
			p.p1.Y -= paddleSpeed
		}
	case 122: // Z - P1 paddle down
		if p.p1.Y < p.height-p.p1.height/2 {
			p.p1.Y += paddleSpeed
		}
	case 65362: // Arrow up
		if p.p2.Y > p.p2.height/2 {
			p.p2.Y -= paddleSpeed
		}
	case 65364: // Arrow down
		if p.p2.Y < p.height-p.p1.height/2 {
			p.p2.Y += paddleSpeed
		}
	}
	p.drawingArea.QueueDraw()
}

func (p *Pong) checkScore() bool {
	if p.X < 0 {
		p.p2.score++
		p.newBallPosition()
		return true
	}
	if p.X > p.width {
		p.p1.score++
		p.newBallPosition()
		return true
	}
	return false
}

func (p *Pong) checkBounce() {
	if p.Y < 0 {
		// Bounce ball
		ballSpeedY = -ballSpeedY
	}
	if p.Y > p.height {
		// Bounce ball
		ballSpeedY = -ballSpeedY
	}
}

func (p *Pong) checkPlayerBounce() {
	if p.X > distanceFromBorder && p.X < distanceFromBorder+paddleWidth {
		if p.Y > p.p1.Y-p.p1.paddle.height/2 && p.Y < p.p1.Y+p.p1.paddle.height/2 {
			// Bounce player 1
			ballSpeedX = -ballSpeedX
			ballSpeedY = -2 * (p.p1.Y - p.Y) / p.p1.paddle.height
			p.X = distanceFromBorder + paddleWidth
		}
	}
	if p.X > p.width-distanceFromBorder-paddleWidth && p.X < p.width-distanceFromBorder {
		if p.Y > p.p2.Y-p.p2.paddle.height/2 && p.Y < p.p2.Y+p.p2.paddle.height/2 {
			// Bounce player 2
			ballSpeedX = -ballSpeedX
			ballSpeedY = -2 * (p.p2.Y - p.Y) / p.p2.paddle.height
			p.X = p.width - distanceFromBorder - paddleWidth
		}
	}
}

func (p *Pong) newBallPosition() {
	p.X = p.width / 2
	p.Y = p.height / 2
	ballSpeedX = -1
	ballSpeedY = 0
}

func (p *Pong) quit() {
	if p.isActive {
		p.isActive = false
		close(p.ticker.tickerQuit) // Stop ticker
	}
}

func (p *Pong) setupPlayer(x float64, c color.Color) *player {
	play := &player{
		paddle: paddle{
			position: position{X: x, Y: p.height / 2},
			size:     size{width: paddleWidth, height: paddleHeight},
			color:    c,
		},
	}
	return play
}

func (p *Pong) onResize(da *gtk.DrawingArea) {
	// Save old size
	h := p.height
	w := p.width

	// Get new size
	p.width = float64(da.GetAllocatedWidth())
	p.height = float64(da.GetAllocatedHeight())

	// Calculate resize factors
	fy := p.height / h
	fx := p.width / w

	// Adjust ballSpeedX
	ballSpeedX = ballSpeedX * fx

	// Adjust ball x and y position
	p.Y = p.Y * fy
	p.X = p.X * fx

	// Adjust y position of player 1 and 2
	p.p1.Y = p.p1.Y * fy
	p.p2.Y = p.p2.Y * fy

	// Adjust x position of player 2
	p.p2.X = p.width - distanceFromBorder
}
