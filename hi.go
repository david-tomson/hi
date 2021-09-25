package hi

import (
	"fmt"

	"github.com/hokaccha/go-prettyjson"
)

const errMarshalError = "Failed to convert object to JSON"

func Print(v interface{}) {
	value, ok := v.([]byte)
	if ok {
		fmt.Println(string(value))
		return
	}

	output, err := prettyjson.Marshal(v)

	if err != nil {
		fmt.Println(errMarshalError)
	}

	fmt.Println(string(output))
}
