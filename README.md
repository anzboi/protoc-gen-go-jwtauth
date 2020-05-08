# protoc-gen-go-jwtauth

Declare your methods required scopes in proto.

## Install

Clone the repo and run

```bash
$ go install .
```

## Add jwtauth options to your proto

Make sure to include the `jwtauthoption.proto` in your proto import paths. Add the following import and options to your proto specs

```proto
import jwtauthoption.proto

service Foo {
    rpc Bar(RequestMessage) returns (ResponseMessage) {
        option (jwtauth.scopes) = "your.scope";
        option (jwtauth.scopes) = "another.scope";
    }
}
```

Currently scopes are intersecting, defining two requires that both scopes are present in the caller token.

## Run

Add jwtauth validation function output to your proto with

```bash
$ protoc args --go-jwtauth_out=./build/path
```
