// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"sync"

	"entgo.io/ent/examples/fs/ent/file"
	"entgo.io/ent/examples/fs/ent/predicate"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeFile = "File"
)

// FileMutation represents an operation that mutates the File nodes in the graph.
type FileMutation struct {
	config
	op              Op
	typ             string
	id              *int
	name            *string
	deleted         *bool
	clearedFields   map[string]struct{}
	parent          *int
	clearedparent   bool
	children        map[int]struct{}
	removedchildren map[int]struct{}
	clearedchildren bool
	done            bool
	oldValue        func(context.Context) (*File, error)
	predicates      []predicate.File
}

var _ ent.Mutation = (*FileMutation)(nil)

// fileOption allows management of the mutation configuration using functional options.
type fileOption func(*FileMutation)

// newFileMutation creates new mutation for the File entity.
func newFileMutation(c config, op Op, opts ...fileOption) *FileMutation {
	m := &FileMutation{
		config:        c,
		op:            op,
		typ:           TypeFile,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withFileID sets the ID field of the mutation.
func withFileID(id int) fileOption {
	return func(m *FileMutation) {
		var (
			err   error
			once  sync.Once
			value *File
		)
		m.oldValue = func(ctx context.Context) (*File, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().File.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withFile sets the old File of the mutation.
func withFile(node *File) fileOption {
	return func(m *FileMutation) {
		m.oldValue = func(context.Context) (*File, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m FileMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m FileMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *FileMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetName sets the "name" field.
func (m *FileMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *FileMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the File entity.
// If the File object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *FileMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *FileMutation) ResetName() {
	m.name = nil
}

// SetDeleted sets the "deleted" field.
func (m *FileMutation) SetDeleted(b bool) {
	m.deleted = &b
}

// Deleted returns the value of the "deleted" field in the mutation.
func (m *FileMutation) Deleted() (r bool, exists bool) {
	v := m.deleted
	if v == nil {
		return
	}
	return *v, true
}

// OldDeleted returns the old "deleted" field's value of the File entity.
// If the File object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *FileMutation) OldDeleted(ctx context.Context) (v bool, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldDeleted is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldDeleted requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDeleted: %w", err)
	}
	return oldValue.Deleted, nil
}

// ResetDeleted resets all changes to the "deleted" field.
func (m *FileMutation) ResetDeleted() {
	m.deleted = nil
}

// SetParentID sets the "parent_id" field.
func (m *FileMutation) SetParentID(i int) {
	m.parent = &i
}

// ParentID returns the value of the "parent_id" field in the mutation.
func (m *FileMutation) ParentID() (r int, exists bool) {
	v := m.parent
	if v == nil {
		return
	}
	return *v, true
}

// OldParentID returns the old "parent_id" field's value of the File entity.
// If the File object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *FileMutation) OldParentID(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldParentID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldParentID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldParentID: %w", err)
	}
	return oldValue.ParentID, nil
}

// ClearParentID clears the value of the "parent_id" field.
func (m *FileMutation) ClearParentID() {
	m.parent = nil
	m.clearedFields[file.FieldParentID] = struct{}{}
}

// ParentIDCleared returns if the "parent_id" field was cleared in this mutation.
func (m *FileMutation) ParentIDCleared() bool {
	_, ok := m.clearedFields[file.FieldParentID]
	return ok
}

// ResetParentID resets all changes to the "parent_id" field.
func (m *FileMutation) ResetParentID() {
	m.parent = nil
	delete(m.clearedFields, file.FieldParentID)
}

// ClearParent clears the "parent" edge to the File entity.
func (m *FileMutation) ClearParent() {
	m.clearedparent = true
}

// ParentCleared reports if the "parent" edge to the File entity was cleared.
func (m *FileMutation) ParentCleared() bool {
	return m.ParentIDCleared() || m.clearedparent
}

// ParentIDs returns the "parent" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// ParentID instead. It exists only for internal usage by the builders.
func (m *FileMutation) ParentIDs() (ids []int) {
	if id := m.parent; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetParent resets all changes to the "parent" edge.
func (m *FileMutation) ResetParent() {
	m.parent = nil
	m.clearedparent = false
}

// AddChildIDs adds the "children" edge to the File entity by ids.
func (m *FileMutation) AddChildIDs(ids ...int) {
	if m.children == nil {
		m.children = make(map[int]struct{})
	}
	for i := range ids {
		m.children[ids[i]] = struct{}{}
	}
}

// ClearChildren clears the "children" edge to the File entity.
func (m *FileMutation) ClearChildren() {
	m.clearedchildren = true
}

// ChildrenCleared reports if the "children" edge to the File entity was cleared.
func (m *FileMutation) ChildrenCleared() bool {
	return m.clearedchildren
}

// RemoveChildIDs removes the "children" edge to the File entity by IDs.
func (m *FileMutation) RemoveChildIDs(ids ...int) {
	if m.removedchildren == nil {
		m.removedchildren = make(map[int]struct{})
	}
	for i := range ids {
		m.removedchildren[ids[i]] = struct{}{}
	}
}

// RemovedChildren returns the removed IDs of the "children" edge to the File entity.
func (m *FileMutation) RemovedChildrenIDs() (ids []int) {
	for id := range m.removedchildren {
		ids = append(ids, id)
	}
	return
}

// ChildrenIDs returns the "children" edge IDs in the mutation.
func (m *FileMutation) ChildrenIDs() (ids []int) {
	for id := range m.children {
		ids = append(ids, id)
	}
	return
}

// ResetChildren resets all changes to the "children" edge.
func (m *FileMutation) ResetChildren() {
	m.children = nil
	m.clearedchildren = false
	m.removedchildren = nil
}

// Op returns the operation name.
func (m *FileMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (File).
func (m *FileMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *FileMutation) Fields() []string {
	fields := make([]string, 0, 3)
	if m.name != nil {
		fields = append(fields, file.FieldName)
	}
	if m.deleted != nil {
		fields = append(fields, file.FieldDeleted)
	}
	if m.parent != nil {
		fields = append(fields, file.FieldParentID)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *FileMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case file.FieldName:
		return m.Name()
	case file.FieldDeleted:
		return m.Deleted()
	case file.FieldParentID:
		return m.ParentID()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *FileMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case file.FieldName:
		return m.OldName(ctx)
	case file.FieldDeleted:
		return m.OldDeleted(ctx)
	case file.FieldParentID:
		return m.OldParentID(ctx)
	}
	return nil, fmt.Errorf("unknown File field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *FileMutation) SetField(name string, value ent.Value) error {
	switch name {
	case file.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case file.FieldDeleted:
		v, ok := value.(bool)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDeleted(v)
		return nil
	case file.FieldParentID:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetParentID(v)
		return nil
	}
	return fmt.Errorf("unknown File field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *FileMutation) AddedFields() []string {
	var fields []string
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *FileMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *FileMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown File numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *FileMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(file.FieldParentID) {
		fields = append(fields, file.FieldParentID)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *FileMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *FileMutation) ClearField(name string) error {
	switch name {
	case file.FieldParentID:
		m.ClearParentID()
		return nil
	}
	return fmt.Errorf("unknown File nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *FileMutation) ResetField(name string) error {
	switch name {
	case file.FieldName:
		m.ResetName()
		return nil
	case file.FieldDeleted:
		m.ResetDeleted()
		return nil
	case file.FieldParentID:
		m.ResetParentID()
		return nil
	}
	return fmt.Errorf("unknown File field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *FileMutation) AddedEdges() []string {
	edges := make([]string, 0, 2)
	if m.parent != nil {
		edges = append(edges, file.EdgeParent)
	}
	if m.children != nil {
		edges = append(edges, file.EdgeChildren)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *FileMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case file.EdgeParent:
		if id := m.parent; id != nil {
			return []ent.Value{*id}
		}
	case file.EdgeChildren:
		ids := make([]ent.Value, 0, len(m.children))
		for id := range m.children {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *FileMutation) RemovedEdges() []string {
	edges := make([]string, 0, 2)
	if m.removedchildren != nil {
		edges = append(edges, file.EdgeChildren)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *FileMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case file.EdgeChildren:
		ids := make([]ent.Value, 0, len(m.removedchildren))
		for id := range m.removedchildren {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *FileMutation) ClearedEdges() []string {
	edges := make([]string, 0, 2)
	if m.clearedparent {
		edges = append(edges, file.EdgeParent)
	}
	if m.clearedchildren {
		edges = append(edges, file.EdgeChildren)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *FileMutation) EdgeCleared(name string) bool {
	switch name {
	case file.EdgeParent:
		return m.clearedparent
	case file.EdgeChildren:
		return m.clearedchildren
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *FileMutation) ClearEdge(name string) error {
	switch name {
	case file.EdgeParent:
		m.ClearParent()
		return nil
	}
	return fmt.Errorf("unknown File unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *FileMutation) ResetEdge(name string) error {
	switch name {
	case file.EdgeParent:
		m.ResetParent()
		return nil
	case file.EdgeChildren:
		m.ResetChildren()
		return nil
	}
	return fmt.Errorf("unknown File edge %s", name)
}
