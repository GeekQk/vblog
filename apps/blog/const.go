package blog

type Status int

const (
	//草稿
	STATUS_ORAFT Status = iota
	//发布
	STATUS_PUBLISHED
)
