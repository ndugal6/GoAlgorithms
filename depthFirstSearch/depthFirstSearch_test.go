package algoExpert

import (
	"reflect"
	"testing"
)

func NewNode(name string) *Node {
	return &Node{
		Name:     name,
		Children: []*Node{},
	}
}
func (n *Node) addChildren(names ...string) *Node {
	for _, name := range names {
		child := Node{Name: name}
		n.Children = append(n.Children, &child)
	}
	return n
}

func TestDepthFirstSearch(t *testing.T) {
	a, b, c := NewNode("A"), NewNode("B"), NewNode("C")
	a.Children = []*Node{b, c}
	a.Children = []*Node{b, c}
	/*
		play around with difference between two following lines
		d := *b
		d := b
		d.Children = []*Node{c}
	*/
	cases := []struct {
		input          Node
		expectedOutput []string
	}{
		{
			*a,
			[]string{"A", "B", "C"},
		},
	}
	for _, c := range cases {
		emptyString := make([]string, 0)
		got := c.input.DepthFirstSearch(emptyString)
		if !reflect.DeepEqual(got, c.expectedOutput) {
			t.Errorf("DepthFirstSearch of Node:%s == %s, want %s. Output length: %d", c.input.Name, got, c.expectedOutput, len(got))
		}
	}
}
