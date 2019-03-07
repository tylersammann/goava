package main

import (
	"github.com/blacklocus/goava/sets"
	"fmt"
)

func main() {
	// new empty set
	emptySet, _ := sets.New()
	printSet("new empty set", emptySet)

	// new set with multiple items
	intSet, _ := sets.New(12, 23, 34)
	printSet("new int set", intSet)

	// new set with multiple items
	badTypeSet, err := sets.New(12, "asdf")
	fmt.Printf("new set (12, 'asdf'): %v - %v\n", badTypeSet, err)

	// Add
	intSet.Add(12)
	printSet("add 12", intSet)
	intSet.Add(13, 5)
	printSet("add (13, 5)", intSet)
	err = intSet.Add(1.2)
	printSet("add 1.2", intSet)
	fmt.Printf("%v\n", err)

	// Has
	fmt.Printf("has 12 %v %v\n", intSet.Values(), intSet.Has(12))
	fmt.Printf("has 14 %v %v\n", intSet.Values(), intSet.Has(14))
	fmt.Printf("has 1.2 %v %v\n", intSet.Values(), intSet.Has(1.2))

	// Remove
	intSet.Remove(13, 12, 23, 5)
	printSet("rm (13, 12, 23, 5)", intSet)
	intSet.Remove(13)
	printSet("rm 13", intSet)
	intSet.Remove(13, 12, 23)
	printSet("rm (13, 12, 23)", intSet)
	intSet.Remove(34)
	printSet("rm 34", intSet)
	err = intSet.Remove(int32(2))
	printSet("rm int32(2)", intSet)
	fmt.Printf("%v\n", err)

	// String set
	strSet, _ := sets.New("this", "here", "is", "cool")
	printSet("new str set", strSet)

	strSet.Add("", "not")
	printSet("add ('', not)", strSet)

	strSet.Add("not")
	printSet("add not", strSet)

	strSet.Remove("")
	printSet("rm ''", strSet)

	// Equals
	fmt.Print("\n== Equality ==\n")

	emptySet2, _ := sets.New()
	fmt.Printf("empty set equals empty set: %v\n", emptySet.Equals(emptySet2))

	strSet2, _ := sets.New(strSet.Values()...)
	fmt.Printf("str set equals str set 2: %v\n", strSet.Equals(strSet2))
	strSet2.Add("non-sense")
	fmt.Printf("should not equal diff str set: %v\n", strSet.Equals(strSet2))

	fmt.Printf("empty set equals empty int set: %v\n", emptySet.Equals(intSet))
	fmt.Printf("str set equals empty int set: %v\n", strSet.Equals(intSet))
	fmt.Printf("empty set equals str set: %v\n", emptySet.Equals(strSet))

	// Contains
	fmt.Print("\n== Contains ==\n")

	c1, _ := sets.New(1, 2, 3)
	c2, _ := sets.New(1, 2, 3)
	c3, _ := sets.New(1, 2)
	c4, _ := sets.New(1, 2, 3, 4, 5)
	c5, _ := sets.New(1, 2, 6, 7)
	c6, _ := sets.New()
	c7, _ := sets.New(1.0)

	printContains(c1, c1)
	printContains(c6, c6)

	printContains(c1, c2)
	printContains(c1, c3)
	printContains(c1, c4)
	printContains(c1, c5)
	printContains(c1, c6)
	printContains(c1, c7)
	printContains(c1, nil)

	// String
	fmt.Println(c1)
	fmt.Println(strSet)
}

func printSet(str string, set sets.Set) {
	fmt.Printf("%s: %v\n", str, set)
}

func printContains(set1 sets.Set, set2 sets.Set) {
	fmt.Printf("%v contains %v: %v\n", set1, set2, set1.Contains(set2))
}
