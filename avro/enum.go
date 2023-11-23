package avro

import (
  "github.com/iancoleman/orderedmap"
  "google.golang.org/protobuf/types/descriptorpb"
)

type Enum struct {
  Name string
  Namespace string
  Symbols []string
  Default string
}

func (t Enum) GetName() string {
  return t.Name
}

func (t Enum) GetNamespace() string {
  return t.Namespace
}

func (t Enum) ToJSON(_ *TypeRepo) (any, error) {
  jsonMap := orderedmap.New()
  jsonMap.Set("type", "enum")
  jsonMap.Set("name", t.Name)
  jsonMap.Set("symbols", t.Symbols)
  jsonMap.Set("default", t.Symbols[0])
  return jsonMap, nil
}

func EnumFromProto(proto *descriptorpb.EnumDescriptorProto, protoPackage string) Enum {
	symbols := make([]string, len(proto.Value))
	for i, value := range proto.Value {
		symbols[i] = *value.Name
	}
  return Enum{
    Name: proto.GetName(),
    Namespace: protoPackage,
    Symbols: symbols,
  }
}
