package main

import (
	"fmt"
	"qiniu/qiniutools/tools"
)

func main() {

	sfh1 := "Bpb/f2I7AABV5/Unj2vUFAY4PzoKAAAAyAwBAAAAAAAaQNfKIAAAAFm8P/S4NmmOC3QSrqFEI9/285ZJ"
	efh2 := "Bpb_f1BIAAA3NmXd_1j9FMIEMEoHAAAAyAwBAAAAAAAMOqmZKwAAAFm8P_S4NmmOC3QSrqFEI9_285ZJ"

	fh1, err := tools.DecodeStdEfh(sfh1)
	if err != nil {
		panic(err)
	}
	fh2, err := tools.DecodeEfh(efh2)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", fh1)
	fmt.Printf("%+v\n", fh2)
}

type Test struct {
	User
	OmitUser `json:"Omit_User"`
}

type User struct {
	Address `json:"address"`
}

type Address struct {
	Name string `json:"name"`
}

type OmitUser struct {
	Address `json:"address"`
}
