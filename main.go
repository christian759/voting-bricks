package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"log"
	"os"
)

// Candidate represents a candidate in the voting system.
type Candidate struct {
	name       string     // Name of the candidate
	genesis    [32]byte   // Genesis hash associated with the candidate
	people     [][32]byte // List of hashes related to this candidate
	supporters []Voter
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
func PreviousHash(name string, cand1, cand2, cand3 Candidate) ([32]byte, Candidate, error) {
	var previousHash [32]byte // Default zero-value of [32]byte for no hash found

	// Check each candidate's name and return the last hash if available
	if name == cand1.name {
		if len(cand1.people) > 0 {
			previousHash = cand1.people[len(cand1.people)-1]
			return previousHash, cand1, nil
		}
	} else if name == cand2.name {
		if len(cand2.people) > 0 {
			previousHash = cand2.people[len(cand2.people)-1]
			return previousHash, cand2, nil
		}
	} else if name == cand3.name {
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
	fmt.Printf("What is your name? ")
	read := bufio.NewReader(os.Stdin)
	name, err := read.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("What is your choice of president? ")
	read1 := bufio.NewReader(os.Stdin)
	choice, err := read1.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
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
	return hash_id
}

func main() {
	// Get user input for name and choice
	name, choice := Questions()

	// Initialize candidates with empty people slices
	President1 := Candidate{name: "tinubu", people: make([][32]byte, 0), supporters: make([]Voter, 0)}
	President1.genesisGenerator()
	President2 := Candidate{name: "atiku", people: make([][32]byte, 0), supporters: make([]Voter, 0)}
	President2.genesisGenerator()
	President3 := Candidate{name: "obi", people: make([][32]byte, 0), supporters: make([]Voter, 0)}
	President3.genesisGenerator()

	// Check if the user's choice matches any of the candidates
	if choice != President1.name && choice != President2.name && choice != President3.name {
		log.Fatal("GET OUT OF HERE, YOU DON'T EVEN KNOW THE NAME OF THE CANDIDATE AHHHH!!!!")
		return
	}

	// Retrieve the previous hash based on the user's name and chosen candidate
	prevhash, cand, err := PreviousHash(name, President1, President2, President3)
	if err != nil {
		// Handle the case where the candidate's previous hash is not found
		log.Fatal(err)
		return
	} // Generate the hash based on user's name, choice, and previous hash
	newHash := HashGenerator(name, cand.name, prevhash)
	voter := Voter{name: name, choice: choice, hashid: newHash, previous_hashid: prevhash}
	if voter.choice == President1.name {
		President1.supporters = append(President1.supporters, voter)
	} else if voter.choice == President2.name {
		President2.supporters = append(President2.supporters, voter)
	} else if voter.choice == President3.name {
		President3.supporters = append(President3.supporters, voter)
	}
}
