/*
Copyright (c) 2020 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

// LabelKind is the name of the type used to represent objects
// of type 'label'.
const LabelKind = "Label"

// LabelLinkKind is the name of the type used to represent links
// to objects of type 'label'.
const LabelLinkKind = "LabelLink"

// LabelNilKind is the name of the type used to nil references
// to objects of type 'label'.
const LabelNilKind = "LabelNil"

// Label represents the values of the 'label' type.
//
// Representation of a label in clusterdeployment.
type Label struct {
	bitmap_ uint32
	id      string
	href    string
	key     string
	value   string
}

// Kind returns the name of the type of the object.
func (o *Label) Kind() string {
	if o == nil {
		return LabelNilKind
	}
	if o.bitmap_&1 != 0 {
		return LabelLinkKind
	}
	return LabelKind
}

// Link returns true if this is a link.
func (o *Label) Link() bool {
	return o != nil && o.bitmap_&1 != 0
}

// ID returns the identifier of the object.
func (o *Label) ID() string {
	if o != nil && o.bitmap_&2 != 0 {
		return o.id
	}
	return ""
}

// GetID returns the identifier of the object and a flag indicating if the
// identifier has a value.
func (o *Label) GetID() (value string, ok bool) {
	ok = o != nil && o.bitmap_&2 != 0
	if ok {
		value = o.id
	}
	return
}

// HREF returns the link to the object.
func (o *Label) HREF() string {
	if o != nil && o.bitmap_&4 != 0 {
		return o.href
	}
	return ""
}

// GetHREF returns the link of the object and a flag indicating if the
// link has a value.
func (o *Label) GetHREF() (value string, ok bool) {
	ok = o != nil && o.bitmap_&4 != 0
	if ok {
		value = o.href
	}
	return
}

// Empty returns true if the object is empty, i.e. no attribute has a value.
func (o *Label) Empty() bool {
	return o == nil || o.bitmap_&^1 == 0
}

// Key returns the value of the 'key' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// the key of the label
func (o *Label) Key() string {
	if o != nil && o.bitmap_&8 != 0 {
		return o.key
	}
	return ""
}

// GetKey returns the value of the 'key' attribute and
// a flag indicating if the attribute has a value.
//
// the key of the label
func (o *Label) GetKey() (value string, ok bool) {
	ok = o != nil && o.bitmap_&8 != 0
	if ok {
		value = o.key
	}
	return
}

// Value returns the value of the 'value' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// the value to set in the label
func (o *Label) Value() string {
	if o != nil && o.bitmap_&16 != 0 {
		return o.value
	}
	return ""
}

// GetValue returns the value of the 'value' attribute and
// a flag indicating if the attribute has a value.
//
// the value to set in the label
func (o *Label) GetValue() (value string, ok bool) {
	ok = o != nil && o.bitmap_&16 != 0
	if ok {
		value = o.value
	}
	return
}

// LabelListKind is the name of the type used to represent list of objects of
// type 'label'.
const LabelListKind = "LabelList"

// LabelListLinkKind is the name of the type used to represent links to list
// of objects of type 'label'.
const LabelListLinkKind = "LabelListLink"

// LabelNilKind is the name of the type used to nil lists of objects of
// type 'label'.
const LabelListNilKind = "LabelListNil"

// LabelList is a list of values of the 'label' type.
type LabelList struct {
	href  string
	link  bool
	items []*Label
}

// Kind returns the name of the type of the object.
func (l *LabelList) Kind() string {
	if l == nil {
		return LabelListNilKind
	}
	if l.link {
		return LabelListLinkKind
	}
	return LabelListKind
}

// Link returns true iif this is a link.
func (l *LabelList) Link() bool {
	return l != nil && l.link
}

// HREF returns the link to the list.
func (l *LabelList) HREF() string {
	if l != nil {
		return l.href
	}
	return ""
}

// GetHREF returns the link of the list and a flag indicating if the
// link has a value.
func (l *LabelList) GetHREF() (value string, ok bool) {
	ok = l != nil && l.href != ""
	if ok {
		value = l.href
	}
	return
}

// Len returns the length of the list.
func (l *LabelList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.items)
}

// Items sets the items of the list.
func (l *LabelList) SetLink(link bool) {
	l.link = link
}

// Items sets the items of the list.
func (l *LabelList) SetHREF(href string) {
	l.href = href
}

// Items sets the items of the list.
func (l *LabelList) SetItems(items []*Label) {
	l.items = items
}

// Items returns the items of the list.
func (l *LabelList) Items() []*Label {
	if l == nil {
		return nil
	}
	return l.items
}

// Empty returns true if the list is empty.
func (l *LabelList) Empty() bool {
	return l == nil || len(l.items) == 0
}

// Get returns the item of the list with the given index. If there is no item with
// that index it returns nil.
func (l *LabelList) Get(i int) *Label {
	if l == nil || i < 0 || i >= len(l.items) {
		return nil
	}
	return l.items[i]
}

// Slice returns an slice containing the items of the list. The returned slice is a
// copy of the one used internally, so it can be modified without affecting the
// internal representation.
//
// If you don't need to modify the returned slice consider using the Each or Range
// functions, as they don't need to allocate a new slice.
func (l *LabelList) Slice() []*Label {
	var slice []*Label
	if l == nil {
		slice = make([]*Label, 0)
	} else {
		slice = make([]*Label, len(l.items))
		copy(slice, l.items)
	}
	return slice
}

// Each runs the given function for each item of the list, in order. If the function
// returns false the iteration stops, otherwise it continues till all the elements
// of the list have been processed.
func (l *LabelList) Each(f func(item *Label) bool) {
	if l == nil {
		return
	}
	for _, item := range l.items {
		if !f(item) {
			break
		}
	}
}

// Range runs the given function for each index and item of the list, in order. If
// the function returns false the iteration stops, otherwise it continues till all
// the elements of the list have been processed.
func (l *LabelList) Range(f func(index int, item *Label) bool) {
	if l == nil {
		return
	}
	for index, item := range l.items {
		if !f(index, item) {
			break
		}
	}
}
