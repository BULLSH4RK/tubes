package main

import "fmt"

const NMAX int = 10000

type player struct{
	nick string
	id,menang,kalah,wr int 
}

type tabPlayer [NMAX]player

func main (){
	var pilih int

	fmt.Println("========== selamat datang ===========")
	fmt.Println("pilih mau mau apa kali ini")
	fmt.Println("ketik 1 untuk menambahkan data nick dan id")
	fmt.Println("ketik 2 untuk menambahkan kemenangan dan kekalahan")
	fmt.Println("ketik 3 untuk menampilkan data keseluruhan")
	fmt.Println("ketik 0 untuk keluar program")
	fmt.Println("=====================================")

	

	menu(&pilih)

}

func menu(pilih *int){
	var n,pilihan int
	var data tabPlayer
	
	*pilih = pilihan

	fmt.Print("pilihan menu: ")
	fmt.Scan(&pilihan)

	if pilihan == 0 {
		fmt.Println("========== TERIMAKASIH ==========")
	}


	switch pilihan{
	case 1:
		tambahdata(&n ,&data,pilihan)
	case 2:
		editwr(n,&data,pilihan)
	case 3:
		tampildata(n,data,pilihan)
	default:
		fmt.Println("pilihan tidak valid")
	}
}

func tambahdata(n *int, data *tabPlayer,pilih int){
	var i int
	var idx int
	var pemain tabPlayer

	


	pemain[i].menang = 0
	pemain[i].kalah = 0
	pemain[i].wr = 0


	fmt.Println("========== menu 1 ===========")
	
	fmt.Print("masukan berapa nickname dan id yang ingin dimasukkan")
	fmt.Scan(&idx)
	fmt.Println()
	fmt.Println("masukan nickname dan id player yang ingin ditambahkan")
	for i=0;i<idx;i++{
		fmt.Scan(&pemain[i].nick, &pemain[i].id)
	}
	*n = idx
	*data = pemain

	fmt.Println("========== input data berhasil ===========")

	menu(&pilih)
}

func tampildata(n int, data tabPlayer,pilih int){

	var dicari,idx int

	fmt.Print("masukkan id player yang mau dicari: ")
	fmt.Scan(&dicari)

	for i:=0 ; i<n;i++{
		if dicari == data[i].id{
			fmt.Println("========== data player ==========")
			fmt.Println(data[i].nick)
			fmt.Println(data[i].id)
			fmt.Println(data[i].menang)
			fmt.Println(data[i].kalah)
			fmt.Println(data[i].wr)
			fmt.Println("=================================")

			idx++
		}
	}

	if idx == 0 {
		fmt.Println("player dengan id ", dicari," tidak ditemukan")
		tampildata(n,data,pilih)
	}
	menu(&pilih)
}

func editwr(n int , data *tabPlayer, pilih int){
	var pemain tabPlayer
	var i,pilihan,idx, dicari int



	fmt.Print("masukkan id player yang mau dicari: ")
	fmt.Scan(&dicari)

	for i = 0; i<n;i++{
		if dicari == pemain[i].id{
			fmt.Println("========== data player ditemukan ==========")
			fmt.Println(pemain[i].nick)
			fmt.Println("========== apa yang ingin diubah ==========")
			fmt.Println(" 1. merubah jumlah menang")
			fmt.Println(" 2. merubah jumlah kalah")
			fmt.Print("masukkan pilihan: ")
			fmt.Scan(&pilihan)

				switch pilihan{
				case 1:
				fmt.Print("berapa jumlah menang: ")
				fmt.Scan(&pemain[i].menang)
				case 2:
				fmt.Print("berapa jumlah kalah: ")
				fmt.Scan(&pemain[i].kalah)
				default:
				fmt.Println("pilihan yang dimasukkan tidak sesuai !!!")
				editwr(n,&pemain, pilih)
				}

			pemain[i].wr = (pemain[i].menang/(pemain[i].menang+pemain[i].kalah))*100
				*data = pemain

			idx++
		}
	}

	if idx == 0{
		fmt.Println("player dengan id ", dicari," tidak ditemukan")
		editwr(n,&pemain,pilih)
	}
	menu(&pilih)
}