package invite

import "sort"

// By handles customer ordering
type By func(c1, c2 *Customer) bool

// Sort sorts customers data according to By
func (by By) Sort(customers []Customer) {
	cs := &customerSorter{
		custs: customers,
		by:    by,
	}
	sort.Sort(cs)
}

type customerSorter struct {
	custs []Customer
	by    func(c1, c2 *Customer) bool
}

func (c *customerSorter) Len() int {
	return len(c.custs)
}

func (c *customerSorter) Swap(i, j int) {
	c.custs[i], c.custs[j] = c.custs[j], c.custs[i]
}

func (c *customerSorter) Less(i, j int) bool {
	return c.by(&c.custs[i], &c.custs[j])
}
