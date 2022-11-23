using Go = import "/go.capnp";
@0x85d3acc39d94e0f8;
$Go.package("benchmark");
$Go.import("benchmark/books");

struct Book {
    title @0 :Text;
    # Title of the book.

    pageCount @1 :Int32;
    # Number of pages in the book.
}
