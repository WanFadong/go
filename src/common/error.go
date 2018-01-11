package common

import (
	"fmt"
	"encoding/json"
)

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func FmtIfError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func OutputResultOrPanic(r interface{}, err error) {
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", r)
}

func OutputJson(r interface{}) {
	rj, err := json.Marshal(r)
	PanicIfError(err)
	fmt.Println(string(rj))
}
