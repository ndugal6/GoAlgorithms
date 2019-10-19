package algoExpert

type BST struct {
	value int

	left  *BST
	right *BST
}

var leaf1 = BST{1, nil, nil}
var leaf5 = BST{5, nil, nil}
var leaf14 = BST{14, nil, nil}
var leaf22 = BST{22, nil, nil}
var d3_2 = BST{2, &leaf1, nil}
var d3_5 = leaf5
var d3_13 = BST{13, nil, &leaf14}
var d3_22 = leaf22
var d2_5 = BST{5, &d3_2, &leaf5}
var d2_15 = BST{15, &d3_13, &leaf22}
var d1_10 = BST{10, &d2_5, &d2_15}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

var closest = 0

func (tree *BST) FindClosestValue(target int) int {
	nextNode := tree.NextNode(target)

	switch {
	case nextNode == nil:
		return tree.value
	case nextNode.value == target:
		return target
	}

	nextClosest := nextNode.FindClosestValue(target)
	if Abs(target-nextClosest) > Abs(target-tree.value) {
		return tree.value
	}
	return nextClosest
}

func (tree *BST) NextNode(target int) *BST {
	switch {
	case tree.value > target:
		return tree.left
	case tree.value < target:
		return tree.right
	}
	return tree
}
