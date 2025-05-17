package main

import (
	"fmt"
)

// Pure function to calculate interest rate based on loan amount
func CalculateInterestRate(nominal int) float64 {
	switch {
	case nominal < 10000000:
		return 0.05
	case nominal <= 50000000:
		return 0.08
	case nominal <= 100000000:
		return 0.15
	default:
		return 0.18
	}
}

// Pure function to calculate monthly installment
func CalculateMonthlyInstallment(amount int, interestRate float64, installments int) int {
	totalAmount := float64(amount) * (1 + interestRate)
	return int(totalAmount) / installments
}

// Pure function to display tenor options based on loan amount
func GetTenorOptions(nominal int) (string, string) {
	var tenorOptions, installmentOptions string

	switch {
	case nominal <= 300000:
		tenorOptions = "7 Hari, 14 Hari, 30 Hari"
		installmentOptions = "1x, 2x, 3x"
	case nominal <= 1000000:
		tenorOptions = "30 Hari, 60 Hari, 90 Hari"
		installmentOptions = "1x, 2x, 3x"
	case nominal <= 3000000:
		tenorOptions = "3 Bulan, 4 Bulan, 6 Bulan"
		installmentOptions = "3x, 4x, 6x"
	case nominal <= 5000000:
		tenorOptions = "3 Bulan, 6 Bulan, 9 Bulan"
		installmentOptions = "3x, 6x, 9x"
	case nominal <= 10000000:
		tenorOptions = "6 Bulan, 9 Bulan, 12 Bulan"
		installmentOptions = "6x, 9x, 12x"
	case nominal <= 20000000:
		tenorOptions = "12 Bulan, 18 Bulan, 24 Bulan"
		installmentOptions = "12x, 18x, 24x"
	case nominal <= 50000000:
		tenorOptions = "24 Bulan, 30 Bulan, 36 Bulan"
		installmentOptions = "24x, 30x, 36x"
	default:
		tenorOptions = "Nominal melebihi batas dukungan"
		installmentOptions = ""
	}

	return tenorOptions, installmentOptions
}

// Function to display tenor options to the user
func DisplayTenorOptions(nominal int) {
	tenorOptions, installmentOptions := GetTenorOptions(nominal)

	fmt.Printf("Nominal: Rp%d\n", nominal)
	fmt.Printf("Pilihan tenor: %s\n", tenorOptions)
	fmt.Printf("Cicilan: %s\n", installmentOptions)
}

// ApplyForLoan creates a new loan application
// Higher-order function that accepts a loan processor function
func ApplyForLoan(user *Pengguna, db *[100]Pinjaman, processor func(Pinjaman) string) {
	var (
		nominal      int
		tenor        int
		installments int
	)

	fmt.Println("\n=== AJUKAN PINJAMAN ===")
	fmt.Print("Masukkan nominal: ")
	fmt.Scan(&nominal)

	DisplayTenorOptions(nominal)

	fmt.Print("Pilih tenor (dalam bulan/hari): ")
	fmt.Scan(&tenor)
	fmt.Print("Pilih jumlah angsuran: ")
	fmt.Scan(&installments)

	// Use pure functions for calculations
	interestRate := CalculateInterestRate(nominal)
	monthlyInstallment := CalculateMonthlyInstallment(nominal, interestRate, installments)

	// Create new loan with immutable pattern
	newLoan := NewPinjaman(
		user.idPengguna,
		nominal,
		tenor,
		installments,
		monthlyInstallment,
		interestRate,
		false,
	)

	// Find empty index in the database
	emptyIndex := FindEmptyLoanIndex(*db)
	if emptyIndex == -1 {
		fmt.Println("Database pinjaman penuh, tidak dapat menambah pinjaman baru.")
		return
	}

	// Add the loan to the database
	(*db)[emptyIndex] = newLoan

	// Process the loan using the provided function
	message := processor(newLoan)
	fmt.Println(message)

	// Display loan details
	fmt.Println("\n=== PINJAMAN BERHASIL DIAJUKAN ===")
	fmt.Printf("Jumlah Pinjaman: Rp%d\n", nominal)
	fmt.Printf("Bunga: %.2f%%\n", interestRate*100)
	fmt.Printf("Total Pinjaman: Rp%d\n", int(float64(nominal)*(1+interestRate)))
	fmt.Printf("Tenor: %d bulan/hari\n", tenor)
	fmt.Printf("Jumlah Angsuran: %d kali\n", installments)
	fmt.Printf("Angsuran per Bulan: Rp%d\n", monthlyInstallment)
}

// Function to display user's loans
func ViewUserLoans(user *Pengguna, db *[100]Pinjaman) {
	fmt.Println("\n=== PINJAMAN SAYA ===")

	// Use a filter function to get user loans
	userLoans := FilterLoans(*db, func(p Pinjaman) bool {
		return p.idPeminjam == user.idPengguna
	})

	if len(userLoans) == 0 {
		fmt.Println("Anda belum memiliki pinjaman.")
		return
	}

	for i, loan := range userLoans {
		fmt.Printf("- Pinjaman ke-%d\n", i+1)
		fmt.Printf("  Jumlah Pinjaman: Rp%d\n", loan.jumlahPinjaman)
		fmt.Printf("  Tenor: %d bulan/hari\n", loan.tenor)
		fmt.Printf("  Bunga: %.2f%%\n", loan.bunga*100)
		fmt.Printf("  Total Pinjaman: Rp%d\n", int(float64(loan.jumlahPinjaman)*(1+loan.bunga)))
		fmt.Printf("  Jumlah Angsuran: %d kali\n", loan.jumlahAngsuran)
		fmt.Printf("  Angsuran per Bulan: Rp%d\n", loan.angsuranBulanan)

		statusStr := "BELUM LUNAS"
		if loan.statusLunas {
			statusStr = "LUNAS"
		}
		fmt.Printf("  Status: %s\n", statusStr)
		fmt.Println("-----------------------------")
	}
}

// Repay a loan
func RepayLoan(user *Pengguna, db *[100]Pinjaman) {
	var (
		loanIndex int
		option    int
	)

	fmt.Println("\n=== PELUNASAN PINJAMAN ===")

	// Filter loans that are not paid yet
	unpaidLoans := FilterLoans(*db, func(p Pinjaman) bool {
		return p.idPeminjam == user.idPengguna && !p.statusLunas
	})

	if len(unpaidLoans) == 0 {
		fmt.Println("Anda tidak memiliki pinjaman yang belum lunas.")
		return
	}

	// Display loans
	for _, loan := range unpaidLoans {
		fmt.Printf("[%d] Pinjaman Rp%d (Tenor: %d)\n", loan.idPeminjam, loan.jumlahPinjaman, loan.tenor)
	}

	fmt.Print("Pilih pinjaman yang akan dilunasi: ")
	fmt.Scan(&loanIndex)

	// Validate index
	found := false
	for i := 0; i < len(*db); i++ {
		if (*db)[i].idPeminjam == user.idPengguna && i == loanIndex {
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Pilihan tidak valid.")
		return
	}

	fmt.Println("\n[1] Lunasi Seluruhnya")
	fmt.Println("[2] Bayar Angsuran")
	fmt.Print("Pilih metode pelunasan: ")
	fmt.Scan(&option)

	switch option {
	case 1:
		// Update loan with immutable pattern
		(*db)[loanIndex] = (*db)[loanIndex].WithStatus(true)
		fmt.Println("Pinjaman berhasil dilunasi seluruhnya.")
	case 2:
		// Implementation for installment payment
		fmt.Println("Pembayaran angsuran berhasil.")
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

// Update loan information
func UpdateLoanData(user *Pengguna, db *[100]Pinjaman) {
	var (
		loanIndex int
		option    int
		newAmount int
		newTenor  int
	)

	fmt.Println("\n=== UBAH DATA PINJAMAN ===")

	// Filter active loans
	activeLoans := FilterLoans(*db, func(p Pinjaman) bool {
		return p.idPeminjam == user.idPengguna && !p.statusLunas
	})

	if len(activeLoans) == 0 {
		fmt.Println("Anda tidak memiliki pinjaman yang dapat diubah.")
		return
	}

	// Display loans
	for i, loan := range activeLoans {
		fmt.Printf("[%d] Pinjaman Rp%d (Tenor: %d)\n", i, loan.jumlahPinjaman, loan.tenor)
	}

	fmt.Print("Pilih pinjaman yang akan diubah: ")
	fmt.Scan(&loanIndex)

	// Validate index
	if loanIndex < 0 || loanIndex >= len(*db) || (*db)[loanIndex].idPeminjam != user.idPengguna {
		fmt.Println("Pilihan tidak valid.")
		return
	}

	fmt.Println("\n[1] Ubah Jumlah Pinjaman")
	fmt.Println("[2] Ubah Tenor")
	fmt.Print("Pilih data yang akan diubah: ")
	fmt.Scan(&option)

	currentLoan := (*db)[loanIndex]
	var updatedLoan Pinjaman

	switch option {
	case 1:
		fmt.Print("Masukkan jumlah pinjaman baru: ")
		fmt.Scan(&newAmount)

		// Calculate new interest rate
		newInterest := CalculateInterestRate(newAmount)

		// Calculate new monthly installment
		newInstallment := CalculateMonthlyInstallment(
			newAmount,
			newInterest,
			currentLoan.jumlahAngsuran,
		)

		// Create updated loan with immutable pattern
		updatedLoan = currentLoan.
			WithJumlahPinjaman(newAmount).
			WithBunga(newInterest).
			WithAngsuranBulanan(newInstallment)

		// Update database
		(*db)[loanIndex] = updatedLoan

		fmt.Println("Jumlah pinjaman berhasil diubah.")

	case 2:
		fmt.Print("Masukkan tenor baru: ")
		fmt.Scan(&newTenor)

		// Update loan with immutable pattern
		(*db)[loanIndex] = currentLoan.WithTenor(newTenor)

		fmt.Println("Tenor berhasil diubah.")

	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

// Delete a loan
func DeleteLoan(user *Pengguna, db *[100]Pinjaman) {
	var loanIndex int

	fmt.Println("\n=== HAPUS DATA PINJAMAN ===")

	// Filter user loans
	userLoans := FilterLoans(*db, func(p Pinjaman) bool {
		return p.idPeminjam == user.idPengguna
	})

	if len(userLoans) == 0 {
		fmt.Println("Anda tidak memiliki pinjaman.")
		return
	}

	// Display loans
	for i, loan := range userLoans {
		statusStr := "BELUM LUNAS"
		if loan.statusLunas {
			statusStr = "LUNAS"
		}
		fmt.Printf("[%d] Pinjaman Rp%d (Status: %s)\n", i, loan.jumlahPinjaman, statusStr)
	}

	fmt.Print("Pilih pinjaman yang akan dihapus: ")
	fmt.Scan(&loanIndex)

	// Validate index
	if loanIndex < 0 || loanIndex >= len(*db) || (*db)[loanIndex].idPeminjam != user.idPengguna {
		fmt.Println("Pilihan tidak valid.")
		return
	}

	// Delete loan by setting to empty value
	(*db)[loanIndex] = Pinjaman{}

	fmt.Println("Pinjaman berhasil dihapus.")
}

// Higher-order function to filter loans based on a predicate
func FilterLoans(loans [100]Pinjaman, predicate func(Pinjaman) bool) []Pinjaman {
	var result []Pinjaman
	for i := 0; i < len(loans); i++ {
		if loans[i].idPeminjam != 0 && predicate(loans[i]) {
			result = append(result, loans[i])
		}
	}
	return result
}
