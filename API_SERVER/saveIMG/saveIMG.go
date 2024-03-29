package saveIMG

import(
	"io"
	"os"
	"log"
	"bytes"
	"image"
	"image/jpeg"
	"image/png"

	"golang.org/x/image/draw"
)

func SaveOrigin(FName string, b *bytes.Buffer) {
	fpath := "/go/Content/ORIGIN/"+FName
	orImg, err := os.Create(fpath)
	if err != nil {
		log.Println(err)
		return
	}
	defer orImg.Close()

	io.Copy(orImg, b)
	os.Chmod(fpath, 0644)
}

func SaveThumbnail(dImg image.Image, FName string, format string) {
	var rect image.Rectangle

	maxSize := 256
	imgWH := dImg.Bounds()
	fImgY := float64(imgWH.Dy())
	fImgX := float64(imgWH.Dx())
	fMaxS := float64(maxSize)

	if imgWH.Dy() > imgWH.Dx() {
		x := int(fImgX/(fImgY/fMaxS))
		rect = image.Rect(0, 0, x, maxSize)
	} else if imgWH.Dy() < imgWH.Dx() {
		y := int(fImgY/(fImgX/fMaxS))
		rect = image.Rect(0, 0, maxSize, y)
	} else if imgWH.Dy() == imgWH.Dx() {
		rect = image.Rect(0, 0, maxSize, maxSize)
	}

	imgScale := image.NewRGBA(rect)
	draw.BiLinear.Scale(imgScale, imgScale.Bounds(), dImg, imgWH, draw.Over, nil)

	thImg, err := os.Create("/go/Content/THUMBNAIL/"+FName)
	if err != nil {
		log.Println(err)
		return
	}

	switch format {
	case "jpeg":
		err := jpeg.Encode(thImg, imgScale, &jpeg.Options{Quality: 100})
		if err != nil {
			log.Println(err)
			return
		}
	case "png":
		err := png.Encode(thImg, imgScale)
		if err != nil {
			log.Println(err)
			return
		}
	}
}
