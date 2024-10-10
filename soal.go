package main

import (
	"fmt"
	"math"
)

/*
Assignment A: Ride-Hailing

1. Hitung tarif perjalanan.
Sebagai Programmer di Startup Ride Hailing App, kamu diminta untuk membuat program hitung tarif perjalanan.

Pertama, setiap jenis kendaraan mempunyai tarif per KM sebagai berikut:
motor: Rp3000 / KM
mobil: Rp6000 / KM
taxi: Rp9000 / KM

Kedua, perhitungan tarif akan dikenakan tambahan biaya sebesar 20% jika waktu pemesanannya ada di dalam jam sibuk.
Terakhir, akan ada voucher diskon yang bisa dipakai untuk mengurangi harga tarif yang dikenakan.

Contoh Kasus:
User memesan perjalanan di dalam Jam Sibuk, untuk jarak 10KM dengan motor, memakai voucher Rp5000. Maka harga tarif yang dikenakan adalah Rp31000.
Penjelasan:
Harga tarif awal = (10 * 3000) = 30000
Dikenakan jam sibuk (+20%) = 30000 + 6000 = 36000
Harga akhir setelah voucher diskon = 36000 - 5000 = 31000 (jika desimal, output 2 angka dibelakang koma)

Notes:
- Pasti ada satu simbol "D"
- Minimal ada satu simbol "U"
*/

// calculateTravelFare, vehicle value one of "motor" | "mobil" | "taxi"

var fee int

func calculateTravelFare(vehicle string, distance int, isBusyHour bool, discount float64) float64 {
	switch vehicle {
	case "motor":
		fee = 3000
		if isBusyHour {
			travelFare := float64(distance * fee)
			addedFare := float64(travelFare * 20 / 100)
			totalFare := float64(travelFare + addedFare - discount)
			return totalFare
		} else {
			travelFare := float64(distance * fee)
			totalFare := float64(travelFare - discount)
			return totalFare
		}
	case "mobil":
		fee = 6000
		if isBusyHour {
			travelFare := float64(distance * fee)
			addedFare := float64(travelFare * 20 / 100)
			totalFare := float64(travelFare + addedFare - discount)
			return totalFare
		} else {
			travelFare := float64(distance * fee)
			totalFare := float64(travelFare - discount)
			return totalFare
		}
	case "taxi":
		fee = 9000
		if isBusyHour {
			travelFare := float64(distance * fee)
			addedFare := float64(travelFare * 20 / 100)
			totalFare := float64(travelFare + addedFare - discount)
			return totalFare
		} else {
			travelFare := float64(distance * fee)
			totalFare := float64(travelFare - discount)
			return totalFare
		}
	default:
		return 0
	}
}

/*
2. Cari User Terdekat dari Driver.
Dari sudut pandang Driver dari Ride-Hailing App, tentu saja kita harus carikan User yang dekat dengan Drivernya.
Anda diminta membuat program untuk mencari User terdekat pada jalan lurus.

Suatu jalan lurus dinyatakan seperti ini:
Contoh: ["", "U", "", "", "D", "", "", "", "U", "", ""]
Posisi user ditandai dengan karakter "U" dan Driver ojol ditandai dgn karakter "D".
Carilah posisi user yang terdekat Driver, dengan output index array user terdekat.

Jawaban dari contoh diatas: 1
Penjelasan:
Ada 2 user pada jalan lurus tersebut.
Untuk user pertama (index 1), jarak antara U dan D adalah 3.
Untuk user kedua (index 8), jarak antara U dan D adalah 4.
karena jarak antara U & D yang terdekat adalah 3, maka output index user pertama = 1.

Notes:
- Pasti ada satu simbol "D"
- Minimal ada satu simbol "U"
*/

// closestUser, output index array of closest User from Driver
func closestUser(road []string) int {

	driverIndex := -1
	closestUserIndex := -1
	minDistance := math.MaxInt32 // Use a large initial value

	// Find the index of the driver
	for i, char := range road {
		if char == "D" {
			driverIndex = i
			break // We can stop searching once we find the driver
		}
	}

	// Iterate to find the closest user
	for i, char := range road {
		if char == "U" {
			distance := int(math.Abs(float64(driverIndex - i))) // Calculate the absolute distance
			if distance < minDistance {
				minDistance = distance
				closestUserIndex = i
			}
		}
	}

	return closestUserIndex

}

func main() {
	fmt.Println(calculateTravelFare("motor", 10, true, 5000)) // contoh soal. output 31000
	fmt.Println(calculateTravelFare("mobil", 10, true, 5000)) // output 67000
	fmt.Println(calculateTravelFare("taxi", 10, true, 5000))  // output 103000

	road1 := []string{"", "U", "", "", "D", "", "", "", "U", "", ""}
	fmt.Println(closestUser(road1)) // contoh soal. output 1

	road2 := []string{"D", "", "", "", "", "", "", "", "", "U", "U"}
	fmt.Println(closestUser(road2)) // output 9

	road3 := []string{"", "", "", "U", "", "", "U", "", "U", "", "D"}
	fmt.Println(closestUser(road3)) // output 8
}
