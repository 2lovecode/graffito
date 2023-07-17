package err

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
)

type MyError struct {
	message    string
	innerError error
}

func (e *MyError) Error() string {
	if e.innerError != nil {
		return fmt.Sprintf("%s: %v", e.message, e.innerError)
	}
	return e.message
}

func (e *MyError) Unwrap() error {
	return e.innerError
}

func DoSomething() error {
	err := DoSomethingElse()
	if err != nil {
		return &MyError{
			message:    "Failed to do something",
			innerError: err,
		}
	}
	return nil
}

func DoSomethingElse() error {
	return errors.New("Something went wrong")
}

func NewCommand() *cobra.Command {
	return &cobra.Command{Use: "err", Run: func(cmd *cobra.Command, args []string) {
		err := DoSomething()
		err2 := &MyError{
			message:    "good",
			innerError: err,
		}
		if err2 != nil {
			fmt.Println(err2)
			fmt.Println(errors.Unwrap(err2))
			fmt.Println(errors.Unwrap(errors.Unwrap(err2)))
		}
	}}

}
