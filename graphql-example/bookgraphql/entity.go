package main

type Book struct {
	Id       string
	Title    string
	Author   string
	Price    float32
	Category string
}

func populate() []Book {
	books := []Book{
		{"102391123", "Harry Potter 1", "JK Rowling", 100000.0, "fiction"},
		{"102391242", "Harry Potter 2", "JK Rowling", 300000.0, "fiction"},
		{"102391111", "Harry Potter 3", "JK Rowling", 300000.0, "fiction"},
		{"102391231", "Harry Potter 4", "JK Rowling", 300000.0, "fiction"},
	}
	return books
}
