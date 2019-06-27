package linklist

type IntLinkList interface {
	append(v int) IntLinkList
	contains(v int) bool
	remove(v int) IntLinkList
}

