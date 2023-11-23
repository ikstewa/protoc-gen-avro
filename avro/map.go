package avro

import (
  "fmt"
  "github.com/iancoleman/orderedmap"
)

type Map struct {
  Name string
  Namespace string
  Values Type
}

func (t Map) GetName() string {
	return t.Name
}

func (t Map) GetNamespace() string {
	return t.Namespace
}

func (t Map) ToJSON(types *TypeRepo) (any, error) {
  valueJson, err := t.Values.ToJSON(types)
  if err != nil {
    return nil, fmt.Errorf("error parsing map value type: %w", err)
  }
  jsonMap := orderedmap.New()
  jsonMap.Set("type", "map")
  jsonMap.Set("values", valueJson)
  return jsonMap, nil
}
