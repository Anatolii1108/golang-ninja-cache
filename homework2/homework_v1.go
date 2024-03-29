package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

var actions = []string{"logged in", "logged out", "created record", "deleted record", "updated account"}

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
	rand.Seed(time.Now().Unix()) // deprecated

	startTime := time.Now()

	const usersCount = 100
	const workersCount = 10
	users := make(chan User, usersCount)

	for i := 0; i < workersCount; i++ {
		go saveUserInfoWorker(i+1, users)
	}

	for i := 0; i < usersCount; i++ {
		generateUser(i+1, users)
	}

	fmt.Printf("DONE! Time Elapsed: %.2f seconds\n", time.Since(startTime).Seconds())
}

func saveUserInfoWorker(id int, users <-chan User) {
	for user := range users {
		fmt.Printf("WRITING FILE FOR UID %d | WORKER %d\n", user.id, id)

		filename := fmt.Sprintf("users/uid%d.txt", user.id)
		file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			log.Fatal(err)
		}

		file.WriteString(user.getActivityInfo())
		time.Sleep(time.Second)
	}
}

func generateUser(id int, users chan<- User) {
	users <- User{
		id:    id,
		email: fmt.Sprintf("user%d@company.com", id),
		logs:  generateLogs(rand.Intn(1000)),
	}
	fmt.Printf("generated user %d\n", id)
	time.Sleep(time.Millisecond * 100)
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
