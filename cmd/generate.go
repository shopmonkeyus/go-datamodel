package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"sort"
	"strings"

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
	PrimaryKeys []string            `json:"primaryKeys"`
}

func stringContains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func stringFind(a []string, x string) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return -1
}

func ensurePosition(a []string, element string, position int) int {
	index := stringFind(a, element)
	if index < 0 {
		return position - 1
	}
	if index > 0 && index != position {
		cur := a[position]
		a[index] = cur
		a[position] = element
	}
	return position
}

func identifier(val string) string {
	switch val {
	case "id":
		return "ID"
	case "url":
		return "URL"
	}
	newval := strcase.ToCamel(val)
	if strings.HasSuffix(newval, "Id") {
		return newval[0:len(newval)-2] + "ID"
	}
	return newval
}

func IsPrimaryKeyField(model Model, name string) bool {
	for _, pk := range model.PrimaryKeys {
		if pk == name {
			return true
		}
	}
	return false
}

func toType(file *jen.File, model Model, modelName string, stmt *jen.Statement, name string, property Property, tags map[string]string) *jen.Statement {
	_stmt := stmt

	var isNullable bool

	if property.Nullable != nil && *property.Nullable {
		_stmt = _stmt.Op("*")
		isNullable = true
	}

	switch property.Type {
	case "string":
		if property.Format != nil && *property.Format == "date-time" {
			_stmt = _stmt.Qual("github.com/shopmonkeyus/go-datamodel/datatypes", "DateTime")
			switch name {
			case "createdDate", "updatedDate":
				tags["gorm"] = "column:" + name
			default:
			}
			return _stmt
		} else if IsPrimaryKeyField(model, name) {
			tags["gorm"] = "primaryKey"
		}
		if property.Enum != nil {
			enumPrefix := modelName + strcase.ToCamel(name)
			enumName := enumPrefix + "Enum"
			file.Type().Id(enumName).String()
			file.Line()
			vals := make([]jen.Code, 0)
			for _, val := range *property.Enum {
				vals = append(vals, jen.Id(enumPrefix+val).Op(enumName).Op("=").Lit(val))
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
				if isNullable {
					return stmt.Qual("github.com/shopmonkeyus/go-datamodel/datatypes", "NullableStringArray")
				}
				return stmt.Qual("github.com/shopmonkeyus/go-datamodel/datatypes", "StringArray")
			case "number":
				if isNullable {
					return stmt.Qual("github.com/shopmonkeyus/go-datamodel/datatypes", "NullableNumberArray")
				}
				return stmt.Qual("github.com/shopmonkeyus/go-datamodel/datatypes", "NumberArray")
			}
		}
	}
	if name == "meta" || name == "metadata" {
		tags["gorm"] = "column:" + name
		tags["json"] = name + ",omitempty"
		return _stmt.Op("*").Qual("github.com/shopmonkeyus/go-datamodel/datatypes", "JSON")
	}
	return stmt.Qual("github.com/shopmonkeyus/go-datamodel/datatypes", "JSON")
}

var generateCmd = &cobra.Command{
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

			initFile.Type().Id("Model").Interface(
				jen.Id("TableName").Params().String(),
				jen.Id("PrimaryKeys").Params().Op("[]").String(),
				jen.Id("String").Params().String(),
			)

			initFile.Line()

			params := make([]jen.Code, 0)
			params2 := make([]jen.Code, 0)
			cases := make([]jen.Code, 0)
			cases2 := make([]jen.Code, 0)
			models := make([]string, 0)

			for name := range schema {
				models = append(models, name)
			}
			sort.Strings(models)

			initFile.Line()

			initFile.Comment("// ModelNames is an array of all the models defined in this package")
			namesf := initFile.Var().Id("ModelNames").Op("=").Op("[]string{\n")
			initFile.Line()

			initFile.Comment("// TableNames is an array of all the tables defined in this package")
			namest := initFile.Var().Id("TableNames").Op("=").Op("[]string{\n")
			initFile.Line()

			initFile.Comment("// ModelInstances is an array of an empty object for each model in this package")
			intf := initFile.Var().Id("ModelInstances").Op("=").Op("[]interface{}{\n")

			for _, name := range models {
				prop := schema[name]
				namesf.Op(`"` + name + `",` + "\n")
				namest.Op(`"` + prop.DBName + `",` + "\n")
				intf.Op("&" + name + "{},\n")
			}
			namesf.Op("}")
			initFile.Line()
			namest.Op("}")
			initFile.Line()
			intf.Op("}")

			initFile.Line()

			for _, name := range models {
				prop := schema[name]
				rt := jen.Return(jen.Op("New" + name).Params(jen.Op("buf")))
				cases = append(cases, jen.Op("case").Lit(name).Op(",").Lit(prop.DBName).Op(":").Block(rt))

				rt2 := jen.Return(jen.Op("New"+name+"FromChangeEvent").Params(jen.Op("buf"), jen.Op("gzip")))
				cases2 = append(cases2, jen.Op("case").Lit(name).Op(",").Lit(prop.DBName).Op(":").Block(rt2))
			}

			params = append(params, jen.Op("switch").Id("name").Block(cases...))
			params = append(params, jen.Line())
			params = append(params, jen.Return(jen.Op("nil,").Qual("errors", "New").Params(jen.Lit("invalid model: ").Op("+").Op("name"))))

			params2 = append(params2, jen.Op("switch").Id("name").Block(cases2...))
			params2 = append(params2, jen.Line())
			params2 = append(params2, jen.Return(jen.Op("nil,").Qual("errors", "New").Params(jen.Lit("invalid model: ").Op("+").Op("name"))))

			initFile.Comment("NewFromModel will return a model from a model name and buffer encoded as EncodingType")
			initFile.Func().Id("NewFromModel").Params(jen.Id("name").String().Op(",").Id("buf").Op("[]").Byte()).Op("(Model, error)").Block(params...)
			initFile.Line()

			initFile.Comment("NewFromChangeEvent will return change event for model from a buffer encoded as EncodingType")
			initFile.Func().Id("NewFromChangeEvent").Params(jen.Id("name").String().Op(",").Id("buf").Op("[]").Byte(), jen.Id("gzip").Op("bool")).Op("(any, error)").Block(params2...)
			initFile.Line()
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

			var pos int

			// re-order these common fields after sorting the remaining
			// columns by alphabetical
			// don't sort if we have multiple primary keys since the field order is important
			if len(model.PrimaryKeys) == 1 && model.PrimaryKeys[0] == "id" {
				pos = ensurePosition(keys, "id", 0)
				pos = ensurePosition(keys, "createdDate", pos+1)
				pos = ensurePosition(keys, "updatedDate", pos+1)
				pos = ensurePosition(keys, "meta", pos+1)
				pos = ensurePosition(keys, "metadata", pos+1)
				pos = ensurePosition(keys, "companyId", pos+1)
				pos = ensurePosition(keys, "locationId", pos+1)
				pos = ensurePosition(keys, "customFields", pos+1)

				// figure out all the keys not included in the special list above
				oldkeys := make([]string, 0)
				for _, k := range keys {
					switch k {
					case "id", "createdDate", "updatedDate", "meta", "metadata", "companyId", "locationId", "customFields":
						continue
					default:
						oldkeys = append(oldkeys, k)
					}
				}

				// sort the remainder of the keys
				sort.SliceStable(oldkeys, func(i, j int) bool {
					return oldkeys[i] < oldkeys[j]
				})

				// add the old keys after our last special keys
				for i, k := range oldkeys {
					newindex := pos + i + 1
					keys[newindex] = k
				}
			}

			fields := make([]jen.Code, 0)

			for i, key := range keys {
				prop := model.Properties[key]
				tags := map[string]string{"json": key}
				field := toType(f, model, name, jen.Id(identifier(key)), key, prop, tags)
				required := stringContains(model.Required, key)
				gorm := tags["gorm"]
				if required && (prop.Nullable == nil || !*prop.Nullable) {
					if gorm == "" {
						gorm = "not null"
					} else {
						gorm += ";not null"
					}
					tags["gorm"] = gorm
				}
				if gorm != "" {
					gorm += ";"
				}
				gorm += "column:" + key
				tags["gorm"] = gorm
				field = field.Tag(tags)
				if prop.Description != "" {
					field = field.Comment(prop.Description)
				}
				if pos == i {
					field = field.Line() // put a line break after our last special fields
				}
				fields = append(fields, field)
			}

			f.Type().Id(name).Struct(
				fields...,
			)

			f.Var().Op("_").Id("Model").Op("=").Params(jen.Op("*" + name)).Params(jen.Nil())

			f.Comment("TableName returns the name of the table for this model which GORM will use when using this model")
			f.Func().Params(
				jen.Id("m").Op("*").Id(name),
			).Id("TableName").Params().String().Block(
				jen.Return(jen.Lit(model.DBName)),
			)
			f.Line()

			pks := make([]string, 0)
			for _, pk := range model.PrimaryKeys {
				pks = append(pks, `"`+pk+`"`)
			}

			f.Comment("PrimaryKeys returns an array of the primary keys for this model")
			f.Func().Params(
				jen.Id("m").Op("*").Id(name),
			).Id("PrimaryKeys").Params().Op("[]").String().Block(
				jen.Return(jen.Op("[]string{" + strings.Join(pks, ",") + "}")),
			)
			f.Line()

			f.Comment("String returns a string representation as JSON for this model")
			f.Func().Params(
				jen.Id("m").Op("*").Id(name),
			).Id("String").Params().String().Block(
				jen.Id("buf").Op(",").Id("_").Op(":=").Qual("encoding/json", "Marshal").Params(jen.Id("m")),
				jen.Return(jen.String().Params(jen.Id("buf"))),
			)
			f.Line()

			f.Comment("New" + name + " returns a new model instance from an encoded buffer")
			f.Func().Id("New"+name).Params(
				jen.Id("buf").Op("[]").Byte(),
			).Op("(*"+name+", error)").Block(
				jen.Var().Id("result").Op(name),
				jen.Id("err").Op(":=").Qual("encoding/json", "Unmarshal").Params(jen.Op("buf"), jen.Op("&result")),
				jen.Op("if err != nil").Block(jen.Return(jen.Op("nil, err"))),
				jen.Return(jen.Op("&result, nil")),
			)
			f.Line()

			f.Comment("New" + name + "FromChangeEvent returns a new model instance from an encoded buffer as change event")
			f.Func().Id("New"+name+"FromChangeEvent").Params(
				jen.Id("buf").Op("[]").Byte(),
				jen.Id("gzip").Bool(),
			).Op("(*").Qual("github.com/shopmonkeyus/go-datamodel/datatypes", "ChangeEvent["+name+"]").Op(", error)").Block(
				jen.Var().Id("result").Qual("github.com/shopmonkeyus/go-datamodel/datatypes", "ChangeEvent["+name+"]"),
				jen.Var().Id("decompressed").Op("=").Op("buf"),
				jen.If(jen.Op("gzip")).Block(
					jen.Id("dec").Op(",").Id("err").Op(":=").Qual("github.com/shopmonkeyus/go-datamodel/datatypes", "Gunzip").Params(jen.Op("buf")),
					jen.Op("if err != nil").Block(jen.Return(jen.Op("nil, err"))),
					jen.Op("decompressed = dec"),
				),
				jen.Id("err").Op(":=").Qual("encoding/json", "Unmarshal").Params(jen.Op("decompressed"), jen.Op("&result")),
				jen.Op("if err != nil").Block(jen.Return(jen.Op("nil, err"))),
				jen.Return(jen.Op("&result, nil")),
			)
			f.Line()

			fn := path.Join(outdir, name+".go")
			if err := f.Save(fn); err != nil {
				fmt.Printf("Error generating file: %s for model: %s", fn, name)
				fmt.Println(err)
				os.Exit(1)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
