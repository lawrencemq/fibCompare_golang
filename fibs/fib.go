package fibs



func FibTwoNumbers(number int) int64 {
	if number <= 1 {
		return 1
	}
	var first, second int64 = 1, 1
	onFirst := true

	for i := 0; i < number -1; i++ {
		if onFirst {
			first = first + second
		}else{
			second = first + second
		}
		onFirst = !onFirst
	}

	if !onFirst {
		return first
	}else{
		return second
	}
}

func FibLinkedList(number int) int64 {
	if number <= 1 {
		return 1
	}
	list := CreateLinkedList()
	list.Append(1)
	list.Append(1)

	for i := 0; i < number-1; i++ {
		newValue := list.rootNode.value + list.rootNode.nextNode.value
		list.Append(newValue)
		list.PopFirst()
	}

	return list.endNode.value

}

func FibRecursive(number int) int64 {
	if number <= 1 {
		return 1
	}
	return FibRecursive(number - 1) + FibRecursive(number-2)
}
