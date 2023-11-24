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

