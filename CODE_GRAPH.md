# Code Graph Architecture

## Overview

The Code Graph is the core data structure in bot-go that represents source code as a graph database. It captures the structural relationships between code elements (functions, classes, variables, etc.) and enables powerful queries for code analysis, navigation, and code smell detection.

## Architecture

```
┌─────────────────────────────────────────────────────────────────┐
│                        CodeGraph                                 │
│  ┌─────────────────────────────────────────────────────────┐   │
│  │                   GraphDatabase                          │   │
│  │         (Neo4j or Kuzu implementation)                   │   │
│  └─────────────────────────────────────────────────────────┘   │
│                              │                                   │
│              ┌───────────────┴───────────────┐                  │
│              ▼                               ▼                  │
│     ┌─────────────┐                 ┌─────────────┐            │
│     │   Neo4j DB  │                 │   Kuzu DB   │            │
│     │  (External) │                 │ (Embedded)  │            │
│     └─────────────┘                 └─────────────┘            │
└─────────────────────────────────────────────────────────────────┘
```

### Components

**CodeGraph** (`internal/service/codegraph/code_graph.go`)
- Main API for creating and querying code graph nodes and relationships
- Supports batch writing for improved performance during indexing
- File-level buffers for parallel processing

**GraphDatabase** (`internal/service/codegraph/graph_db.go`)
- Abstract interface for graph database operations
- Supports both Neo4j (production) and Kuzu (embedded/in-memory)
- All queries use Cypher query language

**Parsers** (`internal/parse/`)
- Tree-sitter based parsers convert source code to AST
- Language-specific visitors (GoVisitor, PythonVisitor, etc.) create graph nodes
- `TranslateFromSyntaxTree` orchestrates the parsing process

## Node Types

All nodes are defined in `internal/model/ast/ast.go` and share common properties:

| Property | Type | Description |
|----------|------|-------------|
| `id` | `NodeID (int64)` | Unique identifier |
| `node_type` | `NodeType (int8)` | Type discriminator |
| `file_id` | `int32` | Reference to source file |
| `name` | `string` | Node name (optional) |
| `range` | `Range` | Source location (start/end line/char) |
| `version` | `int32` | File version for incremental updates |
| `scope_id` | `NodeID` | Parent scope reference |
| `metadata` | `map[string]any` | Additional properties |

### Node Type Reference

| Type ID | Constant | Graph Label | Description |
|---------|----------|-------------|-------------|
| 1 | `NodeTypeModuleScope` | `ModuleScope` | Package or module (e.g., Go package) |
| 2 | `NodeTypeFileScope` | `FileScope` | Source file container |
| 3 | `NodeTypeBlock` | `Block` | Code block (function body, loop body, etc.) |
| 4 | `NodeTypeVariable` | `Variable` | Variable declaration or reference |
| 5 | `NodeTypeExpression` | `Expression` | General expression |
| 6 | `NodeTypeConditional` | `Conditional` | If/switch/match statements |
| 7 | `NodeTypeFunction` | `Function` | Function or method declaration |
| 8 | `NodeTypeClass` | `Class` | Struct, interface, class, or type definition |
| 9 | `NodeTypeField` | `Field` | Struct/class field or accessed property |
| 10 | `NodeTypeFunctionCall` | `FunctionCall` | Function or method invocation |
| 11 | `NodeTypeFileNumber` | `FileNumber` | Internal: file ID allocation tracker |
| 12 | `NodeTypeLoop` | `Loop` | For/while/range loops |
| 13 | `NodeTypeImport` | `Import` | Import declaration |

### Node Metadata

Different node types store additional information in their `metadata` field:

**FileScope**
```go
metadata = {
    "path": "/absolute/path/to/file.go",  // File path
    "language": "go",                      // Source language
    "modified": "2024-01-15T10:30:00Z",   // Modification time
}
```

**Class**
```go
metadata = {
    "struct_type": "struct",    // "struct", "interface", "class"
    "is_receiver": true,        // If this is a method receiver type
}
```

**Variable**
```go
metadata = {
    "type": "string",           // Variable type (if known)
    "return": true,             // If this is a return value
}
```

**Function**
```go
metadata = {
    "path": "/path/to/file.go", // Containing file
}
```

**FunctionCall**
```go
metadata = {
    "external": true,           // If call target is in another file
    "fake": true,              // If function name is synthetic
}
```

**Import**
```go
metadata = {
    "path": "github.com/pkg/name", // Full import path
}
```

**Loop/Conditional**
```go
metadata = {
    "condition": 12345,         // Node ID of condition expression
    "init": 12346,              // Node ID of init statement (for loops)
}
```

## Relationship Types

Relationships connect nodes to represent code structure and data flow.

### Structural Relationships

| Relationship | From | To | Description |
|--------------|------|-----|-------------|
| `CONTAINS` | Parent | Child | Hierarchical containment (file→class, class→method, method→variable) |
| `CONTAINED_BY` | Child | Parent | Inverse of CONTAINS (rarely used) |
| `BODY` | Function/Loop/Conditional | Block | Links to body block |
| `BRANCH` | Conditional | Block | Links to conditional branches (if/else) |

**CONTAINS Examples:**
```cypher
// File contains a class
(FileScope)-[:CONTAINS]->(Class)

// Class contains methods and fields
(Class)-[:CONTAINS]->(Function)
(Class)-[:CONTAINS]->(Field)

// Function contains variables, calls, loops
(Function)-[:CONTAINS]->(Variable)
(Function)-[:CONTAINS]->(FunctionCall)
(Function)-[:CONTAINS]->(Loop)

// Nested containment (variable depth)
(Function)-[:CONTAINS*]->(Field)  // Any depth
```

**BRANCH Properties:**
```go
metadata = {
    "branch_type": "if" | "else" | "else_if" | "case" | "default",
}
```

### Call Relationships

| Relationship | From | To | Description |
|--------------|------|-----|-------------|
| `CALLS` | Caller Function | Callee Function | Function calls another function |
| `CALLS_FUNCTION` | FunctionCall | Function | Resolved call site to definition |

**Call Graph Example:**
```cypher
// Function A calls Function B
(A:Function)-[:CALLS]->(B:Function)

// Call site linked to definition
(call:FunctionCall)-[:CALLS_FUNCTION]->(target:Function)
```

### Data Flow Relationships

| Relationship | From | To | Description |
|--------------|------|-----|-------------|
| `DATA_FLOW` | Source | Target | Data flows from source to target (assignment) |
| `USES_VARIABLE` | Expression | Variable | Expression uses a variable |
| `ALIAS` | Alias | Original | One node is an alias for another |

**DATA_FLOW Example:**
```cypher
// Assignment: x = y + 1
// Data flows from RHS to LHS
(rhs:Expression)-[:DATA_FLOW]->(x:Variable)

// Field write detection
(source)-[:DATA_FLOW]->(field:Field)
```

### Type Relationships

| Relationship | From | To | Description |
|--------------|------|-----|-------------|
| `HAS_FIELD` | Variable/Expression | Field | Selector expression (e.g., `obj.field`) |
| `THIS` | Variable | Class | Receiver variable points to its class |
| `INHERITS` | Child Class | Parent Class | Inheritance/embedding |

**THIS Relationship (Method Receivers):**
```cypher
// In Go: func (s *Server) Start() { ... }
// The receiver 's' has a THIS relationship to Server
(s:Variable)-[:THIS]->(Server:Class)

// Query: Find all fields accessed via 'this' in a method
MATCH (m:Function)-[:CONTAINS*]->(thisVar)-[:THIS]->(c:Class)
MATCH (thisVar)-[:HAS_FIELD*]->(f:Field)
RETURN f
```

### Function Argument Relationships

| Relationship | From | To | Properties |
|--------------|------|-----|------------|
| `FUNCTION_ARG` | Function | Variable | `{position: int}` - Parameter definition |
| `FUNCTION_CALL_ARG` | FunctionCall | Expression | `{position: int}` - Call argument |
| `RETURNS` | Function | Expression | Return value |

### Module Relationships

| Relationship | From | To | Description |
|--------------|------|-----|-------------|
| `IMPORTS` | FileScope | Module | File imports a module |
| `FROM` | Import | Module | Import comes from a module |

## Query Patterns

### Basic Navigation

```cypher
// Get all methods of a class
MATCH (c:Class {id: $classId})-[:CONTAINS]->(m:Function)
RETURN m

// Get all fields of a class
MATCH (c:Class {id: $classId})-[:CONTAINS]->(f:Field)
RETURN f

// Get containing class of a method
MATCH (c:Class)-[:CONTAINS]->(m:Function {id: $methodId})
RETURN c
```

### Field Access Analysis

```cypher
// Get all fields accessed within a method (any depth)
MATCH (m:Function {id: $methodId})-[:CONTAINS*]->(f:Field)
RETURN DISTINCT f

// Get fields written by a method (via DATA_FLOW)
MATCH (m:Function {id: $methodId})-[:CONTAINS*]->(source)-[:DATA_FLOW]->(f:Field)
RETURN DISTINCT f

// Get fields accessed via 'this' receiver
MATCH (m:Function {id: $methodId})-[:CONTAINS*]->(thisVar)-[:THIS]->(c:Class)
MATCH (thisVar)-[:HAS_FIELD*]->(f:Field)
RETURN DISTINCT f
```

### Call Graph Analysis

```cypher
// Get functions called by a function
MATCH (f:Function {id: $functionId})-[:CALLS]->(called:Function)
RETURN called

// Get callers of a function
MATCH (caller:Function)-[:CALLS]->(f:Function {id: $functionId})
RETURN caller

// Get all call sites within a function
MATCH (f:Function {id: $functionId})-[:CONTAINS*]->(call:FunctionCall)
RETURN call
```

### Complexity Analysis

```cypher
// Count conditionals in a function
MATCH (f:Function {id: $functionId})-[:CONTAINS*]->(c:Conditional)
RETURN count(c) AS conditionalCount

// Count loops in a function
MATCH (f:Function {id: $functionId})-[:CONTAINS*]->(l:Loop)
RETURN count(l) AS loopCount

// Get nesting depth (nested blocks)
MATCH path = (f:Function {id: $functionId})-[:CONTAINS*]->(b:Block)
RETURN max(length(path)) AS maxDepth
```

## API Reference

### Node Creation

```go
// Create specific node types
cg.CreateFunction(ctx, node)
cg.CreateClass(ctx, node)
cg.CreateVariable(ctx, node)
cg.CreateField(ctx, node)
cg.CreateFunctionCall(ctx, node)
cg.CreateFileScope(ctx, node)
cg.CreateModuleScope(ctx, node)
cg.CreateConditional(ctx, node)
cg.CreateLoop(ctx, node)
cg.CreateBlock(ctx, node)
cg.CreateExpression(ctx, node)
cg.CreateImport(ctx, node)
```

### Relationship Creation

```go
// Structural
cg.CreateContainsRelation(ctx, parentID, childID, fileID)
cg.CreateBodyRelation(ctx, parentID, bodyID, fileID)
cg.CreateConditionalRelation(ctx, condID, branchID, branchType, fileID)

// Calls
cg.CreateCallsRelation(ctx, callerID, calleeID, fileID)
cg.CreateCallsFunctionRelation(ctx, callID, funcID, fileID)

// Data flow
cg.CreateDataFlowRelation(ctx, sourceID, targetID, fileID)
cg.CreateUsesVariableRelation(ctx, userID, varID, fileID)
cg.CreateAliasRelation(ctx, aliasID, originalID, fileID)

// Types
cg.CreateHasFieldRelation(ctx, parentID, fieldID, fileID)
cg.CreateThisRelation(ctx, thisVarID, classID, fileID)
cg.CreateInheritsRelation(ctx, childID, parentID, fileID)

// Functions
cg.CreateFunctionArgRelation(ctx, funcID, argID, position, fileID)
cg.CreateFunctionCallArgRelation(ctx, callID, argID, position, fileID)
cg.CreateReturnsRelation(ctx, funcID, returnID, fileID)

// Modules
cg.CreateImportsRelation(ctx, importerID, importedID, fileID)
cg.CreateFromRelation(ctx, fromID, toID, fileID)
```

### Query Methods

```go
// Node queries
cg.ReadFunction(ctx, nodeID)
cg.ReadClass(ctx, nodeID)
cg.GetFileScopeByName(ctx, relativePath)
cg.GetAllFileScopes(ctx)

// Relationship queries
cg.GetMethodsOfClass(ctx, classID)
cg.GetFieldsOfClass(ctx, classID)
cg.GetFieldsAccessedByMethod(ctx, methodID)
cg.GetFieldsWrittenByMethod(ctx, methodID)
cg.GetContainingClass(ctx, methodID)
cg.GetFieldOwnerClass(ctx, fieldID)
cg.GetThisClassForMethod(ctx, methodID)

// Call graph queries
cg.GetCalledFunctions(ctx, functionID)
cg.GetCallerFunctions(ctx, functionID)
```

## Batch Writing

For indexing large codebases, batch writing improves performance:

```go
// Enable in config
config.CodeGraph.EnableBatchWrites = true
config.CodeGraph.BatchSize = 100

// Initialize buffers before processing a file
cg.InitializeFileBuffers(fileID)

// Nodes and relations are buffered automatically
cg.CreateFunction(ctx, node)
cg.CreateContainsRelation(ctx, parentID, childID, fileID)

// Flush buffers when file processing is complete
cg.FlushFile(ctx, fileID)
```

## Database Backends

### Neo4j
- Production-ready, external server
- Full Cypher support
- Requires separate deployment

```go
cg, err := codegraph.NewCodeGraph(
    "neo4j://localhost:7687",
    "neo4j",
    "password",
    config,
    logger,
)
```

### Kuzu
- Embedded database
- Supports in-memory or file-based storage
- No external dependencies

```go
// In-memory
config.Kuzu.Path = ":memory:"

// File-based
config.Kuzu.Path = "/path/to/database"

cg, err := codegraph.NewCodeGraphWithKuzu(config, logger)
```

## Language Support

The code graph is populated by language-specific visitors:

| Language | Visitor | Supported Features |
|----------|---------|-------------------|
| Go | `GoVisitor` | Functions, structs, interfaces, methods, imports |
| Python | `PythonVisitor` | Functions, classes, methods, imports |
| JavaScript/TypeScript | `JavaScriptVisitor` | Functions, classes, methods, imports |
| Java | `JavaVisitor` | Classes, methods, fields, imports |

Each visitor implements the `SyntaxTreeVisitor` interface and converts tree-sitter AST nodes to code graph nodes.
