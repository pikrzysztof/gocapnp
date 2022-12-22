package benchmark

import (
	capnp "capnproto.org/go/capnp/v3"
	"fmt"
	"runtime"
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
	msg, seg, err := capnp.NewMessage(arena)
	if err != nil {
		panic(err)
	}
	msg.ResetReadLimit(1 << 63)

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
	bytes, err := b.TitleBytes()
	if err != nil {
		panic(err)
	}
	return 100 < len(bytes)
}

func loop(n int, book *Book, wg *sync.WaitGroup) {
	for i := 0; i < n; i++ {
		onerun(book)
	}
	wg.Done()
}

func BenchmarkAll(b *testing.B) {
	for numcpu := 1; numcpu <= runtime.NumCPU(); numcpu++ {
		b.Run(fmt.Sprintf("%d concurrency", numcpu), func(b *testing.B) {
			book := CreateBook()
			var wg sync.WaitGroup
			wg.Add(numcpu)
			for i := 0; i < numcpu; i++ {
				go loop(b.N, &book, &wg)
			}
			wg.Wait()
		})
	}
}
