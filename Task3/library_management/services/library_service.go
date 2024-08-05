package services

import (
    "errors"
    "library_management/models"
)

type LibraryManager interface {
    AddBook(book models.Book)
    RemoveBook(bookID int) error
    BorrowBook(bookID int, memberID int) error
    ReturnBook(bookID int, memberID int) error
    ListAvailableBooks() []models.Book
    ListBorrowedBooks(memberID int) []models.Book
}

type Library struct {
    Books   map[int]models.Book
    Members map[int]models.Member
}

func NewLibrary() *Library {
    return &Library{
        Books:   make(map[int]models.Book),
        Members: make(map[int]models.Member),
    }
}

func (l *Library) AddBook(book models.Book) {
    l.Books[book.ID] = book
}

func (l *Library) RemoveBook(bookID int) error {
    if _, exists := l.Books[bookID]; !exists {
        return errors.New("book not found")
    }
    delete(l.Books, bookID)
    return nil
}

func (l *Library) BorrowBook(bookID int, memberID int) error {
    book, bookExists := l.Books[bookID]
    if !bookExists {
        return errors.New("book not found")
    }
    if book.Status == "Borrowed" {
        return errors.New("book already borrowed")
    }
    member, memberExists := l.Members[memberID]
    if !memberExists {
        return errors.New("member not found")
    }
    book.Status = "Borrowed"
    l.Books[bookID] = book
    member.BorrowedBooks = append(member.BorrowedBooks, book)
    l.Members[memberID] = member
    return nil
}

func (l *Library) ReturnBook(bookID int, memberID int) error {
    book, bookExists := l.Books[bookID]
    if !bookExists {
        return errors.New("book not found")
    }
    member, memberExists := l.Members[memberID]
    if !memberExists {
        return errors.New("member not found")
    }
    found := false
    for i, borrowedBook := range member.BorrowedBooks {
        if borrowedBook.ID == bookID {
            member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
            found = true
            break
        }
    }
    if !found {
        return errors.New("book not borrowed by this member")
    }
    book.Status = "Available"
    l.Books[bookID] = book
    l.Members[memberID] = member
    return nil
}

func (l *Library) ListAvailableBooks() []models.Book {
    availableBooks := []models.Book{}
    for _, book := range l.Books {
        if book.Status == "Available" {
            availableBooks = append(availableBooks, book)
        }
    }
    return availableBooks
}

func (l *Library) ListBorrowedBooks(memberID int) []models.Book {
    member, memberExists := l.Members[memberID]
    if !memberExists {
        return nil
    }
    return member.BorrowedBooks
}
