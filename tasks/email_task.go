package tasks

import (
	"context"
	"fmt"
	"time"
)

// EmailTask represents a task that sends an email.
// It implements the Task interface.
type EmailTask struct {
	Recipient string
	Subject   string
	Body      string
}

// NewEmailTask creates a new EmailTask with the specified recipient, subject, and body.
func NewEmailTask(recipient, subject, body string) *EmailTask {
	return &EmailTask{
		Recipient: recipient,
		Subject:   subject,
		Body:      body,
	}
}

// Execute sends the email.
func (e *EmailTask) Execute(ctx context.Context) error {
	time.Sleep(2 * time.Second) // Simulate a delay for sending the email.
	// Simulate sending an email.
	fmt.Printf("Sending email to %s with subject '%s' and body '%s'\n", e.Recipient, e.Subject, e.Body)
	return nil
}
