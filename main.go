package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type T struct {
	SRCPATH    string
	TESTSUITES map[string]struct {
		BASEPATH string
		COMMAND  string
		ENV      map[string]string
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	t := T{}
	dat, err := ioutil.ReadFile(".ggtr.yaml")
	check(err)
	err = yaml.Unmarshal(dat, &t)
	check(err)
	fmt.Print(t)
}
