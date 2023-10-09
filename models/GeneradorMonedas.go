package models

import(
	"github.com/faiface/pixel"
_ "image/png"
"math/rand"
    "time"
)

type GeneradorMonedas struct{
	PosicionX float64
	PosicionY float64
	Sprite *pixel.Sprite
}
func generarNumeroAleatorio()(float64, float64){
	rand.Seed(time.Now().UnixNano())
	x1:=50.0
	x2:=950.0
	y1:=50.0
	y2:=450.0
	x:=  x1 + rand.Float64()*(x2-x1)
	y:= y1 + rand.Float64()*(y2-y1)
	return x, y
}
func NuevoGeneradorMonedas(s *pixel.Sprite)*GeneradorMonedas{
	rand.Seed(time.Now().UnixNano())
	x,y := generarNumeroAleatorio()
	return &GeneradorMonedas{PosicionX: x, PosicionY: y, Sprite: s}
}
func (generadorMonedas * GeneradorMonedas)Colision(gM *GeneradorMonedas, Px float64, Py float64, puntuacion *Puntuacion){
	if (generadorMonedas.PosicionY -70 < Py && generadorMonedas.PosicionY + 70 > Py){
		if (generadorMonedas.PosicionX -30 < Px && generadorMonedas.PosicionX + 30 > Px){
		x,y := generarNumeroAleatorio()
		generadorMonedas.PosicionX= x
		generadorMonedas.PosicionY= y
		puntuacion.AumentarPuntuacion(puntuacion)
	}
	}
}