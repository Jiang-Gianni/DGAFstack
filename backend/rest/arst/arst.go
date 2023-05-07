package arst

type Arst struct {
	ID       int    `json:"id" primaryKey:"1"`
	Name     string `json:"name"`
	ThisBool bool   `json:"thisBool"`
}
