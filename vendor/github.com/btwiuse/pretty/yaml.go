package pretty

import "fmt"
import "github.com/ghodss/yaml"

func YAMLString(v interface{}) string {
	b, err := yaml.Marshal(v)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%s", b)
}

func YAML(v interface{}) {
	b, err := yaml.Marshal(v)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", b)
}

func YamlString(v interface{}) string {
	b, err := yaml.Marshal(v)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%s", b)
}

func Yaml(v interface{}) {
	b, err := yaml.Marshal(v)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", b)
}
