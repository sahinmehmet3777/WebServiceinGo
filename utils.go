// utils.go
package main

import "errors"

// validateTaskFields validates the fields of a task
func validateTaskFields(title, description, status string) error {
    if title == "" {
        return errors.New("title cannot be empty")
    }

    if description == "" {
        return errors.New("description cannot be empty")
    }

    if status == "" {
        return errors.New("status cannot be empty")
    }

    return nil
}
