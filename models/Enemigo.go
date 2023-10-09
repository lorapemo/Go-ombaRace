package models

import(
	"github.com/faiface/pixel"
_ "image/png"
)

type Enemigo struct{
	PosicionX float64
	PosicionY float64
	Sprite *pixel.Sprite
}
func NuevoEnemigo(x float64, y float64, s *pixel.Sprite)*Enemigo{
	return &Enemigo{PosicionX: x, PosicionY: y, Sprite: s}
}
func (e * Enemigo)Movimiento(enemigo *Enemigo, MPx float64, MPy float64){
	if (enemigo.PosicionX < MPx){
		enemigo.PosicionX = enemigo.PosicionX + .5
	}
	if (enemigo.PosicionX > MPx){
		enemigo.PosicionX = enemigo.PosicionX - .5
	}
	if (enemigo.PosicionY < MPy){
		enemigo.PosicionY = enemigo.PosicionY + .5
	}
	if (enemigo.PosicionY > MPy){
		enemigo.PosicionY = enemigo.PosicionY - .5
	}
}

