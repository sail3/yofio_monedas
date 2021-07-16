package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/credit-assigment", CreditAssignerHandler)

	r.Run(":8080")
}

type InvestmentPayload struct {
	Investment int32 `json:"investment"`
}

func CreditAssignerHandler(c *gin.Context) {
	investmentPayload := new(InvestmentPayload)
	err := c.BindJSON(&investmentPayload)
	fmt.Println(investmentPayload)
	if err != nil {
		panic(err)
	}
	cm := CreditManager{
		CreditLine:    []int32{700, 500, 300},
		wasCalculated: make([]int32, 0),
	}
	ct300, ct500, ct700, err := cm.CreditAssigner(investmentPayload.Investment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"credit_type_300": ct300,
		"credit_type_500": ct500,
		"credit_type_700": ct700,
	})
}

type CreditManager struct {
	// lineas de credito
	CreditLine []int32
	// distribucion calculada
	wasCalculated []int32
}

type CreditAssigner interface {
	Assign(investment int32) (int32, int32, int32, error)
}

func (cm *CreditManager) CreditAssigner(investment int32) (int32, int32, int32, error) {
	cm.computeCredit(investment, make([]int32, 0), 0)
	if len(cm.wasCalculated) <= 0 {
		return 0, 0, 0, fmt.Errorf("no se ha podido asignar el total de los recursos")
	}
	var linea300, linea500, linea700 int32 = 0, 0, 0
	for _, value := range cm.wasCalculated {
		switch value {
		case 300:
			linea300++
		case 500:
			linea500++
		case 700:
			linea700++
		}
	}
	return linea300, linea500, linea700, nil
}

func (cm *CreditManager) computeCredit(inversion int32, prestamos []int32, suma int32) {
	if suma == inversion && len(cm.wasCalculated) == 0 {
		cm.wasCalculated = prestamos
	} else if suma < inversion {
		for i := 0; i < len(cm.CreditLine); i++ {
			prestamos = append(prestamos, cm.CreditLine[i])
			suma += cm.CreditLine[i]
			cm.computeCredit(inversion, prestamos, suma)
			suma -= cm.CreditLine[i]
			prestamos = prestamos[:len(prestamos)-1]
		}
	}
}
