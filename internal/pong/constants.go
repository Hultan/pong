package pong

import (
	"image/color"
)

var scoreColor = color.White
var netColor = color.RGBA{R: 200, G: 200, B: 200, A: 255}
var ballColor = color.White
var player1PaddleColor = color.RGBA{R: 0, G: 0, B: 200, A: 255}
var player2PaddleColor = color.RGBA{R: 100, G: 200, B: 0, A: 255}
var backgroundColor = color.RGBA{R: 0, G: 200, B: 200, A: 155}
var ballSpeedX = -1.0
var ballSpeedY = 0.0

const gameSpeed = 4
const paddleSpeed = 10
const paddleHeight = 100
const paddleWidth = 20
const distanceFromBorder = 30.0
const middleLineWidth = 10
const ballRadius = 10
