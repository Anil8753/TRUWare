package main

import "errors"

type RegistrationEntry struct {
	Type      string `json:"type"`
	ID        string `json:"id"`
	OwnerId   string `json:"ownerId"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	Contact   string `json:"contact"`
	Email     string `json:"email"`
	GSTNumber string `json:"gst"`
}

func validateRegistrationData(re *RegistrationEntry) error {

	if re.Name == "" {
		return errors.New("name cannot be empty")
	}

	if re.Address == "" {
		return errors.New("address cannot be empty")
	}

	if re.Contact == "" {
		return errors.New("contact details cannot be empty")
	}

	if re.GSTNumber == "" {
		return errors.New("GST number cannot be empty")
	}

	return nil
}
