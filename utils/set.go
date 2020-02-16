package utils

import (
	"fmt"
	"reflect"
	"sort"
)

// NewStringSet creates a new string set with the given contents.
func NewStringSet(contents ...string) StringSet {
	set := StringSet{}
	set.Add(contents...)
	return set
}

// NewStringSet creates a new string set with the given contents.
func NewStringSetFromStringMapKeys(contents map[string]string) StringSet {
	set := StringSet{}
	for k, _ := range contents {
		set.Add(k)
	}
	return set
}

// StringSet is a set data structure for strings
type StringSet map[string]bool

// Add puts one or more elements into the set.
func (set StringSet) Add(str ...string) {
	for _, s := range str {
		set[s] = true
	}
}

// ToSlice converts the set to a slice with sort.Strings order.
func (set StringSet) ToSlice() []string {
	out := []string{}
	for k := range set {
		out = append(out, k)
	}

	sort.Strings(out)
	return out
}

// IsEmpty determines if the set has zero elements.
func (set StringSet) IsEmpty() bool {
	return len(set) == 0
}

// Equals compares the contents of the two sets and returns true if they are
// the same.
func (set StringSet) Equals(other StringSet) bool {
	return reflect.DeepEqual(set, other)
}

// Contains performs a set membership check.
func (set StringSet) Contains(other string) bool {
	_, ok := set[other]
	return ok
}

// Returns a copy of this set with every string in the other removed.
func (set StringSet) Minus(other StringSet) StringSet {
	difference := NewStringSet()

	for k, _ := range set {
		if !other.Contains(k) {
			difference.Add(k)
		}
	}

	return difference
}

// String converts this set to a human readable string.
func (set StringSet) String() string {
	return fmt.Sprintf("%v", set.ToSlice())
}
