package main

import (
	"fmt"
	"time"
)

type Student struct {
	Id           uint `aut:"pk"`
	PersonInfoTH StudentInfo
	PersonInfoEN StudentInfo
}

type StudentInfo struct {
	Prefix    string
	Firstname string
	Lastname  string
	Sex       string
}

func main() {
	start_time := time.Now()
	defer func() {
		fmt.Printf("use times : %f s\n", time.Since(start_time).Seconds())
	}()
	var m Mygodb
	m.Connect(Config{
		User: "postgres",
		Pass: "14032547",
		DBnm: "demo_database",
		Port: "5432",
		Host: "localhost",
	})
	m.AutoMigrate(&Student{})
}
