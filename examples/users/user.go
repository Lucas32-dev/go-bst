package main

import (
	"bst"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type User struct {
	Name  string
	Phone int
}

type UserPhoneIndex struct {
	User *User
}

func (u UserPhoneIndex) Value() int {
	return u.User.Phone
}

func main() {
	b := bst.New()

	users := randUsers(1_000_000)
	for i := range users {
		index := UserPhoneIndex{
			User: &users[i],
		}

		if err := b.Insert(index); err != nil {
			fmt.Printf("Fail to insert: %v\n", err)
			return
		}
	}

	phone1 := users[500_000].Phone

	traditionalSearch(users, phone1)
	bstSearch(b, phone1)
}

func traditionalSearch(users []User, value int) {
	start := time.Now()

	for _, u := range users {
		if u.Phone == value {
			break
		}
	}

	dur := time.Since(start)
	fmt.Printf("Traditional search: %v\n", dur)

}

func bstSearch(b bst.BinarySearchTree, value int) {
	start := time.Now()

	b.Search(value)

	dur := time.Since(start)
	fmt.Printf("BST search: %v\n", dur)
}

func randUsers(size int) []User {
	users := make([]User, size)

	var user User
	for i := 0; i < size; i++ {
		user.Name = "User" + strconv.Itoa(i)
		// there is a small chance of adding duplicated phone numbers,
		// which will result an bst error, but we are not covering that
		user.Phone = int(rand.Int())
		users[i] = user
	}

	return users
}
