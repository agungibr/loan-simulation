package main

import (
	"fmt"
	"strings"
)

// MaxLoginAttempts defines the maximum number of login attempts before the program exits
const MaxLoginAttempts = 3

// RegisterUser handles the registration process (referentially transparent)
// Returns the new user and a success status
func RegisterUser(db [100]Pengguna) (Pengguna, [100]Pengguna, bool) {
	var (
		nama     string
		email    string
		password string
		confirm  string
	)

	fmt.Println("\n=== REGISTER NEW USER ===")
	fmt.Print("Masukkan nama: ")
	fmt.Scan(&nama)

	// Check if the username already exists
	for i := 0; i < len(db); i++ {
		if strings.EqualFold(nama, db[i].nama) && db[i].idPengguna != 0 {
			fmt.Println("Nama pengguna sudah digunakan. Silakan gunakan nama lain.")
			return Pengguna{}, db, false
		}
	}

	fmt.Print("Masukkan email: ")
	fmt.Scan(&email)
	fmt.Print("Masukkan password: ")
	fmt.Scan(&password)
	fmt.Print("Konfirmasi password: ")
	fmt.Scan(&confirm)

	if password != confirm {
		fmt.Println("Password dan konfirmasi tidak cocok!")
		return Pengguna{}, db, false
	}

	// Create new user with immutable pattern
	newId := GetNextUserId(db)
	newUser := NewPengguna(newId, nama, email, password)

	// Find empty slot and add user
	emptyIndex := FindEmptyUserIndex(db)
	if emptyIndex == -1 {
		fmt.Println("Database pengguna penuh!")
		return Pengguna{}, db, false
	}

	// Create a new copy of the database with the new user
	var newDb [100]Pengguna
	copy(newDb[:], db[:])
	newDb[emptyIndex] = newUser

	fmt.Println("Registrasi berhasil! Silakan login.")
	return newUser, newDb, true
}

// LoginUser attempts to log in a user (referentially transparent)
// Returns the authenticated user and a success status
func LoginUser(db [100]Pengguna) (*Pengguna, bool) {
	var (
		nama     string
		password string
	)

	fmt.Println("\n=== LOGIN ===")
	fmt.Print("Masukkan nama: ")
	fmt.Scan(&nama)
	fmt.Print("Masukkan password: ")
	fmt.Scan(&password)

	for i := 0; i < len(db); i++ {
		if nama == db[i].nama && password == db[i].password && db[i].idPengguna != 0 {
			// Create a copy to maintain immutability
			userCopy := db[i]
			return &userCopy, true
		}
	}

	fmt.Println("Login gagal. Username atau password salah.")
	return nil, false
}

// HandleAuth manages the authentication process with limited attempts
// Higher-order function that takes a function for login handling
func HandleAuth(db [100]Pengguna, processLoginResult func(*Pengguna) bool) bool {
	var attempts int

	for {
		fmt.Println("\n=== TRAPINJAMAN ONLINE ===")
		fmt.Println("[1] Login")
		fmt.Println("[2] Register")
		fmt.Println("[0] Keluar")

		var choice int
		fmt.Print("Pilih: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			// Login process
			user, success := LoginUser(db)
			if success {
				return processLoginResult(user)
			}

			attempts++
			remainingAttempts := MaxLoginAttempts - attempts
			if remainingAttempts > 0 {
				fmt.Printf("Sisa percobaan login: %d\n", remainingAttempts)
			} else {
				fmt.Println("Anda telah melebihi batas percobaan login. Program akan berhenti.")
				return false
			}

		case 2:
			// Registration process
			_, newDb, success := RegisterUser(db)
			if success {
				// Update the global database
				for i := 0; i < len(newDb); i++ {
					dbPengguna[i] = newDb[i]
				}
			}

		case 0:
			fmt.Println("Terima kasih telah menggunakan TraPinjaman Online.")
			return false

		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
