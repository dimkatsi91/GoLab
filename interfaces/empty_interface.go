package main


import "fmt"


type DeskEnv string

const (
	Cinnamon DeskEnv = "Cinnamon"
	Kde      DeskEnv = "KDE"
	Xfce     DeskEnv = "XFCE"
	Mate     DeskEnv = "MATE"
	Other    DeskEnv = "NO DESKTOP"
)

type Linux struct {
	desktop DeskEnv
}

type Mint struct {
	Linux
	name string
}

type Debian struct {
	Linux
	name string
}

func info(distro interface{}) {
	//fmt.Println("Linux Distro Name    : ", distro.name)
	//fmt.Println("Desktop Environement : ", distro.desktop)
	fmt.Println(distro)
}



func main() {

	mint20 := Mint{ Linux{Cinnamon}, "Linux Mint 20.1" }
	debian_jessie := Debian{ Linux{Mate}, "Debian Jessie 8" }
	mint17_3 := Mint { Linux{Xfce}, "Linux Mint 17.3"}

	//info(mint20)

	distros := []interface{}{mint20, debian_jessie, mint17_3}
	//fmt.Println(distros)
	for _, distro := range distros {
		info(distro)
		// Same as next print-out
		//
		// fmt.Println(distro)
	}
}


/*
	{{Cinnamon} Linux Mint 20.1}
	{{MATE} Debian Jessie 8}
	{{XFCE} Linux Mint 17.3}
*/