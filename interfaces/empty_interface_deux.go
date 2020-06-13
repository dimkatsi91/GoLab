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
	switch d := distro.(type) {
	case Mint:
		fmt.Printf("Linux Mint --- %v\n", d)
	case Debian:
		fmt.Printf("Linux Debian --- %v\n", d)
	case Linux:
		fmt.Printf("Linux Empty ... \n")
	case string:
		fmt.Println("Oops a string detected -- Printing it out --> ", d)
	default:
		fmt.Printf("Something else, not a Linux Distro for sure ..\n")
	}
}


// ILinux --> Linux Interface | Do Not confuse it with Linux structure
// type Whatever interface {}
//
type ILinux interface {}


// main
//
func main() {

	mint20 := Mint{ Linux{Cinnamon}, "Linux Mint 20.1" }
	debian_jessie := Debian{ Linux{Mate}, "Debian Jessie 8" }
	mint17_3 := Mint { Linux{Xfce}, "Linux Mint 17.3"}

	diman := "Dimos Katsimardos"
	dimanAge := 28

	// a slice of Linux Distros using the empty intrerface ILinux
	// Also this slice has a string which is handled by info() func
	// And also an int which is not handled by info() func, so the default case
	// will be printed-out btw
	// 
	// distros := []Whatever{mint20, debian_jessie, mint17_3, diman, dimanAge}
	distros := []ILinux{mint20, debian_jessie, mint17_3, diman, dimanAge}
	for _, distro := range distros {
		info(distro)
	}
}


/*
	Linux Mint --- {{Cinnamon} Linux Mint 20.1}
	Linux Debian --- {{MATE} Debian Jessie 8}
	Linux Mint --- {{XFCE} Linux Mint 17.3}
	Oops a string detected -- Printing it out -->  Dimos Katsimardos
	Something else, not a Linux Distro for sure ..
*/