package main

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestWholeFile(t *testing.T) {
	ioutil.WriteFile("source.txt", []byte("abcd"), 0666)
	gocopy("source.txt", "dest.txt", 0, 0)
	bytes, _ := ioutil.ReadFile("dest.txt")
	assert.Equal(t, bytes, []byte("abcd"))
}

func TestOffsetAndLimit(t *testing.T) {
	ioutil.WriteFile("source.txt", []byte("abcd"), 0666)
	gocopy("source.txt", "dest.txt", 1, 2)
	bytes, _ := ioutil.ReadFile("dest.txt")
	assert.Equal(t, bytes, []byte("bc"))
}
