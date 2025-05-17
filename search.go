package main

import (
	"fmt"
	"strings"
)

// SearchUser is a higher-order function that takes a search algorithm function
// and applies it to find users by name
func SearchUser(db *[100]Pengguna, dbLoans *[100]Pinjaman, searchAlgorithm func(*[100]Pengguna, string) []int) {
	var namaTarget string

	fmt.Print("Masukkan nama peminjam: ")
	fmt.Scan(&namaTarget)

	// Apply the search algorithm
	hasil := searchAlgorithm(db, namaTarget)

	// Display the results
	DisplaySearchResults(hasil, db, dbLoans)
}

// SequentialSearch implements a linear search algorithm (pure function)
func SequentialSearch(db *[100]Pengguna, namaTarget string) []int {
	var hasil []int

	for i := 0; i < len(db); i++ {
		if db[i].idPengguna != 0 && strings.Contains(
			strings.ToLower(db[i].nama),
			strings.ToLower(namaTarget),
		) {
			hasil = append(hasil, db[i].idPengguna)
		}
	}

	return hasil
}

// BinarySearch implements a binary search algorithm
// Requires sorted data to function correctly
func BinarySearch(db *[100]Pengguna, namaTarget string) []int {
	// First, create a sorted copy of the database
	sortedDb := SortUsersByName(*db)

	var hasil []int
	low := 0
	high := len(sortedDb) - 1

	for low <= high {
		mid := (low + high) / 2

		if mid >= len(sortedDb) || sortedDb[mid].idPengguna == 0 {
			high = mid - 1
			continue
		}

		if strings.Contains(
			strings.ToLower(sortedDb[mid].nama),
			strings.ToLower(namaTarget),
		) {
			hasil = append(hasil, sortedDb[mid].idPengguna)

			// Check elements to the left and right for more matches
			for i := mid - 1; i >= 0; i-- {
				if sortedDb[i].idPengguna != 0 && strings.Contains(
					strings.ToLower(sortedDb[i].nama),
					strings.ToLower(namaTarget),
				) {
					hasil = append(hasil, sortedDb[i].idPengguna)
				} else {
					break
				}
			}

			for i := mid + 1; i < len(sortedDb); i++ {
				if sortedDb[i].idPengguna != 0 && strings.Contains(
					strings.ToLower(sortedDb[i].nama),
					strings.ToLower(namaTarget),
				) {
					hasil = append(hasil, sortedDb[i].idPengguna)
				} else {
					break
				}
			}

			break
		} else if strings.ToLower(sortedDb[mid].nama) < strings.ToLower(namaTarget) {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return hasil
}

// SortUsersByName returns a sorted copy of the users database (immutable)
func SortUsersByName(db [100]Pengguna) []Pengguna {
	// Extract non-empty users
	var users []Pengguna
	for i := 0; i < len(db); i++ {
		if db[i].idPengguna != 0 {
			users = append(users, db[i])
		}
	}

	// Simple bubble sort
	for i := 0; i < len(users)-1; i++ {
		for j := 0; j < len(users)-i-1; j++ {
			if users[j].nama > users[j+1].nama {
				users[j], users[j+1] = users[j+1], users[j]
			}
		}
	}

	return users
}

// DisplaySearchResults shows the search results
func DisplaySearchResults(userIds []int, db *[100]Pengguna, dbLoans *[100]Pinjaman) {
	if len(userIds) == 0 {
		fmt.Println("Tidak ditemukan data peminjam dengan nama tersebut.")
		return
	}

	fmt.Printf("Ditemukan %d peminjam:\n", len(userIds))

	for _, id := range userIds {
		// Find the user
		var user Pengguna
		for i := 0; i < len(db); i++ {
			if (*db)[i].idPengguna == id {
				user = (*db)[i]
				break
			}
		}

		fmt.Printf("\n=== DATA PEMINJAM %s (ID: %d) ===\n", user.nama, user.idPengguna)

		// Find loans for this user
		loans := FilterLoans(*dbLoans, func(p Pinjaman) bool {
			return p.idPeminjam == id
		})

		if len(loans) == 0 {
			fmt.Println("Peminjam belum memiliki pinjaman.")
			continue
		}

		// Display all loans
		for i, loan := range loans {
			fmt.Printf("- Pinjaman ke-%d\n", i+1)
			fmt.Printf("  Jumlah: Rp%d\n", loan.jumlahPinjaman)
			fmt.Printf("  Tenor: %d bulan/hari\n", loan.tenor)
			fmt.Printf("  Bunga: %.2f%%\n", loan.bunga*100)

			statusStr := "BELUM LUNAS"
			if loan.statusLunas {
				statusStr = "LUNAS"
			}
			fmt.Printf("  Status: %s\n", statusStr)
			fmt.Println("-----------------------------")
		}
	}
}

// SearchBorrowerData provides a user interface for searching
func SearchBorrowerData(dbUsers *[100]Pengguna, dbLoans *[100]Pinjaman) {
	var choice int

	fmt.Println("\n=== CARI DATA PEMINJAM ===")
	fmt.Println("[1] Pencarian Sequential")
	fmt.Println("[2] Pencarian Binary")
	fmt.Print("Pilih metode pencarian: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		SearchUser(dbUsers, dbLoans, SequentialSearch)
	case 2:
		SearchUser(dbUsers, dbLoans, BinarySearch)
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}
