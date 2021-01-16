package main

import "fmt"

const maxLEN = 1000

var (
	arrSize int
	array   []int
	tree    [maxLEN]int
)

func build(arr []int, size int) {
	array = arr
	arrSize = size
	buildTree(0, 0, size-1)
}

func printTree(size int) {
	for i := 0; i < size; i++ {
		fmt.Printf("tree[%d] = %d\n", i, tree[i])
	}
}

func buildTree(node int, start int, end int) {
	if start == end {
		tree[node] = array[start]
	} else {
		mid := (start + end) / 2
		leftNode := 2*node + 1
		rightNode := 2*node + 2
		buildTree(leftNode, start, mid)
		buildTree(rightNode, mid+1, end)
		tree[node] = tree[leftNode] + tree[rightNode]
	}
}

func query(left int, right int) int {
	return queryTree(0, 0, arrSize-1, left, right)
}

func queryTree(node int, start int, end int, l int, r int) int {
	if r < start || l > end {
		return 0
	} else if start == end {
		return tree[node]
	} else {
		mid := (start + end) / 2
		leftNode := 2*node + 1
		rightNode := 2*node + 2
		leftSum := queryTree(leftNode, start, mid, l, r)
		rightSum := queryTree(rightNode, mid+1, end, l, r)
		return leftSum + rightSum
	}
}

func update(idx int, val int) {
	array[idx] = val
	updateTree(0, 0, arrSize-1, idx, val)
}

func updateTree(node int, start int, end int, idx int, val int) {
	if start == end {
		tree[node] = val
	} else {
		mid := (start + end) / 2
		leftNode := 2*node + 1
		rightNode := 2*node + 2
		if idx >= start && idx <= mid {
			updateTree(leftNode, start, mid, idx, val)
		} else {
			updateTree(rightNode, mid+1, end, idx, val)
		}
		tree[node] = tree[leftNode] + tree[rightNode]
	}
}
