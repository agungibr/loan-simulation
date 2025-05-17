package main

import (
	"fmt"
)

// SortLoans is a higher-order function that takes a sorting algorithm function
// and applies it to sort loans based on a field selector and comparator
func SortLoans(db *[100]Pinjaman,
	sortAlgorithm func(*[100]Pinjaman, func(Pinjaman) int, func(int, int) bool) [100]Pinjaman,
	fieldSelector func(Pinjaman) int,
	comparator func(int, int) bool) {

	// Apply the sorting algorithm and return a new sorted array
	sortedLoans := sortAlgorithm(db, fieldSelector, comparator)

	// Update the original database with the sorted array
	*db = sortedLoans

	// Display the sorted loans
	DisplayAllLoans(*db)
}

// SelectionSort implements the selection sort algorithm (pure function)
func SelectionSort(db *[100]Pinjaman, fieldSelector func(Pinjaman) int, comparator func(int, int) bool) [100]Pinjaman {
	// Create a copy to maintain immutability
	var result [100]Pinjaman
	copy(result[:], (*db)[:])

	// Count valid loans
	validCount := 0
	for i := 0; i < len(result); i++ {
		if result[i].idPeminjam != 0 {
			validCount++
		}
	}

	// Selection sort on the valid items
	for i := 0; i < validCount-1; i++ {
		minIdx := i
		for j := i + 1; j < len(result); j++ {
			if result[j].idPeminjam != 0 &&
				comparator(fieldSelector(result[j]), fieldSelector(result[minIdx])) {
				minIdx = j
			}
		}
		// Swap
		if minIdx != i {
			result[i], result[minIdx] = result[minIdx], result[i]
		}
	}

	return result
}

// InsertionSort implements the insertion sort algorithm (pure function)
func InsertionSort(db *[100]Pinjaman, fieldSelector func(Pinjaman) int, comparator func(int, int) bool) [100]Pinjaman {
	// Create a copy to maintain immutability
	var result [100]Pinjaman
	copy(result[:], (*db)[:])

	for i := 1; i < len(result); i++ {
		if result[i].idPeminjam == 0 {
			continue
		}

		key := result[i]
		j := i - 1

		for j >= 0 && result[j].idPeminjam != 0 &&
			comparator(fieldSelector(key), fieldSelector(result[j])) {
			result[j+1] = result[j]
			j--
		}

		result[j+1] = key
	}

	return result
}

// Field selectors - pure functions that extract a field from a loan
func AmountSelector(p Pinjaman) int {
	return p.jumlahPinjaman
}

func TenorSelector(p Pinjaman) int {
	return p.tenor
}

// Comparators - pure functions that compare two values
func LessThan(a, b int) bool {
	return a < b
}

func GreaterThan(a, b int) bool {
	return a > b
}

// DisplayAllLoans shows all loans in the database
func DisplayAllLoans(db [100]Pinjaman) {
	fmt.Println("\n=== DAFTAR PINJAMAN TERURUT ===")

	// Filter valid loans
	validLoans := FilterLoans(db, func(p Pinjaman) bool {
		return true // All non-empty loans (already filtered in FilterLoans)
	})

	if len(validLoans) == 0 {
		fmt.Println("Tidak ada data pinjaman.")
		return
	}

	for i, loan := range validLoans {
		fmt.Printf("- Pinjaman ke-%d\n", i+1)
		fmt.Printf("  Peminjam ID: %d\n", loan.idPeminjam)
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

// SortBorrowerData provides a user interface for sorting
func SortBorrowerData(db *[100]Pinjaman) {
	var sortField, sortMethod, sortOrder int

	fmt.Println("\n=== URUTKAN DATA PEMINJAM ===")
	fmt.Println("[1] Urutkan berdasarkan Jumlah Pinjaman")
	fmt.Println("[2] Urutkan berdasarkan Tenor")
	fmt.Print("Pilih kriteria pengurutan: ")
	fmt.Scan(&sortField)

	fmt.Println("\n[1] Selection Sort")
	fmt.Println("[2] Insertion Sort")
	fmt.Print("Pilih metode pengurutan: ")
	fmt.Scan(&sortMethod)

	fmt.Println("\n[1] Menaik (Ascending)")
	fmt.Println("[2] Menurun (Descending)")
	fmt.Print("Pilih arah pengurutan: ")
	fmt.Scan(&sortOrder)

	// Select the field selector based on user choice
	var fieldSelector func(Pinjaman) int
	switch sortField {
	case 1:
		fieldSelector = AmountSelector
	case 2:
		fieldSelector = TenorSelector
	default:
		fmt.Println("Pilihan tidak valid.")
		return
	}

	// Select the comparator based on sort order
	var comparator func(int, int) bool
	if sortOrder == 1 {
		comparator = LessThan // Ascending
	} else {
		comparator = GreaterThan // Descending
	}

	// Select and apply the sort algorithm
	switch sortMethod {
	case 1:
		SortLoans(db, SelectionSort, fieldSelector, comparator)
	case 2:
		SortLoans(db, InsertionSort, fieldSelector, comparator)
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}
