package main

import "fmt"

func main() {
	var nb int
	var note float64
	somme := 0.0
	notes := []float64{}
	//création d'une variable et d'une liste

	fmt.Println("=== GESTIONNAIRE DE NOTES ===")
	fmt.Println()
	fmt.Print("Entrez un nombre de notes :")
	fmt.Scanln(&nb)
	for i := 1; i <= nb; i++ {
		fmt.Println("Note", i, ":")
		fmt.Scanln(&note)
		//affiche note 1 :..., note 2 :..., etc

		if note < 0 || note > 20 {
			break
		}
		// vérifie que la note est comprise entre 0 et 20 sinon il relance le programme

		notes = append(notes, note)
		somme += note
		//ajout des notes dans la liste
	}

	if len(notes) == 0 {
		fmt.Println("Aucune note valide.")
		return
	} // au cas ou on ne rentre aucune note

	moyenne := calculerMoyenne(somme, len(notes))
	nmax := trouverMaximum(notes)
	nmin := trouverMinimum(notes)
	fmt.Println("=== RÉSULTATS ===")
	fmt.Println()
	fmt.Println("Moyenne :", moyenne)
	fmt.Println("Note maximum :", nmax)
	fmt.Println("Note minimum :", nmin)
	afficherResultat(moyenne)
	// calcul de la moyenne, nmax et nmin puis afficher tout

	//BONUS
	fmt.Printf("Nombre au dessu de la moyenne :%.2f\n ", compterNotesAuDessusDeLaMoyenne(notes, moyenne))
	AfficherApprciation(moyenne)

}

func calculerMoyenne(somme float64, nb int) float64 {
	return float64(somme) / float64(nb)
} // calcul de la moyenne

func trouverMaximum(notes []float64) float64 {
	c2 := notes[0]
	for _, note := range notes {
		if note > c2 {
			c2 = note
		}
	}
	return c2
} //recherche de la valeur maximum

func trouverMinimum(notes []float64) float64 {
	c3 := notes[0]
	for _, note := range notes {
		if note < c3 {
			c3 = note
		}
	}
	return c3
} //recherche de la valeur minimum

func afficherResultat(moyenne float64) {
	if moyenne >= 10 {
		fmt.Println("Vous êtes admis !")
	} else {
		fmt.Println("Non admis...")
	}
} //afficher si on est admis ou non selon la moyenne des notes

// BONUS
func compterNotesAuDessusDeLaMoyenne(notes []float64, moyenne float64) float64 {
	compteur := 0.0
	for _, note := range notes {
		if note >= moyenne {
			compteur++
		}
	}
	return compteur
}

func AfficherApprciation(moyenne float64) {
	if moyenne <= 10 {
		fmt.Println("Appréciation : insuffisante")
	} else if moyenne <= 12 {
		fmt.Println("Appréciation : passable")
	} else if moyenne <= 14 {
		fmt.Println("Appréciation : assez bien")
	} else if moyenne <= 16 {
		fmt.Println("Appréciation : bien")
	} else {
		fmt.Println("Appréciation : Très bien")
	}
}
