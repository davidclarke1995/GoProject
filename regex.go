//David Clarke G00335563
//Adapted from https://web.microsoftstream.com/video/9d83a3f3-bc4f-4bda-95cc-b21c8e67675e
//Reference: https://golang.org/
 
package assets
 
 import (
 	"fmt"
)
 
  type nfa struct {
 	initial *state
 	accept  *state
 }
 
 type state struct {
 	symbol rune
 	edge1  *state
 	edge2  *state
 }

 
func Poregtonfa(pofix string) *nfa {
 	nfastack := []*nfa{}
 
 	for _, r := range pofix {
 		switch r {
 		case '.':
 			frag2 := nfastack[len(nfastack)-1]
 			nfastack = nfastack[:len(nfastack)-1]
 			frag1 := nfastack[len(nfastack)-1]
 			nfastack = nfastack[:len(nfastack)-1]
 
 			frag1.accept.edge1 = frag2.initial
 
 			nfastack = append(nfastack, &nfa{initial: frag1.initial, accept: frag2.accept})
 
 		case '|':
 			frag2 := nfastack[len(nfastack)-1]
 			nfastack = nfastack[:len(nfastack)-1]
 			frag1 := nfastack[len(nfastack)-1]
 			nfastack = nfastack[:len(nfastack)-1]
 
 			initial := state{edge1: frag1.initial, edge2: frag2.initial}
 			accept := state{}
 			frag1.accept.edge1 = &accept
 			frag2.accept.edge1 = &accept
 
 			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
 
 		case '*':
 			frag := nfastack[len(nfastack)-1]
 			nfastack = nfastack[:len(nfastack)-1]
 
 			accept := state{}
 			initial := state{edge1: frag.initial, edge2: &accept}
 			frag.accept.edge1 = frag.initial
 			frag.accept.edge2 = &accept
 
 			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
 
 		default:
 			accept := state{}
 			initial := state{symbol: r, edge1: &accept}
 
 			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
 		}
 	}
	//error on the stack
 	if len(nfastack) != 1 {
		fmt.Println("Error, ", len(nfastack), " NFA's have been found, they are: ", nfastack)
 	}
 
 	return nfastack[0]
}
 
 func addState(l []*state, s *state, a *state) []*state {
 	l = append(l, s)
 
 	if s != a && s.symbol == 0 {
 		l = addState(l, s.edge1, a)
 		if s.edge2 != nil {
 			l = addState(l, s.edge2, a)
 		}
 	}
 	return l
 }
 
func Pomatch(po string, s string) bool {
 	ismatch := false

	ponfa := Poregtonfa(po)
 
 	current := []*state{}
 	next := []*state{}
 	current = addState(current[:], ponfa.initial, ponfa.accept)
 
 	for _, r := range s {
 		for _, c := range current {
 			if c.symbol == r {
 				next = addState(next[:], c.edge1, ponfa.accept)
 			}
 		}
 		current, next = next, []*state{}
 	}
 
 	for _, c := range current {
 		if c == ponfa.accept {
 			ismatch = true
 			break
 		}
 	}
 
 	return ismatch
 }