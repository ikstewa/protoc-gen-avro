package avro

import (
  "fmt"
)

type TypeRepo struct {
  Types map[string]NamedType
  seenTypes map[string]bool // go "set"
}

func NewTypeRepo() *TypeRepo {
  return &TypeRepo{Types: make(map[string]NamedType)}
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

func (r *TypeRepo) GetType(name string) (Type, error) {
  if r.seenTypes[name] {
    return Bare(name[1:]), nil
  }
  t, ok := r.Types[name]
  if !ok {
//    r.LogTypes()
    return nil, fmt.Errorf("type %s not found", name)
  }
  r.seenTypes[FullName(t)] = true
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

