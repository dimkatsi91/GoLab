package main

import (
	"fmt"
	"errors"
	"log"
)


// A LinuxUser struct
//
type LinuxUser struct {
	username, passwd string
}

// create function checks if a LinuxUser's username
// is root and returns an error if this is the case
//
func (u *LinuxUser) create() (string, error) {
	if u.username=="root" {
		return "[]", errors.New("ACTION NOT PERMITTED! 'root' user cannot be added! Try again ...")
	}
	return ("Created user: " + u.username + " with password: " + u.passwd), nil
}

func main() {
	fmt.Println("/////////////////////////////////////////////////////////////////////////////////////////////////////")
	newAdmin := LinuxUser{"admin", "covid2020"}
	info, err := newAdmin.create()
	if err != nil {
		log.Fatalln("No error happens in this case ... err: ", err)
	}
	fmt.Println(info)
	fmt.Println("/////////////////////////////////////////////////////////////////////////////////////////////////////")
	admin := LinuxUser{"root", "covid_2020"}
	_, err = admin.create()

	if err != nil {
		log.Fatalln("Error occured: --> ", err)
	}
}
