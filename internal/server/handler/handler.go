package handler

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"

	"sampisal/internal/common/pow"
	"sampisal/internal/repository"
)

var quotes = []string{
	"The only true wisdom is in knowing you know nothing.",
	"The unexamined life is not worth living.",
	"There is only one good, knowledge, and one evil, ignorance.",
	"I cannot teach anybody anything. I can only make them think.",
	"Be kind, for everyone you meet is fighting a hard battle.",
}

func HandleConnection(client repository.TCPClient) {
	defer client.Close()

	challenge := pow.GenerateChallenge()
	err := client.Write(challenge)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Sent challenge:", challenge)

	response, err := client.Read()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Received response:", response)

	if !pow.VerifyChallengeResponse(challenge, strings.TrimSpace(response)) {
		fmt.Println("Invalid response")
		return
	}

	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(quotes))))
	if err != nil {
		fmt.Println(err)
		return
	}
	randomQuote := quotes[n.Int64()]

	err = client.Write(randomQuote)
	if err != nil {
		fmt.Println(err)
		return
	}
}
