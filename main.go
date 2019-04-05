package main

import (
	"fmt"
	"github.com/tylersammann/goava/sets"
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
	fmt.Printf("empty set equals empty set: %v\n", emptySet.Equals(emptySet2))

	strSet2 := strSet.Copy()
	fmt.Printf("str set equals str set 2: %v\n", strSet.Equals(strSet2))
	strSet2.Add("non-sense")
	fmt.Printf("should not equal diff str set: %v\n", strSet.Equals(strSet2))

	fmt.Printf("empty set equals empty int set: %v\n", emptySet.Equals(intSet))
	fmt.Printf("str set equals empty int set: %v\n", strSet.Equals(intSet))
	fmt.Printf("empty set equals str set: %v\n", emptySet.Equals(strSet))

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

	diff := setA.Difference(setB)
	fmt.Printf("diff %v %v: %v\n", setA, setB, diff)

	diff2 := setB.Difference(setA)
	fmt.Printf("diff %v %v: %v\n", setB, setA, diff2)

	diff3 := setA.Difference(setA)
	fmt.Printf("diff %v %v: %v\n", setA, setA, diff3)

	// Intersection
	fmt.Print("\n== Intersection / Union ==\n")

	setC := sets.New(1, 2, 3)
	setD := sets.New(2, 3, 4, 5)

	inter := setC.Intersection(setD)
	fmt.Printf("intersection %v %v: %v\n", setC, setD, inter)

	inter2 := setC.Intersection(setC)
	fmt.Printf("intersection %v %v: %v\n", setC, setC, inter2)

	inter3 := setC.Intersection(sets.New())
	fmt.Printf("intersection %v %v: %v\n", setC, sets.New(), inter3)

	union := setC.Union(setD)
	fmt.Printf("union %v %v: %v\n", setC, setD, union)

	union2 := setC.Union(setC)
	fmt.Printf("union %v %v: %v\n", setC, setC, union2)

	union3 := setC.Union(sets.New())
	fmt.Printf("union %v %v: %v\n", setC, sets.New(), union3)
}

func printSet(str string, set sets.Set) {
	fmt.Printf("%s: %v\n", str, set)
}

func printContains(set1 sets.Set, set2 sets.Set) {
	fmt.Printf("%v contains %v: %v\n", set1, set2, set1.Contains(set2))
}
