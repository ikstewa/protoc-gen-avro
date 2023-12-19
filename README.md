# protoc-gen-avro
Generate Avro schemas from Protobuf files.

## Usage

Download this project from the Releases page. Put the generated binary in your path:

```bash
mv protoc-gen-avro /usr/local/bin
```

Generate Avro schemas from your Protobuf files:

```bash
protoc --avro_out=. *.proto
```

## Options

* `emit_only` - A semicolon-separated list of records to emit. If not specified, all records will be emitted.

```bash
protoc --avro_out=. --avro_opt=emit_only=Foo;Bar *.proto
```

This will generate only `Foo.avsc` and `Bar.avsc` files.

* `namespace_map` - A semicolon-separated list of namespaces to map. If not specified, all namespaces will be mapped.

```bash
protoc --avro_out=. --avro_opt=namespace_map=foo:bar,baz:spam *.proto
```

...will change the output namespace for `foo` to `bar` and `baz` to `spam`.

* `collapse_fields` - A semicolon-separated list of records to collapse. If not specified, no records will be collapsed. Collapsed records should have a single field in them, and they will be replaced in the output by that field. This can be useful to overcome some limitations of Protobuf - e.g. Protobuf doesn't have the ability to have a map of arrays, while Avro does.

```bash
protoc --avro_out=. --avro_opt=collapse_fields=StringList;SomeOtherRecord *.proto
```

If you have the input proto looking like:

```protobuf
message StringList {
  repeated string strings = 1;
}
message MyRecord {
    repeated StringList my_field = 1;
}
```

...the output will look like:

```avro
{
  "type": "record",
  "name": "MyRecord",
  "fields": [
    {
      "name": "my_field",
      "type": {
        "type": "array",
        "items": {
          "type": "string"
        }
      }
    }
  ]
}
```

---

To Do List:

* Add tests
* Need to decide on how to truly differentiate between optional and required fields (technically all fields are optional on Protobuf, but maybe we should use the actual `optional` keyword and only have those be optional in Avro?)
* Split up `main.go` - it's trying to do too much
* Homebrew?

---

This project supported by [Flipp](https://corp.flipp.com/).
