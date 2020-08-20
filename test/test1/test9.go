package main

type Link struct {
	value   interface{}
	subKey  string
	subType string
	subPtr  *Link
}

// 添加链表节点
func AddLink(topLink *Link, subLink *Link) {
	if topLink.subPtr == nil {
		topLink.subPtr = subLink
	} else {
		AddLink(topLink.subPtr, subLink)
	}
}









func main() {

}
