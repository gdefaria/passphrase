package main

import (
	"crypto/rand"
	"embed"
	"flag"
	"fmt"
	"log"
	"math"
	"math/big"
	"strings"
)

//go:embed wordlists/*.txt
var wordlists embed.FS

func loadWordlist(listName string) ([]string, error) {
	data, err := wordlists.ReadFile(fmt.Sprintf("wordlists/%s.txt", listName))
	
	if err != nil {
		return nil, err
	}

	return strings.Split(string(data), "\n"), nil
}

func generate(words []string, length int) ([]string, error) {
	var passphrase []string

	for i := 0; i < length; i++ {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(words))))

		if err != nil {
			return passphrase, fmt.Errorf("Failed to generate random number: %w", err)
		}

		passphrase = append(passphrase, words[index.Int64()])
	}

	return passphrase, nil
}

func main() {
	listName := flag.String("list", "eff-long", "Available wordlists: [eff-long] [bip39] [bip39-pt]")
	length := flag.Int("len", 10, "Length (amount of words) of the passphrase")
	flag.Parse()

	if *length <= 0 {
		log.Fatalf("Length must be greater than 0")
	}

	words, err := loadWordlist(*listName)

	if err != nil {
		log.Fatalf("Error loading wordlist: %v", err)
	}

	pass, _ := generate(words, *length)
	entropy := math.Log2(float64(len(words))) * float64(*length)

	fmt.Printf("%s (%.2f bits)", strings.Join(pass, " "), entropy)

	// Removing slice data from memory
	defer func() {
		for i := range pass {
			pass[i] = ""
		}
	}()
}