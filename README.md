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

* `namespace_map` - A comma-separated list of namespaces to map. If not specified, all namespaces will be mapped.

```bash
protoc --avro_out=. --avro_opt=namespace_map=foo:bar,baz:spam *.proto
```

...will change the output namespace for `foo` to `bar` and `baz` to `spam`.

* `collapse_fields` - A semicolon-separated list of records to collapse. Collapsed records should have a single field in them, and they will be replaced in the output by that field. This can be useful to overcome some limitations of Protobuf - e.g. Protobuf doesn't have the ability to have an array of maps, while Avro does.

```bash
protoc --avro_out=. --avro_opt=collapse_fields=StringList;SomeOtherRecord *.proto
```

If you have the input proto looking like:

```protobuf
message StringList {
  repeated string strings = 1;
}
message MyRecord {
    map<string, StringList> my_field = 1;
}
```

...the output will look like:

```json
{
  "type": "record",
  "name": "MyRecord",
  "fields": [
    {
      "name": "my_field",
      "type": {
        "type": "map",
        "values": {
          "type": {
            "type": "array",
            "items": "string"
          }
        }
      }
    }
  ]
}
```

* `remove_enum_prefixes` - if set to true, will remove the prefixes from enum values. E.g. if you have an enum like:
```protobuf
enum Category {
  CATEGORY_GOOD = 0;
  CATEGORY_BAD = 1;
}
```

...with this option on, it will be translated into:

```json
{
  "type": "enum",
  "name": "CATEGORY",
  "symbols": [
    "GOOD",
    "BAD"
  ]
}
```

* `preserve_non_string_maps` - if set to true, will replace maps with non-string keys with records. E.g. if you have a map like:
```protobuf
message MyRecord {
  map<int32, string> my_field = 1;
}
```

...with this option off, it will be translated into:

```json
{
  "type": "record",
  "name": "MyRecord",
  "fields": [
    {
      "name": "my_field",
      "type": {
        "type": "map",
        "values": "string"
      }
    }
  ]
}
```

...but with it on, it will be translated into:

```json
{
  "type": "record",
  "name": "MyRecord",
  "fields": [
    {
      "name": "my_field",
      "type": {
        "type": "record",
        "fields": [
          {
            "name": "key",
            "type": "int"
          },
          {
            "name": "value",
            "type": "string"
          }
        ]
      }
    }
  ]
}
```

* `prefix_schema_files_with_package` - if set to true, files will be generated into folders matching the proto package. E.g. :
if set to true, will remove the prefixes from enum values. E.g. if you have an enum like:
```protobuf
package my.test.data;
message Yowza {
  float hoo_boy = 1;
}
```

...with this option on, it will be generated as:

`./my.test.data/Yowza.avsc`

* `json_fieldname` - if set to true, field names will use JSON format (camel case). E.g. :
```protobuf
message Yowza {
  int32 hoo_boy = 1;
}
```

...with this option on, it will be translated into:

```json
{
  "type": "record",
  "name": "Yowza",
  "fields": [
    {
      "name": "hooBoy",
      "type": "int"
    }
  ]
}
```

* `retain_oneof_fieldnames` - if set to true, when using a oneof the fields retain their original name instead of using the type name. This is intended to better match the json output from proto. E.g. :
```protobuf
message AOneOf {
  oneof oneof_types {
    string a_string = 1;
    int32 a_num = 2;
  }
}
message Widget {
  AOneOf a_one_of = 6;
}
```

...with this option _off_ (default), it will be translated into:
```json
{
  "type": "record",
  "name": "Widget",
  "fields": [
    {
      "name": "a_one_of",
      "type": {
        "type": "record",
        "name": "AOneOf",
        "namespace": "testdata",
        "fields": [
          {
            "name": "oneof_types",
            "type": [
              "string",
              "int"
            ],
            "default": null
          }
        ]
      }
    }
  ]
}
```

...with this option on, it will be translated into:

```json
{
  "type": "record",
  "name": "Widget",
  "fields": [
    {
      "name": "a_one_of",
      "type": {
        "type": "record",
        "name": "AOneOf",
        "fields": [
          {
            "name": "a_string",
            "type": [
              "null",
              "string"
            ],
            "default": null
          },
          {
            "name": "a_num",
            "type": [
              "null",
              "int"
            ],
            "default": null
          }
        ]
      }
    }
  ]
}
```

* `nullable_arrays` - if set to true (default false), arrays are mapped to union of null. E.g. :

```protobuf
message StringList {
  repeated string strings = 1;
}
```

...the output will look like:

```json
{
  "type": "record",
  "name": "StringList",
  "fields": [
    {
      "name": "strings",
      "type": [
        "null",
        {
          "type": "array",
          "items": "string"
        }
      ],
      "default": null
    }
  ]
}
```

---

To Do List:

* Add tests
* Homebrew?

---

This project supported by [Flipp](https://corp.flipp.com/).
