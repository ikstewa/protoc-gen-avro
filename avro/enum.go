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

// https://golangbyexample.com/longest-common-prefix-golang/
func longestCommonPrefix(strs []string) string {
	lenStrs := len(strs)

	if lenStrs == 0 {
		return ""
	}

	firstString := strs[0]

	lenFirstString := len(firstString)

	commonPrefix := ""
	for i := 0; i < lenFirstString; i++ {
		firstStringChar := string(firstString[i])
		match := true
		for j := 1; j < lenStrs; j++ {
			if (len(strs[j]) - 1) < i {
				match = false
				break
			}

			if string(strs[j][i]) != firstStringChar {
				match = false
				break
			}
		}

		if match {
			commonPrefix += firstStringChar
		} else {
			break
		}
	}

	return commonPrefix
}

func (t Enum) ToJSON(repo *TypeRepo) (any, error) {
  jsonMap := orderedmap.New()
  jsonMap.Set("type", "enum")
  jsonMap.Set("name", t.Name)
	symbols := t.Symbols
	defaultVal := t.Symbols[0]
  if repo.RemoveEnumPrefixes {
    prefix := longestCommonPrefix(t.Symbols)
    for i, symbol := range t.Symbols {
      symbols[i] = symbol[len(prefix):]
    }
		defaultVal = defaultVal[len(prefix):]
  }
	jsonMap.Set("symbols", symbols)
  jsonMap.Set("default", defaultVal)
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
