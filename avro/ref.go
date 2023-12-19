package avro

import "fmt"

// Ref represents a reference to an existing record/enum
type Ref string

func (r Ref) ToJSON(types *TypeRepo) (any, error) {
  switch string(r) {
  case ".google.protobuf.Timestamp":
    return Bare("long").ToJSON(types)
  case ".google.protobuf.Duration":
    return Bare("long").ToJSON(types)
  case ".google.protobuf.Any":
    return Bare("string").ToJSON(types)
  case ".google.protobuf.Struct":
    unionType := Union{Types: []Type{
      Bare("string"),
      Bare("int"),
      Bare("long"),
      Bare("float"),
      Bare("double"),
      Bare("boolean"),
      Bare("bytes"),
    }}

    return Map{Values: unionType}.ToJSON(types)
  case ".google.protobuf.Value":
    return Bare("string").ToJSON(types)
  }
  foundType, err := types.GetType(string(r))
  if err != nil {
    return nil, fmt.Errorf("no type found for %s", r)
  }
  return foundType.ToJSON(types)
}
