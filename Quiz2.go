package main

import (
	"fmt"
	"strings"
)

func hitungParkir(durasi int, isMember bool, isLibur bool) float64 {
	var totalBiaya float64

	if durasi <= 2 {
		totalBiaya = 5000
	} else {
		totalBiaya = 5000 + float64(durasi-2)*2000
	}

	if isMember {
		if durasi <= 5 {
			totalBiaya = totalBiaya * 0.5
		} else {
			totalBiaya = totalBiaya * 0.7
		}
	}

	if isLibur {
		totalBiaya = totalBiaya + 3000
	}

	return totalBiaya
}

func main() {
	var durasiJam int
	var inputMember, inputLibur string
	var statusMember, statusLibur bool

	fmt.Println("=== SISTEM TARIF PARKIR ===")
	fmt.Print("Masukkan durasi parkir (jam): ")
	fmt.Scan(&durasiJam)
	fmt.Print("Apakah Member? (y/n): ")
	fmt.Scan(&inputMember)
	if strings.ToLower(inputMember) == "y" {
		statusMember = true
	}

	fmt.Print("Apakah Hari Libur? (y/n): ")
	fmt.Scan(&inputLibur)
	if strings.ToLower(inputLibur) == "y" {
		statusLibur = true
	}
	totalBayar := hitungParkir(durasiJam, statusMember, statusLibur)
	fmt.Printf("\nTotal Biaya Parkir: Rp %.0f\n", totalBayar)
}
