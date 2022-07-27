package main

import (
	"fmt"
	"reflect"
)

type Reptile interface {
	Lay() ReptileEgg
}

type ReptileCreator func() Reptile

type ReptileEgg struct {
	CreateReptile ReptileCreator
	hatchCount    int
}

func (egg *ReptileEgg) Hatch(reptileType Reptile) Reptile {
	if egg.hatchCount > 0 {
		return nil
	} else {
		egg.hatchCount += 1
		// bagaimana cara ceknya?
		return FireDragon{}
	}
}

type FireDragon struct {
}

func (fd FireDragon) Lay() ReptileEgg {
	return ReptileEgg{}
}

type BlueDragon struct {
}

func (bd BlueDragon) Lay() ReptileEgg {
	return ReptileEgg{}
}

// other reptile dont make firedragon
func main() {
	var fireDragon FireDragon
	var egg ReptileEgg = fireDragon.Lay()
	var childDragon Reptile = egg.Hatch(fireDragon)

	fmt.Println(reflect.TypeOf(childDragon))

	var newChildDragon Reptile = egg.Hatch(fireDragon)
	fmt.Println(newChildDragon)

	var blueDragon BlueDragon
	var eggBlue ReptileEgg = blueDragon.Lay()
	var blueChildDragon Reptile = eggBlue.Hatch(blueDragon)

	fmt.Println(reflect.TypeOf(blueChildDragon))
}
