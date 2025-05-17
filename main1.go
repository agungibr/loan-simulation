// This is the updated main.go file
package main

import (
	"fmt"
)

func main() {
	// Initialize database with sample data
	SeedData()

	// Handle auth with a higher-order function that processes login results
	HandleAuth(dbPengguna, func(user *Pengguna) bool {
		if user == nil {
			return false
		}

		// Run the main application after successful login
		return RunMainApplication(user)
	})
}

// RunMainApplication runs the main application after authentication
// Returns a boolean indicating whether the application should continue
func RunMainApplication(user *Pengguna) bool {
	var choice int

	for {
		// Display main menu
		fmt.Println("\n=== TRAPINJAMAN ONLINE ===")
		fmt.Println("[1] Ajukan Pinjaman")
		fmt.Println("[2] Lihat Pinjaman Saya")
		fmt.Println("[3] Pelunasan")
		fmt.Println("[4] Ubah Data Pinjaman")
		fmt.Println("[5] Hapus Data Pinjaman")
		fmt.Println("[6] Cari Data Peminjam")
		fmt.Println("[7] Urutkan Data Peminjam")
		fmt.Println("[8] Lihat Laporan Pinjaman")
		fmt.Println("[9] Profil Saya")
		fmt.Println("[0] Keluar")
		fmt.Print("Pilih: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			// Apply for a loan using higher-order function
			ApplyForLoan(user, &dbDataPeminjam, func(loan Pinjaman) string {
				return "Pinjaman telah diproses dan disetujui."
			})

		case 2:
			ViewUserLoans(user, &dbDataPeminjam)

		case 3:
			RepayLoan(user, &dbDataPeminjam)

		case 4:
			UpdateLoanData(user, &dbDataPeminjam)

		case 5:
			DeleteLoan(user, &dbDataPeminjam)

		case 6:
			SearchBorrowerData(&dbPengguna, &dbDataPeminjam)

		case 7:
			SortBorrowerData(&dbDataPeminjam)

		case 8:
			GenerateLoanReport(&dbDataPeminjam)

		case 9:
			DisplayUserProfile(user)

		case 0:
			fmt.Println("Terima kasih telah menggunakan TraPinjaman Online!")
			return false

		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
