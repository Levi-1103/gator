package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("command doesn't have any arguments")
	}

	err := s.db.ResetDB(context.Background())
	if err != nil {
		return fmt.Errorf("db could not be reset")
	}

	fmt.Printf("db has been reset")

	return nil
}
