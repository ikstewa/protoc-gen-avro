# protoc-gen-avro
Generate Avro schemas from Protobuf files.

## Usage

Download this project from the Releases page. Put the generated binary in your path:

```bash
mv protoc-gen-avro $GOPATH/bin
```

Generate Avro schemas from your Protobuf files:

```bash
protoc --avro_out=. *.proto
```

By default, this will generate an `.avsc` file for each message that has been read. You can instead specify the records to emit by using the `emit_only` option:

```bash
protoc --avro_out=. --avro_opt=emit_only=Foo,Bar:. *.proto
```

This will generate only `Foo.avsc` and `Bar.avsc` files.

You can also change the namespaces being mapped:

```bash
protoc --avro_out=. --avro_opt=namespace_map=foo:bar,baz:spam *.proto
```

...will change the output namespace for `foo` to `bar` and `baz` to `spam`.

---

To Do List:

* Add tests
* Map is currently outputting as Array<Map> due to how Protobuf handles [maps](https://protobuf.com/docs/descriptors#map-fields) (as repeated entries). Need to fix.
* Need to decide on how to truly differentiate between optional and required fields (technically all fields are optional on Protobuf, but maybe we should use the actual `optional` keyword and only have those be optional in Avro?)
* Split up `main.go` - it's trying to do too much

---

This project supported by [Flipp](https://corp.flipp.com/).
