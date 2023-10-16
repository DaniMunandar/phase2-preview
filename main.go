package main

// Task1
// import (
// 	"fmt"
// 	"time"
// )

// func printNumbers() {
// 	for i := 1; i <= 10; i++ {
// 		fmt.Printf("%d \n", i)
// 		time.Sleep(100 * time.Millisecond) // Menambahkan sleep untuk simulasi pekerjaan
// 	}
// }

// func printLetters() {
// 	for ch := 'a'; ch <= 'j'; ch++ {
// 		fmt.Printf("%c \n", ch)
// 		time.Sleep(100 * time.Millisecond) // Menambahkan sleep untuk simulasi pekerjaan
// 	}
// }

// func main() {
// 	go printNumbers()
// 	go printLetters()

// 	// Menunggu beberapa saat agar goroutine selesai
// 	time.Sleep(2 * time.Second)
// }

// Task 2
// import (
// 	"fmt"
// 	"sync"
// )

// func printNumbers(wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	for i := 1; i <= 10; i++ {
// 		fmt.Println(i)
// 	}
// }

// func printLetters(wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	for char := 'a'; char <= 'j'; char++ {
// 		fmt.Printf("%c\n", char)
// 	}
// }

// func main() {
// 	var wg sync.WaitGroup

// 	// Menambahkan 2 ke WaitGroup untuk kedua goroutine yang akan digunakan
// 	wg.Add(2)

// 	// Memulai goroutine untuk mencetak angka
// 	go printNumbers(&wg)

// 	// Memulai goroutine untuk mencetak huruf
// 	go printLetters(&wg)

// 	// Menunggu sampai kedua goroutine selesai
// 	wg.Wait()
// }

// Task 3
// import (
// 	"fmt"
// 	"sync"
// )

// func produksi(ch chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := 1; i <= 10; i++ {
// 		ch <- i // Mengirim angka ke saluran
// 	}
// 	close(ch) // Menutup saluran setelah selesai mengirimkan data
// }

// func konsumsi(ch chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for angka := range ch {
// 		fmt.Printf("%d \n", angka) // Menerima dan mencetak angka dari saluran
// 	}
// }

// func main() {
// 	// Membuat saluran (channel)
// 	ch := make(chan int)

// 	var wg sync.WaitGroup
// 	wg.Add(2) // Menambahkan 2 goroutine ke WaitGroup

// 	// Menjalankan goroutine "produksi"
// 	go produksi(ch, &wg)

// 	// Menjalankan goroutine "konsumsi"
// 	go konsumsi(ch, &wg)

// 	// Menunggu kedua goroutine selesai
// 	wg.Wait()
// }

// Task 4
// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func produksiBuffered(ch chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := 1; i <= 5; i++ {
// 		ch <- i // Mengirim data ke saluran buffered
// 		fmt.Printf("Pengirim Buffered: Mengirim %d\n", i)
// 	}
// 	close(ch)
// }

// func konsumsiBuffered(ch chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for data := range ch {
// 		time.Sleep(1 * time.Second) // Simulasi pemrosesan data
// 		fmt.Printf("Penerima Buffered: Menerima %d\n", data)
// 	}
// }

// func produksiUnbuffered(ch chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := 1; i <= 5; i++ {
// 		ch <- i // Mengirim data ke saluran tanpa buffer
// 		fmt.Printf("Pengirim Unbuffered: Mengirim %d\n", i)
// 	}
// 	close(ch)
// }

// func konsumsiUnbuffered(ch chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for data := range ch {
// 		time.Sleep(1 * time.Second) // Simulasi pemrosesan data
// 		fmt.Printf("Penerima Unbuffered: Menerima %d\n", data)
// 	}
// }

// func main() {
// 	// Saluran buffered dengan buffer berukuran 2
// 	bufferedCh := make(chan int, 2)
// 	// Saluran tanpa buffer
// 	unbufferedCh := make(chan int)

// 	var wg sync.WaitGroup
// 	wg.Add(4) // Menambahkan 4 goroutine ke WaitGroup

// 	// Menjalankan goroutine "produksiBuffered"
// 	go produksiBuffered(bufferedCh, &wg)

// 	// Menjalankan goroutine "konsumsiBuffered"
// 	go konsumsiBuffered(bufferedCh, &wg)

// 	// Menjalankan goroutine "produksiUnbuffered"
// 	go produksiUnbuffered(unbufferedCh, &wg)

// 	// Menjalankan goroutine "konsumsiUnbuffered"
// 	go konsumsiUnbuffered(unbufferedCh, &wg)

// 	wg.Wait() // Menunggu semua goroutine selesai
// }

// Task 5
// import (
// 	"fmt"
// )

// func main() {
// 	// Membuat saluran untuk nomor genap dan ganjil
// 	genapCh := make(chan int)
// 	ganjilCh := make(chan int)

// 	// Goroutine untuk mengirim nomor ke saluran genap atau ganjil
// 	go func() {
// 		for i := 1; i <= 20; i++ {
// 			if i%2 == 0 {
// 				genapCh <- i // Mengirim nomor genap ke saluran genap
// 			} else {
// 				ganjilCh <- i // Mengirim nomor ganjil ke saluran ganjil
// 			}
// 		}
// 		close(genapCh)
// 		close(ganjilCh)
// 	}()

// 	// Menerima dan mencetak nomor genap atau ganjil dari saluran
// 	for {
// 		select {
// 		case genap, ok := <-genapCh:
// 			if !ok {
// 				genapCh = nil // Menonaktifkan saluran
// 				break
// 			}
// 			fmt.Printf("Received even number: %d\n", genap)
// 		case ganjil, ok := <-ganjilCh:
// 			if !ok {
// 				ganjilCh = nil // Menonaktifkan saluran
// 				break
// 			}
// 			fmt.Printf("Received odd number: %d\n", ganjil)
// 		}

// 		// Keluar dari loop jika kedua saluran telah ditutup
// 		if genapCh == nil && ganjilCh == nil {
// 			break
// 		}
// 	}
// }

// 	Task 6
import "fmt"

func main() {
	oddCh := make(chan int)
	evenCh := make(chan int)
	errorCh := make(chan int)
	go numbers(oddCh, evenCh, errorCh)

	for i := 0; i < 22; i++ {
		select {
		case even := <-evenCh:
			fmt.Printf("Received an even number\t: %v\n", even)

		case odd := <-oddCh:
			fmt.Printf("Received an odd number\t: %v\n", odd)

		case error := <-errorCh:
			fmt.Printf("Error: number %v is greater than 20\n", error)
		}
	}
}

func numbers(oddCh chan<- int, evenCh chan<- int, errorCh chan<- int) {
	for i := 1; i <= 22; i++ {
		if i > 20 {
			errorCh <- i
		} else if i%2 == 0 {
			evenCh <- i
		} else {
			oddCh <- i
		}
	}
	close(errorCh)
	close(oddCh)
	close(evenCh)
}
