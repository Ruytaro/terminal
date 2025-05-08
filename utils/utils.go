package utils

import (
	"image/color"
	"log"
)

func MapValue(value, inMin, inMax, outMin, outMax float64) float64 {
	return outMin + (value-inMin)*(outMax-outMin)/(inMax-inMin)
}

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func RGB565ToComponents(color uint16) (r, g, b uint8) {
	r = uint8(color >> 11 & 0x1F)
	g = uint8(color >> 5 & 0x3F)
	b = uint8(color & 0x1F)
	return r, g, b
}

func ColorToComponents(color color.Color) (r, g, b int) {
	cr, cg, cb, _ := color.RGBA()
	r = int(cr >> 8)
	g = int(cg >> 8)
	b = int(cb >> 8)
	return r, g, b
}

func RGBAToRGB565(r, g, b, _ uint32) uint16 {
	r5 := uint16((r >> 11) & 0x1F) // 5 bits
	g6 := uint16((g >> 10) & 0x3F) // 6 bits
	b5 := uint16((b >> 11) & 0x1F) // 5 bits
	return (r5 << 11) | (g6 << 5) | b5
}

func ColorToRGB565(c color.Color) uint16 {
	r, g, b, _ := c.RGBA()
	r5 := uint16((r >> 11) & 0x1F) // 5 bits
	g6 := uint16((g >> 10) & 0x3F) // 6 bits
	b5 := uint16((b >> 11) & 0x1F) // 5 bits
	return (r5 << 11) | (g6 << 5) | b5
}

func RGBAtoColor(r, g, b, a uint8) color.Color {
	return color.RGBA{r, g, b, a}
}
func SplitChunks(s string, size int) []string {
	var chunks []string
	runes := []rune(s)
	for i := 0; i < len(runes); i += size {
		end := i + size
		if end > len(runes) {
			end = len(runes)
		}
		chunks = append(chunks, string(runes[i:end]))
	}
	return chunks
}
