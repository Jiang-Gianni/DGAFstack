//go:generate go run astra_gen.go

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strings"
	"text/template"

	"github.com/Jiang-Gianni/DGAFstack/rest/mystruct"
	"github.com/Jiang-Gianni/DGAFstack/rest/user"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	typeList = []interface{}{
		&mystruct.MyStruct{},
		&user.User{},
	}
	packages = []string{
		"github.com/Jiang-Gianni/DGAFstack/rest/mystruct",
		"github.com/Jiang-Gianni/DGAFstack/rest/user",
	}
	keyspace            = "test"
	sqlFilesPath        = "../astra/tables/"
	sqlTemplateFile     = "astra_tmpl_table.tmpl"
	restApiFilesPath    = "../astra/"
	restApiTemplateFile = "astra_tmpl_api.tmpl"
	TypesGoDB           = map[string]string{
		"bool":   "boolean",
		"float":  "float",
		"int":    "int",
		"string": "text",
	}
	ConvertFunctionMap = func(fieldName string, fieldType string, isUuid bool) []string {
		stringList := []string{
			"client.ToUUID(val)",
			fieldName + ".String()",
			"%s",
		}
		if !isUuid {
			stringList = []string{
				"client.ToString(val)",
				fieldName,
				"'%s'",
			}
		}
		stringMap := map[string][3]string{
			"string": [3]string(stringList),
			"int":    {"client.ToInt(val)", "int(" + fieldName + ")", "%d"},
			"bool":   {"client.ToBoolean(val)", fieldName, "%t"},
		}
		result := [3]string(stringMap[fieldType])
		return result[:]
	}
)

func main() {
	os.RemoveAll(sqlFilesPath)
	os.Mkdir(sqlFilesPath, os.ModePerm)
	for _, typeElement := range typeList {
		GenerateTableSql(typeElement, keyspace)
	}
	GenerateRestApi(typeList, keyspace)
}

func GenerateRestApi[T any](structList []T, keyspace string) {

	var structs []map[string]any

	for _, singleStruct := range structList {
		elem := reflect.ValueOf(singleStruct).Elem()
		structName := elem.Type().Name()
		packageStruct := strings.ToLower(structName) + "." + structName
		n := elem.NumField()
		fields := make([]map[string]string, n)
		var idField string
		fieldList := []string{}
		snakeFieldList := []string{}
		percList := []string{}
		for i := 0; i < n; i++ {
			indexField := elem.Type().Field(i)
			fieldName := indexField.Name
			fieldType := indexField.Type
			_, isPrimaryKey := indexField.Tag.Lookup("primaryKey")
			if isPrimaryKey {
				idField = fieldName
			}
			funcStrings := ConvertFunctionMap(fieldName, fieldType.String(), isPrimaryKey)
			if !isPrimaryKey {
				fieldList = append(fieldList, "row."+fieldName)
				snakeFieldList = append(snakeFieldList, PascalToSnake(fieldName))
				percList = append(percList, funcStrings[2])
			}
			field := map[string]string{
				"FieldName":      fieldName,
				"FirstFunction":  funcStrings[0],
				"SecondFunction": funcStrings[1],
			}
			fields[i] = field
		}
		structs = append(structs, map[string]any{
			"Name":             structName,
			"Pointer":          "*" + packageStruct,
			"PackageStruct":    packageStruct,
			"Fields":           fields,
			"TableName":        PascalToSnake(structName),
			"IdField":          idField,
			"IdFieldSnake":     PascalToSnake(idField),
			"SnakeFieldJoined": strings.Join(snakeFieldList, ", "),
			"PercJoined":       strings.Join(percList, ", "),
			"FieldList":        strings.Join(fieldList, ", "),
			"SnakeFieldList":   snakeFieldList,
			"PercList":         percList,
		})
	}

	utilAstarApi := restApiFilesPath + "util_astra_api.gen.go"
	var helpers template.FuncMap = map[string]interface{}{
		"isNotLast": func(index int, len int) bool {
			return index+1 != len
		},
	}
	tmpl := template.Must(template.New("").Funcs(helpers).ParseFiles(restApiTemplateFile))
	var processed bytes.Buffer
	tmpData := map[string]any{
		"Packages": packages,
		"Structs":  structs,
	}
	tmpl.ExecuteTemplate(&processed, restApiTemplateFile, tmpData)
	fmt.Println("Writing file: ", utilAstarApi)
	f, _ := os.Create(utilAstarApi)
	w := bufio.NewWriter(f)
	w.WriteString(processed.String())
	w.Flush()
	f.Close()
}

func GenerateTableSql[T any](myInstance T, keyspace string) {
	elem := reflect.ValueOf(myInstance).Elem()
	tableName := PascalToSnake(elem.Type().Name())
	outputPath := sqlFilesPath + tableName + ".gen.sql"
	var columns []map[string]any
	n := elem.NumField()
	for i := 0; i < n; i++ {
		field := elem.Type().Field(i)
		_, isPrimaryKey := field.Tag.Lookup("primaryKey")
		columnType := "uuid"
		if !isPrimaryKey {
			columnType = TypesGoDB[field.Type.String()]
		}
		var isLast int
		if n == i+1 {
			isLast = 1
		}
		columns = append(columns, map[string]any{
			"Name":         PascalToSnake(field.Name),
			"Type":         columnType,
			"isPrimaryKey": isPrimaryKey,
			"isLast":       isLast,
		})
	}
	tmpl := template.Must(template.New("").ParseFiles(sqlTemplateFile))
	var processed bytes.Buffer
	tmpData := map[string]any{
		"Keyspace":  PascalToSnake(keyspace),
		"TableName": tableName,
		"Columns":   columns,
	}
	tmpl.ExecuteTemplate(&processed, sqlTemplateFile, tmpData)
	fmt.Println("Writing file: ", outputPath)
	f, _ := os.Create(outputPath)
	w := bufio.NewWriter(f)
	w.WriteString(processed.String())
	w.Flush()
	f.Close()
}

func PascalToSnake(text string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := matchFirstCap.ReplaceAllString(text, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func SnakeToPascal(str string) string {
	list := strings.Split(str, "_")
	for i, v := range list {
		list[i] = cases.Title(language.English).String(v)
	}
	return strings.Join(list, "")
}
