// Ciphers with Arbitrary Finite Domain (CAFD)
package cafd

import "math/rand"

type CAFD[T any] struct {
	elements []T
	perms    []int
}

func NewCAFD[T any](e []T) *CAFD[T] {
	output := CAFD[T]{
		elements: e,
	}

	perms := rand.Perm(len(e))

	output.perms = perms

	return &output
}

func (c *CAFD[T]) Next() interface{} {
	// Get next index and return element
	i := c.perms[0]
	elem := c.elements[i]

	// Remove used index
	c.perms = c.perms[1:]

	return elem
}
