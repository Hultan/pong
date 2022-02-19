package pong

import (
	"image/color"
	"math"
	"strconv"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gtk"
)

// onDraw : The onDraw signal handler
func (p *Pong) onDraw(da *gtk.DrawingArea, ctx *cairo.Context) {
	p.drawBackground(da, ctx)
	p.drawNet(da, ctx)
	p.drawPlayerPaddle(da, ctx, p.p1)
	p.drawPlayerPaddle(da, ctx, p.p2)
	p.drawScore(da, ctx)
	p.drawBall(da, ctx)
}

//
// HELPER FUNCTIONS
//

// drawBackground : Draws the background
func (p *Pong) drawBackground(da *gtk.DrawingArea, ctx *cairo.Context) {
	width := float64(da.GetAllocatedWidth())
	height := float64(da.GetAllocatedHeight())
	p.setColor(ctx, backgroundColor)
	ctx.Rectangle(0, 0, width, height)
	ctx.Fill()
}

func (p *Pong) drawPlayerPaddle(_ *gtk.DrawingArea, ctx *cairo.Context, play *player) {
	p.setColor(ctx, play.paddle.color)
	ctx.Rectangle(play.X-play.paddle.width/2, play.Y-play.paddle.height/2,
		play.paddle.width, play.paddle.height)
	ctx.Fill()
}

func (p *Pong) drawNet(_ *gtk.DrawingArea, ctx *cairo.Context) {
	p.setColor(ctx, netColor)
	ctx.Rectangle(p.width/2-middleLineWidth/2, 0.0, middleLineWidth, p.height)
	ctx.Fill()
}

func (p *Pong) drawBall(_ *gtk.DrawingArea, ctx *cairo.Context) {
	p.setColor(ctx, ballColor)
	ctx.Arc(p.X, p.Y, ballRadius, 0, math.Pi*2)
	ctx.Fill()
}

func (p *Pong) drawScore(_ *gtk.DrawingArea, ctx *cairo.Context) {
	p.setColor(ctx, scoreColor)
	ctx.SelectFontFace("C059 Bold", cairo.FONT_SLANT_NORMAL, cairo.FONT_WEIGHT_BOLD)
	ctx.SetFontSize(70)
	ctx.MoveTo(p.width/4, 70)
	ctx.ShowText(strconv.Itoa(p.p1.score))
	ctx.MoveTo(3*p.width/4, 70)
	ctx.ShowText(strconv.Itoa(p.p2.score))
}

func (p *Pong) setColor(ctx *cairo.Context, c color.Color) {
	r, g, b, a := c.RGBA()
	ctx.SetSourceRGBA(col(r), col(g), col(b), col(a))
}
