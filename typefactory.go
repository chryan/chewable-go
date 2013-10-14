// Copyright (C) 2013 Ryan Chew. All rights reserved.
// Use of this source code is governed by a Apache 2.0
// license that can be found in the LICENSE file.

package cbl

import (
	"reflect"
)

type typeFactory struct {
	typeMap map[string]reflect.Type
}

var TypeFactory *typeFactory = &typeFactory{typeMap: make(map[string]reflect.Type)}

func (t *typeFactory) Add(addtype reflect.Type) *typeFactory {
	if kind := addtype.Kind(); kind == reflect.Ptr || kind == reflect.Interface {
		addtype = addtype.Elem()
	}

	name := addtype.String()
	if _, ok := t.typeMap[name]; !ok {
		t.typeMap[name] = addtype
	}

	return t
}

func (t *typeFactory) New(typename, pkgname string) interface{} {
	if typ, ok := t.typeMap[pkgname+"."+typename]; ok {
		return reflect.New(typ).Elem().Interface()
	}
	return nil
}

func (t *typeFactory) Get(typename, pkgname string) reflect.Type {
	if typ, ok := t.typeMap[pkgname+"."+typename]; ok {
		return typ
	}
	return nil
}

func (t *typeFactory) Del(typename, pkgname string) {
	delete(t.typeMap, pkgname+"."+typename)
}

func (t *typeFactory) Clear() {
	t.typeMap = make(map[string]reflect.Type)
}
