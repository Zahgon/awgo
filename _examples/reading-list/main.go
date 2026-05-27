// Copyright (c) 2018 Dean Jackson <deanishe@deanishe.net>
// MIT Licence - http://opensource.org/licenses/MIT

/*
Workflow reading-list is a more advanced example of fuzzy filtering.

It contains a custom implementation of fuzzy.Sortable and its own filtering
function to keep books sorted by unread/unpublished/read status.

Regular fuzzy sorting only considers match quality, so with the query
"kingkiller", the Kingkiller Chronicle series would be sorted based on where
the term "kingkiller" appears in the title, i.e. shortest title first:

	The Doors of Stone (The Kingkiller Chronicle, #3) [unpublished]
	The Wise Man's Fear (The Kingkiller Chronicle, #2) [unread]
	The Name of the Wind (The Kingkiller Chronicle, #1) [read]

The custom implementation sorts by status then match quality, thus keeping
unread books before unpublished and read ones:

	The Wise Man's Fear (The Kingkiller Chronicle, #2) [unread]
	The Doors of Stone (The Kingkiller Chronicle, #3) [unpublished]
	The Name of the Wind (The Kingkiller Chronicle, #1) [read]
*/
package main

import (
	aw "github.com/deanishe/awgo"
)

// Reading status. We're going to use these to sort books, so unread books
// are first, then unpublished books, with read books at the bottom.
const (
	Unread = iota
	Unpublished
	Read
)

// Book is a book on the reading list.
type Book struct {
	ID     int64  // Goodreads ID of book
	Title  string // Book title
	Author string // Author name
	Status int    // read/unread/unpublished
}

// URL returns the Goodreads URL for book.
func (b Book) URL() string { _ = "STUB: not implemented"; return "" }

// Icon returns a workflow icon for Book.
func (b Book) Icon() *aw.Icon { _ = "STUB: not implemented"; return nil }

// Books sorts books by status then by title. It implements fuzzy.Sortable
// and therefore also sort.Interface.
type Books []Book

// Implement sort.Interface
func (s Books) Len() int      { _ = "STUB: not implemented"; return 0 }
func (s Books) Swap(i, j int) { _ = "STUB: not implemented"; return }

// Less sorts by status and then by title.
func (s Books) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// Keywords implements fuzzy.Sortable.
func (s Books) Keywords(i int) string { _ = "STUB: not implemented"; return "" }

// filterBooks preserves by-status sorting when fuzzy-sorting books, so unread
// books are always before unpublished books, with read books last.
func filterBooks(books []Book, query string) []Book {
	_ = "STUB: not implemented"
	// Per-status groups
	return nil
}

// Fuzzy-sort books, then add them to the appropriate groups

// Matching items sort to start, so ignore all books from here

// Merge groups, which are each fuzzy-sorted by match quality

var (
	wf *aw.Workflow // Our Workflow struct

	iconRead        = &aw.Icon{Value: "read.png"}
	iconUnread      = &aw.Icon{Value: "icon.png"}
	iconUnpublished = &aw.Icon{Value: "unpublished.png"}

	// Reading list
	books = []Book{
		{365, "Dirk Gently's Holistic Detective Agency (Dirk Gently, #1)", "Douglas Adams", Unread},
		{4982, "The Sirens of Titan", "Kurt Vonnegut Jr.", Unread},
		{7604, "Lolita", "Vladimir Nabokov", Unread},
		{7624, "Lord of the Flies", "William Golding", Read},
		{34492, "Wintersmith (Discworld, #35; Tiffany Aching, #3)", "Terry Pratchett", Read},
		{34517, "Reaper Man (Discworld, #11; Death, #2)", "Terry Pratchett", Read},
		{34532, "Hogfather (Discworld, #20; Death, #4)", "Terry Pratchett", Unread},
		{38447, "The Handmaid's Tale (The Handmaid's Tale, #1)", "Margaret Atwood", Unread},
		{66788, "Ring (Xeelee Sequence, #4)", "Stephen Baxter", Unread},
		{186074, "The Name of the Wind (The Kingkiller Chronicle, #1)", "Patrick Rothfuss", Read},
		{310612, "A Confederacy of Dunces", "John Kennedy Toole", Unread},
		{332613, "One Flew Over the Cuckoo's Nest", "Ken Kesey", Unread},
		{736131, "The Valley of Fear (Sherlock Holmes, #7)", "Arthur Conan Doyle", Read},
		{754713, "His Last Bow (Sherlock Holmes, #8)", "Arthur Conan Doyle", Unread},
		{1215032, "The Wise Man's Fear (The Kingkiller Chronicle, #2)", "Patrick Rothfuss", Unread},
		{10814687, "Whispers Under Ground (Rivers of London, #3)", "Ben Aaronovitch", Unread},
		{12111823, "The Winds of Winter (A Song of Ice and Fire, #6)", "George R.R. Martin", Unpublished},
		{18190723, "The Causal Angel (Jean le Flambeur, #3)", "Hannu Rajaniemi", Unread},
		{21032488, "The Doors of Stone (The Kingkiller Chronicle, #3)", "Patrick Rothfuss", Unpublished},
		{42086897, "How To: Absurd Scientific Advice for Common Real-World Problems", "Randall Munroe", Unread},
	}
)

func main() {
	wf = aw.New()
	wf.Run(run)
}

func run() {
	_ = "STUB: not implemented"
	// Use wf.Args so magic actions are handled
	return
}

// Disable UIDs so Alfred respects our sort order. Without this,
// it may bump read/unpublished books to the top of results, but
// we want to force them to always be below unread books.

// Sort by status

// Filter and keep by-status sorting

// Script Filter results
