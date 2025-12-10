package util

import (
	"context"

	"bot-go/internal/model/ast"
	"bot-go/internal/service/codegraph"
	"bot-go/internal/signals"
)

// FieldAccess represents a single field access
type FieldAccess struct {
	FieldID      ast.NodeID
	FieldName    string
	OwnerClassID ast.NodeID
	IsForeign    bool
	AccessType   signals.AccessType
}

// FieldAccessMatrix maps methods to their accessed fields
type FieldAccessMatrix struct {
	Methods map[ast.NodeID][]ast.NodeID // methodID -> []fieldID
	Fields  map[ast.NodeID][]ast.NodeID // fieldID -> []methodID (inverse)
}

// FieldAccessAnalyzer tracks field access patterns
type FieldAccessAnalyzer struct {
	codeGraph *codegraph.CodeGraph
}

// NewFieldAccessAnalyzer creates a new field access analyzer
func NewFieldAccessAnalyzer(cg *codegraph.CodeGraph) *FieldAccessAnalyzer {
	return &FieldAccessAnalyzer{
		codeGraph: cg,
	}
}

// GetMethodFieldAccesses returns fields accessed by a method
// It finds fields by traversing CONTAINS relationships from method to Field nodes
// and determines access type (read/write) by checking DATA_FLOW relationships
func (a *FieldAccessAnalyzer) GetMethodFieldAccesses(ctx context.Context, methodID ast.NodeID) ([]FieldAccess, error) {
	// Get fields accessed within the method body
	fieldNodes, err := a.codeGraph.GetFieldsAccessedByMethod(ctx, methodID)
	if err != nil {
		return nil, err
	}

	// Get fields that are written to (targets of DATA_FLOW)
	writtenFields, err := a.codeGraph.GetFieldsWrittenByMethod(ctx, methodID)
	if err != nil {
		return nil, err
	}

	// Build a set of written field IDs for quick lookup
	writtenFieldSet := make(map[ast.NodeID]bool)
	for _, wf := range writtenFields {
		writtenFieldSet[wf.ID] = true
	}

	// Get the class that contains this method (to determine if access is local or foreign)
	containingClass, err := a.codeGraph.GetContainingClass(ctx, methodID)
	if err != nil {
		return nil, err
	}

	var containingClassID ast.NodeID
	if containingClass != nil {
		containingClassID = containingClass.ID
	}

	var accesses []FieldAccess
	for _, fieldNode := range fieldNodes {
		// Determine the owner class of this field
		ownerClass, err := a.codeGraph.GetFieldOwnerClass(ctx, fieldNode.ID)
		if err != nil {
			continue
		}

		var ownerClassID ast.NodeID
		if ownerClass != nil {
			ownerClassID = ownerClass.ID
		}

		// Determine if this is a foreign access
		isForeign := ownerClassID != ast.InvalidNodeID &&
			containingClassID != ast.InvalidNodeID &&
			ownerClassID != containingClassID

		// Determine access type based on DATA_FLOW relationship
		accessType := signals.AccessTypeRead
		if writtenFieldSet[fieldNode.ID] {
			accessType = signals.AccessTypeWrite
		}

		accesses = append(accesses, FieldAccess{
			FieldID:      fieldNode.ID,
			FieldName:    fieldNode.Name,
			OwnerClassID: ownerClassID,
			IsForeign:    isForeign,
			AccessType:   accessType,
		})
	}

	return accesses, nil
}

// GetClassFieldAccessMatrix returns method->fields access matrix for a class
func (a *FieldAccessAnalyzer) GetClassFieldAccessMatrix(ctx context.Context, classID ast.NodeID) (*FieldAccessMatrix, error) {
	// Get all methods of the class
	methods, err := a.codeGraph.GetMethodsOfClass(ctx, classID)
	if err != nil {
		return nil, err
	}

	matrix := NewFieldAccessMatrix()

	for _, method := range methods {
		// Get fields accessed by this method
		accesses, err := a.GetMethodFieldAccesses(ctx, method.ID)
		if err != nil {
			continue
		}

		for _, access := range accesses {
			matrix.AddAccess(method.ID, access.FieldID)
		}
	}

	return matrix, nil
}

// GetLocalFieldAccesses returns accesses to fields within the same class
func (a *FieldAccessAnalyzer) GetLocalFieldAccesses(ctx context.Context, methodID ast.NodeID, classID ast.NodeID) ([]FieldAccess, error) {
	allAccesses, err := a.GetMethodFieldAccesses(ctx, methodID)
	if err != nil {
		return nil, err
	}

	var localAccesses []FieldAccess
	for _, access := range allAccesses {
		// Local access: field belongs to the same class
		if access.OwnerClassID == classID {
			access.IsForeign = false
			localAccesses = append(localAccesses, access)
		}
	}

	return localAccesses, nil
}

// GetForeignFieldAccesses returns accesses to fields from other classes
func (a *FieldAccessAnalyzer) GetForeignFieldAccesses(ctx context.Context, methodID ast.NodeID, classID ast.NodeID) ([]FieldAccess, error) {
	allAccesses, err := a.GetMethodFieldAccesses(ctx, methodID)
	if err != nil {
		return nil, err
	}

	var foreignAccesses []FieldAccess
	for _, access := range allAccesses {
		// Foreign access: field belongs to a different class (and has a known owner)
		if access.OwnerClassID != ast.InvalidNodeID && access.OwnerClassID != classID {
			access.IsForeign = true
			foreignAccesses = append(foreignAccesses, access)
		}
	}

	return foreignAccesses, nil
}

// NewFieldAccessMatrix creates a new field access matrix
func NewFieldAccessMatrix() *FieldAccessMatrix {
	return &FieldAccessMatrix{
		Methods: make(map[ast.NodeID][]ast.NodeID),
		Fields:  make(map[ast.NodeID][]ast.NodeID),
	}
}

// AddAccess adds a method-field access relationship
func (m *FieldAccessMatrix) AddAccess(methodID ast.NodeID, fieldID ast.NodeID) {
	if m.Methods == nil {
		m.Methods = make(map[ast.NodeID][]ast.NodeID)
	}
	if m.Fields == nil {
		m.Fields = make(map[ast.NodeID][]ast.NodeID)
	}

	// Add field to method's list (avoid duplicates)
	found := false
	for _, f := range m.Methods[methodID] {
		if f == fieldID {
			found = true
			break
		}
	}
	if !found {
		m.Methods[methodID] = append(m.Methods[methodID], fieldID)
	}

	// Add method to field's list (avoid duplicates)
	found = false
	for _, mid := range m.Fields[fieldID] {
		if mid == methodID {
			found = true
			break
		}
	}
	if !found {
		m.Fields[fieldID] = append(m.Fields[fieldID], methodID)
	}
}

// GetSharedFields returns fields accessed by multiple methods
func (m *FieldAccessMatrix) GetSharedFields() map[ast.NodeID][]ast.NodeID {
	if m.Fields == nil {
		return nil
	}

	shared := make(map[ast.NodeID][]ast.NodeID)
	for fieldID, methods := range m.Fields {
		if len(methods) > 1 {
			shared[fieldID] = methods
		}
	}
	return shared
}

// GetConnectedMethodPairs returns method pairs that share field access
func (m *FieldAccessMatrix) GetConnectedMethodPairs() [][2]ast.NodeID {
	if m.Fields == nil {
		return nil
	}

	var pairs [][2]ast.NodeID
	seen := make(map[[2]ast.NodeID]bool)

	for _, methods := range m.Fields {
		// For each field, create pairs of all methods that access it
		for i := 0; i < len(methods); i++ {
			for j := i + 1; j < len(methods); j++ {
				// Normalize pair order for deduplication
				pair := [2]ast.NodeID{methods[i], methods[j]}
				if methods[i] > methods[j] {
					pair = [2]ast.NodeID{methods[j], methods[i]}
				}
				if !seen[pair] {
					seen[pair] = true
					pairs = append(pairs, pair)
				}
			}
		}
	}
	return pairs
}

// GetMethodCount returns the number of methods in the matrix
func (m *FieldAccessMatrix) GetMethodCount() int {
	if m.Methods == nil {
		return 0
	}
	return len(m.Methods)
}

// GetFieldCount returns the number of fields in the matrix
func (m *FieldAccessMatrix) GetFieldCount() int {
	if m.Fields == nil {
		return 0
	}
	return len(m.Fields)
}

// GetFieldsAccessedByMethod returns fields accessed by a specific method
func (m *FieldAccessMatrix) GetFieldsAccessedByMethod(methodID ast.NodeID) []ast.NodeID {
	if m.Methods == nil {
		return nil
	}
	return m.Methods[methodID]
}

// GetMethodsAccessingField returns methods that access a specific field
func (m *FieldAccessMatrix) GetMethodsAccessingField(fieldID ast.NodeID) []ast.NodeID {
	if m.Fields == nil {
		return nil
	}
	return m.Fields[fieldID]
}

// SharesFieldWith checks if two methods share access to at least one field
func (m *FieldAccessMatrix) SharesFieldWith(methodID1 ast.NodeID, methodID2 ast.NodeID) bool {
	if m.Methods == nil {
		return false
	}

	fields1 := m.Methods[methodID1]
	fields2 := m.Methods[methodID2]

	if len(fields1) == 0 || len(fields2) == 0 {
		return false
	}

	// Create set of fields for method2
	fields2Set := make(map[ast.NodeID]bool)
	for _, f := range fields2 {
		fields2Set[f] = true
	}

	// Check if any field from method1 is in method2's set
	for _, f := range fields1 {
		if fields2Set[f] {
			return true
		}
	}
	return false
}
