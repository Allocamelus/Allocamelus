package post

type ListPosts map[int64]*Post

type List struct {
	Posts ListPosts       `json:"posts"`
	Order map[int64]int64 `json:"order"`
}

func NewList() *List {
	l := new(List)
	l.Posts = ListPosts{}
	l.Order = map[int64]int64{}
	return l
}
