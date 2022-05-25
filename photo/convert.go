package photo


import (
	"os"
	"image"
	"color"
	_ "image/png"
	_ "image/jpeg"
)

func GrayWeight(r, b, g uint32) color.Gray {
	//oldPixel := img.At(x, y)
	//pixel := color.GrayModel.Convert(oldPixel)

	lum := (19595*r + 38470*g + 7471*b + 1<<15) >> 24
	return color.Gray{uint8(lum)}
}

