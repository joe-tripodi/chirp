package utils

import "strings"

func cleanChirp(chirp string) string {
	tokens := strings.Split(chirp, " ")
	for i, token := range tokens {
		lower := strings.ToLower(token)
		if lower == "sharbert" || lower == "kerfuffle" || lower == "fornax" {
			tokens[i] = "****"
		}
	}

	return strings.Join(tokens, " ")
}
