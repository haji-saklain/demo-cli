package commands

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

func Hello(db *sql.DB, message string) error {
	output := fmt.Sprintf("Hello, %s!", message)
	if err := insertOutput(db, "hello", output); err != nil {
		return err
	}
	return nil
}

func Goodbye(db *sql.DB, message string) error {
	output := fmt.Sprintf("Goodbye, %s!", message)
	if err := insertOutput(db, "goodbye", output); err != nil {
		return err
	}
	return nil
}

func CountNumbers(db *sql.DB, n int) error {
	fmt.Println("Inside CountNumbers function with n =", n)
	if n < 0 {
		return errors.New("count cannot be negative")
	}
	output := ""
	for i := 1; i <= n; i++ {
		output += fmt.Sprintf("%d\n", i)
	}
	if err := insertOutput(db, "count", output); err != nil {
		return err
	}
	return nil
}

func PrintCurrentTime(db *sql.DB) error {
	output := fmt.Sprintf("Current time: %s", time.Now().Format(time.RFC3339))
	if err := insertOutput(db, "time", output); err != nil {
		return err
	}
	return nil
}

func insertOutput(db *sql.DB, command, output string) error {
	stmt, err := db.Prepare("INSERT INTO output_table (command, output) VALUES ($1, $2)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(command, output)
	if err != nil {
		return err
	}
	return nil
}
