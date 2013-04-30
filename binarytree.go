package main
import ("fmt"; "math")

type Node struct {
	value float64
	height float64
	left *Node
	right *Node
}


func left_count(root *Node) int {
	switch {
		case root != nil && root.left != nil:
			return 1 + left_count(root.left) + left_count(root.right)
		case root != nil && root.right != nil:
			return left_count(root.right)
	}
	return 0
} 

func right_count(root *Node) int {
	switch {
		case root != nil && root.right != nil:
			return 1 + left_count(root.left) + left_count(root.right)
		case root != nil && root.left != nil:
			return left_count(root.left)
	}
	return 0
}

func is_balanced(root *Node) bool {

	//fmt.Println(root.left.height, "Left")
	//fmt.Println(root.right.height, "Right")
	get_heights(root)

	if 1 < math.Max(root.left.height, root.right.height) - math.Min(root.left.height, root.right.height) {
		return false
	} 
	return true

}

func get_heights(root *Node) float64 {
	// get heights of left and right nodes
	// return -> (left_height, right_height)

	/*fmt.Println(root.left.height, "Left")
	fmt.Println(root.right.height, "Right")
	fmt.Println("------------------")*/

	switch { 

	 	case root != nil && root.right != nil && root.left != nil:
	 		
	 		root.left.height = 1 + get_heights(root.left)
			root.left.height = 1 + get_heights(root.right)

			return math.Max(root.left.height, root.right.height)

		case root != nil && root.right != nil && root.left == nil:
			
			root.left.height = 1 + get_heights(root.left)
			return root.left.height

		case root != nil && root.right == nil && root.left != nil:
			
			root.right.height = 1 + get_heights(root.right)			
			return root.right.height
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

	//x := left_count(root)
	//fmt.Println(x)

	j := is_balanced(root)
	fmt.Println(j)
}