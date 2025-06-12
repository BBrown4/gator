package main

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't retrieve users: %w", err)
	}

	if len(users) == 0 {
		fmt.Println("No users found.")
		return nil
	}

	for _, user := range users {
		str := fmt.Sprintf(" * %s", user.Name)
		if user.Name == s.cfg.CurrentUserName {
			str += " (current)"
		}

		fmt.Println(str)
	}
	return nil
}
