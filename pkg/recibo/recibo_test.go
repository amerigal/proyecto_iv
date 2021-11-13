package recibo

import (
	"testing"
	"time"
)

func TestCantidadesNoNulas(t *testing.T) {
	articulo, _ := NewArticulo("descripcion", "tipo", 2.50, 'A')
	articuloRecibo := ArticuloRecibo{0, articulo}
	_, err := NewRecibo([]ArticuloRecibo{articuloRecibo}, time.Now(), "u", "l", "e")
	if err == nil {
		t.Fatalf("Positividad de las cantidades no validada.")
	}
}

func TestFechaNoFutura(t *testing.T) {
	articulo, _ := NewArticulo("descripcion", "tipo", 2.50, 'A')
	articuloRecibo := ArticuloRecibo{1, articulo}
	fecha := time.Now().Add(time.Hour * 24)
	_, err := NewRecibo([]ArticuloRecibo{articuloRecibo}, fecha, "u", "l", "e")
	if err == nil {
		t.Fatalf("Validez de la fecha no comprobada.")
	}
}
