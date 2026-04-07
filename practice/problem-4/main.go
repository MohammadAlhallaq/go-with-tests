package main

import (
	"errors"
	"fmt"
)

type NotFoundError struct {
	ID int
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("User with ID '%d' not found", e.ID)
}

type Database interface {
	FindUser(id int) (string, error)
}

type MemoryDB struct {
	users map[int]string
}

func (m MemoryDB) FindUser(id int) (string, error) {
	user, ok := m.users[id]
	if !ok {
		return "", NotFoundError{ID: id}
	}
	return user, nil
}

func loadUser(d Database, id int) {
	user, err := d.FindUser(id)
	if err != nil {
		var notFoundErr NotFoundError
		if errors.As(err, &notFoundErr) {
			fmt.Println("Not found:", notFoundErr)
			return
		}
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Loaded:", user)
}

func main() {
	db := MemoryDB{
		users: map[int]string{1: "Alice", 2: "Bob"},
	}

	loadUser(db, 1)
	loadUser(db, 99)
}
