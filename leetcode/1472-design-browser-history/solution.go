package designbrowserhistory

type ListNode struct {
	val  string
	prev *ListNode
	next *ListNode
}
type BrowserHistory struct {
	cur *ListNode
}

func Constructor(homepage string) BrowserHistory {
	return BrowserHistory{
		cur: &ListNode{val: homepage},
	}
}

func (this *BrowserHistory) Visit(url string) {
	this.cur.next = &ListNode{
		val:  url,
		prev: this.cur,
	}
	this.cur = this.cur.next
}

func (this *BrowserHistory) Back(steps int) string {
	for this.cur.prev != nil && steps > 0 {
		this.cur = this.cur.prev
		steps--
	}
	return this.cur.val
}

func (this *BrowserHistory) Forward(steps int) string {
	for this.cur.next != nil && steps > 0 {
		this.cur = this.cur.next
		steps--
	}
	return this.cur.val
}
