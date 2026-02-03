package main

import (
	"fmt"
	"time"
)

type Tanks struct {
	Id    uint8
	Name  string
	Power int
	HP    int
}

func battle(pilihan int, tanks []Tanks) {

	fmt.Println("Anda Memilih :", tanks[pilihan-1].Name, ", Musuh anda adalah T34, Lawan ? \n1. Ya , 2. Mundur")
	var confirm int
	fmt.Scan(&confirm)

	if confirm == 1 {

		index := pilihan - 1
		fmt.Println("Bersiaplah, Anda adalah Komandan", tanks[index].Name)

		time.Sleep(1 * time.Second)

		var (
			myName     string = tanks[index].Name
			enemyName  string = tanks[1].Name
			enemyPower int    = tanks[1].Power
			myPower    int    = tanks[index].Power
			enemyHP    int    = tanks[1].HP
			myHP       int    = tanks[index].HP
		)

		myFinalHP := myHP - enemyPower
		enemyFinalHP := enemyHP - myPower

		if myFinalHP > 0 && enemyFinalHP <= 0 {
			fmt.Println("Anda Menembus Baja", enemyName, "Enemy HP : ", enemyHP-myPower)
		} else if myFinalHP <= 0 && enemyFinalHP > 0 {
			fmt.Println("Musuh Menembus Baja", myName, "Your HP : ", myHP-enemyPower)
		} else if myFinalHP > 0 && enemyFinalHP > 0 {
			fmt.Println("Kedua Tank Berdiri Kokoh")
		} else {
			fmt.Println("Kedua Tank Hancur")
		}

	} else if confirm == 2 {
		fmt.Println("Anda Mundur")
	} else {
		fmt.Println("Invalid")
	}
}
