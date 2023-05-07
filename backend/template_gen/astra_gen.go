//go:generate go run astra_gen.go astra_tmpl_table.tmpl

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"os"
	"reflect"
	"regexp"
	"strings"

	"github.com/Jiang-Gianni/DGAFstack/rest/arst"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	typeList := []interface{}{
		&arst.Arst{},
		// &user.User{},
	}

	keyspace := "test"

	for _, typeElement := range typeList {
		generateTableSql(typeElement, keyspace)
	}

}

var (
	TypesGoDB = map[string]string{
		"bool":   "boolean",
		"float":  "float",
		"int":    "int",
		"string": "text",
	}
)

func generateTableSql[T any](myInstance T, keyspace string) {

	elem := reflect.ValueOf(myInstance).Elem()

	var columns []map[string]any

	n := elem.NumField()
	for i := 0; i < n; i++ {
		field := elem.Type().Field(i)
		_, isPrimaryKey := field.Tag.Lookup("primaryKey")
		var isLast int
		if n == i+1 {
			isLast = 1
		}
		columns = append(columns, map[string]any{
			"Name":         pascalToSnake(field.Name),
			"Type":         TypesGoDB[field.Type.String()],
			"isPrimaryKey": isPrimaryKey,
			"isLast":       isLast,
		})
		fmt.Println(field.Name)
		fmt.Println(isLast)
	}

	fileName := os.Args[1]
	tmpl := template.Must(template.New("").ParseFiles(fileName))
	var processed bytes.Buffer

	tableName := elem.Type().Name()
	tmpData := map[string]any{
		"Keyspace":  pascalToSnake(keyspace),
		"TableName": pascalToSnake(tableName),
		"Columns":   columns,
	}

	tmpl.ExecuteTemplate(&processed, fileName, tmpData)

	sqlFilesPath := "../astra/tables/"
	os.Mkdir(sqlFilesPath, os.ModePerm)
	outputPath := sqlFilesPath + tableName + ".gen.sql"
	fmt.Println("Writing file: ", outputPath)

	f, _ := os.Create(outputPath)
	w := bufio.NewWriter(f)
	w.WriteString(processed.String())
	w.Flush()

}

func pascalToSnake(text string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := matchFirstCap.ReplaceAllString(text, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func snakeToPascal(str string) string {
	list := strings.Split(str, "_")
	for i, v := range list {
		list[i] = cases.Title(language.English).String(v)
	}
	return strings.Join(list, "")
}
