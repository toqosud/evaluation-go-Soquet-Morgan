package main

import "fmt"

func main() {
	for {
		afficherMenu()
		if choix == 0 {
			fmt.Println("Au revoir !")
			break
		}
		afficherBoisson(choix)
		obtenirPrix(argent)
	}
}

var choix int
var argent int

func afficherMenu() {
	fmt.Println("===== DISTRIBUTEUR =====")
	fmt.Println("1.  Eau = 1.00€")
	fmt.Println("2.  Soda = 2.00€")
	fmt.Println("3.  Café = 2.00€")
	fmt.Println("4.  Chocolat = 3.00€")
	fmt.Println("0.  Quitter")
	fmt.Println("========================")
	fmt.Scanln(&choix)
	fmt.Println("Votre choix :", choix)
}

func afficherBoisson(choix int) {
	switch choix {
	case 1:
		fmt.Println("Vous avez choisi : Eau.")
		fmt.Println("Prix : 1.00€")
	case 2:
		fmt.Println("Vous avez choisi : Soda.")
		fmt.Println("Prix : 2.00€")
	case 3:
		fmt.Println("Vous avez choisi : Café.")
		fmt.Println("Prix : 2.00€")
	case 4:
		fmt.Println("Vous avez choisi : Chocolat.")
		fmt.Println("Prix : 3.00€")
	case 0:
		fmt.Println("Au revoir !")
	default:
		fmt.Println("Choix invalide !")
	}
}

func obtenirPrix(argent int) int {
	fmt.Println("Insérez argent :")
	fmt.Scanln(&argent)
	fmt.Println("Montant inséré :", argent, "€")
	switch choix {
	case 1:
		if argent < 1 {
			fmt.Println("Montant insuffisant !")
			return argent
		}
		fmt.Println("Merci voici votre monnaie :", argent-1)
		return argent - 1
	case 2:
		if argent < 2 {
			fmt.Println("Montant insuffisant !")
			return argent
		}
		fmt.Println("Merci voici votre monnaie :", argent-2)
		return argent - 2
	case 3:
		if argent < 2 {
			fmt.Println("Montant insuffisant !")
			return argent
		}
		fmt.Println("Merci voici votre monnaie :", argent-2)
		return argent - 2
	case 4:
		if argent < 3 {
			fmt.Println("Montant insuffisant !")
			return argent
		}
		fmt.Println("Merci voici votre monnaie :", argent-3)
		return argent - 3
	case 0:
		return argent
	default:
		return 0
	}
}
