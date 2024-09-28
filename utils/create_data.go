package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/agtabesh/lsh"
	"github.com/agtabesh/lsh/types"
)

// Define a Contact struct
type Contact struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// CreateRandomContacts generates a specified number of random contacts
func CreateRandomContacts(count int, firstNameData []string) []Contact {
	contacts := make([]Contact, count)
	rand.NewSource(time.Now().UnixNano()) // Seed random number generator

	for i := 0; i < count; i++ {
		firstNameIdx := rand.Intn(len(firstNameData))
		lastNameIdx := rand.Intn(len(firstNameData))
		contacts[i] = Contact{
			FirstName: firstNameData[firstNameIdx],
			LastName:  firstNameData[lastNameIdx],
		}
	}
	return contacts
}

// ReadFirstNameData reads first names from a JSON file
func ReadFirstNameData(filePath string) ([]string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var names []string
	err = json.Unmarshal(data, &names)
	return names, err
}

// AddContactsToLSH adds a slice of contacts to the LSH instance
func AddContactsToLSH(instance *lsh.LSH, contacts []Contact) {
	ctx := context.Background()
	for _, contact := range contacts {
		vector := TransformContact(contact)
		fmt.Println("Adding vector:", vector)
		vectorID := types.VectorID(fmt.Sprintf("%s %s", contact.FirstName, contact.LastName))
		err := instance.Add(ctx, vectorID, vector)
		if err != nil {
			fmt.Println("Error adding contact to LSH:", err)
		}
	}
}

// TransformContact transforms a Contact struct into a types.Vector
func TransformContact(contact Contact) types.Vector {
	fullName := fmt.Sprintf("%s %s", contact.FirstName, contact.LastName)
	vectorID := types.VectorID(fullName) // Use the VectorID constructor
	return types.Vector{vectorID: 1.0}
}
