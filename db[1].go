package main

type Pinjaman struct {
	idPeminjam      int
	jumlahPinjaman  int
	tenor           int
	bunga           float64
	jumlahAngsuran  int
	angsuranBulanan int // Menambahkan field untuk angsuran bulanan
	statusLunas     bool
}

type Pengguna struct {
	idPengguna int
	nama       string
	email      string
	password   string
}

var dbDataPeminjam [100]Pinjaman //variabel globalnya penyimpanan data
var dbPengguna [100]Pengguna     //database sederhana yg nyimpan array 100

func indeksKosongDbPeminjam(db *[100]Pinjaman) int { //karna ngirim alamatnya raray lgsg, jadi bukan duplikatnya
	for i := 0; i < len(db); i++ {
		if db[i].idPeminjam == 0 { //kalau idpeminjam si indeks i itu mai 0berarti masih kosong..belum diisi
			return i //dapat di indeks brp, kembaliin di nomor indeksnya gtu
		}
	}
	return -1 //kalo ga ada yg koosng kembaliin -1
}

// misal i=0 ... id peminjam=1 ...berrti lanjut
// i=1.... id peminjam=2.. lanjut
//i=2...idpeminjam =0.COCOK...RETURN 2

//ini sebagai laman sistem yang menyimpan dat abaru,jadi dia cari indeks ke brp dh ada id nya belum kalo blm ada. buatkan id sesuai indeks
