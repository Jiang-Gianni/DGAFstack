package mytable

type MyTable struct {
	Uuid   string `json:"uuid" primaryKey:"1"`
	Name   string `json:"name"`
	Number int    `json:"number"`
}
