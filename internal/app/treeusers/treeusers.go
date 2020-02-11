package treeusers

import (
	"github.com/MishaNiki/go-login-signup/internal/app/model"
	"github.com/emirpasic/gods/tree/master/trees/redblacktree"
)

// TreeUsers red-black tree of users
type TreeUsers redblacktree.Tree

// New tree
func New() *TreeUsers {
	return redblacktree.NewWithIntComparator()
}

// AddUser ...
func (tu *TreeUsers) AddUser(key int, u *model.User) {
	tu.Put(key, u)
}

// DeleteUser ...
func (tu *TreeUsers) DeleteUser(key int) {
	tu.Remove(key)
}

// GetUser ...
func (tu *TreeUsers) GetUser(key int) (*model.User, bool) {
	u, ok := tu.Get(key)
	if !ok {
		return nil, ok
	}
	user := u.(*model.User)
	return user, ok
}

// ф: Количество узлов в дереве
// ф: Чистка всего дерева
