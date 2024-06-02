package avro

import (
	"fmt"
	"github.com/iancoleman/orderedmap"
	"google.golang.org/protobuf/types/descriptorpb"
)

type noDefaultType struct{}
var noDefault = noDefaultType{}

type Field struct {
	Name string
	Type Type
	Default any
}

func (t Field) ToJSON(types *TypeRepo) (any, error) {
	typeJson, err := t.Type.ToJSON(types)
	if err != nil {
		return nil, fmt.Errorf("error parsing field type %s %w", t.Name, err)
	}
  jsonMap := orderedmap.New()
  jsonMap.Set("name", t.Name)
  jsonMap.Set("type", typeJson)
	var defaultValue any
  if t.Default != "" {
		defaultValue = t.Default
  } else {
		defaultValue = DefaultValue(typeJson)
	}
	// Avro can't actually handle defaults for records
	if defaultValue != noDefault {
		jsonMap.Set("default", defaultValue)
	}

  return jsonMap, nil
}

func FieldFromProto(proto *descriptorpb.FieldDescriptorProto, typeRepo *TypeRepo) Field {
  var name string
  if typeRepo.JsonFieldnames {
    name = proto.GetJsonName()
  } else {
    name = proto.GetName()
  }
  return Field{
    Name: name,
    Type: FieldTypeFromProto(proto, typeRepo),
    Default: proto.GetDefaultValue(),
  }
}

func FieldTypeFromProto(proto *descriptorpb.FieldDescriptorProto, typeRepo *TypeRepo) Type {
	basicType := BasicFieldTypeFromProto(proto)
	if proto.GetLabel() == descriptorpb.FieldDescriptorProto_LABEL_REPEATED {
		if typeRepo.NullableArrays {
			return Union{Types: []Type{Bare("null"), Array{Items: basicType}}}
		} else {
			return Array{Items: basicType}
		}
	} else if proto.GetProto3Optional() || (proto.OneofIndex != nil && typeRepo.RetainOneofFieldnames) {
		return Union{Types: []Type{Bare("null"), basicType}}
	} else {
		return basicType
	}
}

func BasicFieldTypeFromProto(proto *descriptorpb.FieldDescriptorProto) Type {
	switch proto.GetType() {
	case descriptorpb.FieldDescriptorProto_TYPE_FLOAT:
		return Bare("float")
	case descriptorpb.FieldDescriptorProto_TYPE_DOUBLE:
		return Bare("double")
	case descriptorpb.FieldDescriptorProto_TYPE_INT64:
		return Bare("long")
	case descriptorpb.FieldDescriptorProto_TYPE_UINT64:
		return Bare("long")
	case descriptorpb.FieldDescriptorProto_TYPE_FIXED64:
		return Bare("long")
	case descriptorpb.FieldDescriptorProto_TYPE_SINT64:
		return Bare("long")
	case descriptorpb.FieldDescriptorProto_TYPE_SFIXED64:
		return Bare("long")
	case descriptorpb.FieldDescriptorProto_TYPE_INT32:
		return Bare("int")
	case descriptorpb.FieldDescriptorProto_TYPE_UINT32:
		return Bare("int")
	case descriptorpb.FieldDescriptorProto_TYPE_FIXED32:
		return Bare("int")
	case descriptorpb.FieldDescriptorProto_TYPE_SINT32:
		return Bare("int")
	case descriptorpb.FieldDescriptorProto_TYPE_SFIXED32:
		return Bare("int")
	case descriptorpb.FieldDescriptorProto_TYPE_BOOL:
		return Bare("boolean")
	case descriptorpb.FieldDescriptorProto_TYPE_STRING:
		return Bare("string")
	case descriptorpb.FieldDescriptorProto_TYPE_MESSAGE:
		return Ref(proto.GetTypeName())
	case descriptorpb.FieldDescriptorProto_TYPE_BYTES:
		return Bare("bytes")
	case descriptorpb.FieldDescriptorProto_TYPE_ENUM:
		return Ref(proto.GetTypeName())
	}
	return Bare(proto.GetName())
}

func DefaultValue(t any) any {
	switch t {
	case "null":
		return nil
	case "boolean":
		return false
	case "int":
		return 0
	case "long":
		return 0
	case "float":
		return 0.0
	case "double":
		return 0.0
	case "map":
		return map[string]any{}
	case "record":
		return noDefault
	case "array":
		return []any{}
	case "string":
		return ""
	}

	switch typedT := t.(type) {
		case []any:
			return DefaultValue(typedT[0])
		case *orderedmap.OrderedMap:
			val, _ := typedT.Get("type")
			if val == "enum" {
				defaultVal, _ := typedT.Get("default")
				return defaultVal
			}
			return DefaultValue(val)
	}

	return noDefault
}
