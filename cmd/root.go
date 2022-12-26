package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"sort"

	"github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
	"github.com/spf13/cobra"
)

type Property struct {
	Description string    `json:"description"`
	Type        string    `json:"type"`
	Format      *string   `json:"format"`
	Nullable    *bool     `json:"nullable"`
	Enum        *[]string `json:"enum"`
	Items       *struct {
		Description string `json:"description"`
		Type        string `json:"type"`
	}
}

type Model struct {
	Type        string              `json:"type"`
	Required    []string            `json:"required"`
	Properties  map[string]Property `json:"properties"`
	Description string              `json:"description"`
	DBName      string              `json:"dbName"`
}

func identifier(val string) string {
	switch val {
	case "id":
		return "ID"
	}
	return strcase.ToCamel(val)
}

func toType(file *jen.File, modelName string, stmt *jen.Statement, name string, property Property, tags map[string]string) *jen.Statement {
	_stmt := stmt
	if property.Nullable != nil && *property.Nullable {
		_stmt = _stmt.Op("*")
	}

	switch property.Type {
	case "string":
		if property.Format != nil && *property.Format == "date-time" {
			_stmt = _stmt.Qual("time", "Time")
			switch name {
			case "createdDate", "updatedDate":
				tags["gorm"] = "column:" + name
			default:
			}
			return _stmt
		} else if name == "id" {
			tags["gorm"] = "primaryKey"
		}
		if property.Enum != nil {
			enumPrefix := modelName + strcase.ToCamel(name)
			enumName := enumPrefix + "Enum"
			file.Type().Id(enumName).String()
			file.Line()
			vals := make([]jen.Code, 0)
			for i, val := range *property.Enum {
				if i == 0 {
					vals = append(vals, jen.Id(enumPrefix+val).Op(enumName).Op("=").Lit(val))
				} else if val != "" {
					vals = append(vals, jen.Id(enumPrefix+val).Op("=").Lit(val))
				}
			}
			file.Const().Defs(vals...)
			return _stmt.Id(enumName)
		}
		return _stmt.String()
	case "integer":
		return _stmt.Int64()
	case "number":
		return _stmt.Float64()
	case "boolean":
		return _stmt.Bool()
	case "array":
		if property.Items != nil {
			switch (*property.Items).Type {
			case "string":
				return _stmt.Op("[]").String()
			case "number":
				return _stmt.Op("[]").Float64()
			}
		}
	}
	if name == "meta" || name == "metadata" {
		tags["gorm"] = "embedded;column:" + name
		tags["json"] = name + ",omitempty"
		if name == "meta" {
			return _stmt.Op("*").Id("Meta")
		}
	}
	return _stmt.Any()
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "generate [input-file] [output-dir]",
	Short: "Generate datamodel files from a schema",
	Long: `This command will process a JSON schema and generate Go datamodel
files for each model. You can download the schema by accessing the API endpoint such as:

curl https://api.shopmonkey.io/v3/schema -O schema.json

Make sure you use the correct version to generate the schema.  Once you have the schema, you
can now generate a set of go source files for each model:

	$ go run . schema.json v3

Where v3 is the output folder to generate the files into.
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
		fn := args[0]
		if _, err := os.Stat(fn); err != nil {
			fmt.Printf("File %s doesn't exist\n", fn)
			os.Exit(1)
		}
		outdir := args[1]
		if _, err := os.Stat(outdir); err != nil {
			if err := os.MkdirAll(outdir, 0750); err != nil {
				panic(err)
			}
		}
		jsonFile, err := os.Open(fn)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer jsonFile.Close()
		var schema map[string]Model
		if err := json.NewDecoder(jsonFile).Decode(&schema); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		initFile := jen.NewFile(outdir)
		initFile.PackageComment("Code generated. DO NOT EDIT.")
		initFile.Line()

		{
			initFile.Type().Id("Meta").Struct(
				jen.Id("UserID").Op("*").String().Tag(map[string]string{"json": "userId,omitempty"}),
				jen.Id("SessionID").Op("*").String().Tag(map[string]string{"json": "sessionId,omitempty"}),
				jen.Id("Version").Op("*").Int64().Tag(map[string]string{"json": "version,omitempty"}),
			)
		}

		ifn := path.Join(outdir, "init.go")
		if err := initFile.Save(ifn); err != nil {
			fmt.Printf("Error generating file: %s", ifn)
			fmt.Println(err)
			os.Exit(1)
		}

		for name, model := range schema {
			f := jen.NewFile(outdir)
			f.PackageComment("Code generated. DO NOT EDIT.")
			f.Line()

			if model.Description != "" {
				f.Comment(model.Description)
			}

			keys := make([]string, 0)
			for key := range model.Properties {
				keys = append(keys, key)
			}

			sort.SliceStable(keys, func(i, j int) bool {
				if keys[i] == keys[j] {
					return false
				}
				// make the primary key first
				if keys[i] == "id" {
					return true
				}
				return keys[i] < keys[j]
			})

			fields := make([]jen.Code, 0)

			for _, key := range keys {
				prop := model.Properties[key]
				tags := map[string]string{"json": key}
				field := toType(f, name, jen.Id(identifier(key)), key, prop, tags)
				if prop.Nullable == nil {
					gorm := tags["gorm"]
					if gorm == "" {
						gorm = "not null"
					} else {
						gorm += ";not null"
					}
					tags["gorm"] = gorm
				}
				field = field.Tag(tags)
				if prop.Description != "" {
					field = field.Comment(prop.Description)
				}
				fields = append(fields, field)
			}

			f.Type().Id(name).Struct(
				fields...,
			)

			f.Comment("TableName returns the name of the table for this model which GORM will use when using this model")
			f.Func().Params(
				jen.Id("m").Op("*").Id(name),
			).Id("TableName").Params().String().Block(
				jen.Return(jen.Lit(model.DBName)),
			)

			f.Line()

			f.Func().Params(
				jen.Id("m").Op("*").Id(name),
			).Id("String").Params().String().Block(
				jen.Id("buf").Op(",").Id("_").Op(":=").Qual("encoding/json", "Marshal").Params(jen.Id("m")),
				jen.Return(jen.String().Params(jen.Id("buf"))),
			)

			fn := path.Join(outdir, name+".go")
			if err := f.Save(fn); err != nil {
				fmt.Printf("Error generating file: %s for model: %s", fn, name)
				fmt.Println(err)
				os.Exit(1)
			}
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
