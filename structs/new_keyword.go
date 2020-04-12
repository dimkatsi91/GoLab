/*
	Enum & new keyword in GoLang example
	No 'enum' in Go -> Use a type (like everything else) and const values instead
*/
package main

import (
	"fmt"
)

// Something like enum in GoLang
//
type DeskEnv string

const (
	Cinnamon DeskEnv = "Cinnamon"
	Kde      DeskEnv = "KDE"
	Xfce     DeskEnv = "XFCE"
	Mate     DeskEnv = "MATE"
	Other    DeskEnv = "NO DESKTOP"
)

type Mint struct {
	// Desktop Environment : Cinnamon | KDE | XFCE | MATE
	desktop                                   DeskEnv
	os_version, shell_version, kernel_version string
}

func (mint *Mint) echo() {
	fmt.Println("********************** INFO **********************")
	fmt.Printf("\tDesktop Environment: %s\n\tOS Version: %s\n\tKernel Version: %s\n\tShell Version: %s\n", mint.desktop, mint.os_version, mint.kernel_version, mint.shell_version)
	fmt.Println("**************************************************")
}

func main() {
	// new keyword usage in GoLang
	//
	mint19_3 := new(Mint)

	mint19_3.desktop = Cinnamon
	mint19_3.os_version = "19.3"
	mint19_3.shell_version = "3.28.3"
	mint19_3.kernel_version = "4.16"
	mint19_3.echo()
}
