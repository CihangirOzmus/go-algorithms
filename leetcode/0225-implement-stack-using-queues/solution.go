package implementstackusingqueues

type MyStack struct {
	q []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.q = append(this.q, x)
	for i := 0; i < len(this.q)-1; i++ {
		front := this.q[0]
		this.q = this.q[1:]
		this.q = append(this.q, front)
	}
}

func (this *MyStack) Pop() int {
	if len(this.q) > 0 {
		pop := this.q[0]
		this.q = this.q[1:]
		return pop
	}
	return -1
}

func (this *MyStack) Top() int {
	if len(this.q) > 0 {
		top := this.q[0]
		return top
	}
	return -1
}

func (this *MyStack) Empty() bool {
	return len(this.q) == 0
}
