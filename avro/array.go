package avro

import (
  "fmt"
  "github.com/iancoleman/orderedmap"
)

type Array struct {
  Items Type
}

func (t Array) ToJSON(types *TypeRepo) (any, error) {
  itemJson, err := t.Items.ToJSON(types)
  if err != nil {
    return nil, fmt.Errorf("error parsing item type: %w", err)
  }
  jsonMap := orderedmap.New()
  jsonMap.Set("type", "array")
  jsonMap.Set("items", itemJson)
  return jsonMap, nil
}
