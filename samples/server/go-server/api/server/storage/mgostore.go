// Copyright 2017 The Grafeas Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package storage

import (
	"fmt"
	"net/http"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/grafeas/grafeas/samples/server/go-server/api"
	"github.com/grafeas/grafeas/samples/server/go-server/api/server/name"
	"github.com/grafeas/grafeas/server-go"
	"github.com/grafeas/grafeas/server-go/errors"
)

// mgoStore is mongoDb based storage backend for grafeas
type mgoStore struct {
	session     *mgo.Session
	occurrences *mgo.Collection
	notes       *mgo.Collection
}

// NewMgoStore initiates a mongodb session
func NewMgoStore() server.Storager {
	m := &mgoStore{Dbinit(), nil, nil}
	m.occurrences = m.session.DB("grafeas").C("occurrences")
	m.notes = m.session.DB("grafeas").C("notes")
	return m
}

// CreateOccurrence adds the specified occurrence to the mem store
func (m *mgoStore) CreateOccurrence(o *swagger.Occurrence) *errors.AppError {
	//m.occurrences := m.session.DB("grafeas").C("occurrences")
	count, err := m.occurrences.Find(bson.M{"name": o.Name}).Count()
	if count > 0 {
		return &errors.AppError{Err: fmt.Sprintf("Occurrence with name %q already exists", o.Name),
			StatusCode: http.StatusBadRequest}
	}
	if err != nil {
		return &errors.AppError{Err: fmt.Sprintf("error searching for occurrence %q", o.Name),
			StatusCode: http.StatusBadRequest}
	}
	m.occurrences.Insert(o)
	return nil
}

// DeleteOccurrence deletes the occurrence with the given pID and oID from the memStore
func (m *mgoStore) DeleteOccurrence(pID, oID string) *errors.AppError {
	oName := name.OccurrenceName(pID, oID)
	//m.occurrences := m.session.DB("grafeas").C("occurrences")
	err := m.occurrences.Remove(bson.M{"name": oName})
	if err != nil {
		return &errors.AppError{Err: fmt.Sprintf("Occurrence with name %q not found", oName),
			StatusCode: http.StatusBadRequest}
	}
	// 	delete(m.occurrencesByID, oName)
	return nil
}

//
// // UpdateOccurrence updates the existing occurrence with the given projectID and occurrenceID
func (m *mgoStore) UpdateOccurrence(pID, oID string, o *swagger.Occurrence) *errors.AppError {
	// 	oName := name.OccurrenceName(pID, oID)
	// 	if _, ok := m.occurrencesByID[oName]; !ok {
	// 		return &errors.AppError{Err: fmt.Sprintf("Occurrence with oName %q does not Exist", oName),
	// 			StatusCode: http.StatusBadRequest}
	// 	}
	// 	m.occurrencesByID[oName] = *o
	return nil
}

//
// // GetOccurrence returns the occurrence with pID and oID
func (m *mgoStore) GetOccurrence(pID, oID string) (*swagger.Occurrence, *errors.AppError) {
	oName := name.OccurrenceName(pID, oID)
	//	m.occurrences := m.session.DB("grafeas").C("occurrences")
	var o swagger.Occurrence
	err := m.occurrences.Find(bson.M{"name": oName}).One(&o)
	if err != nil {
		return nil, &errors.AppError{Err: fmt.Sprintf("Occurrence with name %q not found", oName),
			StatusCode: http.StatusBadRequest}
	}
	return &o, nil
}

//
func (m *mgoStore) ListOccurrences(pID, filters string) []swagger.Occurrence {
	os := []swagger.Occurrence{}
	// 	for _, o := range m.occurrencesByID {
	// 		if strings.HasPrefix(o.Name, fmt.Sprintf("projects/%v", pID)) {
	// 			os = append(os, o)
	// 		}
	// 	}
	return os
}

// CreateNote adds the specified note to the mem store
func (m *mgoStore) CreateNote(n *swagger.Note) *errors.AppError {
	//m.notes := m.session.DB("grafeas").C("notes")
	count, err := m.notes.Find(bson.M{"name": n.Name}).Count()
	if count > 0 {
		return &errors.AppError{Err: fmt.Sprintf("Note with name %q already exists", n.Name),
			StatusCode: http.StatusBadRequest}
	}
	if err != nil {
		return &errors.AppError{Err: fmt.Sprintf("error searching for note %q", n.Name),
			StatusCode: http.StatusBadRequest}
	}
	m.notes.Insert(n)
	return nil
}

// DeleteNote deletes the note with the given pID and nID from the memStore
func (m *mgoStore) DeleteNote(pID, nID string) *errors.AppError {
	// 	nName := name.NoteName(pID, nID)
	// 	if _, ok := m.notesByID[nName]; !ok {
	// 		return &errors.AppError{Err: fmt.Sprintf("Note with name %q does not Exist", nName),
	// 			StatusCode: http.StatusBadRequest}
	// 	}
	// 	delete(m.notesByID, nName)
	return nil
}

//
// // UpdateNote updates the existing note with the given pID and nID
func (m *mgoStore) UpdateNote(pID, nID string, n *swagger.Note) *errors.AppError {
	// 	nName := name.NoteName(pID, nID)
	// 	if _, ok := m.notesByID[nName]; !ok {
	// 		return &errors.AppError{Err: fmt.Sprintf("Note with name %q does not Exist", nName),
	// 			StatusCode: http.StatusBadRequest}
	// 	}
	// 	m.notesByID[nName] = *n
	return nil
}

//
// // GetNote returns the note with pID and nID
func (m *mgoStore) GetNote(pID, nID string) (*swagger.Note, *errors.AppError) {
	//	nName := name.NoteName(pID, nID)
	var n swagger.Note
	// 	n, ok := m.notesByID[nName]
	// 	if !ok {
	// 		return nil, &errors.AppError{Err: fmt.Sprintf("Note with name %q does not Exist", nName),
	// 			StatusCode: http.StatusBadRequest}
	// 	}
	return &n, nil
}

//
// // GetNoteByOccurrence returns the note attached to occurrence with pID and oID
func (m *mgoStore) GetNoteByOccurrence(pID, oID string) (*swagger.Note, *errors.AppError) {
	oName := name.OccurrenceName(pID, oID)
	//m.occurrences := m.session.DB("grafeas").C("occurrences")
	var o swagger.Occurrence
	err := m.occurrences.Find(bson.M{"name": oName}).One(&o)
	if err != nil {
		return nil, &errors.AppError{Err: fmt.Sprintf("Occurrence with name %q does not Exist", oName),
			StatusCode: http.StatusBadRequest}
	}
	//m.notes := m.session.DB("grafeas").C("notes")
	var n swagger.Note
	err = m.notes.Find(bson.M{"name": o.NoteName}).One(&n)
	if err != nil {
		return nil, &errors.AppError{Err: fmt.Sprintf("Note with name %q does not Exist", o.NoteName),
			StatusCode: http.StatusBadRequest}
	}
	return &n, nil
}

//
func (m *mgoStore) ListNotes(pID, filters string) []swagger.Note {
	ns := []swagger.Note{}
	// 	for _, n := range m.notesByID {
	// 		if strings.HasPrefix(n.Name, fmt.Sprintf("projects/%v", pID)) {
	// 			ns = append(ns, n)
	// 		}
	// 	}
	return ns
}

func (m *mgoStore) ListNoteOccurrences(pID, nID, filters string) ([]swagger.Occurrence, *errors.AppError) {
	// 	// TODO: use filters
	// 	// Verify that note exists
	// 	if _, err := m.GetNote(pID, nID); err != nil {
	// 		return nil, err
	// 	}
	// 	nName := name.FormatNote(pID, nID)
	os := []swagger.Occurrence{}
	// 	for _, o := range m.occurrencesByID {
	// 		if o.NoteName == nName {
	// 			os = append(os, o)
	// 		}
	// 	}
	return os, nil
}

//
// // GetOperation returns the operation with pID and oID
func (m *mgoStore) GetOperation(pID, opID string) (*swagger.Operation, *errors.AppError) {
	// 	oName := name.OperationName(pID, opID)
	// 	o, ok := m.opsByID[oName]
	var o swagger.Operation
	// 	if !ok {
	// 		return nil, &errors.AppError{Err: fmt.Sprintf("Operation with name %q does not Exist", oName),
	// 			StatusCode: http.StatusBadRequest}
	// 	}
	return &o, nil
}

//
// CreateOperation adds the specified operation to the mem store
func (m *mgoStore) CreateOperation(o *swagger.Operation) *errors.AppError {
	// 	if _, ok := m.opsByID[o.Name]; ok {
	// 		return &errors.AppError{Err: fmt.Sprintf("Operation with name %q already exists", o.Name),
	// 			StatusCode: http.StatusBadRequest}
	// 	}
	// 	m.opsByID[o.Name] = *o
	return nil
}

//
// // DeleteOperation deletes the operation with the given pID and oID from the memStore
func (m *mgoStore) DeleteOperation(pID, opID string) *errors.AppError {
	// 	opName := name.OperationName(pID, opID)
	// 	if _, ok := m.opsByID[opName]; !ok {
	// 		return &errors.AppError{Err: fmt.Sprintf("Operation with name %q does not Exist", opName),
	// 			StatusCode: http.StatusBadRequest}
	// 	}
	// 	delete(m.occurrencesByID, opName)
	return nil
}

//
// // UpdateOperation updates the existing operation with the given pID and nID
func (m *mgoStore) UpdateOperation(pID, opID string, op *swagger.Operation) *errors.AppError {
	// 	opName := name.OperationName(pID, opID)
	// 	if _, ok := m.opsByID[opName]; !ok {
	// 		return &errors.AppError{Err: fmt.Sprintf("Operation with name %q does not Exist", opName),
	// 			StatusCode: http.StatusBadRequest}
	// 	}
	// 	m.opsByID[opName] = *op
	return nil
}

//
func (m *mgoStore) ListOperations(pID, filters string) []swagger.Operation {
	ops := []swagger.Operation{}
	// 	for _, op := range m.opsByID {
	// 		if strings.HasPrefix(op.Name, fmt.Sprintf("projects/%v", pID)) {
	// 			ops = append(ops, op)
	// 		}
	// 	}
	return ops
}
