package main

import (
	"fmt"
)

// LoanStats represents statistics about loans
type LoanStats struct {
	TotalLoans       int
	TotalBorrowers   int
	TotalPaidLoans   int
	TotalUnpaidLoans int
	TotalLoanValue   int
	PaidPercentage   float64
}

// CalculateLoanStats computes statistics from loan data (pure function)
func CalculateLoanStats(db [100]Pinjaman) LoanStats {
	var stats LoanStats
	borrowerMap := make(map[int]bool)

	for i := 0; i < len(db); i++ {
		if db[i].idPeminjam != 0 {
			stats.TotalLoans++
			stats.TotalLoanValue += db[i].jumlahPinjaman

			if db[i].statusLunas {
				stats.TotalPaidLoans++
			} else {
				stats.TotalUnpaidLoans++
			}

			// Track unique borrowers
			borrowerMap[db[i].idPeminjam] = true
		}
	}

	stats.TotalBorrowers = len(borrowerMap)

	if stats.TotalLoans > 0 {
		stats.PaidPercentage = float64(stats.TotalPaidLoans) / float64(stats.TotalLoans) * 100
	}

	return stats
}

// DisplayLoanStats prints loan statistics
func DisplayLoanStats(stats LoanStats) {
	fmt.Println("\n=== LAPORAN PINJAMAN ===")
	fmt.Printf("Jumlah Peminjam: %d orang\n", stats.TotalBorrowers)
	fmt.Printf("Total Pinjaman: %d\n", stats.TotalLoans)
	fmt.Printf("Total Nilai Pinjaman: Rp%d\n", stats.TotalLoanValue)
	fmt.Printf("Status Lunas: %d pinjaman\n", stats.TotalPaidLoans)
	fmt.Printf("Status Belum Lunas: %d pinjaman\n", stats.TotalUnpaidLoans)
	fmt.Printf("Persentase Pinjaman Lunas: %.2f%%\n", stats.PaidPercentage)
}

// GenerateLoanReport creates and displays a loan report
func GenerateLoanReport(db *[100]Pinjaman) {
	// Calculate statistics using a pure function
	stats := CalculateLoanStats(*db)

	// Display the statistics
	DisplayLoanStats(stats)

	// Additional breakdown by loan amount ranges
	ShowLoanRangeBreakdown(*db)
}

// ShowLoanRangeBreakdown provides a breakdown of loans by amount ranges
func ShowLoanRangeBreakdown(db [100]Pinjaman) {
	// Define loan amount ranges
	ranges := []struct {
		Min   int
		Max   int
		Count int
		Value int
	}{
		{0, 1000000, 0, 0},
		{1000001, 5000000, 0, 0},
		{5000001, 10000000, 0, 0},
		{10000001, 50000000, 0, 0},
		{50000001, -1, 0, 0}, // -1 means no upper limit
	}

	// Count loans in each range
	for i := 0; i < len(db); i++ {
		if db[i].idPeminjam != 0 {
			amount := db[i].jumlahPinjaman
			for j := 0; j < len(ranges); j++ {
				if (amount >= ranges[j].Min) &&
					(ranges[j].Max == -1 || amount <= ranges[j].Max) {
					ranges[j].Count++
					ranges[j].Value += amount
					break
				}
			}
		}
	}

	// Display the breakdown
	fmt.Println("\n=== BREAKDOWN PINJAMAN BERDASARKAN JUMLAH ===")
	for i, r := range ranges {
		if r.Max == -1 {
			fmt.Printf("Range %d: Rp%d ke atas\n", i+1, r.Min)
		} else {
			fmt.Printf("Range %d: Rp%d - Rp%d\n", i+1, r.Min, r.Max)
		}
		fmt.Printf("  Jumlah Pinjaman: %d\n", r.Count)
		fmt.Printf("  Total Nilai: Rp%d\n", r.Value)
	}
}

// DisplayUserProfile shows the profile of the current user
func DisplayUserProfile(user *Pengguna) {
	fmt.Println("\n=== PROFIL SAYA ===")
	fmt.Printf("ID: %d\n", user.idPengguna)
	fmt.Printf("Nama: %s\n", user.nama)
	fmt.Printf("Email: %s\n", user.email)
}
