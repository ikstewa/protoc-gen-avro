package avro

import "fmt"

type Union struct {
  Types []Type
}

func (u Union) ToJSON(types *TypeRepo) (any, error) {
  flattened := u.flatten()
  jsonSlice := make([]any, len(flattened))
  for i, unionType := range flattened {
    typeJson, err := unionType.ToJSON(types)
    if err != nil {
      return nil, fmt.Errorf("error parsing union type: %w", err)
    }
    jsonSlice[i] = typeJson
  }
  return jsonSlice, nil
}

func (u Union) flatten() []Type {
  var flattened []Type
  for _, unionType := range u.Types {
    switch unionType.(type) {
    case Union:
      flattened = append(flattened, unionType.(Union).flatten()...)
    default:
      flattened = append(flattened, unionType)
    }
  }
  return flattened
}
