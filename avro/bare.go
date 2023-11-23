package avro

type Bare string

func (b Bare) ToJSON(_ *TypeRepo) (any, error) {
  return string(b), nil
}
