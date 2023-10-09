package models

type Puntuacion struct{
	Valor int
}
func NuevaPuntuacion(valor int)*Puntuacion{
	return &Puntuacion{Valor: valor}
}
func (p * Puntuacion)AumentarPuntuacion(puntuacion * Puntuacion){
	puntuacion.Valor = puntuacion.Valor + 1
}