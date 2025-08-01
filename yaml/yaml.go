package yaml

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v2"
)

// String -> Yaml
var data = `
a: Easy!
b:
  c: 2
  d: [3, 4]
  connectionStatus: true
`

// Typed | Yaml -> String ?
type T struct {
	A string
	B struct {
		RenamedC int   `yaml:"c"`
		D        []int `yaml:",flow"`
		F        bool  `yaml:"connectionStatus"`
	}
}

// Runner Yaml marshalling / unmarshalling example
// Ref: https://pkg.go.dev/gopkg.in/yaml.v2#section-readme
func Runner() {
	t := T{}

	err := yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t:\n%v\n\n", t)

	d, err := yaml.Marshal(&t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t dump:\n%s\n\n", string(d))

	m := make(map[interface{}]interface{})

	err = yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m:\n%v\n\n", m)

	d, err = yaml.Marshal(&m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m dump:\n%s\n\n", string(d))
}
