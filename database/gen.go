//go:build ignore

package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"

	_ "github.com/lib/pq"

	"github.com/go-jet/jet/v2/generator/metadata"
	"github.com/go-jet/jet/v2/generator/postgres"
	"github.com/go-jet/jet/v2/generator/template"
	postgres2 "github.com/go-jet/jet/v2/postgres"
)

func main() {
	err := generate()

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func generate() (err error) {
	err = postgres.Generate(
		"./gen", // or GenerateDSN(...)
		postgres.DBConnection{
			Host:       "db",
			Port:       5432,
			User:       os.Getenv("DB_USER"),
			Password:   os.Getenv("DB_PASSWORD"),
			DBName:     os.Getenv("DB_NAME"),
			SchemaName: "public",
			SslMode:    "disable",
		},
		template.Default(postgres2.Dialect).
			UseSchema(func(schemaMetaData metadata.Schema) template.Schema {
				return template.DefaultSchema(schemaMetaData).
					UseModel(template.DefaultModel().
						UseTable(func(table metadata.Table) template.TableModel {
							return template.DefaultTableModel(table).
								UseField(func(columnMetaData metadata.Column) template.TableModelField {
									defaultTableModelField := template.DefaultTableModelField(columnMetaData)
									return defaultTableModelField.UseTags(
										fmt.Sprintf(`json:"%s"`, SnakeToCamel(columnMetaData.Name, false)),
									)
								})
						}),
					)
			}),
	)

	if err != nil {
		return fmt.Errorf("could not generate database files: %w", err)
	}

	return err
}

// SnakeToCamel returns a string converted from snake case to uppercase
func SnakeToCamel(s string, firstLetterUppercase ...bool) string {
	upperCase := true
	if len(firstLetterUppercase) > 0 {
		upperCase = firstLetterUppercase[0]
	}
	return snakeToCamel(s, upperCase)
}

func snakeToCamel(s string, upperCase bool) string {
	if len(s) == 0 {
		return s
	}
	var result string

	words := strings.Split(s, "_")

	for i, word := range words {
		if upperCase || i > 0 {
			result += camelizeWord(word, len(words) > 1)
		} else {
			result += word
		}
	}

	return result
}

func camelizeWord(word string, force bool) string {
	runes := []rune(word)

	for i, r := range runes {
		if i == 0 {
			runes[i] = unicode.ToUpper(r)
		} else {
			if !force && unicode.IsLower(r) { // already camelCase
				return string(runes)
			}

			runes[i] = unicode.ToLower(r)
		}
	}

	return string(runes)
}
