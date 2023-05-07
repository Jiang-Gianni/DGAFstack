package mystruct

type MyStruct struct {
	Uuid      string `json:"uuid" primaryKey:"1"`
	Name      string `json:"name"`
	Number    int    `json:"number"`
	MyBoolean bool   `json:"myBoolean"`
}
