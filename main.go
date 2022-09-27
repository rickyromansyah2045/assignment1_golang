package main

import (
	"biodata-friend/helpers"
	"os"
)

func main() {

	var absen = os.Args[1]

	helpers.ShowBiodata(absen)
}
