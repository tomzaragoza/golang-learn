package main
import ("fmt"; "math")

type Node struct {
	/* Node structure for binary trees */
	value float64
	height float64
	left *Node
	right *Node
}

func left_count(root *Node) int {
	/* Return the total number of left nodes in the binary tree rooted at root.*/

	switch {
		case root != nil && root.left != nil:
			return 1 + left_count(root.left) + left_count(root.right)
		case root != nil && root.right != nil:
			return left_count(root.right)
	}
	return 0
} 

func right_count(root *Node) int {
	/* Return the total number of right nodes in the binary tree rooted at root*/
	switch {
		case root != nil && root.right != nil:
			return 1 + left_count(root.left) + left_count(root.right)
		case root != nil && root.left != nil:
			return left_count(root.left)
	}
	return 0
}

/* following are from the answer section of the exercise book */
func minDepth(root *Node) float64 {
	if root == nil {
		return 0
	}
	return 1 + math.Min(minDepth(root.left), minDepth(root.right))
}

func maxDepth(root *Node) float64 {
	if root == nil {
		return 0
	}
	return 1 + math.Max(maxDepth(root.left), maxDepth(root.right))
}

func balanced(root *Node) bool {
	return (maxDepth(root) - minDepth(root) <= 1)
}
/* End Answer from the book */
func is_balanced(root *Node) bool {
	/* Return if the binary tree rooted at root is balance. By definition,
	balanced means no two leaves differ no more than one node from the root*/
	get_heights(root)

	fmt.Println("Left node: ", root.left.height, "Right node: ", root.right.height)

	if 1 < math.Max(root.left.height, root.right.height) - math.Min(root.left.height, root.right.height) {
		return false
	} 
	return true

}

func get_heights(root *Node) float64 {
	/* Recursively get the heights of trees rooted at each node in the binary
	tree rooted at root. */
	switch { 

	 	case root != nil && root.right != nil && root.left != nil:
	 		
	 		root.left.height += 1 + get_heights(root.left)
			root.right.height += 1 + get_heights(root.right)

			return math.Max(root.left.height, root.right.height)

		case root != nil && root.right != nil && root.left == nil:
			
			root.right.height += 1 + get_heights(root.right)
			return root.right.height

		case root != nil && root.right == nil && root.left != nil:
			
			root.left.height += 1 + get_heights(root.left)			
			return root.left.height
	}
	
	return 0
}

func main() {
	root := new(Node) // Remember, new() creates a new pointer of size Node for root
	left1 := new(Node)
	left2 := new(Node)
	left3 := new(Node)
	left4 := new(Node)

	right1 := new(Node)
	right2 := new(Node)
	right3 := new(Node)

	right4 := new(Node)
	right5 := new(Node)

	root.left = left1
	left1.left = left2

	root.right = right1
	right1.left = left3
	right1.right = right2
	right2.left = left4
	right2.right = right3
	right3.right = right4
	right4.right = right5
	// this creates an unbalanced binary tree

	j := is_balanced(root)
	fmt.Println(j)

	k := balanced(root)
	fmt.Println(k)
}