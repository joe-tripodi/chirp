package sanitize

import "testing"

func TestCleaningUpChirp(t *testing.T) {
	cases := []struct {
		input string
		want  string
	}{
		{
			input: "Sharbert",
			want:  "****",
		},
		{
			input: "I had something interesting for breakfast",
			want:  "I had something interesting for breakfast",
		},
		{
			input: "I hear Mastodon is better than Chirpy. sharbert I need to migrate",
			want:  "I hear Mastodon is better than Chirpy. **** I need to migrate",
		},
	}
	badWords := map[string]struct{}{
		"kerfuffle": {},
		"sharbert":  {},
		"fornax":    {},
	}
	for _, c := range cases {
		got := CleanChirp(c.input, badWords)
		if c.want != got {
			t.Errorf("want: %s, got: %s", c.want, got)
		}
	}
}
