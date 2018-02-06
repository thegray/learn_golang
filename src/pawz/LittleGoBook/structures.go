package main

import (
	"fmt"
)

//field in structure
type Saiyan struct {
	Name   string
	Power  int
	Father *Saiyan
}

//function in structure
func (s *Saiyan) SuperSaiyan() {
	s.Power += 10000
	fmt.Println("power after supersaiyan: ", s.Power)
}

//composition
type SaiyanPeople struct {
	*Saiyan
	Address string
}

//overloading
func (sp *SaiyanPeople) SuperSaiyan() {
	sp.Power += 15000
	fmt.Println("supersaiyan special: ", sp.Power)
}

func main() {
	gohan := &Saiyan{
		Name:  "Gohan",
		Power: 1000,
		Father: &Saiyan{
			Name:   "Goku",
			Power:  9000,
			Father: nil,
		},
	}

	fmt.Println("gohan power: ", gohan.Power)
	gohan.SuperSaiyan()

	//composition in golang
	raditz := &SaiyanPeople{
		Saiyan: &Saiyan{
			Name:   "Raditz",
			Power:  8000,
			Father: nil,
		},
		Address: "Planet Saiya",
	}

	fmt.Println("raditz power: ", raditz.Power)
	raditz.SuperSaiyan()

	fmt.Println(raditz.Name)
	fmt.Println(raditz.Saiyan.Name)

	//overloading
	raditz.SuperSaiyan()        //call the overloaded function
	raditz.Saiyan.SuperSaiyan() //call the composed function
}
