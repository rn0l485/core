package photo


import (
	"os"
	"image"
	"color"
	_ "image/png"
	_ "image/jpeg"
)



func ReadPhoto(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err !=nil{
		 return nil,err
	}
	fi, _ := f.Stat()

	defer f.Close()

	img, format, err := image.Decode(f)
	if err != nil {
		 return nil,err
	}
	if format == "jpeg" || format == "png" {
		return img, nil
	} else {
		return nil, errors.New("format-error")
	}
}

func Converting2Tensor(img image.Image) [][]color.Color {
	size:= img.Bounds().Size()
	var pixels [][]color.Color
	for i:=0; i<size.X; i++ {
		var y []color.Color
		for j:=0; j<size.Y;j++{
			y = append(y,img.At(i,j))
		}
		pixels = append(pixels,y)
	}
	return pixels
}

func Converting2Image(pixels [][]color.Color) image.RGBA {
	rect := image.Rect(0,0,len(pixels),len(pixels[0]))
	nImg := image.NewRGBA(rect)

	for x:=0; x<len(pixels); x++ {
		if q:=pixels[x]; q == nil {
			continue
		}
		
		for y:=0; y<len(pixels[0]); y++ {
			
			p := pixels[x][y]
			if p == nil {
				continue
			}
			
			if original,ok := color.RGBAModel.Convert(p).(color.RGBA); ok {
				nImg.Set(x,y,original)
			}
		}
	}
}