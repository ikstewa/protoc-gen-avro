package avro

import (
	"fmt"
	"os"
)

type Type interface {
	ToJSON(*TypeRepo) (any, error)
}

type NamedType interface {
	Type
	GetName() string
	GetNamespace() string
}

func LogMsg(msg string, args ...any) {
	fmt.Fprintf(os.Stderr, fmt.Sprintf(msg, args))
	fmt.Fprintln(os.Stderr)
}

func LogObj(arg any) {
	fmt.Fprintf(os.Stderr, fmt.Sprintf("%v", arg))
	fmt.Fprintln(os.Stderr)
}

func FullName(t NamedType) string {
	return fmt.Sprintf(".%s.%s", t.GetNamespace(), t.GetName())
}

func DefaultValue(t Type) any {
	switch t {
	case Bare("null"):
		return nil
	case Bare("boolean"):
		return false
	case Bare("int"):
		return 0
	case Bare("long"):
		return 0
	case Bare("float"):
		return 0.0
	case Bare("double"):
		return 0.0
	}

	switch typedT := t.(type) {
		case Record:
			return map[string]any{}
		case Map:
			return map[string]any{}
		case Array:
			return []string{}
		case Union:
			if typedT.Types[0] == Bare("null") {
				return nil
			}
			return DefaultValue(typedT.Types[0])
	}

	return ""
}
