package main

import (
	"fmt"
)

type Data struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func biodata() {

	listData := []Data{}

	Data1 := Data{}
	Data1.Nama = "andi"
	Data1.Alamat = "jalan 1"
	Data1.Pekerjaan = "pegawai"
	Data1.Alasan = "belajar lebih dalam"
	listData = append(listData, Data1)

	Data2 := Data{}
	Data2.Nama = "budi"
	Data2.Alamat = "jalan 2"
	Data2.Pekerjaan = "pegawai"
	Data2.Alasan = "belajar lebih dalam"
	listData = append(listData, Data2)

	Data3 := Data{}
	Data3.Nama = "bono"
	Data3.Alamat = "jalan 3"
	Data3.Pekerjaan = "pegawai"
	Data3.Alasan = "belajar lebih dalam"
	listData = append(listData, Data3)

	Data4 := Data{}
	Data4.Nama = "eko"
	Data4.Alamat = "jalan 4"
	Data4.Pekerjaan = "pegawai"
	Data4.Alasan = "belajar lebih dalam"
	listData = append(listData, Data4)

	for _, v := range listData {
		fmt.Printf("%s\n%s\n%s\n%s\n", v.Nama, v.Alamat, v.Pekerjaan, v.Alasan)
	}

}
