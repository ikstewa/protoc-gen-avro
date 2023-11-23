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

By default, this will generate an `.avsc` file for each `.proto` file. You can instead specify the records to emit by using the `emit_only` option:

```bash
protoc --avro_out=. --avro_opt=emit_only=Foo,Bar:. *.proto
```

This will generate only `Foo.avsc` and `Bar.avsc` files.

---

This project supported by [Flipp](https://corp.flipp.com/).
