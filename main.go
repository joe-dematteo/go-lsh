package main

import (
	"context"
	"fmt"
	"time"

	"go-lsh/utils" // Import the utils package

	"github.com/agtabesh/lsh"
)

func main() {
	start := time.Now()

	// LSH configuration
	config := lsh.LSHConfig{
		SignatureSize: 128, // It is better to be a power of 2
	}

	hashFamily := lsh.NewXXHASH64HashFamily(config.SignatureSize)
	similarityMeasure := lsh.NewHammingSimilarity()
	store := lsh.NewInMemoryStore()

	// Create LSH instance
	instance, err := lsh.NewLSH(config, hashFamily, similarityMeasure, store)
	if err != nil {
		fmt.Println("Error initing LSH instance:", err)
		return
	}

	// Define number of contacts and search count
	contactsToCreate := 50000
	searchCount := 50000

	// Load and create contacts (using utils functions)
	firstNameData, err := utils.ReadFirstNameData("utils/data/names.json")
	if err != nil {
		fmt.Println("Error reading first names:", err)
		return
	}
	contacts := utils.CreateRandomContacts(contactsToCreate, firstNameData)

	// Add contacts to LSH
	utils.AddContactsToLSH(instance, contacts)

	// Search logic
	searchTerm := "Chantalle"                                                       // Example: Search for contacts with last name "Chantalle"
	searchTermVector := utils.TransformContact(utils.Contact{LastName: searchTerm}) // Create a vector for the search term

	fmt.Println("searchTermVector:", searchTermVector)

	similarContacts, err := instance.QueryByVector(context.Background(), searchTermVector, searchCount)
	if err != nil {
		fmt.Println("Error querying vector:", err)
		return
	}

	elapsed := time.Since(start)

	fmt.Printf("\nSimilar Contacts (%d) for '%s':\n", len(similarContacts), searchTerm)
	for _, contactID := range similarContacts {
		// Assuming you have a way to retrieve contact details by ID
		fmt.Println("- ", contactID) // Replace with actual contact information retrieval based on ID
	}

	fmt.Printf("\nExecution time: %s\n", elapsed)

}
