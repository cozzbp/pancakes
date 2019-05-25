package pancakes

// Pancakes represents a stack of pancakes
type Pancakes struct {
	Stack []rune
}

const (
	// HAPPY represents a pancake happy-face up
	HAPPY = '+'
	// BLANK represents a pancake happy-face down
	BLANK = '-'
)

// New creates a new stack of pancakes
func New(s string) *Pancakes {
	return &Pancakes{
		Stack: []rune(s),
	}
}

// Flip Iterates through the pancake stack and flips as needed
// and returns the number of flips
func (p *Pancakes) Flip() int {
	flips := 0

	for i, v := range p.Stack {
		if i == len(p.Stack)-1 && v == BLANK {
			p.Stack = FlipAtIndex(p.Stack, i)
			flips++
			break
		}
		if i < len(p.Stack)-1 {
			if p.Stack[i+1] != v {
				p.Stack = FlipAtIndex(p.Stack, i)
				flips++
			}
		}
	}

	return flips
}

// FlipAtIndex Flips part of the stack at index i
func FlipAtIndex(s []rune, i int) []rune {
	subStack := s[:i+1]
	flip := make([]rune, len(subStack))

	for i, v := range subStack {
		if v == HAPPY {
			flip[len(subStack)-1-i] = BLANK
		} else if v == BLANK {
			flip[len(subStack)-1-i] = HAPPY
		} else {
			panic("unsupported pancake face type")
		}
	}

	return append(flip, s[i+1:]...)
}
