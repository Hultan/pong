package pong

// Convert 0-65535 color to 0-1 color
func col(c uint32) float64 {
	return float64(c) / 65535
}
