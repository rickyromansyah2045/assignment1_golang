package helpers

import "fmt"

func ShowBiodata(args string) {

	var absen int

	if args == "1" {
		absen = 0
	} else if args == "2" {
		absen = 1
	} else if args == "3" {
		absen = 2
	} else if args == "4" {
		absen = 3
	} else {
		absen = 5
	}

	var biodata = [4]Biodata{
		{Name: "Ricky", Address: "Jl. Kebayoran", Job: "Software Engineer", Reason: "Ingin masuk gojek"},
		{Name: "Aji", Address: "Jl Ciawi Barat", Job: "Programmer", Reason: "Ingin Mempelajari Hal Baru"},
		{Name: "Rifki", Address: "Jl Malioboro Timur", Job: "Marketing", Reason: "Ingin Jadi lebih baik"},
		{Name: "Winda", Address: "Jl Kebayoran", Job: "Finance", Reason: "Ingin Switch Karir"},
	}

	if absen == 5 {
		fmt.Println("Data hanya tersedia dari 1 sampai 4")
	} else {
		fmt.Println("Nama: ", biodata[absen].Name)
		fmt.Println("Address: ", biodata[absen].Address)
		fmt.Println("Job: ", biodata[absen].Job)
		fmt.Println("Reason: ", biodata[absen].Reason)
	}

}
