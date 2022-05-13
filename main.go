package main

import "fmt"

func main() {
	//condition
	var point = 8

	if point == 10 {
		fmt.Println("Lulus dengan nilai sempurna")
	} else if point > 5 {
		fmt.Println("Lulus")

	} else if point == 4 {
		fmt.Println("hampir lulus")

	} else {
		fmt.Printf("Tidak lulus, nilai anda %d\n", point)
	}

	//switch case
	switch point {
	case 7:
		fmt.Println("Lulus")

	case 5:
		fmt.Println("so so")
	case 2, 3, 4:
		fmt.Println("Awesome")
	default:
		fmt.Println("bad")

	}

	switch {
	case (point < 8) && (point > 3):
		fmt.Println("Great")
	case (point == 8):
		fmt.Println("Luar biasa")
	default:
		fmt.Println("coba lagi")
	}

	//Perulangan

	for i := 0; i < 5; i++ {
		fmt.Printf("Nilainya : %d \n", i)
	}

	var a = 0
	for a < 5 {
		fmt.Println("Angka", a)
		a++
	}

	var b = 0
	for {
		fmt.Println("Nilai:", b)
		b++

		if b == 5 {
			break
		}
	}

	for c := 0; c < 5; c++ {
		for d := c; d < 5; d++ {
			fmt.Print(d, " ")
		}
		fmt.Println()
	}

	//Array
	var names = [4]string{"nina", "noni", "nani", "neno"}
	var fruits [3]string

	fruits[0] = "grape"
	fruits[1] = "banana"
	fruits[2] = "Apple"

	fmt.Println("nama  : ", names)
	fmt.Println("banyak array :", len(names))
	fmt.Println("isi semua element : ", fruits)

	//inisialisasi array tanpa jumlah elemen

	var angka = [...]int{1, 2, 3, 4}
	fmt.Println("angka : ", angka)
	fmt.Println(len(angka))

	//array multidimensi
	var number1 = [2][3]int{[3]int{1, 2, 3}, [3]int{4, 5, 6}}
	var number2 = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("numbers1 :", number1)
	fmt.Println("numbers2 :", number2)

	//perulangan elemen array menggunakan for
	for g := 0; g < len(names); g++ {
		fmt.Printf("elemen ke %d : %s\n", g, names[g])

	}

	//for range
	var tempVar string
	for i, name := range names {
		fmt.Printf("elemen ke %d : %s\n", i, name)
		if i == 3 {
			tempVar = name
			fmt.Printf("temp variable: %s", tempVar)
		}
	}
	fmt.Println()

	//Penggunaan Variabel Underscore _ Dalam for - range
	for _, name := range names {
		fmt.Printf("nama  : %s\n", name)
	}

	//slice
	var foods = []string{"apple", "banana", "melon", "grape"}
	var newFoods = foods[0:2]
	fmt.Println(newFoods)

	//fungsi append pada slice
	var newFoods1 = append(foods, "papaya")
	fmt.Println(newFoods1)

	


}
