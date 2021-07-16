package main

import (
	"testing"
)

func TestCreditAssignerStruct(t *testing.T) {
	cm := CreditManager{
		CreditLine:    []int32{700, 500, 300},
		wasCalculated: make([]int32, 0),
	}
	montoInversiones := []int32{3000, 6000}
	for _, montoInversion := range montoInversiones {
		cm.wasCalculated = make([]int32, 0)
		line300, line500, line700, err := cm.CreditAssigner(montoInversion)
		if err != nil {
			t.Fatalf("error en la asignacion de de credito.")
		}

		distribucionOntenida := (line300*300 + line500*500 + line700*700)
		if distribucionOntenida != montoInversion {
			t.Fatalf("error en el calculo de lineas de prestamo, se esperaba %d se tiene %d", montoInversion, distribucionOntenida)
		}
	}
	montoInversion := int32(400)
	cm.wasCalculated = make([]int32, 0)
	_, _, _, err := cm.CreditAssigner(montoInversion)
	if err == nil {
		t.Fatalf("error en la asignacion de de credito.")
	}

}
