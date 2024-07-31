package main

import (
	"crypto/sha256"
	"fmt"
	"log"
	"strings"
)

// Candidate represents a candidate in the voting system.
type Candidate struct {
	name       string     // Name of the candidate
	genesis    [32]byte   // Genesis hash associated with the candidate
	people     [][32]byte // List of hashes related to this candidate
	supporters []string   // list of the candidate supporters
}

// Voter represents a voter in the system.
type Voter struct {
	previous_hashid [32]byte // Previous hash ID as a string
	hashid          [32]byte // Current hash ID of the voter
	name            string   // Name of the voter
	choice          string   // Voter's choice of candidate
}

// HashGenerator creates a SHA-256 hash based on the name, choice, and previous hash ID.
func HashGenerator(name, choice string, previous_hashid [32]byte) [32]byte {
	// Convert previous_hashid to a string to concatenate with name and choice
	previousHashStr := string(previous_hashid[:])

	// Concatenate name, choice, and previous hash string to form the voting ID
	votingid := name + choice + previousHashStr

	// Generate SHA-256 hash of the concatenated string
	votinghash_id := sha256.Sum256([]byte(votingid))

	// Return the generated hash
	return votinghash_id
}

// PreviousHash retrieves the last hash of a candidate's people list based on the candidate's name.
// Returns the previous hash, the candidate, and an error if no match is found.
func PreviousHash(choice string, cand1, cand2, cand3 Candidate) ([32]byte, Candidate, error) {
	var previousHash [32]byte // Default zero-value of [32]byte for no hash found

	// Check each candidate's name and return the last hash if available
	if choice == cand1.name {
		if len(cand1.people) > 0 {
			previousHash = cand1.people[len(cand1.people)-1]
			return previousHash, cand1, nil
		}
	} else if choice == cand2.name {
		if len(cand2.people) > 0 {
			previousHash = cand2.people[len(cand2.people)-1]
			return previousHash, cand2, nil
		}
	} else if choice == cand3.name {
		if len(cand3.people) > 0 {
			previousHash = cand3.people[len(cand3.people)-1]
			return previousHash, cand3, nil
		}
	}

	// If no candidate matches, return the default previous hash and error
	return previousHash, cand1, fmt.Errorf("candidate not found or has no previous hashes")
}

// Questions prompts the user to enter their name and choice of president.
// Returns the user's name and choice as strings.
func Questions() (string, string) {
	var name, choice string

	fmt.Print("Enter your name: ")
	_, err := fmt.Scan(&name)
	if err != nil {
		fmt.Println("Error reading name:", err)
	}
	fmt.Printf("What is your choice of president? ")
	_, err = fmt.Scan(&choice)
	if err != nil {
		fmt.Println("Error reading your choice:", err)
	}
	name = strings.TrimSpace(name)
	choice = strings.TrimSpace(choice)
	return name, choice
}

// genesisGenerator generates and sets the genesis hash for the candidate.
// Returns the generated genesis hash.
func (c *Candidate) genesisGenerator() [32]byte {
	// Concatenate candidate's name with itself to create a unique string for hashing
	votingid := c.name + c.name

	// Generate SHA-256 hash of the concatenated string
	hash_id := sha256.Sum256([]byte(votingid))

	// Set the candidate's genesis hash and return it
	c.genesis = hash_id
	c.people = append(c.people, hash_id)
	return hash_id
}

func Contains(slice []string, person string) bool {
	for _, item := range slice {
		if item == person {
			return false
		}
	}
	return true
}

func hashchecker(list [][32]byte, hash [32]byte) bool {
	if hash == list[len(list)-1] {
		return hash == list[len(list)-1]
	}
	return false
}

func main() {

	// Initialize candidates with empty people slices
	President1 := Candidate{name: "tinubu", people: make([][32]byte, 0), supporters: make([]string, 0)}
	President1.genesisGenerator()
	President2 := Candidate{name: "atiku", people: make([][32]byte, 0), supporters: make([]string, 0)}
	President2.genesisGenerator()
	President3 := Candidate{name: "obi", people: make([][32]byte, 0), supporters: make([]string, 0)}
	President3.genesisGenerator()

	// using a for loop to take votes and determine who wins the election
	i := 20
	for i > 0 {
		// Get user input for name and choice
		name, choice := Questions()

		// Check if the user's choice matches any of the candidates
		if (choice != President1.name) && (choice != President2.name) && (choice != President3.name) {
			log.Fatal("GET OUT OF HERE, YOU DON'T EVEN KNOW THE NAME OF THE CANDIDATES AHHHH!!!!")
			return
		}

		// Retrieve the previous hash based on the user's name and chosen candidate
		prevhash, cand, err := PreviousHash(choice, President1, President2, President3)
		if err != nil {
			// Handle the case where the candidate's previous hash is not found
			log.Fatal(err)
			return
		} // Generate the hash based on user's name, choice, and previous hash
		newHash := HashGenerator(name, cand.name, prevhash)
		voter := Voter{name: name, choice: choice, hashid: newHash, previous_hashid: prevhash}

		//validation and adding voters to the supporter's list
		if voter.choice == President1.name {
			if Contains(President1.supporters, voter.name) && hashchecker(President1.people, voter.previous_hashid) {
				President1.supporters = append(President1.supporters, voter.name)
				President1.people = append(President1.people, voter.hashid)
			} else {
				fmt.Printf("This blockchain system has been compromised")
				return
			}
		} else if voter.choice == President2.name {
			if Contains(President2.supporters, voter.name) && hashchecker(President2.people, voter.previous_hashid) {
				President2.supporters = append(President2.supporters, voter.name)
				President2.people = append(President2.people, voter.hashid)
			} else {
				fmt.Printf("This blockchain system has been compromised")
				return
			}
		} else if voter.choice == President3.name {
			if Contains(President3.supporters, voter.name) && hashchecker(President3.people, voter.previous_hashid) {
				President3.supporters = append(President3.supporters, voter.name)
				President3.people = append(President3.people, voter.hashid)
			} else {
				fmt.Printf("This blockchain system has been compromised")
				return
			}
		}

		i--
	}

	// finding the winner and th election i.e the candidate with the most supporters
	maxSupporters := float64(len(President1.people))
	winner := President1.name

	if float64(len(President2.people)) > maxSupporters {
		maxSupporters = float64(len(President2.people))
		winner = President2.name
	}
	if float64(len(President3.people)) > maxSupporters {
		maxSupporters = float64(len(President3.people))
		winner = President3.name
	}

	fmt.Printf("The candidate with the most supporters is %s with %d supporters.\n", winner, int(maxSupporters))

}
