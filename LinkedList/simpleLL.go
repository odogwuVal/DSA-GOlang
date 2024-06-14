package main

type linkedlist struct {
	number int
	next   *LinkedList
}

func(node *LinkedList) insert(num int) {
	var temp := &linkedlist{}
	temp.number = num
	temp.next = nil

	if node == nil {
		node := temp
	}else {
		node.number
	}

}


func main() {
	head := new(LinkedList)
}
