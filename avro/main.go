package main

import (
	"fmt"
	"reflect"
	"time"

	"encoding/json"
	"github.com/linkedin/goavro"
	"github.com/mitchellh/mapstructure"

	// "git.chotot.org/ad-service/spine/apps/types"
	"git.chotot.org/badboyd/structs"
	"git.chotot.org/data/kstream/utils"
	"git.chotot.org/kelvin/navro"
	// "git.chotot.org/ad-service/spine/pkg/types/ad"
)

var (
	typeMap = map[reflect.Kind]string{
		reflect.Uint8:   "int",
		reflect.Int8:    "int",
		reflect.Int32:   "int",
		reflect.Int16:   "int",
		reflect.Int:     "long",
		reflect.Int64:   "long",
		reflect.String:  "string",
		reflect.Float32: "float",
		reflect.Float64: "double",
		reflect.Map:     "map",
		reflect.Bool:    "boolean",
		reflect.Array:   "array",
		reflect.Slice:   "array",
	}
)

type astruct struct {
	I         *int64 `json:"i,omitempty"`
	String    string `json:"string,omitempty"`
	StringPtr *string
	InStruct  *InnerStruct
}

type arrayStruct struct {
	Structs  *[]*astruct
	StructND *astruct
}

type InnerStruct struct {
	Number      *int
	InnerStruct *InnerOfInnerStruct
}

type InnerOfInnerStruct struct {
	INumberPtr *int
	INumber    []int
}

type arrayOfBytesPtr struct {
	Bytes *[]byte
}

type StructHasIntPointer struct {
	IntPtr *int
	Map    map[string]astruct
}

func main() {
	//
	numberThirteen64 := int64(13)
	numberThirteen := 13
	numberThirteenString := "13"
	// tmp, err := convertToAvroMapV2(numberThirteenString)
	// fmt.Println(tmp, err)
	nubmerThriteenPtr := &numberThirteenString
	// tmp, err = convertToAvroMapV2(nubmerThriteenPtr)
	// fmt.Println(tmp, err)
	//
	// str := InnerStruct{
	// 	Number: &numberThirteen,
	// 	InnerStruct: InnerOfInnerStruct{
	// 		INumber: 13,
	// 	},
	// }
	// tmp, err = convertToAvroMapV2(str)
	// fmt.Println(tmp, err)
	aStruct := astruct{
		I:         &numberThirteen64,
		String:    "string",
		StringPtr: &numberThirteenString,
		// InStruct:  nil,
	}
	normalizeAvroMap := genMapV2(aStruct)

	tmp := astruct{}
	mapcodec, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		Result:           &tmp,
	})
	mapcodec.Decode(normalizeAvroMap)
	fmt.Printf("map decode: %+v\n", *tmp.I)
	fmt.Printf("map decode: %+v\n", tmp.String)
	fmt.Printf("map decode: %+v\n", *tmp.StringPtr)
	fmt.Printf("map decode: %+v\n", tmp.InStruct)

	deeper := InnerOfInnerStruct{
		INumber:    []int{numberThirteen, 12, 13},
		INumberPtr: &numberThirteen,
	}

	deep := InnerStruct{
		Number:      nil,
		InnerStruct: &deeper,
	}

	genMapV2(deep)
	aStructNo2 := astruct{
		I:         &numberThirteen64,
		String:    "string",
		StringPtr: &numberThirteenString,
		InStruct:  &deep,
	}
	arrayStructInstance := arrayStruct{
		Structs:  &[]*astruct{&aStruct, &aStructNo2},
		StructND: &aStructNo2,
	}
	normalizeAvroMap = genMapV2(arrayStructInstance)
	arrayStructInstance2nd := arrayStruct{}
	mapcodec, _ = mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		Result:           &arrayStructInstance2nd,
	})
	mapcodec.Decode(normalizeAvroMap)
	b, _ := json.Marshal(arrayStructInstance2nd)
	fmt.Printf("map decode: %+v\n", string(b))
	// genMapV2(aStructNo2)
	var numberThirteenStringNil *string
	genMapV2(numberThirteen)
	genMapV2(numberThirteenString)
	genMapV2(numberThirteenStringNil)
	genMapV2(nubmerThriteenPtr)
	// genMapV2(deeper)
	// genMapV2(deep)
	//
	// bytes := []byte{1, 2, 3}
	// genMapV2(arrayOfBytesPtr{
	// 	Bytes: &bytes,
	// })
	// genMapV2(types.LoadActionResponse{})
	// genMapV2(arrayStructInstance)
	//
	// // schema, err := navro.GenerateV2(arrayStruct{})
	// // fmt.Println(schema, err)
	// var nilInt *int
	// genMapV2(nilInt)
	// genMapV2(StructHasIntPointer{
	// 	IntPtr: &numberThirteen,
	// 	Map:    nil},
	// )
}

func genMapV2(i interface{}) interface{} {
	schema, names, err := navro.GenerateV2(i)
	if err != nil {
		panic(err)
	}
	fmt.Println(schema)
	fmt.Println("record names: ", names)
	fmt.Println("record names: ", GetMapOfRecordNames(schema))

	codec, err := goavro.NewCodec(schema)
	if err != nil {
		panic(err)
	}

	iMap, err := utils.ConvertToAvroDatum(i)
	if err != nil {
		panic(err)
	}
	fmt.Println(iMap)

	iBytes, err := codec.BinaryFromNative(nil, iMap)
	if err != nil {
		panic(err)
	}

	originMap, _, err := codec.NativeFromBinary(iBytes)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Decode: %+v\n", originMap)

	normalizedDatum, err := normalizeAvroDatum(originMap, names)
	fmt.Printf("Normalized: %+v \n Error: %+v\n", normalizedDatum, err)
	return normalizedDatum
}

func genMap(i interface{}) {
	schema, recordNames, err := navro.GenerateV2(i)
	if err != nil {
		panic(err)
	}
	fmt.Println(recordNames)
	codec, err := goavro.NewCodec(schema)
	if err != nil {
		panic(err)
	}

	aStructMap := structs.Map(i)
	fmt.Printf("Before: %+v\n", aStructMap)
	convertToAvroMap(aStructMap)
	fmt.Printf("After: %+v\n", aStructMap)

	aStructMapInBytes, err := codec.BinaryFromNative(nil, aStructMap)
	if err != nil {
		panic(err)
	}

	originMap, _, err := codec.NativeFromBinary(aStructMapInBytes)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Decode: %+v\n", originMap)
}

func convertDaytimeToEpoch(data map[string]interface{}) (int64, error) {
	if data == nil {
		return 0, fmt.Errorf("NIL_ARG")
	}

	unixTime := time.Date(int(data["year"].(int32)),
		time.Month(data["month"].(int32)),
		int(data["day"].(int32)),
		int(data["hour"].(int32)),
		int(data["minute"].(int32)),
		int(data["second"].(int32)),
		int(data["micro"].(int32)),
		time.Local).Unix()
	return unixTime, nil
}

func normalizeAvroDatum(i interface{}, types map[string]bool) (interface{}, error) {
	if i == nil {
		return nil, nil
	}

	iTyp := reflect.TypeOf(i)
	iVal := reflect.ValueOf(i)

	if reflect.ValueOf(nil) == iVal || !iVal.IsValid() {
		return nil, nil
	}

	switch iTyp.Kind() {
	case reflect.Map:
		normalizedMap := make(map[string]interface{})
		for _, key := range iVal.MapKeys() {
			val := iVal.MapIndex(key).Interface()
			if key.String() == "DateTime" {
				return convertDaytimeToEpoch(val.(map[string]interface{}))
			}

			elem, err := normalizeAvroDatum(val, types)
			if err != nil {
				return nil, err
			}
			elemTyp := reflect.TypeOf(elem)
			// union type for primitive types
			if elem != nil && key.String() == getAvroName(elemTyp) {
				return elem, nil
			}
			// union of record type
			if types[key.String()] {
				return elem, nil
			}
			normalizedMap[key.String()] = elem
		}
		return normalizedMap, nil
	case reflect.Slice, reflect.Array:
		normalizedArray := make([]interface{}, 0, iVal.Len())
		for i := 0; i < iVal.Len(); i++ {
			elem, err := normalizeAvroDatum(iVal.Index(i).Interface(), types)
			if err != nil {
				return nil, err
			}
			normalizedArray = append(normalizedArray, elem)
		}
		return normalizedArray, nil
	default:
		return i, nil
	}
}

func convertToAvroDatum(i interface{}) (interface{}, error) {
	if i == nil {
		return nil, nil
	}

	return getAvroDatumByType(reflect.TypeOf(i), reflect.ValueOf(i))
}

func getAvroName(t reflect.Type) string {
	switch t.Kind() {
	case reflect.Invalid:
		return "null"
	case reflect.Bool:
		return "boolean"
	case reflect.Uint8, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Uint16: // <= 32 bits will be mapped to Int avro
		return "int"
	case reflect.Int, reflect.Uint32, reflect.Int64, reflect.Uint64: // cause in Go, int could be 32 or 64 bit
		return "long"
	case reflect.String:
		return "string"
	case reflect.Float32:
		return "float"
	case reflect.Float64:
		return "double"
	case reflect.Array, reflect.Slice:
		if t.Elem().Kind() == reflect.Uint8 {
			return "bytes"
		}
		return "array"
	case reflect.Map:
		if t.Key().Kind() == reflect.String {
			return "map"
		}
		return ""
	default:
		return t.Name()
	}
}

func getAvroDatumByType(t reflect.Type, v reflect.Value) (interface{}, error) {
	if zeroVal := reflect.Zero(t); zeroVal == v || !v.IsValid() {
		return nil, nil
	}

	switch t.Kind() {
	case reflect.Invalid:
		return nil, nil
	case reflect.Ptr:
		elem, err := getAvroDatumByType(t.Elem(), v.Elem())
		if err != nil {
			return nil, err
		}
		if elem == nil {
			return nil, nil
		}
		return map[string]interface{}{getAvroName(t.Elem()): elem}, nil
	case reflect.Struct:
		record := make(map[string]interface{})
		for i := 0; i < t.NumField(); i++ {
			fieldTyp := t.Field(i)
			fieldVal := v.Field(i)

			fieldMap, err := getAvroDatumByType(fieldTyp.Type, fieldVal)
			if err != nil {
				return nil, err
			}
			record[fieldTyp.Name] = fieldMap
		}
		return record, nil
	case reflect.Array, reflect.Slice:
		if t.Elem().Kind() == reflect.Uint8 {
			elems := make([]byte, 0, v.Len())
			for i := 0; i < v.Len(); i++ {
				elem := v.Index(i).Interface().(byte)
				elems = append(elems, elem)
			}
			return elems, nil
		}
		elems := make([]interface{}, 0, v.Len())
		for i := 0; i < v.Len(); i++ {
			elem, err := getAvroDatumByType(v.Index(i).Type(), v.Index(i))
			if err != nil {
				return nil, err
			}
			elems = append(elems, elem)
		}
		return elems, nil
	case reflect.Map:
		elems := make(map[string]interface{})
		for _, key := range v.MapKeys() {
			elem, err := getAvroDatumByType(t.Elem(), v.MapIndex(key))
			if err != nil {
				return nil, err
			}
			elems[key.String()] = elem
		}
		return elems, nil
	default:
		return v.Interface(), nil
	}
}

func convertToAvroMap(in map[string]interface{}) {
	for k, v := range in {
		if v == nil {
			continue
		}

		typOfv := reflect.TypeOf(v)
		valOfv := reflect.ValueOf(v)

		fmt.Printf("Key of : %+v\n", k)
		fmt.Printf("Type of Ptr: %+v\n", typOfv.Kind().String())
		fmt.Printf("Val of Ptr: %+v\n", valOfv)

		switch typOfv.Kind() {
		case reflect.Ptr:
			if valOfv.IsNil() {
				in[k] = nil
			} else {
				elemVal := valOfv.Elem().Interface()
				elemType := typeMap[typOfv.Elem().Kind()]
				in[k] = goavro.Union(elemType, elemVal)
			}
		case reflect.Map:
			// fmt.Printf("Type of: %+v\n", typOfv)
			convertToAvroMap(v.(map[string]interface{}))
		case reflect.Array, reflect.Slice:
			for i := 0; i < valOfv.Len(); i++ {
				convertToAvroMap(valOfv.Index(i).Interface().(map[string]interface{}))
			}
		default:
			in[k] = v
		}
	}
}

// GetMapOfRecordNames returns set of record names from input schema
func GetMapOfRecordNames(schema string) map[string]bool {
	recordNameSet := make(map[string]bool)
	schemaMap := make(map[string]interface{})
	if err := json.Unmarshal([]byte(schema), &schemaMap); err != nil {
		return recordNameSet
	}

	getRecordNameFromMap(schemaMap, recordNameSet)
	return recordNameSet
}

func getRecordNameFromMap(schemaMap map[string]interface{}, recordNameSet map[string]bool) {
	if schemaMap["type"] == nil {
		return
	}

	schemaType := reflect.TypeOf(schemaMap["type"])
	switch schemaType.Kind() {
	case reflect.String:
		if schemaMap["type"].(string) != "record" {
			return
		}
		recordNameSet[schemaMap["name"].(string)] = true
		if schemaMap["fields"] == nil {
			return
		}

		for _, field := range schemaMap["fields"].([]interface{}) {
			if fieldMap, ok := field.(map[string]interface{}); ok {
				getRecordNameFromMap(fieldMap, recordNameSet)
			}
		}
	case reflect.Array, reflect.Slice:
		for _, v := range schemaMap["type"].([]interface{}) {
			if tmp, ok := v.(map[string]interface{}); ok {
				getRecordNameFromMap(tmp, recordNameSet)
			}
		}
	}
}
