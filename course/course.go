package course

import "fmt"

type Curse struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIDs []uint
	Classes map[uint]string
}

/*
* PrintClasses	: Imprime las clases que tiene el curso
* Return		: void
* Author		: Yeison Medina
* Date			: 2023-09-10
 */
func (c Curse) PrintClasses() {
	text := "Las clases son estas:"
	for _, class := range c.Classes {
		text += class + ", "
	}
	fmt.Println(text)
}

/*
 * ChangePrice	: Cambia el valor de la variable Price
 * Param price	: float64
 * Return		: void
 * Author		: Yeison Medina
 */
func (c *Curse) ChangePrice(price float64) {
	c.Price = price
}
