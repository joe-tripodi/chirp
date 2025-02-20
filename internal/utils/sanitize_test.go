package utils

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
	for _, c := range cases {
		got := cleanChirp(c.input)
		if c.want != got {
			t.Errorf("want: %s, got: %s", c.want, got)
		}
	}
}
