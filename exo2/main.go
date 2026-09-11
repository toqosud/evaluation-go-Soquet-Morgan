package main

import "fmt"

var choix int

func afficherMenu() {
	fmt.Scanln(&choix)
	fmt.Println("===== DISTRIBUTEUR =====")
	fmt.Println("1. Eau          = 1.00€")
	fmt.Println("2. Soda 	     = 2.00€")
	fmt.Println("3. Café		 = 2.00€")
	fmt.Println("4. Chocolat	 = 3.00€")
	fmt.Println("0. Quitter")
	fmt.Println("========================")
	fmt.Println("Votre choix :")
	fmt.Println("/n")
}
