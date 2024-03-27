package pow

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const (
	ChallengeDifficulty = 2
	ChallengeSymbol     = "0"
)

func GenerateChallenge() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return strconv.Itoa(r.Int())
}

func VerifyChallengeResponse(challenge, response string) bool {
	hash := sha256.Sum256([]byte(challenge + response))
	return strings.HasPrefix(hex.EncodeToString(hash[:]), strings.Repeat(ChallengeSymbol, ChallengeDifficulty))
}

func SolveChallenge(challenge string) string {
	var response string
	for i := 0; ; i++ {
		response = fmt.Sprintf("%d", i)

		if VerifyChallengeResponse(challenge, response) {
			break
		}
	}

	return response
}
