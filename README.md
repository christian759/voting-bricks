**Blockchain-Based Voting System**

- This repository contains a Go-based implementation of a simple blockchain voting system. The system allows voters to vote for one of three presidential candidates using a SHA-256 hash to ensure the integrity of the voting process.

**Table of Contents**
- Overview
- Features
- Requirements
- Installation
- Usage

**Overview**
- This application simulates a voting process where voters can cast their votes for three presidential candidates: Tinubu, Atiku, and Obi. Each vote is secured using SHA-256 hashing, ensuring that the voting process is tamper-proof and verifiable.

**Features**
- Candidate and Voter Representation: Structs to represent candidates and voters.
- Hash Generation: SHA-256 hash generation to secure votes.
- Previous Hash Retrieval: Function to retrieve the last hash of a candidate's people list.
- Genesis Hash Generation: Function to generate the genesis hash for each candidate.
- User Input: Prompts for voter name and candidate choice.
- Validation: Checks to ensure voters vote only once and validate hash chains.
- Winner Determination: Calculates and displays the candidate with the most votes.

**Requirements**
- Go 1.16 or higher

**Installation--Clone the repository:**
- git clone https://github.com/christian759/voting-bricks.git
- cd voting-bricks

**Usage**

*Build the project:*
- go build
- go install

*Run the application:*
- ./voting-bricks
- Follow the prompts to enter your name and vote for a candidate.
- Repeat the voting process until all votes are cast.
- The program will display the candidate with the most supporters.
