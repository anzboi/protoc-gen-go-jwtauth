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
import "jwtauthoption.proto";

service Foo {
    rpc Bar(RequestMessage) returns (ResponseMessage) {
        option (jwtauth.scopes) = {
            and: "your.scope"
            and: "another.scope"
        };
        option (jwtauth.scopes) = {
            and: "why.not.a.third.scope"
        };
    };
}
```

The structure of scopes options allows for arbitrary combinations of scopes joined under intersection and union (AND and OR logic). Scopes separated by an `option (jwtauth.Scopes)` block obey `OR` logic, while scopes inside an `option` block obey `AND` logic.

The above example will validate scopes if they have
```
your.scopes AND another.scope
OR
why.not.a.third.scope
```

## Run

Add jwtauth validation function output to your proto with

```bash
$ protoc args --go-jwtauth_out=./build/path
```
