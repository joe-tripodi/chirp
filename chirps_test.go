package main

import (
	"errors"
	"testing"
)

func TestChirpValidation(t *testing.T) {
	cases := []struct {
		in   string
		want struct {
			out string
			err error
		}
	}{
		{
			in: "Helloooo this is valid",
			want: struct {
				out string
				err error
			}{
				out: "Helloooo this is valid",
				err: nil,
			},
		},
		{
			in: "Kerfuffle",
			want: struct {
				out string
				err error
			}{
				out: "****",
				err: nil,
			},
		},
		{
			in: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla convallis egestas rhoncus. Donec facilisis fermentum sem, ac viverra pepe nunc.",
			want: struct {
				out string
				err error
			}{
				out: "",
				err: ErrChirpTooLong,
			},
		},
	}
	for _, c := range cases {
		got, err := validateChirp(c.in)
		if !errors.Is(err, c.want.err) {
			t.Errorf("test: in: %v\nfailed test: errors did not match\nwant: %v, got: %v", c.in, c.want.err, err)
		}
		if got != c.want.out {
			t.Errorf("test: in: %v\nfailed test: output did not match\nwant: %s, got:%s", c.in, c.want.out, got)
		}
	}
}
