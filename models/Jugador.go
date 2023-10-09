package models
import(
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
_ "image/png"
)
type Jugador struct{
	PosicionX float64
	PosicionY float64
	Sprite *pixel.Sprite
}
func NuevoJugador(x float64, y float64, s *pixel.Sprite)*Jugador{
	return &Jugador{PosicionX: x, PosicionY: y, Sprite: s}
}
func (j * Jugador)Movimiento(jugador *Jugador,win *pixelgl.Window){
	if win.Pressed(pixelgl.KeyLeft) {
		jugador.PosicionX -= 1 // Mueve el sprite a la izquierda
	}
	if win.Pressed(pixelgl.KeyRight) {
		jugador.PosicionX += 1 // Mueve el sprite a la derecha
	}
	if win.Pressed(pixelgl.KeyUp) {
		jugador.PosicionY += 1 // Mueve el sprite hacia arriba
	}
	if win.Pressed(pixelgl.KeyDown) {
		jugador.PosicionY -= 1 // Mueve el sprite hacia abajo
	}
}


