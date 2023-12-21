package hangmanweb

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
	"time"

)

var listAscii [][]string
var standardFile string
var shadowFile string
var thinkertoyFile string



func init() {
	rand.Seed(time.Now().UnixNano())
}

func play_hangman(numletters int) (playagain string, is_winner bool) {
	stage_of_death := 0
	var updateddashes string
	gamemode := 0
	has_guessed_1_letter := false
	has_won := false
	guess := ""
	guessed_letters := ""
	again := ""
	dashes := ""
	newdashes := ""
	guessedLetters := []string{}
	asciiArt := `
	
┓┏  ┏┓  ┳┓  ┏┓  ┳┳┓  ┏┓  ┳┓  
┣┫  ┣┫  ┃┃  ┃┓  ┃┃┃  ┣┫  ┃┃  
┛┗  ┛┗  ┛┗  ┗┛  ┛ ┗  ┛┗  ┛┗  
	`
	asciiwon := `
	
┏┓  ┏┓  ┳┓  ┏┓  ┳┓  ┏┓  ┏┳┓  ┳┳  ┓   ┏┓  ┏┳┓  ┳  ┏┓  ┳┓
┃   ┃┃  ┃┃  ┃┓  ┣┫  ┣┫   ┃   ┃┃  ┃   ┣┫   ┃   ┃  ┃┃  ┃┃
┗┛  ┗┛  ┛┗  ┗┛  ┛┗  ┛┗   ┻   ┗┛  ┗┛  ┛┗   ┻   ┻  ┗┛  ┛┗
	`
	fmt.Println(asciiArt)
	word := random_word(numletters, gamemode)
	for {
		fmt.Println("Select game mode:")
		fmt.Println("1. Only use Common words (easy mode)")
		fmt.Println("2. Use all words (hard mode)")
		fmt.Scanln(&gamemode)
		if (gamemode == 1) || (gamemode == 2) {
			clearscreen()
			break
		} else {
			fmt.Println("Please type 1 or 2")
		}
	} 
    word = random_word(numletters, gamemode)

	for {
		if has_won == true {
            clearscreen()
            fmt.Println(asciiwon)
            fmt.Printf("You won the game! The word was %s\n", word)
            return "y", true
		}

		draw_hangman(stage_of_death)
		if stage_of_death >= 10 {
			fmt.Printf("Oh dear hangman is dead\n")
			fmt.Printf("The word that could have saved him was %s\n", word)
			for {
				fmt.Printf("Play again? (y/n) \n")
				fmt.Scanln(&again)
				isYorN, somekindoferror := regexp.MatchString("^y|Y|n|N", again)
				if somekindoferror != nil {
					fmt.Printf("Something has gone horribly wrong. ")
					fmt.Printf("exiting with error can not regex match %v", again)
					return
				}
				if isYorN == false {
					fmt.Printf("You didn't type 'y' or 'n'! Try again\n")
				} else if len(again) > 1 {
					fmt.Printf("You entered more than 1 character! Try again\n")
				} else if strings.ToLower(again) == "y" {
					return "y", false
				} else {
					return "n", false
				}
			}
		}
		
		alreadyGuessed := false
for _, letter := range guessedLetters {
    if letter == guess {
        alreadyGuessed = true
        break
    }
	
}
		if !alreadyGuessed {
			guessedLetters = append(guessedLetters, guess)
		}
		guessedString := strings.Join(guessedLetters, " ")
		fmt.Printf("Guessed Letters: %s\n", guessedString)
		if has_guessed_1_letter == false {
			dashes = hideword(word)
			fmt.Printf("%s\n", dashes)
		} else {
			fmt.Printf("%s\n", newdashes)
		}
    fmt.Printf("Guess a word or letter: ")
    fmt.Scan(&guess)
		isALetter, somekindoferror := regexp.MatchString("^[a-zA-Z]", guess)
		if somekindoferror != nil {
			clearscreen()
			fmt.Printf("Something has gone horribly wrong. ")
			fmt.Printf("exiting with error can not regex match %v", guess)
			return
		}
		if isALetter == false {
			clearscreen()
			fmt.Printf("That's not a letter! Try again\n")
		} else if strings.Contains(guessed_letters, guess) {
			clearscreen()
			fmt.Println("You have already guessed that letter! Try again")
		} else if len(guess) > 1 {
			clearscreen()
			if guess == word{
				has_won = true	
			} else if guess != word {
				fmt.Println("Wrong word")
				stage_of_death +=2
			}
		} else if strings.Contains(word, guess) {
			clearscreen()
			fmt.Println("The letter you guessed is in the word")
			guessed_letters += guess
			if has_guessed_1_letter == false {
				updateddashes, has_won = revealdashes(word, guess, dashes)
				newdashes = updateddashes
			} else {
				updateddashes, has_won = revealdashes(word, guess, newdashes)
				newdashes = updateddashes
			}
			has_guessed_1_letter = true
			if has_won {
				clearscreen()
				fmt.Println(asciiwon)
				fmt.Printf("You won the game! The word was %s\n", word)
				for {
					fmt.Printf("Play again? (y/n) \n")
					fmt.Scanln(&again)
					isYorN, somekindoferror := regexp.MatchString("^y|Y|n|N", again)
					if somekindoferror != nil {
						fmt.Printf("Something has gone horribly wrong. ")
						fmt.Printf("exiting with error can not regex match %v", again)
						return
					}
					if isYorN == false {
						fmt.Printf("You didn't type 'y' or 'n'! Try again\n")
					} else if len(again) > 1 {
						fmt.Printf("You entered more than 1 character! Try again\n")
					} else if strings.ToLower(again) == "y" {
						return "y", true
					} else {
						return "n", true
					}
				}
			}
		} else {
			clearscreen()
			fmt.Printf("The letter you guessed is not in the word\n")
			stage_of_death++
			guessed_letters += guess
		}
		if has_won {
			clearscreen()
			fmt.Println(asciiwon)
			fmt.Printf("You won the game! The word was %s\n", word)
			for {
				fmt.Printf("Play again? (y/n) \n")
				fmt.Scanln(&again)
				isYorN, somekindoferror := regexp.MatchString("^y|Y|n|N", again)
				if somekindoferror != nil {
					fmt.Printf("Something has gone horribly wrong. ")
					fmt.Printf("exiting with error can not regex match %v", again)
					return
				}
				if isYorN == false {
					fmt.Printf("You didn't type 'y' or 'n'! Try again\n")
				} else if len(again) > 1 {
					fmt.Printf("You entered more than 1 character! Try again\n")
				} else if strings.ToLower(again) == "y" {
					return "y", true
				} else {
					return "n", true
				}
			}
		}
	}
}
func draw_hangman(stage_of_death int) {
	switch stage_of_death {
	case 0:
		fmt.Printf("  +---+\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("========\n")
	case 1:
		fmt.Printf("  +---+\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("========\n")
	case 2:
		fmt.Printf("  +---+\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("  O   |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("========\n")
	case 3:
		fmt.Printf("  +---+\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("  O   |\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("========\n")
	case 4:
		fmt.Printf("  +---+\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("  O   |\n")
		fmt.Printf(" /|   |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("========\n")
	case 5:
		fmt.Printf("  +---+\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("  O   |\n")
		fmt.Printf("_/|   |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("========\n")
	case 6:
		fmt.Printf("  +---+\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("  O   |\n")
		fmt.Printf("_/|\\  |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("========\n")
	case 7:
		fmt.Printf("  +---+\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("  O   |\n")
		fmt.Printf("_/|\\_ |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("========\n")
	case 8:
		fmt.Printf("  +---+\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("  O   |\n")
		fmt.Printf("_/|\\_ |\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("========\n")
	case 9:
		fmt.Printf("  +---+\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("  O   |\n")
		fmt.Printf("_/|\\_ |\n")
		fmt.Printf("  |   |\n")
		fmt.Printf(" /    |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("========\n")
	case 10:
		fmt.Printf("  +---+\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("  O   |\n")
		fmt.Printf("_/|\\_ |\n")
		fmt.Printf("  |   |\n")
		fmt.Printf(" / \\  |\n")
		fmt.Printf("      |\n")
		fmt.Printf("R.I.P |\n")
		fmt.Printf("========\n")
	default :
		fmt.Printf("  +---+\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("  O   |\n")
		fmt.Printf("_/|\\_ |\n")
		fmt.Printf("  |   |\n")
		fmt.Printf(" / \\  |\n")
		fmt.Printf("      |\n")
		fmt.Printf("R.I.P |\n")
		fmt.Printf("========\n")
	}
}
func hideword(word string) string {
	dashes := ""
	for i, letter := range word {
		if i < len(word)/2-1 {
			dashes += string(letter)
		} else {
			dashes += "_"
		}
	}
	return dashes
}

func revealdashes(word string, guess string, dashes string) (string, bool) {
    newdashes := ""
	for i, r := range dashes {
        if c := string(r); c != "_" {
            newdashes += c
        } else {
            var letter = string(word[i])
            if guess == letter {
                newdashes += guess
            } else {
                newdashes += "_"
            }
        }
    }
    return newdashes, newdashes == word
}

func random_word(numletters int, gamemode int) string {
	switch gamemode { 
	case 1:
		var dataletters []byte
		var err error
		if numletters == 4 {
			dataletters, err = ioutil.ReadFile("words/common4l.txt")
		} else if numletters == 5 {
			dataletters, err = ioutil.ReadFile("words/common5l.txt")
		} else if numletters >= 6 {
			dataletters, err = ioutil.ReadFile("words/common6l.txt")
		}
		if err != nil {
			panic(err)
		}
		datastr := string(dataletters)
		somewords := strings.Split(datastr, " ")
		randnum := rand.Intn(len(somewords) - 1)
		chosenword := somewords[randnum]
		return chosenword
	case 2:
		var dataletters []byte
		var err error
		if numletters >= 4 {
			dataletters, err = ioutil.ReadFile("words/all4l.txt")
		}
		if err != nil {
			panic(err)
		}
		datastr := string(dataletters)
		somewords := strings.Split(datastr, " ")
		randnum := rand.Intn(len(somewords) - 1)
		chosenword := somewords[randnum]
		return chosenword
	}
	return "omgthisisabugyoushouldntseethisever"
}
func clearscreen() {
	if runtime.GOOS != "windows" {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
