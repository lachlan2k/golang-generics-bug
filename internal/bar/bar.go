package bar

import (
	"encoding/json"
	"fmt"
)

// Only seems to replicate when I use the 2 different fields, with the data in the latter field.
type template[TypeA any, TypeB any] struct {
	A TypeA
	B TypeB
}

type Specialisation struct {
	template[struct{}, struct {
		JsonField string `json:"json_field"`
	}]
}

type BarInterface interface {
	UnmarshalJsonIntoB(reqJsonStr string)
}

func (p *template[ReqT, ResT]) UnmarshalJsonIntoB(jsonStr string) {
    err := json.Unmarshal([]byte(jsonStr), &p.B)
    fmt.Printf("Is there an error? %v\n", err)

    // This crashes:
    fmt.Printf("p.B: %v", p.B)

    return
}
