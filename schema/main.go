package main

import (
	"fmt"
	"reflect"

	"github.com/landoop/schema-registry"
	"github.com/linkedin/goavro"
	avro "gopkg.in/avro.v0"

	"git.chotot.org/kelvin/navro"

	"git.chotot.org/ad-service/spine/apps/types"
	"git.chotot.org/ad-service/spine/pkg/types/ad"
)

var typeMap = map[reflect.Kind]string{
	reflect.Uint8:   "int",
	reflect.Int:     "int",
	reflect.Int8:    "int",
	reflect.Int32:   "int",
	reflect.Int16:   "int",
	reflect.Int64:   "long",
	reflect.String:  "string",
	reflect.Array:   "array",
	reflect.Slice:   "array",
	reflect.Float32: "float",
	reflect.Float64: "double",
	reflect.Struct:  "record",
	reflect.Map:     "map",
	reflect.Bool:    "boolean",
}

type astruct struct {
	I      *int   `json:"i,omitempty"`
	String string `json:"string,omitempty"`
	Bytes  []byte
}

type testStruct struct {
	TMP astruct
	T   astruct
}

type testReason struct {
	Name   *string
	Ad     []ad.SimpleAd
	Reason []ad.ReasonData
	Params ad.Params
}

func main() {

	intF := avro.SchemaField{
		Name: "A",
		Type: &avro.IntSchema{},
	}

	stringF := avro.SchemaField{
		Name: "String",
		Type: &avro.StringSchema{},
	}

	unionInt := avro.UnionSchema{
		Types: []avro.Schema{&avro.NullSchema{}, &avro.IntSchema{}},
	}

	unionF := avro.SchemaField{
		Name: "NullableString",
		Type: &unionInt,
	}

	schema := avro.SchemaField{
		Name: "LoadActionResponse",
		Type: &avro.RecordSchema{
			Name:   "LoadActionResponse",
			Fields: []*avro.SchemaField{&intF, &stringF, &unionF},
		},
	}

	schema1 := avro.SchemaField{
		Name: "LoadAd",
		Type: schema.Type,
	}

	biggerSchema := avro.RecordSchema{
		Name:      "BiggerRecord",
		Namespace: "com.chotot",
		Fields:    []*avro.SchemaField{&schema, &schema1},
	}

	fmt.Println(registerNewSchema("bigStruct", biggerSchema.String()))
	fmt.Println(biggerSchema.String())

	fmt.Println()
	testGenSchema()
}

func testGenSchema() {
	// tmp := testStruct{}
	// schema, err := navro.GenerateV2(tmp)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(schema)

	newschema := `{
  "fields": [
    {
      "name": "TMP",
      "type": {
        "fields": [
          {
            "default": null,
            "name": "I",
            "type": [
              "null",
              "long"
            ]
          },
          {
            "name": "String",
            "type": "string"
          },
          {
            "name": "Bytes",
            "type": "bytes"
          }
        ],
        "name": "astruct",
        "type": "record"
      }
    },
    {
      "name": "T",
      "type": "astruct"
    }
  ],
  "name": "testStruct",
  "type": "record"
}`

	byteArr := types.LoadActionResponse{}
	schema, err := navro.GenerateV2(byteArr)

	codec, err := goavro.NewCodec(schema)
	if err != nil {
		panic(err)
	}
	fmt.Println(registerNewSchema("Daa", schema))
	fmt.Println(registerNewSchema("Data", newschema))
	fmt.Println(codec.Schema(), err)
}

func registerNewSchema(name, schema string) error {
	var schemaBaseURL = `http://10.60.3.50:8081`
	client, err := schemaregistry.NewClient(schemaBaseURL)
	if err != nil {
		panic(err)
	}
	id, err := client.RegisterNewSchema(name, schema)
	fmt.Println(id, err)
	return err
}

func generate(i interface{}) (string, error) {
	iTyp := reflect.TypeOf(i)
	iVal := reflect.ValueOf(i)
	schema, err := schemaByType(iTyp, iVal)
	if err != nil {
		return "", err
	}
	return schema.String(), nil
}

func schemaByType(t reflect.Type, v reflect.Value) (avro.Schema, error) {
	fmt.Println(t.Name())
	fmt.Println(v.Type().Name())

	switch t.Kind() {
	case reflect.Invalid:
		return new(avro.NullSchema), nil
	case reflect.Bool:
		return new(avro.BooleanSchema), nil
	case reflect.Uint8, reflect.Int, reflect.Int16, reflect.Int8:
		return new(avro.IntSchema), nil
	case reflect.Int64:
		return new(avro.LongSchema), nil
	case reflect.String:
		return new(avro.StringSchema), nil
	case reflect.Float32:
		return new(avro.FloatSchema), nil
	case reflect.Float64:
		return new(avro.DoubleSchema), nil
	case reflect.Array, reflect.Slice:
		elemSchema, err := schemaByType(t.Elem(), v)
		if err != nil {
			return nil, err
		}
		schema := &avro.ArraySchema{
			Items: elemSchema,
		}
		return schema, nil
	case reflect.Map:
		elemSchema, err := schemaByType(t.Elem(), v)
		if err != nil {
			return nil, err
		}
		schema := &avro.MapSchema{
			Values: elemSchema,
		}
		return schema, nil
	case reflect.Struct:
		schema := &avro.RecordSchema{
			Name:   t.Name(),
			Fields: make([]*avro.SchemaField, 0, t.NumField()),
		}
		for i := 0; i < t.NumField(); i++ {
			structField := t.Field(i)
			fieldTypeSchema, err := schemaByType(structField.Type, v)
			if err != nil {
				return nil, err
			}
			fieldSchema := &avro.SchemaField{
				Name: structField.Name,
				Type: fieldTypeSchema,
			}
			schema.Fields = append(schema.Fields, fieldSchema)
		}
		return schema, nil
	case reflect.Ptr:
		elemSchema, err := schemaByType(t.Elem(), v)
		if err != nil {
			return nil, err
		}
		schema := &avro.UnionSchema{
			Types: []avro.Schema{new(avro.NullSchema), elemSchema},
		}
		return schema, nil
	case reflect.Interface:
	}
	return nil, nil
}
