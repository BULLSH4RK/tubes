package main

import "fmt"

const NMAX int = 10000
// cari siapa aja di rank tersebut dengan sequential 
// sort wr dengan selection sort descendings
// hapus data

type player struct {
    nick, rank            string
    id, menang, kalah, wr int
}

type tabPlayer [NMAX]player

func main() {
    var data tabPlayer
    var n int = 0
    var pilih int

    fmt.Println("========== SELAMAT DATANG ===========")


// ubah menu utama jadi ada tambah data,edit data,hapus data,tampilkan data,keluar

    for {
        // TAMPILAN MENU
        fmt.Println("\n=====================================")
        fmt.Println("pilih mau apa kali ini")
        fmt.Println("ketik 1 untuk menambahkan data nick dan id")// menu 1 dan 2 jadi ada dalam satu menu di menu tambah data
        fmt.Println("ketik 2 untuk menambahkan kemenangan dan kekalahan")
        fmt.Println("ketik 3 untuk menampilkan data keseluruhan")// menu 3 dan 4 jadi ada dalam satu menu di menu tampil datas
        fmt.Println("ketik 4 untuk menampilkan semua player (urut berdasarkan id)")
        fmt.Println("ketik 0 untuk keluar program")
        fmt.Println("=====================================")
        fmt.Print("pilihan menu: ")
        fmt.Scan(&pilih)

        // KELUAR DARI PROGRAM
        if pilih == 0 {
            fmt.Println("========== TERIMAKASIH ==========")
            break
        }

        // PROSES BERDASARKAN PILIHAN (menggunakan if-else if)
        if pilih == 1 {
            n = tambahdata(&data, n)
        } else if pilih == 2 {
            editwr(n, &data)
        } else if pilih == 3 {
            tampildata(n, data)
        } else if pilih == 4 {
            // Urutkan data berdasarkan id sebelum ditampilkan
            for i := 0; i < n-1; i++ {
                for j := i + 1; j < n; j++ {
                    if data[i].id > data[j].id {
                        data[i], data[j] = data[j], data[i]
                    }
                }
            }
            tampilsemua(n, data)
        } else {
            fmt.Println("pilihan tidak valid")
        }
    }
}



func hapus(data *tabPlayer,n int ){
	// lanjutkan , ada pilihan hapus data keseluruhan, hapus menang atau kalahs
	var target string
	
	fmt.Println("\n========== HAPUS DATA ===========")
    fmt.Print("masukan nickname yang ingin dirubah: ")
	fmt.Scan(&target)
	
	for i:= 0 ; i <n;i++{
		if data[i].nick == target {
			fmt.Println("\n========== PILIH ===========")
		}
	}
	
}



func tambahdata(data *tabPlayer, n int) int {
    var tambah int

    fmt.Println("\n========== MENU 1 ===========")
    fmt.Print("masukan berapa nickname dan id yang ingin dimasukkan: ")
    fmt.Scan(&tambah)
    fmt.Println()

    // Cek apakah masih cukup kapasitas
    if n+tambah > NMAX {
        fmt.Printf("ERROR: kapasitas tidak mencukupi! (max %d data, tersisa %d data)\n", NMAX, NMAX-n)
        return n
    }

    if tambah <= 0 {
        fmt.Println("ERROR: jumlah data harus lebih dari 0!")
        return n
    }

    fmt.Println("masukan nickname dan id player yang ingin ditambahkan")
    fmt.Println("(format: nickname id) contoh: Budi 123")
    fmt.Println()

    for i := n; i < n+tambah; i++ {
        fmt.Printf("data ke-%d: ", i+1)
        fmt.Scan(&data[i].nick, &data[i].id)

        // Cek duplikasi id
        for j := 0; j < i; j++ {
            if data[j].id == data[i].id {
                fmt.Println("ERROR: id sudah terdaftar! masukkan data lagi")
                i-- // ulangi input untuk index ini
                break
            }
        }

        data[i].menang = 0
        data[i].kalah = 0
        data[i].wr = 0
    }

    // Urutkan data berdasarkan id setelah menambah data
    for i := 0; i < n+tambah-1; i++ {
        for j := i + 1; j < n+tambah; j++ {
            if data[i].id > data[j].id {
                data[i], data[j] = data[j], data[i]
            }
        }
    }

    fmt.Printf("\n========== BERHASIL MENAMBAH %d DATA ===========\n", tambah)
    return n + tambah
}

// Binary search untuk mencari index player berdasarkan id
func binarySearch(data tabPlayer, n int, id int) int {
    left := 0
    right := n - 1

    for left <= right {
        mid := (left + right) / 2
        if data[mid].id == id {
            return mid
        } else if data[mid].id < id {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return -1
}

func tampildata(n int, data tabPlayer) {
    var dicari int

    if n == 0 {
        fmt.Println("\n========== ERROR ==========")
        fmt.Println("belum ada data! silakan tambah data terlebih dahulu")
        return
    }

    fmt.Println("\n========== MENU 3 ===========")
    fmt.Print("masukkan id player yang mau dicari: ")
    fmt.Scan(&dicari)

    // Gunakan binary search
    idx := binarySearch(data, n, dicari)

    if idx != -1 {
        fmt.Println("\n========== DATA PLAYER ==========")
        fmt.Printf("Nickname : %s\n", data[idx].nick)
        fmt.Printf("ID       : %d\n", data[idx].id)
        fmt.Printf("Menang   : %d\n", data[idx].menang)
        fmt.Printf("Kalah    : %d\n", data[idx].kalah)
        fmt.Printf("Winrate  : %d %%\n", data[idx].wr)
        fmt.Printf("Rank     : %s\n", data[idx].rank)
        fmt.Println("=================================")
    } else {
        fmt.Printf("\nplayer dengan id %d tidak ditemukan\n", dicari)
    }
}

func editwr(n int, data *tabPlayer) {
    var dicari int

    if n == 0 {
        fmt.Println("\n========== ERROR ==========")
        fmt.Println("belum ada data! silakan tambah data terlebih dahulu")
        return
    }

    fmt.Println("\n========== MENU 2 ===========")
    fmt.Print("masukkan id player yang mau dicari: ")
    fmt.Scan(&dicari)

    // Gunakan binary search
    idx := binarySearch(*data, n, dicari)

    if idx != -1 {
        fmt.Println("\n========== DATA PLAYER DITEMUKAN ==========")
        fmt.Printf("Nickname       : %s\n", data[idx].nick)
        fmt.Printf("ID             : %d\n", data[idx].id)
        fmt.Printf("Menang saat ini: %d\n", data[idx].menang)
        fmt.Printf("Kalah saat ini : %d\n", data[idx].kalah)
        fmt.Printf("Winrate saat ini: %d %%\n", data[idx].wr)
        fmt.Printf("Rank saat ini   : %s\n", data[idx].rank)

        var pilihan int
        fmt.Println("\n========== APA YANG INGIN DIUBAH ==========")
        fmt.Println("1. mengubah jumlah menang")
        fmt.Println("2. mengubah jumlah kalah")
        fmt.Println("3. mengubah kedua-duanya")
        fmt.Print("masukkan pilihan (1/2/3): ")
        fmt.Scan(&pilihan)
        var kosong, kosong2 int

        if pilihan == 1 {
            fmt.Print("masukkan jumlah menang baru: ")
            fmt.Scan(&kosong)
            if kosong < 0 {
                fmt.Println("ERROR: jumlah menang tidak boleh negatif!")
                data[idx].menang = 0
            } else if kosong < data[idx].menang {
                fmt.Println("ERROR: jumlah menang baru tidak boleh lebih kecil dari jumlah menang saat ini!")
                fmt.Println("jumlah menang saat ini: ", data[idx].menang)
                // Tidak mengubah nilai
            } else {
                data[idx].menang = kosong
            }
        } else if pilihan == 2 {
            fmt.Print("masukkan jumlah kalah baru: ")
            fmt.Scan(&data[idx].kalah)
            if data[idx].kalah < 0 {
                fmt.Println("ERROR: jumlah kalah tidak boleh negatif!")
                data[idx].kalah = 0
            }
        } else if pilihan == 3 {
            fmt.Print("masukkan jumlah menang baru: ")
            fmt.Scan(&kosong)
            fmt.Print("masukkan jumlah kalah baru: ")
            fmt.Scan(&kosong2)

            fmt.Println()

            if kosong < 0 {
                fmt.Println("ERROR: jumlah menang tidak boleh negatif!")
                data[idx].menang = 0
            } else if kosong < data[idx].menang {
                fmt.Println("ERROR: jumlah menang baru tidak boleh lebih kecil dari jumlah menang saat ini!")
                fmt.Println("jumlah menang saat ini: ", data[idx].menang)
                // Tidak mengubah nilai
            } else {
                data[idx].menang = kosong
            }

            fmt.Println()

            if kosong2 < 0 {
                fmt.Println("ERROR: jumlah kalah tidak boleh negatif!")
                data[idx].kalah = 0
            } else {
                data[idx].kalah = kosong2
            }
        } else {
            fmt.Println("ERROR: pilihan yang dimasukkan tidak sesuai!")
            return
        }

        rank := (data[idx].menang * 3) - (data[idx].kalah * 2)
        if rank >= 250 {
            data[idx].rank = "diamond"
        } else if rank >= 200 {
            data[idx].rank = "platinum"
        } else if rank >= 150 {
            data[idx].rank = "gold"
        } else if rank >= 100 {
            data[idx].rank = "silver"
        } else if rank >= 50 {
            data[idx].rank = "bronze"
        } else {
            data[idx].rank = "unranked"
        }
        // Hitung winrate
        total := data[idx].menang + data[idx].kalah
        if total > 0 {
            data[idx].wr = (data[idx].menang * 100) / total
        } else {
            data[idx].wr = 0
        }

        fmt.Println("\n========== UPDATE BERHASIL ==========")
        fmt.Printf("Winrate baru: %d %%\n", data[idx].wr)
        fmt.Printf("Rank baru: %s\n", data[idx].rank)
        fmt.Println("Data berhasil diupdate!")

        // Tampilkan data terbaru
        fmt.Println("\n========== DATA TERBARU ==========")
        fmt.Printf("Nickname: %s\n", data[idx].nick)
        fmt.Printf("ID: %d\n", data[idx].id)
        fmt.Printf("Menang: %d\n", data[idx].menang)
        fmt.Printf("Kalah: %d\n", data[idx].kalah)
        fmt.Printf("Winrate: %d %%\n", data[idx].wr)
        fmt.Printf("Rank: %s\n", data[idx].rank)
        fmt.Println("=================================")
    } else {
        fmt.Printf("\nplayer dengan id %d tidak ditemukan\n", dicari)
    }
}

func tampilsemua(n int, data tabPlayer) {
    if n == 0 {
        fmt.Println("\n========== ERROR ==========")
        fmt.Println("belum ada data! silakan tambah data terlebih dahulu")
        return
    }

    fmt.Println("\n========== SEMUA DATA PLAYER ==========")
    fmt.Printf("Total data: %d player\n\n", n)

    // Tampilkan data (sudah terurut karena disorting di menu 4 atau saat tambah data)
    fmt.Println("No\tID\tNickname\tMenang\tKalah\tWinrate\tRank")
    fmt.Println("==========================================================")
    for i := 0; i < n; i++ {
        fmt.Printf("%d || \t |id| %d || \t |nick| %s || \t |menang| %d|| \t |kalah| %d || \t |wr| %d%% || \t |rank| %s\n", // dirapikan lagi
            i+1, data[i].id, data[i].nick, data[i].menang, data[i].kalah, data[i].wr, data[i].rank)
    }
    fmt.Println("==========================================================")
}