package mcontents

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Prefix []string

// String returns the Prefix as a string (for use as a map key).
func (p Prefix) String() string {
	return strings.Join(p, " ")
}

// Shift removes the first word from the Prefix and appends the given word.
func (p Prefix) Shift(word string) {
	copy(p, p[1:])
	p[len(p)-1] = word
}

type Chain struct {
	chain     map[string][]string
	prefixLen int
}

func NewChain(prefixLen int) *Chain {
	return &Chain{make(map[string][]string), prefixLen}
}

func (c *Chain) Write(b []byte) {
	br := bytes.NewReader(b)
	p := make(Prefix, c.prefixLen)
	for {
		var s string
		if _, err := fmt.Fscan(br, &s); err != nil {
			break
		}
		key := p.String()
		c.chain[key] = append(c.chain[key], s)
		p.Shift(s)
	}
}

func (c *Chain) Generate(n int, keyword string) string {
	p := make(Prefix, c.prefixLen)
	var words []string
	for i := 0; i < n; i++ {
		choices := c.chain[p.String()]
		if len(choices) == 0 {
			break
		}
		next := choices[rand.Intn(len(choices))]
		//		if i == 0 {
		//			words = append(words, comutils.UpcaseInitial(keyword))
		//		} else {

		words = append(words, next)
		//		}
		p.Shift(next)
	}
	return strings.Join(words, " ")
}

func Generate(contents []string) string {

	rand.Seed(time.Now().UnixNano())
	var buffer bytes.Buffer
	prefixLen := 1
	numWords := 1000
	c := NewChain(prefixLen)
	
	for _, phrase := range contents {
		
		buffer.WriteString(phrase)
				
	}	
	c.Write(buffer.Bytes())
	text := c.Generate(numWords, "job")
	
	return text
}
