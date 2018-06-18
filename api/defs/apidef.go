package defs

type UserCreatedential struct {
	Username string `json:"user_name"`
	Pwd string `hson:"pwd"`
}
type VideoInfo struct {
	Id string
	AuthorId int
	Name string
	DisplayCtime string
}
type Comment struct {
	Id string
	VideoId string
	Author string
	Content string
}

type SimpleSession struct {
	Username string
	TTL int64
}