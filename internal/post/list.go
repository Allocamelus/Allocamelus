package post

type ListPosts map[int64]*Post

type List struct {
	Posts ListPosts `json:"posts"`
}
