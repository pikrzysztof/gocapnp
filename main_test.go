package benchmark

import (
	capnp "capnproto.org/go/capnp/v3"
	"sync"
	"testing"
)

func CreateBook() Book {

	arena := capnp.SingleSegment(nil)

	// Make a brand new empty message.  A Message allocates Cap'n Proto structs within
	// its arena.  For convenience, NewMessage also returns the root "segment" of the
	// message, which is needed to instantiate the Book struct.  You don't need to
	// understand segments and roots yet (or maybe ever), but if you're curious, messages
	// and segments are documented here:  https://capnproto.org/encoding.html
	_, seg, err := capnp.NewMessage(arena)
	if err != nil {
		panic(err)
	}

	// Create a new Book struct.  Every message must have a root struct.  Again, it is
	// not important to understand "root structs" at this point.  For now, just understand
	// that every type you instantiate needs to be a "root", unless you plan on assigning
	// it to another object.  When in doubt, use NewRootXXX.
	//
	// If you're insatiably curious, see:  https://capnproto.org/encoding.html#messages
	book, err := NewRootBook(seg)
	if err != nil {
		panic(err)
	}
	err = book.SetTitle("War and Peace")
	if err != nil {
		panic("Failed")
	}
	return book
}

func onerun(b *Book) bool {
	bytes, _ := b.TitleBytes()
	return 100 < len(bytes)
}

func loop(n int, book *Book, wg *sync.WaitGroup) {
	for i := 0; i < n; i++ {
		onerun(book)
	}
	wg.Done()
}

func spinGoroutines(ngos int, b *testing.B) {
	book := CreateBook()
	var wg sync.WaitGroup
	wg.Add(ngos)
	for i := 0; i < ngos; i++ {
		go loop(b.N, &book, &wg)
	}
	wg.Wait()
}

func BenchmarkOne(b *testing.B) {
	spinGoroutines(1, b)
}
func BenchmarkTwo(b *testing.B) {
	spinGoroutines(2, b)
}
func BenchmarkThree(b *testing.B) {
	spinGoroutines(3, b)
}
func BenchmarkFour(b *testing.B) {
	spinGoroutines(4, b)
}
func BenchmarkFive(b *testing.B) {
	spinGoroutines(5, b)
}
func BenchmarkSix(b *testing.B) {
	spinGoroutines(6, b)
}
func BenchmarkSeven(b *testing.B) {
	spinGoroutines(7, b)
}
func BenchmarkEight(b *testing.B) {
	spinGoroutines(8, b)
}
func BenchmarkNine(b *testing.B) {
	spinGoroutines(9, b)
}
func BenchmarkTen(b *testing.B) {
	spinGoroutines(10, b)
}
func BenchmarkEleven(b *testing.B) {
	spinGoroutines(11, b)
}
func BenchmarkTwelve(b *testing.B) {
	spinGoroutines(12, b)
}
func BenchmarkThirteen(b *testing.B) {
	spinGoroutines(13, b)
}
func BenchmarkFourteen(b *testing.B) {
	spinGoroutines(14, b)
}
func BenchmarkFifteen(b *testing.B) {
	spinGoroutines(15, b)
}
func BenchmarkSixteen(b *testing.B) {
	spinGoroutines(16, b)
}
func BenchmarkSeventeen(b *testing.B) {
	spinGoroutines(17, b)
}
func BenchmarkEighteen(b *testing.B) {
	spinGoroutines(18, b)
}
func BenchmarkNineteen(b *testing.B) {
	spinGoroutines(19, b)
}
func BenchmarkTwenty(b *testing.B) {
	spinGoroutines(20, b)
}
func BenchmarkTwentyOne(b *testing.B) {
	spinGoroutines(21, b)
}
func BenchmarkTwentyTwo(b *testing.B) {
	spinGoroutines(22, b)
}
func BenchmarkTwentyThree(b *testing.B) {
	spinGoroutines(23, b)
}
func BenchmarkTwentyFour(b *testing.B) {
	spinGoroutines(24, b)
}
func BenchmarkTwentyFive(b *testing.B) {
	spinGoroutines(25, b)
}
func BenchmarkTwentySix(b *testing.B) {
	spinGoroutines(26, b)
}
func BenchmarkTwentySeven(b *testing.B) {
	spinGoroutines(27, b)
}
func BenchmarkTwentyEight(b *testing.B) {
	spinGoroutines(28, b)
}
