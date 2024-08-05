package services

import (
    "library_management/models"
    "testing"
)

func TestLibrary_AddBook(t *testing.T) {
    library := NewLibrary()
    book := models.Book{ID: 1, Title: "Test Book", Author: "Test Author", Status: "Available"}
    library.AddBook(book)

    if len(library.Books) != 1 {
        t.Errorf("expected 1 book, got %d", len(library.Books))
    }
}

func TestLibrary_RemoveBook(t *testing.T) {
    library := NewLibrary()
    book := models.Book{ID: 1, Title: "Test Book", Author: "Test Author", Status: "Available"}
    library.AddBook(book)
    err := library.RemoveBook(1)
    if err != nil {
        t.Errorf("unexpected error: %v", err)
    }

    if len(library.Books) != 0 {
        t.Errorf("expected 0 books, got %d", len(library.Books))
    }
}

func TestLibrary_BorrowBook(t *testing.T) {
    library := NewLibrary()
    book := models.Book{ID: 1, Title: "Test Book", Author: "Test Author", Status: "Available"}
    member := models.Member{ID: 1, Name: "Test Member"}

    library.AddBook(book)
    library.Members[1] = member

    err := library.BorrowBook(1, 1)
    if err != nil {
        t.Errorf("unexpected error: %v", err)
    }

    if library.Books[1].Status != "Borrowed" {
        t.Errorf("expected book status to be 'Borrowed', got '%s'", library.Books[1].Status)
    }

    if len(library.Members[1].BorrowedBooks) != 1 {
        t.Errorf("expected 1 borrowed book, got %d", len(library.Members[1].BorrowedBooks))
    }
}

func TestLibrary_ReturnBook(t *testing.T) {
    library := NewLibrary()
    book := models.Book{ID: 1, Title: "Test Book", Author: "Test Author", Status: "Available"}
    member := models.Member{ID: 1, Name: "Test Member"}

    library.AddBook(book)
    library.Members[1] = member
    library.BorrowBook(1, 1)

    err := library.ReturnBook(1, 1)
    if err != nil {
        t.Errorf("unexpected error: %v", err)
    }

    if library.Books[1].Status != "Available" {
        t.Errorf("expected book status to be 'Available', got '%s'", library.Books[1].Status)
    }

    if len(library.Members[1].BorrowedBooks) != 0 {
        t.Errorf("expected 0 borrowed books, got %d", len(library.Members[1].BorrowedBooks))
    }
}

func TestLibrary_ListAvailableBooks(t *testing.T) {
    library := NewLibrary()
    book1 := models.Book{ID: 1, Title: "Book 1", Author: "Author 1", Status: "Available"}
    book2 := models.Book{ID: 2, Title: "Book 2", Author: "Author 2", Status: "Borrowed"}

    library.AddBook(book1)
    library.AddBook(book2)

    availableBooks := library.ListAvailableBooks()
    if len(availableBooks) != 1 {
        t.Errorf("expected 1 available book, got %d", len(availableBooks))
    }
}

func TestLibrary_ListBorrowedBooks(t *testing.T) {
    library := NewLibrary()
    book := models.Book{ID: 1, Title: "Test Book", Author: "Test Author", Status: "Available"}
    member := models.Member{ID: 1, Name: "Test Member"}

    library.AddBook(book)
    library.Members[1] = member
    library.BorrowBook(1, 1)

    borrowedBooks := library.ListBorrowedBooks(1)
    if len(borrowedBooks) != 1 {
        t.Errorf("expected 1 borrowed book, got %d", len(borrowedBooks))
    }
}
