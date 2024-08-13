# Library Management System Documentation

## Introduction
This document outlines the functionality and structure of the Library Management System, a command-line tool built with Go. The system provides features to manage books, including adding, removing, borrowing, and returning books, as well as viewing available and borrowed books.

## Project Structure
- **`main.go`**: The main entry point for the application.
- **`controllers/`**: Houses the logic for processing user commands.
  - **`library_controller.go`**: Oversees core application logic and user interactions.
- **`models/`**: Defines the data structures used in the application.
  - **`book.go`**: Contains the `Book` struct.
  - **`member.go`**: Contains the `Member` struct.
- **`services/`**: Implements the business logic and core operations.
  - **`library_service.go`**: Features the `Library` struct and `LibraryManager` interface.
- **`docs/`**: Includes documentation and supplementary materials.

## How to Use
1. **Add a New Book**: Input details to include a new book in the library catalog.
2. **Remove an Existing Book**: Delete a book from the library by specifying its ID.
3. **Borrow a Book**: Allow a member to check out a book using its ID.
4. **Return a Book**: Facilitate the return of a borrowed book by its ID.
5. **List All Available Books**: Show all books currently available for borrowing.
6. **List All Borrowed Books by a Member**: Display a list of books borrowed by a particular member.

## Running the Application
1. Open a terminal and navigate to the project directory.
2. Execute the application with the command `go run main.go`.
