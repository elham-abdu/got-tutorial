# Go Quiz Game 🎮

A simple command-line quiz game written in **Go**.  
The game asks the user for their name and age, then gives them a short quiz and calculates their score and percentage.

## Features

- Takes user **name** and **age**
- Validates age (must be 10 or above)
- Asks **3 quiz questions**
- Accepts multi-word and case-insensitive answers
- Calculates:
  - Total score
  - Percentage
- Prints a special message for perfect score

## Quiz Topics

1. First person who invented the computer  
2. Two largest religions in the world  
3. Most famous person in the world

## Project Structure

tutorial.go  
README.md

## How to Run the Project

Make sure you have Go installed (Go 1.18+ recommended).

### 1. Clone the repository
git clone https://github.com/elham-abdu/got-tutorial.git  
cd got-tutorial

### 2. Run the program
go run tutorial.go

## Example Output

Welcome to the quiz game!  
Enter your name here: Elham  
Hello Elham, Welcome to have fun!  
Enter your age: 20  
Yay Elham, you are ready to play!  
Who is the first person who invented computer?  
Enter your answer: Charles Babbage  
Correct!  
You scored 3 out of 3  
You scored a total of 100.00%  
Round of applause! You scored a perfect 100!

## Future Improvements

- Add more quiz questions
- Add categories or difficulty levels
- Add scoring history
- Randomize question order

## License

This project is free to use for learning and practice.
