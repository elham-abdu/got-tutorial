package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Welcome to the quiz game!")

    
    fmt.Printf("Enter your name here: ")
    name, _ := reader.ReadString('\n')
    name = strings.TrimSpace(name)
    fmt.Printf("Hello %v, Welcome to have fun!\n", name)

    
    fmt.Printf("Enter your age: ")
    var age int
    fmt.Scanln(&age)
    if age >= 10 {
        fmt.Printf("Yay %v, you are ready to play!\n", name)
    } else {
        fmt.Printf("Sorry %v, you can't play.\n", name)
        return
    }

    score := 0
    numQuestions := 3

    
    fmt.Printf("Who is the first person who invented computer?\nEnter your answer: ")
    answer1, _ := reader.ReadString('\n')
    answer1 = strings.TrimSpace(answer1)
    if strings.EqualFold(answer1, "Bashir Rameev") || strings.EqualFold(answer1, "Charles Babbage") {
        fmt.Println("Correct!")
        score += 1
    } else {
        fmt.Println("Incorrect!")
    }

    
    fmt.Printf("What are the two largest religions in the world? (separate by comma)\nEnter your answer: ")
    answer2, _ := reader.ReadString('\n')
    answer2 = strings.TrimSpace(answer2)
    parts := strings.Split(answer2, ",")
    for i := range parts {
        parts[i] = strings.TrimSpace(parts[i])
    }
    if (strings.EqualFold(parts[0], "Islam") && strings.EqualFold(parts[1], "Christianity")) ||
       (strings.EqualFold(parts[0], "Christianity") && strings.EqualFold(parts[1], "Islam")) {
        fmt.Println("Correct!")
        score += 1
    } else {
        fmt.Println("Incorrect!")
    }

    
    fmt.Printf("Who is the most famous person in the world?\nEnter your answer: ")
    answer3, _ := reader.ReadString('\n')
    answer3 = strings.TrimSpace(answer3)
    if strings.EqualFold(answer3, "Christiano Ronaldo") {
        fmt.Println("Correct!")
        score += 1
    } else {
        fmt.Println("Incorrect!")
    }

    // Score
    fmt.Printf("You scored %v out of %v\n", score, numQuestions)
    percent := (float64(score) / float64(numQuestions)) * 100
    fmt.Printf("You scored a total of %.2f%%\n", percent)

    if percent == 100 {
        fmt.Println("Round of applause! You scored a perfect 100!")
    }
}
