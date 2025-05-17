package main

// SeedData initializes the database with sample data
func SeedData() {
	// Initialize user data using the immutable pattern
	dbPengguna[0] = NewPengguna(1, "Agha", "agha@example.com", "password123")
	dbPengguna[1] = NewPengguna(2, "Elfan", "elfangamtenk@example.com", "elfan123")
	dbPengguna[2] = NewPengguna(3, "Citra", "citra@example.com", "citra456")
	dbPengguna[3] = NewPengguna(4, "Farel", "farel67@example.com", "farel456")
	dbPengguna[4] = NewPengguna(5, "Geby", "geby22@example.com", "geby456")
	dbPengguna[5] = NewPengguna(6, "Caca", "cacaboom@example.com", "caca456")
	dbPengguna[6] = NewPengguna(7, "Geebry", "geebry@example.com", "1234567")

	// Initialize loan data with predefined interest rates
	bunga1 := 0.05 // 5%
	bunga2 := 0.04 // 4%
	bunga3 := 0.06 // 6%

	// Calculate installments using pure functions
	angsuran1 := CalculateMonthlyInstallment(1000000, bunga1, 12)
	angsuran2 := CalculateMonthlyInstallment(2000000, bunga2, 12)
	angsuran3 := CalculateMonthlyInstallment(1500000, bunga3, 12)
	angsuran4 := CalculateMonthlyInstallment(2000000, bunga3, 12)
	angsuran5 := CalculateMonthlyInstallment(2500000, bunga3, 12)
	angsuran6 := CalculateMonthlyInstallment(1000000, bunga1, 12)
	angsuran7 := CalculateMonthlyInstallment(1200000, bunga1, 12)

	// Create loans using the immutable pattern
	dbDataPeminjam[0] = NewPinjaman(1, 1000000, 12, 12, angsuran1, bunga1, false)
	dbDataPeminjam[1] = NewPinjaman(2, 2000000, 24, 12, angsuran2, bunga2, true)
	dbDataPeminjam[2] = NewPinjaman(3, 1500000, 18, 12, angsuran3, bunga3, false)
	dbDataPeminjam[3] = NewPinjaman(4, 2000000, 18, 12, angsuran4, bunga3, false)
	dbDataPeminjam[4] = NewPinjaman(5, 2500000, 18, 12, angsuran5, bunga3, false)
	dbDataPeminjam[5] = NewPinjaman(6, 1000000, 12, 12, angsuran6, bunga1, true)
	dbDataPeminjam[6] = NewPinjaman(7, 1200000, 12, 12, angsuran7, bunga1, false)
}
