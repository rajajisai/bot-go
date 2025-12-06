package util

import (
	"context"

	"bot-go/internal/model/ast"
	"bot-go/internal/service/codegraph"
	"bot-go/pkg/lsp/base"
)

// TypeInfo contains type information
type TypeInfo struct {
	Name      string     // Type name
	ClassID   ast.NodeID // Resolved class node (if applicable)
	IsBuiltin bool       // Is a built-in type
	IsArray   bool       // Is an array/slice type
	IsPointer bool       // Is a pointer type
	IsGeneric bool       // Is a generic/parameterized type
}

// TypeResolver resolves variable/expression types
type TypeResolver struct {
	codeGraph *codegraph.CodeGraph
	lspClient base.LSPClient // Optional LSP for accurate resolution
}

// NewTypeResolver creates a new type resolver
func NewTypeResolver(cg *codegraph.CodeGraph, lspClient base.LSPClient) *TypeResolver {
	return &TypeResolver{
		codeGraph: cg,
		lspClient: lspClient,
	}
}

// ResolveVariableType returns the type of a variable
func (r *TypeResolver) ResolveVariableType(ctx context.Context, variableID ast.NodeID) (TypeInfo, error) {
	return TypeInfo{}, nil
}

// ResolveExpressionType returns the type of an expression
func (r *TypeResolver) ResolveExpressionType(ctx context.Context, exprID ast.NodeID) (TypeInfo, error) {
	return TypeInfo{}, nil
}

// ResolveFieldOwner returns the class that owns a field
func (r *TypeResolver) ResolveFieldOwner(ctx context.Context, fieldID ast.NodeID) (ast.NodeID, error) {
	return ast.InvalidNodeID, nil
}

// ResolveMethodOwner returns the class that owns a method
func (r *TypeResolver) ResolveMethodOwner(ctx context.Context, methodID ast.NodeID) (ast.NodeID, error) {
	return ast.InvalidNodeID, nil
}

// ResolveCallTarget returns the target method of a function call
func (r *TypeResolver) ResolveCallTarget(ctx context.Context, callID ast.NodeID) (ast.NodeID, error) {
	return ast.InvalidNodeID, nil
}

// ResolveCallTargetClass returns the class containing the target method
func (r *TypeResolver) ResolveCallTargetClass(ctx context.Context, callID ast.NodeID) (ast.NodeID, error) {
	return ast.InvalidNodeID, nil
}

// IsBuiltinType checks if a type name is a built-in type
func (r *TypeResolver) IsBuiltinType(typeName string, language string) bool {
	return false
}

// GetClassByName finds a class by name in the code graph
func (r *TypeResolver) GetClassByName(ctx context.Context, className string, fileID int32) (ast.NodeID, error) {
	return ast.InvalidNodeID, nil
}
