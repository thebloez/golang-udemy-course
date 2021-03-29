package channel

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestBuatChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)
	go func() {
		time.Sleep(10 * time.Millisecond)
		channel <- "Ryan Safary Hidayat"
	}()

	data := <- channel
	fmt.Println(data)
}

func swapStringToChannel(channel chan string, str string){
	time.Sleep(10 * time.Millisecond)
	channel <- str
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go swapStringToChannel(channel, "Ryan")
	fmt.Println(<-channel)
	go swapStringToChannel(channel, "Dewi")
	fmt.Println(<-channel)
	go swapStringToChannel(channel, "Kanaya")
	fmt.Println(<-channel)
}


// Hanya bisa untuk men-set channel
func onlyIn(channel chan<- string) {
	time.Sleep(5 * time.Millisecond)
	channel <- "Ryan Safary Hidayat"
}


// Hanya bisa get(mengambil) data dari channel
func onlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data + " dari onlyOut")
}

func TestChannelInAndOut(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	// set into channel
	go onlyIn(channel)
	// get from channel
	go onlyOut(channel)

	time.Sleep(5 * time.Millisecond)
}

/*
	Secara default channel hanya menerima 1 data saja.
	Jika dibutuhkan karena case tertentu channel bisa di set untuk menerima lebih dari 1 data.
	Buffered Channel cocok sekali JIKA PENERIMA LEBIH LAMBAT DIBANDING PENGIRIM DATANYA.
*/
func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	// set channel
	channel <- "Ryan"
	channel <- "Dewi"
	channel <- "Kanaya"

	fmt.Println("jml data : ", len(channel))

	// ambil data dari channel
	println(<- channel)
	println(<- channel)
	println(<- channel)


	fmt.Println("------------")
	fmt.Println("Panjang Buffer : ", cap(channel)) // panjang buffer
	fmt.Println("Jumlah Data : ", len(channel)) // jumlah data di buffer
	//fmt.Println(<- channel)

}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i:=0; i < 10; i++ {
			channel <- "Iterasi ke " + strconv.Itoa(i)
		}
		// channel harus di close setelah pengulangan set ke channel
		close(channel)
	}()

	for x := range channel {
		fmt.Println(x)
	}
}




