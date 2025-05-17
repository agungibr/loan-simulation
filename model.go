package main

// We'll use existing struct definitions from db[1].go, so we only need methods here

// Create a new Pinjaman (immutable pattern)
func NewPinjaman(idPeminjam, jumlahPinjaman, tenor, jumlahAngsuran, angsuranBulanan int, bunga float64, statusLunas bool) Pinjaman {
	return Pinjaman{
		idPeminjam:      idPeminjam,
		jumlahPinjaman:  jumlahPinjaman,
		tenor:           tenor,
		bunga:           bunga,
		jumlahAngsuran:  jumlahAngsuran,
		angsuranBulanan: angsuranBulanan,
		statusLunas:     statusLunas,
	}
}

// Create a new Pengguna (immutable pattern)
func NewPengguna(idPengguna int, nama, email, password string) Pengguna {
	return Pengguna{
		idPengguna: idPengguna,
		nama:       nama,
		email:      email,
		password:   password,
	}
}

// Returns a copy of a Pinjaman with modifications (immutable pattern)
func (p Pinjaman) WithJumlahPinjaman(jumlah int) Pinjaman {
	newPinjaman := p
	newPinjaman.jumlahPinjaman = jumlah
	return newPinjaman
}

func (p Pinjaman) WithTenor(tenor int) Pinjaman {
	newPinjaman := p
	newPinjaman.tenor = tenor
	return newPinjaman
}

func (p Pinjaman) WithStatus(statusLunas bool) Pinjaman {
	newPinjaman := p
	newPinjaman.statusLunas = statusLunas
	return newPinjaman
}

func (p Pinjaman) WithBunga(bunga float64) Pinjaman {
	newPinjaman := p
	newPinjaman.bunga = bunga
	return newPinjaman
}

func (p Pinjaman) WithAngsuranBulanan(angsuran int) Pinjaman {
	newPinjaman := p
	newPinjaman.angsuranBulanan = angsuran
	return newPinjaman
}

// Find the next available index in the loan database
// This is pure function (referential transparency)
func FindEmptyLoanIndex(db [100]Pinjaman) int {
	for i := 0; i < len(db); i++ {
		if db[i].idPeminjam == 0 {
			return i
		}
	}
	return -1
}

// Find the next available index in the user database
// This is pure function (referential transparency)
func FindEmptyUserIndex(db [100]Pengguna) int {
	for i := 0; i < len(db); i++ {
		if db[i].idPengguna == 0 {
			return i
		}
	}
	return -1
}

// Get the next available user ID
func GetNextUserId(db [100]Pengguna) int {
	maxId := 0
	for i := 0; i < len(db); i++ {
		if db[i].idPengguna > maxId {
			maxId = db[i].idPengguna
		}
	}
	return maxId + 1
}
