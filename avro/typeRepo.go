package avro

import (
  "fmt"
  "strings"
)

type TypeRepo struct {
  Types map[string]NamedType
  seenTypes map[string]bool // go "set"
  NamespaceMap map[string]string
}

func NewTypeRepo(namespaceMap map[string]string) *TypeRepo {
  return &TypeRepo{
    Types: make(map[string]NamedType),
    NamespaceMap: namespaceMap,
  }
}

func (r *TypeRepo) AddType(t NamedType) {
  fullName := FullName(t)
  r.Types[fullName] = t
}

func (r *TypeRepo) GetTypeByBareName(name string) Type {
  for _, t := range r.Types {
    if t.GetName() == name {
      return t
    }
  }
  return nil
}

func (r *TypeRepo) SeenType(t NamedType) {
  r.seenTypes[FullName(t)] = true
}

func (r *TypeRepo) GetType(name string) (Type, error) {
  if r.seenTypes[name] {
    return Bare(r.MappedNamespace(name[1:])), nil
  }
  t, ok := r.Types[name]
  if !ok {
    return nil, fmt.Errorf("type %s not found", name)
  }
  r.SeenType(t)
  return t, nil
}

func (r *TypeRepo) Start() {
  r.seenTypes = map[string]bool{}
}

func (r *TypeRepo) LogTypes() {
	var keys []string
	for k := range r.Types {
		keys = append(keys, k)
	}
	LogObj(keys)
}

func (r *TypeRepo) MappedNamespace(namespace string) string {
  out := namespace
  for k, v := range r.NamespaceMap {
    out = strings.Replace(out, k, v, -1)
  }
  return out
}
