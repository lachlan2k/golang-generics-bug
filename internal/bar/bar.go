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

type ConcreteStruct struct {
	template[struct{}, struct {
		JsonField string `json:"json_field"`
	}]
}

type BarInterface interface {
	UnmarshalBroken(jsonStr string)
	UnmarshalWorking(jsonStr string)
}

func (p *template[TypeA, TypeB]) UnmarshalBroken(jsonStr string) {
    err := json.Unmarshal([]byte(jsonStr), &p.B)
    fmt.Printf("Is there an error? %v\n", err)

    // This crashes:
    fmt.Printf("p.B: %v", p.B)

    return
}

func (p *template[TypeA, TypeB]) UnmarshalWorking(jsonStr string) {
	var result TypeB

    err := json.Unmarshal([]byte(jsonStr), &result)
    fmt.Printf("Is there an error? %v\n", err)

	p.B = result

    // This works fine
    fmt.Printf("p.B: %v", p.B)

    return
}