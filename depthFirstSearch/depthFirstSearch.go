package algoExpert

/*
Description: Given a Node class that has a name and an array of optional children Nodes,. When put otgether, Nodes form a simple tree like structure. Implement the depthFirstSerach method on the Node class, which takes in an empty
array, traverses the tree using the DF search approach (nav left to right), stores all the of the Nodes' names in the input array, and returns it
*/
type Node struct {
	Name     string
	Children []*Node
}

func (n *Node) DepthFirstSearch(array []string) []string {
	result := make([]string, 0)
	result = append(result, n.Name)

	for _, v := range n.Children {
		var a []string
		result = append(result, v.DepthFirstSearch(a)...)
	}
	return result
}
