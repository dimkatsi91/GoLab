package main

import (
	"fmt"
)


type Joueur struct {
	// Un joueur a un nome et des competences
	// competences est un map competence(string) -> valeur(int)
	//
	nom, equipe string
	comps map[string]int
}


// fonction attache a la struct 'Joueur'
// Imprimer de les informations du joueur
//
func (joueur *Joueur) info() {
	//var joueur_info string
	joueur_info := fmt.Sprintf("Joueur Nom: %s\nJoueur equipe: %s", joueur.nom, joueur.equipe)
	fmt.Println(joueur_info)
	for key, value := range joueur.comps {
		fmt.Printf("%s = %d\n", key, value)
	}
}


// Modele de Conception d'Usine
//
func creer_Joueur(comps map[string]int, nom, equipe string) *Joueur {
	joueur := &Joueur{}
	joueur.nom = nom
	joueur.equipe = equipe
	for key, value := range joueur.comps {
		joueur.comps[key] = value
	}
	return &Joueur{nom, equipe, comps}
}


func main() {
	// GK
	mandanda_comps := make(map[string]int)
	mandanda_comps["charactere"] = 86
	mandanda_comps["agilite"] = 80
	mandanda_comps["equilibre"] = 85
	mandanda := creer_Joueur(mandanda_comps, "Steve Mandanda", "Marseille")
	// DF
	vanBuyten_comps := make(map[string]int)
	vanBuyten_comps["defense"] = 84
	vanBuyten_comps["equilibre"] = 86
	vanBuyten_comps["tacler"] = 81
	vanBuyten := creer_Joueur(vanBuyten_comps, "Van Buyten", "Marseille")
	// MF
	lucho_comps := make(map[string]int)
	lucho_comps["vitesse"] = 78
	lucho_comps["technique"] = 80
	lucho_comps["creativite"] = 86
	lucho := creer_Joueur(lucho_comps, "Lucho Gonzalez", "Marseille")
	// CF
	didier_comps := make(map[string]int)
	didier_comps["vitesse"] = 79
	didier_comps["technique"] = 77
	didier_comps["finition"] = 88
	didier := creer_Joueur(didier_comps, "Didier Drogba", "Marseille")
	//didier.info()

	marseille := make(map[string]*Joueur)
	marseille["GK"] = mandanda
	marseille["DF"] = vanBuyten
	marseille["MF"] = lucho
	marseille["CF"] = didier
	fmt.Println("############## GK ##############")
	marseille["GK"].info()
	fmt.Println("############## DF ##############")
	marseille["DF"].info()
	fmt.Println("############## MF ##############")
	marseille["MF"].info()
	fmt.Println("############## CF ##############")
	marseille["CF"].info()
}

/* OUTPUT ::
$ go run ex_une.go
############## GK ##############
Joueur Nom: Steve Mandanda
Joueur equipe: Marseille
charactere = 86
agilite = 80
equilibre = 85
############## DF ##############
Joueur Nom: Van Buyten
Joueur equipe: Marseille
defense = 84
equilibre = 86
tacler = 81
############## MF ##############
Joueur Nom: Lucho Gonzalez
Joueur equipe: Marseille
vitesse = 78
technique = 80
creativite = 86
############## CF ##############
Joueur Nom: Didier Drogba
Joueur equipe: Marseille
vitesse = 79
technique = 77
finition = 88


*/