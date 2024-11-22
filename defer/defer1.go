// blok yang ditambahkan defer akan selelu dieksekussi terkahir, walapun program error. defer kan tetap diekesekusi
// defer akan menunda blok kode untuk dieksekusi

package main

import "fmt"

func main() {
	fmt.Println("Minum tuak")
	fmt.Println("Minum bir")
	defer fmt.Println("Mabuk, tidurrr.")
	fmt.Println("Minum arak 1 botol.")
}
