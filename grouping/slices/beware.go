package main

import "fmt"

func print(slice []string) {
	fmt.Println("#################################")
	for _, player := range slice {
		fmt.Println(player)
	}
	fmt.Println("#################################")
}

// CAUTION :: Errors will Occur for sure at first !
//
func main() {
	// Names slice of strings ...
	//
	brazilians := []string{"Rivaldo", "Giovanni", "Ze Elias", "Luciano", "Alvez"}
	slice := brazilians[2:5]
	print(brazilians)
	fmt.Println("")
	slice[0] = "Karembue -- Ooops -- Cahnged ???"
	print(brazilians)
}

/* Sample Output ::

#################################
Rivaldo
Giovanni
Ze Elias
Luciano
Alvez
#################################

#################################
Rivaldo
Giovanni
Karembue -- Ooops -- Cahnged ???
Luciano
Alvez
#################################

*/
