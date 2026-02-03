package main

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
)

func main() {
	myFigure := figure.NewColorFigure("PMC AEVIATHAN", "small", "blue", true)

	myFigure.Blink(500, 20, 20)
	myFigure.Print()
	color.Blue("%s", "><<<<<<><<<<<<><<<<<<><<<<<<><<<<<<><<<<<<><<<<<<><<<<<<><<<<<<><<<<<<><<<<<<><<<<<<><<<<<<")
	fmt.Println(">>> SYS  : ЧВК АЭВИАФАН (PMC AEVIATHAN)		Url : Nawvalovsky.vercel.app")
	color.White(">>> TOOL : AVT PanzerATK V0			By  : Nawvalovsky")
	color.Blue("%s", "><<<<<<><<<<<<><<<<<<><<<<<<><<<<<<><<<<<<><<<<<<><<<<<<><<<<<<><<<<<<><<<<<<><<<<<<><<<<<<")

	var listank = []Tanks{
		{Id: 1, Name: "JagdPanzer", Power: 150, HP: 200},
		{Id: 2, Name: "T34", Power: 100, HP: 80},
		{Id: 3, Name: "PZ_IV", Power: 120, HP: 110},
	}
	var pilihan int

	fmt.Println("Menu Tank")
	fmt.Println(listank)

	fmt.Print("Pilihan > ")
	fmt.Scan(&pilihan)

	if pilihan >= 1 && pilihan <= len(listank) {
		battle(pilihan, listank)
	} else {
		fmt.Println("Invalid")
	}

}
