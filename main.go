package main

import (
	"fmt"
	"time"
	"math/rand"
	"os"
	"log"
)

var actions = []string {
	"logged in", 
	"logged out", 
	"created record", 
	"deleted record", 
	"updated account",
}

type logItem struct {
	action    string
	timestamp time.Time
}

type User struct {
	id    int
	email string
	logs  []logItem
}

func (u User) getActivityInfo() string {
	output := fmt.Sprintf("UID: %d; Email: %s;\nActivity Log:\n", u.id, u.email)
	for index, item := range u.logs {
		output += fmt.Sprintf("%d. [%s] at %s\n", index, item.action, item.timestamp.Format(time.RFC3339))
	}

	return output
}

func main() {
	startTime := time.Now()

	inputUsers := make(chan User, 100)
	outputUsers := make(chan int, 100)

	for i := 0; i < 100; i++ {
		go saveUserInfo(inputUsers, outputUsers)	
	}

	go generateUsers(100, inputUsers)

	for i := 0; i < 100; i++ {
		fmt.Printf("WRITING FILE FOR UID %d\n", <- outputUsers)
	} 
	close(outputUsers)

	fmt.Printf("DONE! Time Elapsed: %.2f seconds\n", time.Since(startTime).Seconds())
}

func saveUserInfo(input <-chan User, output chan<- int) {
	for j := range input {
		filename := fmt.Sprintf("users/uid%d.txt", j.id)
		file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			log.Fatal(err)
		}

		file.WriteString(j.getActivityInfo())
		output <- j.id
		time.Sleep(time.Second)
	}
}

func generateUsers(count int, users chan<- User) {

	for i := 0; i < count; i++ {
		users <- User{
			id:    i + 1,
			email: fmt.Sprintf("user%d@company.com", i+1),
			logs:  generateLogs(rand.Intn(100)),
		}
		fmt.Printf("generated user %d\n", i+1)
		time.Sleep(time.Millisecond * 100)
	}
	close(users)
}

func generateLogs(count int) []logItem {
	logs := make([]logItem, count)

	for i := 0; i < count; i++ {
		logs[i] = logItem{
			action:    actions[rand.Intn(len(actions)-1)],
			timestamp: time.Now(),
		}
	}

	return logs
}