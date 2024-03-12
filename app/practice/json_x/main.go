package json_x

import (
	"encoding/json"
	"fmt"
)

type address struct {
	City string
	Country string
}

type vcard struct {
	Name string
	Addr []*address
}

func Run1() {
	addr1 := &address{"beijing", "china"}
	addr2 := &address{"tianjing", "china"}

	myVcard := vcard{"liuhao", []*address{addr1, addr2}}

	js, _ := json.Marshal(myVcard)

	fmt.Printf("Json: %s\n", js)

	b := []byte(`{"Name": "Wednesday", "Age": 6, "Parents": ["Gomez", "Morticia"]}`)
	var f interface{}
	err := json.Unmarshal(b, &f)
	m := f.(map[string]interface{})
	fmt.Println(m)
	if err == nil {
		for k, v := range m {
			switch vv := v.(type) {
			case string:
				fmt.Println(k, "is string", vv)
			case int:
				fmt.Println(k, "is int", vv)
			case []interface{}:
				fmt.Println(k, "is array", vv)
			default:
				fmt.Println(k, "unknown type", vv)
			}
		}
	}

}
