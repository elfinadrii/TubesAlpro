package main

import (
	"fmt"
	
)

const NMAX = 100

type Proyek struct {
	NamaProyek string
	Klien      string
	Deadline   int
	Bayaran    int
	Status     string
}
type TabProyek [NMAX]Proyek

func main() {
	var daftarProyek TabProyek
	var jumlahProyek int
	var pilihan int

	for {
		fmt.Println("\n╔════════════════════════════════════╗")
		fmt.Println("║  MANAJEMEN PROYEK FREELANCE        ║")
		fmt.Println("╠════════════════════════════════════╣")
		fmt.Println("║ 1. Tambah Proyek                   ║")
		fmt.Println("║ 2. Ubah Proyek                     ║")
		fmt.Println("║ 3. Hapus Proyek                    ║")
		fmt.Println("║ 4. Update Status                   ║")
		fmt.Println("║ 5. Cari Proyek                     ║")
		fmt.Println("║ 6. Lihat Semua Proyek              ║")
		fmt.Println("║ 7. Laporan Proyek Selesai          ║")
		fmt.Println("║ 8. Laporan Proyek Berjalan         ║")
		fmt.Println("║ 9. Tampilkan Nilai Bayaran Tertinggi║")
		fmt.Println("║ 0. Keluar                          ║")
		fmt.Println("╚════════════════════════════════════╝")
		fmt.Print(">> Pilih menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tambahProyek(&daftarProyek, &jumlahProyek)
		case 2:
			ubahProyek(&daftarProyek, jumlahProyek)
		case 3:
			hapusProyek(&daftarProyek, &jumlahProyek)
		case 4:
			updateStatusProyek(&daftarProyek, jumlahProyek)
		case 5:
			cariProyekMenu(&daftarProyek, jumlahProyek)
		case 6:
			lihatSemuaProyek(&daftarProyek, jumlahProyek)
		case 7:
			raportProyekSelesai(&daftarProyek, jumlahProyek)
		case 8:
			raportProyekBerjalan(&daftarProyek, jumlahProyek)
		case 9:
			tampilkanNilaiMaksBayaran(&daftarProyek, jumlahProyek)
		case 0:
			fmt.Println("Terima kasih sudah menggunakan aplikasi!")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func tambahProyek(P *TabProyek, N *int) {
	if *N >= NMAX {
		fmt.Println("Kapasitas penuh.")
		return
	}

	var p Proyek

	fmt.Print("Nama Proyek: ")
	fmt.Scanln(&p.NamaProyek)

	fmt.Print("Klien: ")
	fmt.Scanln(&p.Klien)

	fmt.Print("Deadline (YYYYMMDD): ")
	fmt.Scanln(&p.Deadline)

	fmt.Print("Bayaran: ")
	fmt.Scanln(&p.Bayaran)

	var inputStatus string
	fmt.Print("Status (Pending/Sedang Dikerjakan/Selesai): ")
	fmt.Scanln(&inputStatus)

	if p.Status == "" {
		fmt.Println("Status tidak valid.")
		return
	}

	P[*N] = p
	(*N)++
	fmt.Println("Proyek berhasil ditambahkan.")
}


func cariIndexProyek(P *TabProyek, N int, nama string) int {
	for i := 0; i < N; i++ {
		if P[i].NamaProyek == nama {
			return i
		}
	}
	return -1
}

func ubahProyek(P *TabProyek, N int) {
	if N == 0 {
		fmt.Println("Belum ada proyek.")
		return
	}

	var nama string
	fmt.Print("Masukkan nama proyek yang ingin diubah: ")
	fmt.Scan(&nama)

	idx := cariIndexProyek(P, N, nama)
	if idx == -1 {
		fmt.Println("Proyek tidak ditemukan.")
		return
	}

	fmt.Print("Nama Proyek Baru: ")
	fmt.Scanln(&P[idx].NamaProyek)

	fmt.Print("Klien Baru: ")
	fmt.Scanln(&P[idx].Klien)

	fmt.Print("Deadline Baru (YYYYMMDD): ")
	fmt.Scanln(&P[idx].Deadline)

	fmt.Print("Bayaran Baru: ")
	fmt.Scanln(&P[idx].Bayaran)

	var inputStatus string
	fmt.Print("Status Baru (Pending/Sedang Dikerjakan/Selesai): ")
	fmt.Scanln(&inputStatus)

	if P[idx].Status == "" {
		fmt.Println("Status tidak valid.")
		return
	}

	fmt.Println("Proyek berhasil diubah.")
}

func hapusProyek(P *TabProyek, N *int) {
	if *N == 0 {
		fmt.Println("Belum ada proyek.")
		return
	}

	var namaHapus string
	fmt.Print("Masukkan nama proyek yang ingin dihapus: ")
	fmt.Scanln(&namaHapus)

	idx := cariIndexProyek(P, *N, namaHapus)
	if idx == -1 {
		fmt.Println("Proyek tidak ditemukan.")
		return
	}

	for i := idx; i < *N-1; i++ {
		P[i] = P[i+1]
	}
	(*N)--
	fmt.Println("Proyek berhasil dihapus.")
}


func updateStatusProyek(P *TabProyek, N int) {
	if N == 0 {
		fmt.Println("Belum ada proyek.")
		return
	}

	var nama, status string
	fmt.Print("Masukkan nama proyek (tanpa spasi): ")
	fmt.Scan(&nama)

	idx := cariIndexProyek(P, N, nama)
	if idx == -1 {
		fmt.Println("Proyek tidak ditemukan.")
		return
	}

	fmt.Print("Masukkan status baru (Pending/Sedang Dikerjakan/Selesai): ")
	fmt.Scan(&status)

	if status != "Pending" && status != "Sedang Dikerjakan" && status != "Selesai" {
    fmt.Println("Status tidak valid.")
    return
	}
	P[idx].Status = status
}


func cariProyekMenu(P *TabProyek, N int) {
	if N == 0 {
		fmt.Println("Belum ada proyek.")
		return
	}

	var nama string
	fmt.Print("Masukkan nama proyek yang dicari (tanpa spasi): ")
	fmt.Scan(&nama)

	idx := cariIndexProyek(P, N, nama)
	if idx == -1 {
		fmt.Println("Proyek tidak ditemukan.")
	} else {
		p := P[idx]
		fmt.Printf("Ditemukan: %s | Klien: %s | Deadline: %d | Bayaran: %d | Status: %s\n",
			p.NamaProyek, p.Klien, p.Deadline, p.Bayaran, p.Status)
	}
}
	

func urutkanProyek(P *TabProyek, N int) {
	if N == 0 {
		fmt.Println("Belum ada proyek.")
		return
	}

	fmt.Println("1. Urutkan berdasarkan Deadline (Ascending)")
	fmt.Println("2. Urutkan berdasarkan Bayaran (Descending)")
	fmt.Print("Pilih cara urutkan: ")
	var cara int
	fmt.Scan(&cara)

	if cara == 1 {
		for i := 0; i < N-1; i++ {
			for j := 0; j < N-1-i; j++ {
				if P[j].Deadline > P[j+1].Deadline {
					P[j], P[j+1] = P[j+1], P[j]
				}
			}
		}
		fmt.Println("Berhasil diurutkan berdasarkan deadline.")
	} else if cara == 2 {
		for i := 0; i < N-1; i++ {
			for j := 0; j < N-1-i; j++ {
				if P[j].Bayaran < P[j+1].Bayaran {
					P[j], P[j+1] = P[j+1], P[j]
				}
			}
		}
		fmt.Println("Berhasil diurutkan berdasarkan bayaran tertinggi.")
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
}


func raportProyekSelesai(P *TabProyek, N int) {
	fmt.Println("=== Laporan Proyek Selesai ===")
	found := false
	for i := 0; i < N; i++ {
		if P[i].Status == "Selesai" {
			fmt.Printf("%s | Klien: %s | Deadline: %d | Bayaran: %d\n",
				P[i].NamaProyek, P[i].Klien, P[i].Deadline, P[i].Bayaran)
			found = true
		}
	}
	if !found {
		fmt.Println("Tidak ada proyek yang selesai.")
	}
}

func raportProyekBerjalan(P *TabProyek, N int) {
	fmt.Println("=== Laporan Proyek Berjalan ===")
	found := false
	for i := 0; i < N; i++ {
		if P[i].Status == "Pending" || P[i].Status == "Sedang Dikerjakan" {
			fmt.Printf("%s | Klien: %s | Deadline: %d | Status: %s | Bayaran: %d\n",
				P[i].NamaProyek, P[i].Klien, P[i].Deadline, P[i].Status, P[i].Bayaran)
			found = true
		}
	}
	if !found {
		fmt.Println("Tidak ada proyek yang masih berjalan.")
	}
}

func lihatSemuaProyek(P *TabProyek, N int) {
	if N == 0 {
		fmt.Println("Belum ada proyek.")
		return
	}

	fmt.Println("Daftar Proyek:")
	for i := 0; i < N; i++ {
		p := P[i]
		fmt.Printf("%d. %s | Klien: %s | Deadline: %d | Bayaran: %d | Status: %s\n",
			i+1, p.NamaProyek, p.Klien, p.Deadline, p.Bayaran, p.Status)
	}
}

func nilaiMaksBayaran(P *TabProyek, N int) int {
	if N == 0 {
		return 0
	}
	maks := P[0].Bayaran
	for i := 1; i < N; i++ {
		if P[i].Bayaran > maks {
			maks = P[i].Bayaran
		}
	}
	return maks
}

func tampilkanNilaiMaksBayaran(P *TabProyek, N int) {
	maks := nilaiMaksBayaran(P, N)
	if maks == 0 {
		fmt.Println("Belum ada proyek atau nilai bayaran masih 0.")
	} else {
		fmt.Printf("Nilai bayaran tertinggi adalah: %d\n", maks)
	}
}
