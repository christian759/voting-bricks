//
//

package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"log"
	"os"
)

type Candidate struct {
	name    string
	genesis [32]byte
	people  [][32]byte
}

type Voter struct {
	previous_hashid string
	hashid          [32]byte
	name            string
	choice          string
}

func (v Voter) HashGenerator() [32]byte {
	votingid := v.name + v.choice + v.previous_hashid
	votinghash_id := sha256.Sum256([]byte(votingid))
	return votinghash_id
}

// PreviousHash function
func PreviousHash(name string, cand1, cand2, cand3 Candidate) ([32]byte, error) {
	var previousHash [32]byte // Default zero-value of [32]byte

	if name == cand1.name {
		if len(cand1.people) > 0 {
			previousHash = cand1.people[len(cand1.people)-1]
			return previousHash, nil
		}
	} else if name == cand2.name {
		if len(cand2.people) > 0 {
			previousHash = cand2.people[len(cand2.people)-1]
			return previousHash, nil
		}
	} else if name == cand3.name {
		if len(cand3.people) > 0 {
			previousHash = cand3.people[len(cand3.people)-1]
			return previousHash, nil
		}
	}
	return previousHash, nil
}

func (v *Voter) Voting(c Candidate) {
	v.hashid = v.HashGenerator()
}

func Questions() (string, string) {
	fmt.Printf("What is your name ?")
	read := bufio.NewReader(os.Stdin)
	name, err := read.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("What is choice of president ?")
	read1 := bufio.NewReader(os.Stdin)
	choice, err := read1.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return name, choice
}

func (c *Candidate) genesisGenerator() [32]byte {
	votingid := c.name + c.name
	hash_id := sha256.Sum256([]byte(votingid))
	c.genesis = hash_id
	return hash_id
}

func main() {
	name, choice := Questions()

	President1 := Candidate{name: "tinubu", people: make([][32]byte, 0)}
	President1.genesisGenerator()
	President2 := Candidate{name: "atiku", people: make([][32]byte, 0)}
	President2.genesisGenerator()
	President3 := Candidate{name: "obi", people: make([][32]byte, 0)}
	President3.genesisGenerator()

	if choice != President1.name || choice != President2.name || choice != President3.name {
		log.Fatal("GET OUT OF HERE, YOU DON'T EVEN KNOW THE NAME OF THE CANDIDATE AHHHH!!!!")
		return
	}
	PreviousHash(name, President1, President2, President3)

}
