package router

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type TreeNode struct {
	m         map[string]*TreeNode
	leafNode  bool
	paramNode bool
	Params    *ParamNode
	Handle    Handler
}

type ParamNode struct {
	name  string
	value interface{}
}

func NewTreeNode() *TreeNode {
	return &TreeNode{
		m: make(map[string]*TreeNode),
	}
}

func (t *TreeNode) Insert(s string, handle Handler) {
	nodes := strings.Split(s, "/")
	reg := regexp.MustCompile("{.*}")

	for _, node := range nodes {
		if _, ok := t.m[node]; !ok {
			if reg.Match([]byte(node)) {
				params := reg.FindStringSubmatch(node)
				t.m[node] = &TreeNode{
					m: make(map[string]*TreeNode),
					Params: &ParamNode{
						name: params[0],
					},
					paramNode: true,
				}
			} else {
				t.m[node] = &TreeNode{
					m: make(map[string]*TreeNode),
				}
			}
		}
		t = t.m[node]
	}

	t.leafNode = true
	t.Handle = handle

	// for i := 0; i < len(s); i++ {
	// 	if s[i] == '/' {
	// 		continue
	// 	}

	// 	if _, ok := t.m[s[i]]; !ok {
	// 		t.m[s[i]] = &TreeNode{
	// 			m: make(map[uint8]*TreeNode),
	// 		}
	// 	}
	// 	t = t.m[s[i]]
	// }
	// t.leafNode = true
	// t.Handle = handle
}

func (t *TreeNode) Search(s string) (Handler, error) {
	// for i := 0; i < len(s); i++ {
	// 	if _, ok := t.m[s[i]]; !ok {
	// 		return nil, errors.New("subString not found")
	// 	}
	// 	t = t.m[s[i]]
	// }

	// if !t.leafNode {
	// 	return nil, errors.New("Not LeafNode")
	// }

	// return t.Handle, nil

	nodes := strings.Split(s, "/")
	fmt.Println("++++++ Search:", nodes)
	for _, node := range nodes {
		if t.paramNode {
			fmt.Println(t.Params.name, node)
			t = t.m[node]
			continue
		}

		if _, ok := t.m[node]; !ok {
			fmt.Println("+++++++ Search not found:", node)
			return nil, errors.New("subString not found")
		}

		t = t.m[node]
	}

	if !t.leafNode {
		return nil, errors.New("Not LeafNode")
	}

	return t.Handle, nil
}

func withContext(ctx context.Context, f Handler) func(c *context.Context) {
	return func(c *context.Context) {

	}
}
