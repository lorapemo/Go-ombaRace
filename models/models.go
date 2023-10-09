package models

import(
	"os"
	"image"
	_ "image/png"
	"github.com/faiface/pixel"
)
func CreateSprite(ruta string) *pixel.Sprite{
	spriteImage, err :=loadPicture(ruta)
    if err != nil {
        panic(err)
    }
    sprite:=pixel.NewSprite(spriteImage, spriteImage.Bounds())
	return sprite
}
func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}


