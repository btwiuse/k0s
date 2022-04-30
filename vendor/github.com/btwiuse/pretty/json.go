package pretty

import "fmt"
import "encoding/json"

func JSON(v interface{}) {
	fmt.Print(jStr(v))
}

func JSONLine(v interface{}) {
	fmt.Println(jStr(v))
}

func JSONString(v interface{}) string {
	return jStr(v)
}

func JSONStringLine(v interface{}) string {
	return fmt.Sprintf("%s\n", jStr(v))
}

func Json(v interface{}) {
	fmt.Print("%s\n", jstr(v))
}

func JsonLine(v interface{}) {
	fmt.Println("%s\n", jstr(v))
}

func JsonString(v interface{}) string {
	return jstr(v)
}

func JsonStringLine(v interface{}) string {
	return fmt.Sprintf("%s\n", jstr(v))
}

func jStr(v interface{}) string {
	b, err := json.MarshalIndent(v, "", "  ")

	if err != nil {
		panic(err)
	}

	return string(b)
}

func jstr(v interface{}) string {
	var (
		b   []byte
		err error
	)

	switch v.(type) {
	case *json.RawMessage:
		b, err = v.(*json.RawMessage).MarshalJSON()
	case json.RawMessage:
		b, err = v.(json.RawMessage).MarshalJSON()
	default:
		b, err = json.Marshal(v)
	}

	if err != nil {
		panic(err)
	}

	return string(b)
}
