package main

import (
	"github.com/blacklocus/goava/sets"
	"fmt"
)

func main() {
	// new empty set
	emptySet := sets.New()
	printSet("new empty set", emptySet)

	// new set with multiple items
	intSet := sets.New(12, 23, 34)
	printSet("new int set", intSet)

	// new set with multiple bad items
	//badTypeSet := sets.New(12, "asdf")
	//fmt.Printf("new set (12, 'asdf'): %v - %v\n", badTypeSet, err)

	// Add
	intSet.Add(12)
	printSet("add 12", intSet)
	intSet.Add(13, 5)
	printSet("add (13, 5)", intSet)
	//err = intSet.Add(1.2)
	//printSet("add 1.2", intSet)
	//fmt.Printf("%v\n", err)

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
	//err = intSet.Remove(int32(2))
	//printSet("rm int32(2)", intSet)
	//fmt.Printf("%v\n", err)

	// String set
	strSet := sets.New("this", "here", "is", "cool")
	printSet("new str set", strSet)

	strSet.Add("", "not")
	printSet("add ('', not)", strSet)

	strSet.Add("not")
	printSet("add not", strSet)

	strSet.Remove("")
	printSet("rm ''", strSet)

	// Equals
	fmt.Print("\n== Equality ==\n")

	emptySet2 := sets.New()
	fmt.Printf("empty set equals empty set: %v\n", sets.Equals(emptySet, emptySet2))

	strSet2 := strSet.Copy()
	fmt.Printf("str set equals str set 2: %v\n", sets.Equals(strSet, strSet2))
	strSet2.Add("non-sense")
	fmt.Printf("should not equal diff str set: %v\n", sets.Equals(strSet, strSet2))

	fmt.Printf("empty set equals empty int set: %v\n", sets.Equals(emptySet, intSet))
	fmt.Printf("str set equals empty int set: %v\n", sets.Equals(strSet, intSet))
	fmt.Printf("empty set equals str set: %v\n", sets.Equals(emptySet, strSet))

	// Contains
	fmt.Print("\n== Contains ==\n")

	c1 := sets.New(1, 2, 3)
	c2 := sets.New(1, 2, 3)
	c3 := sets.New(1, 2)
	c4 := sets.New(1, 2, 3, 4, 5)
	c5 := sets.New(1, 2, 6, 7)
	c6 := sets.New()
	c7 := sets.New(1.0)

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

	// Complex types
	fmt.Print("\n== Complex Types ==\n")

	ct1 := sets.New([3]int{1, 2, 3}, [3]int{4, 5, 6})
	fmt.Println(ct1)

	ct2 := sets.New([3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, [3][3]int{{1}, {4}, {7}})
	fmt.Println(ct2)

	ct3 := sets.New([2]string{"hi", "there"}, [2]string{"it's", "me"})
	fmt.Println(ct3)

	//ct4 := sets.New([2]string{"hi", "there"}, [3]string{"it's", "me"})
	//fmt.Println(ct4)
	//
	//ct5 := sets.New([2]string{"hi", "there"}, [2]int{2, 3})
	//fmt.Printf("%v %v\n", ct5, err)

	copy1 := ct1.Copy()
	fmt.Println(copy1)

	// Difference
	fmt.Print("\n== Difference ==\n")

	setA := sets.New(1, 2, 3)
	setB := sets.New(2, 3, 4, 5)

	diff := sets.Difference(setA, setB)
	fmt.Printf("diff %v %v: %v\n", setA, setB, diff)

	diff2 := sets.Difference(setB, setA)
	fmt.Printf("diff %v %v: %v\n", setB, setA, diff2)

	diff3 := sets.Difference(setA, setA)
	fmt.Printf("diff %v %v: %v\n", setA, setA, diff3)

	// Intersection
	fmt.Print("\n== Intersection / Union ==\n")

	setC := sets.New(1, 2, 3)
	setD := sets.New(2, 3, 4, 5)

	inter := sets.Intersection(setC, setD)
	fmt.Printf("intersection %v %v: %v\n", setC, setD, inter)

	inter2 := sets.Intersection(setC, setC)
	fmt.Printf("intersection %v %v: %v\n", setC, setC, inter2)

	inter3 := sets.Intersection(setC, sets.New())
	fmt.Printf("intersection %v %v: %v\n", setC, sets.New(), inter3)


	union := sets.Union(setC, setD)
	fmt.Printf("union %v %v: %v\n", setC, setD, union)

	union2 := sets.Union(setC, setC)
	fmt.Printf("union %v %v: %v\n", setC, setC, union2)

	union3 := sets.Union(setC, sets.New())
	fmt.Printf("union %v %v: %v\n", setC, sets.New(), union3)
}

func printSet(str string, set sets.Set) {
	fmt.Printf("%s: %v\n", str, set)
}

func printContains(set1 sets.Set, set2 sets.Set) {
	fmt.Printf("%v contains %v: %v\n", set1, set2, sets.Contains(set1, set2))
}
