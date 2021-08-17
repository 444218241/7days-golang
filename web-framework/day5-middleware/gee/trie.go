package gee

import "strings"

/**
1、开源的路由实现gorouter支持在路由规则中嵌入正则表达式，
	例如/p/[0-9A-Za-z]+，
	即路径中的参数仅匹配数字和字母；
2、另一个开源实现httprouter就不支持正则表达式。
	著名的Web开源框架gin 在早期的版本，并没有实现自己的路由，
	而是直接使用了httprouter，
	后来不知道什么原因，放弃了httprouter，自己实现了一个版本。
*/
type node struct {
	pattern  string  // 路由，例如 /p/:lang
	part     string  // 路由中的一部分，例如 :lang
	children []*node // 子节点，例如 [doc, tutorial, intro]
	isWild   bool    // 是否模糊匹配，part 含有 : 或 * 时为true
}

// 第一个匹配成功的节点，用于插入
// 因为插入的话，第一个匹配到就行。
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// 所有匹配成功的节点，用于查找
// 由于使用了正则，而不是精确匹配，所以要取出所有符合到节点用于查找。
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{
			part:   part,
			isWild: part[0] == ':' || part[0] == '*',
		}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)

	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}
