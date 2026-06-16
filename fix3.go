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
        fmt.Println("ketik 1 untuk menambahkan data")        // menu 1 dan 2 jadi ada dalam satu menu di menu tambah data
        fmt.Println("ketik 2 untuk mengedit data")
        fmt.Println("ketik 3 untuk hapus data")
        fmt.Println("ketik 4 untuk menampilkan data")
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
            editData(n, &data)
        } else if pilih == 3 {
            n = hapus(&data, n)
        } else if pilih == 4 {
            // menu 3 dan 4 jadi ada dalam satu menu di menu tampil data
            tampilanData(n, data)
        } else {
            fmt.Println("pilihan tidak valid")
        }
    }
}

// ============ FUNGSI SORTING ============

// Fungsi untuk sorting berdasarkan ID (ascending) menggunakan INSERTION SORT
func sortByIdAscending(data *tabPlayer, n int) {
    for i := 1; i < n; i++ {
        key := data[i]
        j := i - 1
        
        // Geser elemen yang lebih besar ke kanan
        for j >= 0 && data[j].id > key.id {
            data[j+1] = data[j]
            j--
        }
        data[j+1] = key
    }
}

// Fungsi untuk sorting berdasarkan Winrate (descending) menggunakan SELECTION SORT
func sortByWrDescending(data *tabPlayer, n int) {
    for i := 0; i < n-1; i++ {
        maxIdx := i
        for j := i + 1; j < n; j++ {
            if data[j].wr > data[maxIdx].wr {
                maxIdx = j
            }
        }
        if maxIdx != i {
            // Menggunakan temp untuk swap
            temp := data[i]
            data[i] = data[maxIdx]
            data[maxIdx] = temp
        }
    }
}

// Fungsi untuk menyalin data (copy array)
func copyData(source tabPlayer, dest *tabPlayer, n int) {
    for i := 0; i < n; i++ {
        dest[i] = source[i]
    }
}

// ============ FUNGSI SEARCHING ============

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

// Sequential search untuk mencari player berdasarkan nickname
func sequentialSearchByNick(data tabPlayer, n int, target string) int {
    for i := 0; i < n; i++ {
        if data[i].nick == target {
            return i
        }
    }
    return -1
}

// Sequential search untuk mencari player berdasarkan rank
func sequentialSearchByRank(data tabPlayer, n int, targetRank string) {
    found := false
    fmt.Printf("\n========== PLAYER DENGAN RANK %s ==========\n", targetRank)
    fmt.Println("No\tID\tNickname\tMenang\tKalah\tWinrate")
    fmt.Println("==========================================================")
    
    for i := 0; i < n; i++ {
        if data[i].rank == targetRank {
            fmt.Printf("%d\t%d\t%s\t\t%d\t%d\t%d%%\n",
                i+1, data[i].id, data[i].nick, data[i].menang, data[i].kalah, data[i].wr)
            found = true
        }
    }
    
    if !found {
        fmt.Printf("Tidak ada player dengan rank %s\n", targetRank)
    }
    fmt.Println("==========================================================")
}

// ============ FUNGSI UTILITY ============

// Fungsi untuk mengupdate rank dan winrate
func updateRankAndWr(player *player) {
    // Hitung rank
    rank := (player.menang * 3) - (player.kalah * 2)
    if rank >= 250 {
        player.rank = "diamond"
    } else if rank >= 200 {
        player.rank = "platinum"
    } else if rank >= 150 {
        player.rank = "gold"
    } else if rank >= 100 {
        player.rank = "silver"
    } else if rank >= 50 {
        player.rank = "bronze"
    } else {
        player.rank = "unranked"
    }
    
    // Hitung winrate
    total := player.menang + player.kalah
    if total > 0 {
        player.wr = (player.menang * 100) / total
    } else {
        player.wr = 0
    }
}

// Fungsi untuk menampilkan detail player
func displayPlayerDetail(player player) {
    fmt.Println("\n========== DATA PLAYER ==========")
    fmt.Printf("Nickname : %s\n", player.nick)
    fmt.Printf("ID       : %d\n", player.id)
    fmt.Printf("Menang   : %d\n", player.menang)
    fmt.Printf("Kalah    : %d\n", player.kalah)
    fmt.Printf("Winrate  : %d %%\n", player.wr)
    fmt.Printf("Rank     : %s\n", player.rank)
    fmt.Println("=================================")
}

// ============ FUNGSI UTAMA ============

func hapus(data *tabPlayer, n int) int {
    // lanjutkan , ada pilihan hapus data keseluruhan, hapus menang atau kalahs
    var target string
    var pilihan int

    if n == 0 {
        fmt.Println("\n========== ERROR ==========")
        fmt.Println("belum ada data! silakan tambah data terlebih dahulu")
        return n
    }

    fmt.Println("\n========== HAPUS DATA ===========")
    fmt.Println("1. Hapus seluruh data player")
    fmt.Println("2. Hapus data menang player")
    fmt.Println("3. Hapus data kalah player")
    fmt.Println("4. Hapus player tertentu")
    fmt.Print("Pilihan: ")
    fmt.Scan(&pilihan)

    if pilihan == 1 {
        // Hapus semua data
        for i := 0; i < n; i++ {
            data[i] = player{}
        }
        fmt.Println("\n========== BERHASIL ==========")
        fmt.Printf("Semua data (%d player) berhasil dihapus\n", n)
        return 0
    } else if pilihan == 2 {
        // Hapus data menang player tertentu
        fmt.Print("masukan nickname player: ")
        fmt.Scan(&target)

        // Gunakan sequential search untuk mencari nickname
        idx := sequentialSearchByNick(*data, n, target)
        
        if idx != -1 {
            data[idx].menang = 0
            // Update winrate dan rank
            updateRankAndWr(&data[idx])
            fmt.Printf("\nData menang player %s berhasil direset\n", target)
            return n
        } else {
            fmt.Printf("Player dengan nickname %s tidak ditemukan\n", target)
        }
    } else if pilihan == 3 {
        // Hapus data kalah player tertentu
        fmt.Print("masukan nickname player: ")
        fmt.Scan(&target)

        // Gunakan sequential search untuk mencari nickname
        idx := sequentialSearchByNick(*data, n, target)
        
        if idx != -1 {
            data[idx].kalah = 0
            // Update winrate dan rank
            updateRankAndWr(&data[idx])
            fmt.Printf("\nData kalah player %s berhasil direset\n", target)
            return n
        } else {
            fmt.Printf("Player dengan nickname %s tidak ditemukan\n", target)
        }
    } else if pilihan == 4 {
        // Hapus player tertentu
        fmt.Print("masukan nickname player yang ingin dihapus: ")
        fmt.Scan(&target)

        // Gunakan sequential search untuk mencari nickname
        idx := sequentialSearchByNick(*data, n, target)
        
        if idx != -1 {
            // Geser data ke kiri
            for j := idx; j < n-1; j++ {
                data[j] = data[j+1]
            }
            data[n-1] = player{} // Kosongkan data terakhir
            fmt.Printf("\nPlayer %s berhasil dihapus\n", target)
            return n - 1
        } else {
            fmt.Printf("Player dengan nickname %s tidak ditemukan\n", target)
        }
    } else {
        fmt.Println("Pilihan tidak valid")
    }

    return n
}

func tambahdata(data *tabPlayer, n int) int {
    var tambah int

    fmt.Println("\n========== TAMBAH DATA ===========")
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
        data[i].rank = "unranked"
    }

    // Urutkan data berdasarkan id setelah menambah data (INSERTION SORT)
    sortByIdAscending(data, n+tambah)

    fmt.Printf("\n========== BERHASIL MENAMBAH %d DATA ===========\n", tambah)
    return n + tambah
}

// Fungsi untuk mengedit data (menang/kalah dan nickname)
func editData(n int, data *tabPlayer) {
    var dicari int

    if n == 0 {
        fmt.Println("\n========== ERROR ==========")
        fmt.Println("belum ada data! silakan tambah data terlebih dahulu")
        return
    }

    fmt.Println("\n========== EDIT DATA ===========")
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
        fmt.Println("1. mengubah data menang dan kalah")
        fmt.Println("2. mengubah nickname")
        fmt.Print("masukkan pilihan (1/2): ")
        fmt.Scan(&pilihan)

        if pilihan == 1 {
            // Edit menang dan kalah
            var subPilihan int
            var kosong, kosong2 int
            
            fmt.Println("\n========== UBAH DATA MENANG/KALAH ==========")
            fmt.Println("1. mengubah jumlah menang")
            fmt.Println("2. mengubah jumlah kalah")
            fmt.Println("3. mengubah kedua-duanya")
            fmt.Print("masukkan pilihan (1/2/3): ")
            fmt.Scan(&subPilihan)

            if subPilihan == 1 {
                fmt.Print("masukkan jumlah menang baru: ")
                fmt.Scan(&kosong)
                if kosong < 0 {
                    fmt.Println("ERROR: jumlah menang tidak boleh negatif!")
                    data[idx].menang = 0
                } else {
                    data[idx].menang = kosong
                }
            } else if subPilihan == 2 {
                fmt.Print("masukkan jumlah kalah baru: ")
                fmt.Scan(&data[idx].kalah)
                if data[idx].kalah < 0 {
                    fmt.Println("ERROR: jumlah kalah tidak boleh negatif!")
                    data[idx].kalah = 0
                }
            } else if subPilihan == 3 {
                fmt.Print("masukkan jumlah menang baru: ")
                fmt.Scan(&kosong)
                fmt.Print("masukkan jumlah kalah baru: ")
                fmt.Scan(&kosong2)

                fmt.Println()

                if kosong < 0 {
                    fmt.Println("ERROR: jumlah menang tidak boleh negatif!")
                    data[idx].menang = 0
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

            // Update rank dan winrate
            updateRankAndWr(&data[idx])

            fmt.Println("\n========== UPDATE BERHASIL ==========")
            fmt.Printf("Winrate baru: %d %%\n", data[idx].wr)
            fmt.Printf("Rank baru: %s\n", data[idx].rank)
            fmt.Println("Data berhasil diupdate!")

            // Tampilkan data terbaru
            displayPlayerDetail(data[idx])
            
        } else if pilihan == 2 {
            // Edit nickname
            var newNick string
            fmt.Print("\nmasukkan nickname baru: ")
            fmt.Scan(&newNick)
            
            oldNick := data[idx].nick
            data[idx].nick = newNick
            
            fmt.Println("\n========== UPDATE BERHASIL ==========")
            fmt.Printf("Nickname berhasil diubah dari '%s' menjadi '%s'\n", oldNick, newNick)
            fmt.Println("Data berhasil diupdate!")
            
            // Tampilkan data terbaru
            displayPlayerDetail(data[idx])
        } else {
            fmt.Println("ERROR: pilihan yang dimasukkan tidak sesuai!")
            return
        }
    } else {
        fmt.Printf("\nplayer dengan id %d tidak ditemukan\n", dicari)
    }
}

func tampilanData(n int, data tabPlayer) {
    var pilihan int

    if n == 0 {
        fmt.Println("\n========== ERROR ==========")
        fmt.Println("belum ada data! silakan tambah data terlebih dahulu")
        return
    }

    fmt.Println("\n========== TAMPIL DATA ==========")
    fmt.Println("1. Tampilkan semua player (urut berdasarkan id)")
    fmt.Println("2. Tampilkan player berdasarkan ID tertentu")
    fmt.Println("3. Tampilkan player berdasarkan rank (sequential search)")
    fmt.Println("4. Tampilkan player berdasarkan winrate (descending)")
    fmt.Print("Pilihan: ")
    fmt.Scan(&pilihan)

    if pilihan == 1 {
        // Buat salinan data untuk sorting agar tidak mengubah data asli
        var tempData tabPlayer
        copyData(data, &tempData, n)
        
        // Urutkan data berdasarkan id (INSERTION SORT)
        sortByIdAscending(&tempData, n)
        
        fmt.Println("\n========== SEMUA DATA PLAYER ==========")
        fmt.Printf("Total data: %d player\n\n", n)
        fmt.Println("No\tID\tNickname\tMenang\tKalah\tWinrate\tRank")
        fmt.Println("==========================================================")
        for i := 0; i < n; i++ {
            fmt.Printf("%d\t%d\t%s\t\t%d\t%d\t%d%%\t%s\n",
                i+1, tempData[i].id, tempData[i].nick, tempData[i].menang, tempData[i].kalah, tempData[i].wr, tempData[i].rank)
        }
        fmt.Println("==========================================================")
    } else if pilihan == 2 {
        var dicari int
        fmt.Print("masukkan id player yang mau dicari: ")
        fmt.Scan(&dicari)

        idx := binarySearch(data, n, dicari)

        if idx != -1 {
            displayPlayerDetail(data[idx])
        } else {
            fmt.Printf("\nplayer dengan id %d tidak ditemukan\n", dicari)
        }
    } else if pilihan == 3 {
        var targetRank string
        fmt.Print("Masukkan rank yang dicari (diamond/platinum/gold/silver/bronze/unranked): ")
        fmt.Scan(&targetRank)
        // cari siapa aja di rank tersebut dengan sequential
        sequentialSearchByRank(data, n, targetRank)
    } else if pilihan == 4 {
        // Buat salinan data untuk sorting
        var tempData tabPlayer
        copyData(data, &tempData, n)
        
        // sort wr dengan SELECTION SORT DESCENDING
        sortByWrDescending(&tempData, n)
        
        fmt.Println("\n========== DATA PLAYER (URUTAN WINRATE TERTINGGI) ==========")
        fmt.Println("No\tID\tNickname\tMenang\tKalah\tWinrate\tRank")
        fmt.Println("==========================================================")
        for i := 0; i < n; i++ {
            fmt.Printf("%d\t%d\t%s\t\t%d\t%d\t%d%%\t%s\n",
                i+1, tempData[i].id, tempData[i].nick, tempData[i].menang, tempData[i].kalah, tempData[i].wr, tempData[i].rank)
        }
        fmt.Println("==========================================================")
    } else {
        fmt.Println("Pilihan tidak valid")
    }
}