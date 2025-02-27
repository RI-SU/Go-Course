package main

import "fmt"

func (a *analytics) handleEmailBounce(em email) error {
	user := em.recipient.updateStatus(em.status)
	if user != nil {
		return fmt.Errorf("error updating user status: %w", user)
	}
	track := a.track(em.status)
	if track != nil {
		return fmt.Errorf("error tracking user bounce: %w", track)
	}
	return nil
}
