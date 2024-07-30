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
func (v Voter) PreviousHash(cand1, cand2, cand3) {
	if v.choice == cand1.name {
		previoushash := cand1.people[len(cand1.people-1)]
		return previoushash
	} else if v.choice == cand2.name {
		previoushash := cand2.people[len(cand2.people-1)]
		return previoushash
	} else if v.choice == cand3.name {
		previoushash := cand3.people[len(cand3.people-1)]
		return previoushash
	}
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

func main() {
	name, choice := Questions()
	President1 := Candidate{name: "tinubu", people: make([][32]byte, 0)}
	President2 := Candidate{name: "atiku", people: make([][32]byte, 0)}
	President3 := Candidate{name: "obi", people: make([][32]byte, 0)}

	if choice != President1.name || choice != President2.name || choice != President3.name {
		log.Fatal("GET OUT OF HERE, YOU DON'T EVEN KNOW THE NAME OF THE CANDIDATE AHHHH!!!!")
		return
	}

}
