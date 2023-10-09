package views
import(
"github.com/faiface/pixel"
"github.com/faiface/pixel/pixelgl"


"proyecto/models"
"fmt"
)
func Principal(){
	    cfg := pixelgl.WindowConfig{
        Title:  "Mi Juego",
        Bounds: pixel.R(0, 0, 1000, 500), // Tamaño de la ventana 500x500
    }
    win, err := pixelgl.NewWindow(cfg)
    if err != nil {
        panic(err)
    }
    
    //enemy := models.EnemySpawn()
    //enemy.Draw(win, pixel.IM.Moved(win.Bounds().Center()))
    enemigo := models.NuevoEnemigo(500,250, models.CreateSprite("./assets/enemy.png"))
    enemigo.Sprite.Draw(win, pixel.IM.Moved(pixel.V(enemigo.PosicionX, enemigo.PosicionY)))
    //jugadorPosicion := pixel.V(50,250)
    //jugador := models.JugadorSpawn()
    //jugador.Draw(win, pixel.IM.Moved(jugadorPosicion))
    jugador := models.NuevoJugador(50,250, models.CreateSprite("./assets/jugador.png"))
    jugador.Sprite.Draw(win, pixel.IM.Moved(pixel.V(jugador.PosicionX, jugador.PosicionY)))

    generadorMonedas:= models.NuevoGeneradorMonedas(models.CreateSprite("./assets/moneda.png"))
    generadorMonedas.Sprite.Draw(win, pixel.IM.Moved(pixel.V(generadorMonedas.PosicionX, generadorMonedas.PosicionY)))

    puntosJugador, puntosEnemigo := models.NuevaPuntuacion(0), models.NuevaPuntuacion(0)

    for !win.Closed() {
        // Aquí puedes agregar la lógica de tu juego y la representación visual
        win.Clear(pixel.RGB(1, 1, 1))

        go jugador.Movimiento(jugador,win)
        jugador.Sprite.Draw(win, pixel.IM.Moved(pixel.V(jugador.PosicionX, jugador.PosicionY)))

        go enemigo.Movimiento(enemigo, generadorMonedas.PosicionX, generadorMonedas.PosicionY)
        enemigo.Sprite.Draw(win, pixel.IM.Moved(pixel.V(enemigo.PosicionX, enemigo.PosicionY)))

        go generadorMonedas.Colision(generadorMonedas, jugador.PosicionX, jugador.PosicionY, puntosJugador) 
        go generadorMonedas.Colision(generadorMonedas, enemigo.PosicionX, enemigo.PosicionY, puntosEnemigo)
        generadorMonedas.Sprite.Draw(win, pixel.IM.Moved(pixel.V(generadorMonedas.PosicionX, generadorMonedas.PosicionY)))

        win.Update()
        
    }
    fmt.Println("puntosEnemigo: ")
    fmt.Println(puntosEnemigo.Valor)
    fmt.Println("puntosJugador: ")
    fmt.Println(puntosJugador.Valor)
}

 func V() {
	pixelgl.Run(Principal)
    }

    