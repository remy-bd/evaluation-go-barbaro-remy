package main

import "fmt"

func main() {
	for {
		afficherMenu() //afficher le menu des boissons disponible

		var choix int
		fmt.Print("Votre choix : ")
		fmt.Scan(&choix) //récupère le choix de l'utilisateur

		if choix == 0 {
			fmt.Println("Au revoir")
			break //Sort de la boucle si on choisi 0
		}

		if choix < 0 || choix > 4 {
			fmt.Println("Choix invalide !")
			fmt.Println()
			continue //si l'utilisateur rentre un nombre invalide on recommence
		}

		fmt.Println()
		afficherBoisson(choix)
		prix := obtenirPrix(choix)
		fmt.Println("Prix :", prix, "EUR")
		// on affiche le prix de la boisson que l'utilisateur a choisi

		var montant int
		fmt.Print("Inséré des fonds : ")
		fmt.Scan(&montant)
		fmt.Println("Montant inséré:", montant, "$")
		//récupère les fonds que l'utilisateur a inséré

		if montant < prix {
			manque := prix - montant
			fmt.Println("Montant insuffisant !")
			fmt.Println("Il manque", manque, "EUR.")
			fmt.Println()
			continue
			//on vérifie si il y assez de monnaie

		}

		rendu := montant - prix
		fmt.Println()
		fmt.Println("Merci !")
		fmt.Println("Votre monnaie :", rendu, "EUR")
		fmt.Println()
		//on calcule la monnaie a rendre si il y en a

	}
}

func afficherMenu() {
	fmt.Println("=== DISTRIBUTEUR ===")
	fmt.Println()
	fmt.Println("1 - Eau       : 1 EUR")
	fmt.Println("2 - Soda      : 2 EUR")
	fmt.Println("3 - Cafe      : 2 EUR")
	fmt.Println("4 - Chocolat  : 3 EUR")
	fmt.Println("0 - Quitter")
	fmt.Println()
	//menu avec les différente boissons

}

func obtenirPrix(choix int) int {
	if choix == 1 {
		return 1
	} else if choix == 2 {
		return 2
	} else if choix == 3 {
		return 2
	} else if choix == 4 {
		return 3
	} else {
		return 0
	}
} //on associe le prix a chacune des boissons

func afficherBoisson(choix int) {
	if choix == 1 {
		fmt.Println("Vous avez choisi : Eau")
	} else if choix == 2 {
		fmt.Println("Vous avez choisi : Soda")
	} else if choix == 3 {
		fmt.Println("Vous avez choisi : Cafe")
	} else if choix == 4 {
		fmt.Println("Vous avez choisi : Chocolat")
	} else {
		fmt.Println("Choix invalide !")
	}
} //Affiche la boisson correspondante au choix de l'utilisateur
