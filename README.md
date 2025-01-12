<p align="center">
  <a href="https://mdxjs.com">
    <img alt="kubectl tree logo" src="https://github.com/bedirhangull/protolex-web/blob/main/assets/logo/protolex.png" width="100%" />
  </a>
</p>

# Protolex

[![website](https://img.shields.io/badge/website-000000?style=for-the-badge&logo=About.me&logoColor=white)](https://bedirhangull.github.io/protolex-web/)

Protolex is a powerful Go package for parsing and analyzing Protocol Buffer (protobuf) files. It provides a comprehensive set of tools to extract and analyze various components of proto files, including messages, enums, services, packages, and more.

## Features

### Core Functionality

- Parse and extract message blocks
- Analyze enum definitions and fields
- Extract service definitions and RPCs
- Handle nested package structures
- Parse syntax and package declarations
- Analyze message fields and types
- Handle reserved fields
- Process nested enums within messages

### Package Management

- Support for multiple package declarations
- Handle nested package hierarchies
- Parent-child package relationship analysis

## Installation

```bash
go get github.com/bedirhangull/protolex
```

## Usage

### Basic Usage

```go
package main

import (
    "fmt"
    "github.com/bedirhangull/protolex"
)

func main() {
    // Read and parse proto file
    proto, err := protolex.ReadProtoFile("./path/to/your.proto")
    if err != nil {
        fmt.Printf("Error reading proto file: %v\n", err)
        return
    }

    // Get all message blocks
    messages := proto.GetAllMessageBlocks()
    for _, msg := range messages {
        fmt.Println("Message:", msg)
    }

    // Get all enum blocks
    enums := proto.GetAllEnumBlocks()
    for _, enum := range enums {
        fmt.Println("Enum:", enum)
    }
}
```

### Working with Messages

```go
// Get a specific message by name
message := proto.GetMessageByName("User")

// Get all fields in a message
fields, err := proto.GetAllFieldsByMessageName("User")
if err != nil {
    // Handle error
}

// Get all types used in a message
types, err := proto.GetAllTypesByMessageName("User")
if err != nil {
    // Handle error
}
```

### Working with Enums

```go
// Get a specific enum by name
enum := proto.GetEnumByName("Status")

// Get all enum fields
fields, err := proto.GetAllEnumFieldsByEnumName("Status")
if err != nil {
    // Handle error
}

// Get all enums in a message
enums, err := proto.GetAllEnumsByMessageName("User")
if err != nil {
    // Handle error
}
```

### Package Management

```go
// Get all packages
packages, err := proto.GetAllPackages()
if err != nil {
    // Handle error
}

// Get a specific package
pkg, err := proto.GetPackageByPath("com.example.service")
if err != nil {
    // Handle error
}

// Get child packages
children, err := proto.GetChildPackagesByPath("com.example")
if err != nil {
    // Handle error
}

// Get parent package
parent, err := proto.GetParentPackageByPath("com.example.service")
if err != nil {
    // Handle error
}
```

### Service Analysis

```go
// Get all service blocks
services := proto.GetAllServiceBlocks()

// Get a specific service
service := proto.GetServiceByName("UserService")

// Get specific RPC
rpc := proto.GetRPCByServiceName("UserService", "CreateUser")
```

## API Reference

### Message Operations

| Function                                                       | Return Type         | Description                                  |
| -------------------------------------------------------------- | ------------------- | -------------------------------------------- |
| `GetAllMessageBlocks()`                                        | `[]string`          | Returns all message blocks in the proto file |
| `GetMessageByName(name string)`                                | `string`            | Returns a specific message block by its name |
| `GetAllFieldsByMessageName(messageName string)`                | `([]string, error)` | Returns all fields defined in a message      |
| `GetAllTypesByMessageName(messageName string)`                 | `([]string, error)` | Returns all types used in a message          |
| `GetAllReservedFieldsByMessageName(messageName string)`        | `([]string, error)` | Returns all reserved fields in a message     |
| `GetAllMessageBlocksByRPCName(service string, RPCname string)` | `([]string, error)` | Returns all message blocks used in an RPC    |

### Enum Operations

| Function                                       | Return Type         | Description                               |
| ---------------------------------------------- | ------------------- | ----------------------------------------- |
| `GetAllEnumBlocks()`                           | `[]string`          | Returns all enum blocks in the proto file |
| `GetEnumByName(name string)`                   | `string`            | Returns a specific enum block by its name |
| `GetAllEnumFieldsByEnumName(enumName string)`  | `([]string, error)` | Returns all fields defined in an enum     |
| `GetAllEnumsByMessageName(messageName string)` | `([]string, error)` | Returns all enums used in a message       |

### Service Operations

| Function                                           | Return Type | Description                                           |
| -------------------------------------------------- | ----------- | ----------------------------------------------------- |
| `GetAllServiceBlocks()`                            | `[]string`  | Returns all service blocks in the proto file          |
| `GetServiceByName(name string)`                    | `string`    | Returns a specific service block by its name          |
| `GetRPCByServiceName(name string, RPCname string)` | `string`    | Returns a specific RPC definition from a service      |
| `GetServerStreamingServices()`                     | `[]string`  | Returns the names of server streaming services        |
| `GetClientStreamingServices()`                     | `[]string`  | Returns the names of client streaming services        |
| `GetBidirectionalStreamingServices()`              | `[]string`  | Returns the names of bidirectional streaming services |

### Package Operations

| Function                                    | Return Type           | Description                                    |
| ------------------------------------------- | --------------------- | ---------------------------------------------- |
| `GetAllPackages()`                          | `([]*Package, error)` | Returns all packages defined in the proto file |
| `GetPackageByPath(path string)`             | `(*Package, error)`   | Returns a specific package by its full path    |
| `GetChildPackagesByPath(parentPath string)` | `([]*Package, error)` | Returns all child packages of a given package  |
| `GetParentPackageByPath(childPath string)`  | `(*Package, error)`   | Returns the parent package of a given package  |

### Other Operations

| Function           | Return Type | Description                                       |
| ------------------ | ----------- | ------------------------------------------------- |
| `GetSyntax()`      | `string`    | Returns the syntax declaration of the proto file  |
| `GetPackageName()` | `string`    | Returns the package declaration of the proto file |

## TODO

- Implement tests
- CI CD workflows
