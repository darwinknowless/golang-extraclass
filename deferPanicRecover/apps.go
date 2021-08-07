package main

import (
	"fmt"
)

// recover
func catch() {
	var r = recover()
	if r != nil {
		fmt.Println("error", r)
	} else {
		fmt.Println("aplikasi berjalan dengan lancar")
	}
}

func main() {

	// defer : proses menjalankan perintah paling terakhir (contoh : untuk close connect ke DB)
	// defer fmt.Println("Ini di jalankan terakhir")
	// fmt.Println("Ini di jalankan pertama")
	defer catch()
	// panic : proses untuk menghetikan go routine(banyak proses), trigger error
	panic("error panic")

	// recover : mirip catch di try catch
}
