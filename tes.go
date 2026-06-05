package main

import "fmt"

const NMAX int = 10000

type player struct {
    nick,rank            string
    id, menang, kalah, wr int
}

type tabPlayer [NMAX]player

func main() {
    var data tabPlayer
    var n int = 0
    var pilih int

    fmt.Println("========== SELAMAT DATANG ===========")

    for {
        // TAMPILAN MENU
        fmt.Println("\n=====================================")
        fmt.Println("pilih mau apa kali ini")
        fmt.Println("ketik 1 untuk menambahkan data nick dan id")
        fmt.Println("ketik 2 untuk menambahkan kemenangan dan kekalahan")
        fmt.Println("ketik 3 untuk menampilkan data keseluruhan")
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

        // PROSES BERDASARKAN PILIHAN
        switch pilih {
        case 1:
            n = tambahdata(&data, n)
        case 2:
            editwr(n, &data)
        case 3:
            tampildata(n, data)
        case 4:
            tampilsemua(n, data)
        default:
            fmt.Println("pilihan tidak valid")
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

    fmt.Printf("\n========== BERHASIL MENAMBAH %d DATA ===========\n", tambah)
    return n + tambah
}

func tampildata(n int, data tabPlayer) {
    var dicari int
    var ditemukan bool = false

    if n == 0 {
        fmt.Println("\n========== ERROR ==========")
        fmt.Println("belum ada data! silakan tambah data terlebih dahulu")
        return
    }

    fmt.Println("\n========== MENU 3 ===========")
    fmt.Print("masukkan id player yang mau dicari: ")
    fmt.Scan(&dicari)

    for i := 0; i < n; i++ {
        if dicari == data[i].id {
            fmt.Println("\n========== DATA PLAYER ==========")
            fmt.Printf("Nickname : %s\n", data[i].nick)
            fmt.Printf("ID       : %d\n", data[i].id)
            fmt.Printf("Menang   : %d\n", data[i].menang)
            fmt.Printf("Kalah    : %d\n", data[i].kalah)
            fmt.Printf("Winrate  : %d %%\n", data[i].wr)
            fmt.Printf("Rank     : %d\n", data[i].rank)
            fmt.Println("=================================")
            ditemukan = true
            break
        }
    }

    if !ditemukan {
        fmt.Printf("\nplayer dengan id %d tidak ditemukan\n", dicari)
    }
}

func editwr(n int, data *tabPlayer) {
    var dicari int
    var ditemukan bool = false

    if n == 0 {
        fmt.Println("\n========== ERROR ==========")
        fmt.Println("belum ada data! silakan tambah data terlebih dahulu")
        return
    }

    fmt.Println("\n========== MENU 2 ===========")
    fmt.Print("masukkan id player yang mau dicari: ")
    fmt.Scan(&dicari)

    for i := 0; i < n; i++ {
        if dicari == data[i].id {
            ditemukan = true
            fmt.Println("\n========== DATA PLAYER DITEMUKAN ==========")
            fmt.Printf("Nickname       : %s\n", data[i].nick)
            fmt.Printf("ID             : %d\n", data[i].id)
            fmt.Printf("Menang saat ini: %d\n", data[i].menang)
            fmt.Printf("Kalah saat ini : %d\n", data[i].kalah)
            fmt.Printf("Winrate saat ini: %d %%\n", data[i].wr)
            fmt.Printf("Rank saat ini   : %s\n", data[i].rank)

            var pilihan int
            fmt.Println("\n========== APA YANG INGIN DIUBAH ==========")
            fmt.Println("1. mengubah jumlah menang")
            fmt.Println("2. mengubah jumlah kalah")
            fmt.Println("3. mengubah kedua-duanya")
            fmt.Print("masukkan pilihan (1/2/3): ")
            fmt.Scan(&pilihan)

            switch pilihan {
            case 1:
                fmt.Print("masukkan jumlah menang baru: ")
                fmt.Scan(&data[i].menang)
                if data[i].menang < 0 {
                    fmt.Println("ERROR: jumlah menang tidak boleh negatif!")
                    data[i].menang = 0
                }
            case 2:
                fmt.Print("masukkan jumlah kalah baru: ")
                fmt.Scan(&data[i].kalah)
                if data[i].kalah < 0 {
                    fmt.Println("ERROR: jumlah kalah tidak boleh negatif!")
                    data[i].kalah = 0
                }
            case 3:
                fmt.Print("masukkan jumlah menang baru: ")
                fmt.Scan(&data[i].menang)
                fmt.Print("masukkan jumlah kalah baru: ")
                fmt.Scan(&data[i].kalah)
                if data[i].menang < 0 {
                    fmt.Println("ERROR: jumlah menang tidak boleh negatif!")
                    data[i].menang = 0
                }
                if data[i].kalah < 0 {
                    fmt.Println("ERROR: jumlah kalah tidak boleh negatif!")
                    data[i].kalah = 0
                }
            default:
                fmt.Println("ERROR: pilihan yang dimasukkan tidak sesuai!")
                return
            }

            rank := (data[i].menang * 3) - (data[i].kalah * 2)
            if rank >= 250 {
                data[i].rank = "diamond"
            } else if rank >= 200 {
                data[i].rank = "platinum"
            } else if rank >= 150 {
                data[i].rank = "gold"
            } else if rank >= 100 {
                data[i].rank = "silver"
            } else if rank >= 50 {
                data[i].rank = "bronze"
            } else {
                data[i].rank = "unranked"
            }
            // Hitung winrate
            total := data[i].menang + data[i].kalah
            if total > 0 {
                data[i].wr = (data[i].menang * 100) / total
            } else {
                data[i].wr = 0
            }

            fmt.Println("\n========== UPDATE BERHASIL ==========")
            fmt.Printf("Winrate baru: %d %%\n", data[i].wr)
            fmt.Printf("Rank baru: %s\n", data[i].rank)
            fmt.Println("Data berhasil diupdate!")
            
            // Tampilkan data terbaru
            fmt.Println("\n========== DATA TERBARU ==========")
            fmt.Printf("Nickname: %s\n", data[i].nick)
            fmt.Printf("ID: %d\n", data[i].id)
            fmt.Printf("Menang: %d\n", data[i].menang)
            fmt.Printf("Kalah: %d\n", data[i].kalah)
            fmt.Printf("Winrate: %d %%\n", data[i].wr)
            fmt.Printf("Rank: %s\n", data[i].rank)
            fmt.Println("=================================")
            break
        }
    }

    if !ditemukan {
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
    
    // Buat salinan data untuk sorting
    var temp tabPlayer
    for i := 0; i < n; i++ {
        temp[i] = data[i]
    }
    
    // Sorting berdasarkan id (ascending)
    for i := 0; i < n-1; i++ {
        for j := i + 1; j < n; j++ {
            if temp[i].id > temp[j].id {
                temp[i], temp[j] = temp[j], temp[i]
            }
        }
    }
    
    // Tampilkan data
    fmt.Println("No\tID\tNickname\tMenang\tKalah\tWinrate\tRank")
    fmt.Println("==========================================================")
    for i := 0; i < n; i++ {
        fmt.Printf("%d || \t |id| %d || \t |nick| %s || \t |menang| %d|| \t |kalah| %d || \t |wr| %d%% || \t |rank| %s\n ", 
            i+1, temp[i].id, temp[i].nick, temp[i].menang, temp[i].kalah, temp[i].wr, temp[i].rank)
    }
    fmt.Println("==========================================================")
}

// Fungsi tambahan: menghapus data berdasarkan id
func hapusdata(n int, data *tabPlayer) int {
    var dicari int
    var ditemukan bool = false
    var idx int

    if n == 0 {
        fmt.Println("\nbelum ada data yang bisa dihapus!")
        return n
    }

    fmt.Println("\n========== HAPUS DATA ==========")
    fmt.Print("masukkan id player yang akan dihapus: ")
    fmt.Scan(&dicari)

    for i := 0; i < n; i++ {
        if dicari == data[i].id {
            ditemukan = true
            idx = i
            break
        }
    }

    if !ditemukan {
        fmt.Printf("player dengan id %d tidak ditemukan\n", dicari)
        return n
    }

    // Geser data ke kiri
    for i := idx; i < n-1; i++ {
        data[i] = data[i+1]
    }

    fmt.Printf("player dengan id %d berhasil dihapus!\n", dicari)
    return n - 1
}

// Fungsi tambahan: update winrate otomatis setelah menambah menang/kalah
func updateWinrate(n int, data *tabPlayer) {
    for i := 0; i < n; i++ {
        total := data[i].menang + data[i].kalah
        if total > 0 {
            data[i].wr = (data[i].menang * 100) / total
        } else {
            data[i].wr = 0
        }
    }
}