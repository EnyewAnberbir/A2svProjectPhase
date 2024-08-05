package controllers

import (
    "bufio"
    "fmt"
    "library_management/models"
    "library_management/services"
    "os"
    "strconv"
    "strings"
)

func StartLibrarySystem() {
    library := services.NewLibrary()
    scanner := bufio.NewScanner(os.Stdin)

    for {
        fmt.Println("\nLibrary Management System")
        fmt.Println("1. Add a new book")
        fmt.Println("2. Remove an existing book")
        fmt.Println("3. Borrow a book")
        fmt.Println("4. Return a book")
        fmt.Println("5. List all available books")
        fmt.Println("6. List all borrowed books by a member")
        fmt.Println("7. Exit")

        fmt.Print("Enter your choice: ")
        scanner.Scan()
        choice := scanner.Text()

        switch choice {
        case "1":
            fmt.Print("Enter book ID: ")
            scanner.Scan()
            bookID, _ := strconv.Atoi(scanner.Text())

            fmt.Print("Enter book title: ")
            scanner.Scan()
            bookTitle := scanner.Text()

            fmt.Print("Enter book author: ")
            scanner.Scan()
            bookAuthor := scanner.Text()

            book := models.Book{
                ID:     bookID,
                Title:  bookTitle,
                Author: bookAuthor,
                Status: "Available",
            }
            library.AddBook(book)
            fmt.Println("Book added successfully!")

        case "2":
            fmt.Print("Enter book ID to remove: ")
            scanner.Scan()
            bookID, _ := strconv.Atoi(scanner.Text())

            err := library.RemoveBook(bookID)
            if err != nil {
                fmt.Println("Error:", err)
            } else {
                fmt.Println("Book removed successfully!")
            }

        case "3":
            fmt.Print("Enter book ID to borrow: ")
            scanner.Scan()
            bookID, _ := strconv.Atoi(scanner.Text())

            fmt.Print("Enter member ID: ")
            scanner.Scan()
            memberID, _ := strconv.Atoi(scanner.Text())

            err := library.BorrowBook(bookID, memberID)
            if err != nil {
                fmt.Println("Error:", err)
            } else {
                fmt.Println("Book borrowed successfully!")
            }

        case "4":
            fmt.Print("Enter book ID to return: ")
            scanner.Scan()
            bookID, _ := strconv.Atoi(scanner.Text())

            fmt.Print("Enter member ID: ")
            scanner.Scan()
            memberID, _ := strconv.Atoi(scanner.Text())

            err := library.ReturnBook(bookID, memberID)
            if err != nil {
                fmt.Println("Error:", err)
            } else {
                fmt.Println("Book returned successfully!")
            }

        case "5":
            availableBooks := library.ListAvailableBooks()
            fmt.Println("Available Books:")
            for _, book := range availableBooks {
                fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
            }

        case "6":
            fmt.Print("Enter member ID: ")
            scanner.Scan()
            memberID, _ := strconv.Atoi(scanner.Text())

            borrowedBooks := library.ListBorrowedBooks(memberID)
            fmt.Println("Borrowed Books:")
            for _, book := range borrowedBooks {
                fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
            }

        case "7":
            fmt.Println("Exiting...")
            return

        default:
            fmt.Println("Invalid choice, please try again.")
        }
    }
}
