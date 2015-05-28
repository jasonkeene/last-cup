package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

// Grab the last cup date out of env var.
func getLastCupDate() (time.Time, error) {
	zero_date := time.Time{}
	last_cup := os.Getenv("LAST_CUP")

	if len(last_cup) == 0 {
		return zero_date, errors.New("LAST_CUP env var not set.")
	}

	// parse into time.Time
	format := "2006-1-2"
	date, err := time.Parse(format, last_cup)
	if err != nil {
		return zero_date, errors.New("Error in parsing date set in LAST_CUP," +
			" please make sure it follows the format " + format + ".")
	}
	return date, nil
}

// Return how many days have passed since last cup.
func daysSince() (int, error) {
	date, err := getLastCupDate()
	if err != nil {
		return 0, err
	}
	return int(time.Now().Sub(date) / time.Hour / 24), nil
}

// Map days to messages to dispaly to user.
func message(days int) string {
	switch {
	case days < 0:
		return "You are planning on kicking your caffeine habit, congrats!"
	case days < 5:
		return fmt.Sprintf("Great job! It's been %d days since your last "+
			"cup!\nYou've just started a long journey of caffeine "+
			"abstinence.\nLook forward to not having those fun withdrawl "+
			"headaches I'm sure you're experiencing.", days)
	default:
		months := days / 30
		days = days % 30
		if months != 0 {
			return fmt.Sprintf("Great job! It's been %d months %d days since"+
				" your last cup! You're a Zen master!", months, days)
		} else {
			return fmt.Sprintf("Great job! It's been %d days since your last"+
				" cup!", days)
		}
	}
}

func main() {
	days, err := daysSince()
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
	} else {
		fmt.Println(message(days))
	}
}
