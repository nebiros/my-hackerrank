package binary_tree

import "github.com/pkg/errors"

type errTreeNodeNil struct {
}

func (e *errTreeNodeNil) Error() string {
	return "tree node seems empty"
}

type treeNode struct {
	value string
	data  string
	left  *treeNode
	right *treeNode
}

func (n *treeNode) insert(value string, data string) error {
	if n == nil {
		return errors.Wrap(new(errTreeNodeNil), "cannot insert a value into a nil node")
	}

	switch {
	case value == n.value:
		return nil

	case value < n.value:
		if n.left == nil {
			n.left = &treeNode{
				value: value,
				data:  data,
				left:  nil,
				right: nil,
			}

			return nil
		}

		return n.left.insert(value, data)

	case value > n.value:
		if n.right == nil {
			n.right = &treeNode{
				value: value,
				data:  data,
				left:  nil,
				right: nil,
			}

			return nil
		}

		return n.right.insert(value, data)
	}

	return nil
}

func (n *treeNode) find(value string) (string, bool) {
	if n == nil {
		return "", false
	}

	switch {
	case value == n.value:
		return n.data, true

	case value < n.value:
		return n.left.find(value)

	case value > n.value:
		return n.right.find(value)
	}

	return "", false
}

func (n *treeNode) findMax(aNode *treeNode) (self *treeNode, parent *treeNode) {
	if n == nil {
		self = nil
		parent = aNode
		return
	}

	if n.right == nil {
		self = n
		parent = aNode
		return
	}

	return n.right.findMax(n)
}

func (n *treeNode) replaceNode(parent *treeNode, replacement *treeNode) error {
	if n == nil {
		return errors.Wrap(new(errTreeNodeNil), "cannot replace a node into a nil node")
	}

	if n == parent.left {
		parent.left = replacement
		return nil
	}

	parent.right = replacement
	return nil
}

func (n *treeNode) delete(value string, parent *treeNode) error {
	if n == nil {
		return errors.Wrap(new(errTreeNodeNil), "value to be deleted does not exist in the tree")
	}

	switch {
	case value < n.value:
		return n.left.delete(value, n)

	case value > n.value:
		return n.right.delete(value, n)

	default:
		if n.left == nil && n.right == nil {
			return n.replaceNode(parent, nil)
		}

		if n.left == nil {
			return n.replaceNode(parent, n.right)
		}

		if n.right == nil {
			return n.replaceNode(parent, n.left)
		}

		replacement, parentReplacement := n.left.findMax(n)

		n.value = replacement.value
		n.data = replacement.data

		return replacement.delete(replacement.value, parentReplacement)
	}
}

type errTreeNil struct {
}

func (e *errTreeNil) Error() string {
	return "tree seems empty"
}

// Binary tree, https://en.wikipedia.org/wiki/Binary_search_tree.
// Based on https://appliedgo.net/bintree/.
type tree struct {
	root *treeNode
}

func (t *tree) insert(value string, data string) error {
	if t.root == nil {
		t.root = &treeNode{
			value: value,
			data:  data,
			left:  nil,
			right: nil,
		}

		return nil
	}

	return t.root.insert(value, data)
}

func (t *tree) find(value string) (string, bool) {
	if t.root == nil {
		return "", false
	}

	return t.root.find(value)
}

func (t *tree) delete(value string) error {
	if t.root == nil {
		return errors.Wrap(new(errTreeNil), "cannot delete from an empty tree")
	}

	fakeParent := &treeNode{
		value: "",
		data:  "",
		left:  nil,
		right: t.root,
	}

	if err := t.root.delete(value, fakeParent); err != nil {
		return err
	}

	if fakeParent.right == nil {
		t.root = nil
	}

	return nil
}

func (t *tree) traverse(n *treeNode, f func(node *treeNode)) {
	if n == nil {
		return
	}

	t.traverse(n.left, f)
	f(n)
	t.traverse(n.right, f)
}

func main() {

}
