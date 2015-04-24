package system

import (
	"reflect"

	"kego.io/json"
)

/***********************************************************
*** @array ***
***********************************************************/

// Restriction rules for arrays
type Array_rule struct {
	*Object

	// This is a rule object, defining the type and restrictions on the value of the items
	Items Rule

	// This is the maximum number of items alowed in the array
	MaxItems Number

	// This is the minimum number of items alowed in the array
	MinItems Number

	// If this is true, each item must be unique
	UniqueItems Bool `kego:"{\"default\": false}"`
}

func (o *Array_rule) Base() *Object {
	return o.Object
}

/***********************************************************
*** @bool ***
***********************************************************/

// Restriction rules for bools
type Bool_rule struct {
	*Object

	// Default value of this is missing or null
	Default Bool
}

func (o *Bool_rule) Base() *Object {
	return o.Object
}

/***********************************************************
*** @import ***
***********************************************************/

// Restriction rules for import
type Import_rule struct {
	*Object
}

func (o *Import_rule) Base() *Object {
	return o.Object
}

/***********************************************************
*** @imports ***
***********************************************************/

// Restriction rules for imports
type Imports_rule struct {
	*Object
}

func (o *Imports_rule) Base() *Object {
	return o.Object
}

/***********************************************************
*** @map ***
***********************************************************/

// Restriction rules for maps
type Map_rule struct {
	*Object

	// This is a rule object, defining the type and restrictions on the value of the items.
	Items Rule

	// This is the maximum number of items alowed in the array
	MaxItems Number

	// This is the minimum number of items alowed in the array
	MinItems Number
}

func (o *Map_rule) Base() *Object {
	return o.Object
}

/***********************************************************
*** @number ***
***********************************************************/

// Restriction rules for numbers
type Number_rule struct {
	*Object

	// Default value if this property is omitted
	Default Number

	// If this is true, the value must be less than maximum. If false or not provided, the value must be less than or equal to the maximum.
	ExclusiveMaximum Bool `kego:"{\"default\": false}"`

	// If this is true, the value must be greater than minimum. If false or not provided, the value must be greater than or equal to the minimum.
	ExclusiveMinimum Bool `kego:"{\"default\": false}"`

	// This provides an upper bound for the restriction
	Maximum Number

	// This provides a lower bound for the restriction
	Minimum Number

	// This restricts the number to be a multiple of the given number
	MultipleOf Number
}

func (o *Number_rule) Base() *Object {
	return o.Object
}

/***********************************************************
*** @object ***
***********************************************************/

// Restriction rules for objects
type Object_rule struct {
	*Object
}

func (o *Object_rule) Base() *Object {
	return o.Object
}

/***********************************************************
*** @property ***
***********************************************************/

// Restriction rules for properties
type Property_rule struct {
	*Object
}

func (o *Property_rule) Base() *Object {
	return o.Object
}

/***********************************************************
*** @reference ***
***********************************************************/

// Restriction rules for references
type Reference_rule struct {
	*Object

	// Default value of this is missing or null
	Default Reference
}

func (o *Reference_rule) Base() *Object {
	return o.Object
}

/***********************************************************
*** @rule ***
***********************************************************/

// Restriction rules for rules
type Rule_rule struct {
	*Object
}

func (o *Rule_rule) Base() *Object {
	return o.Object
}

/***********************************************************
*** @string ***
***********************************************************/

// Restriction rules for strings
type String_rule struct {
	*Object

	// Default value of this is missing or null
	Default String

	// The value of this string is restricted to one of the provided values
	Enum []String

	// This restricts the value to one of several built-in formats
	Format String

	// The value must be shorter or equal to the provided maximum length
	MaxLength Number

	// The value must be longer or equal to the provided minimum length
	MinLength Number

	// This is a regex to match the value to
	Pattern String
}

func (o *String_rule) Base() *Object {
	return o.Object
}

/***********************************************************
*** @type ***
***********************************************************/

// Restriction rules for types
type Type_rule struct {
	*Object
}

func (o *Type_rule) Base() *Object {
	return o.Object
}

/***********************************************************
*** array ***
***********************************************************/

// This is the native json array data type
type Array struct {
	*Object
}

/***********************************************************
*** bool ***
***********************************************************/

/***********************************************************
*** import ***
***********************************************************/

// This is a single import.
type Import struct {
	*Object

	// The prefix used by the package. If omitted, we use the last part of the package URL.
	Name String

	// The package path.
	Package String
}

func (o *Import) Base() *Object {
	return o.Object
}

/***********************************************************
*** imports ***
***********************************************************/

// This contains a list of imports.
type Imports struct {
	*Object

	// A list of imports.
	Imports []*Import
}

func (o *Imports) Base() *Object {
	return o.Object
}

/***********************************************************
*** map ***
***********************************************************/

// This is the native json object data type.
type Map struct {
	*Object
}

/***********************************************************
*** number ***
***********************************************************/

/***********************************************************
*** object ***
***********************************************************/

// This is the most basic type.
type Object struct {

	// Unmarshaling context. This should not be in the json - it's added by the unmarshaler.
	Context *json.Context

	// Description for the developer
	Description String

	// All global objects should have an id.
	Id String

	// Type of the object.
	Type Reference
}

/***********************************************************
*** property ***
***********************************************************/

// This is a property of a type
type Property struct {
	*Object

	// This is a rule object, defining the type and restrictions on the value of the this property
	Item Rule

	// This specifies that the field is optional
	Optional Bool `kego:"{\"default\": false}"`
}

func (o *Property) Base() *Object {
	return o.Object
}

/***********************************************************
*** reference ***
***********************************************************/

/***********************************************************
*** rule ***
***********************************************************/

/***********************************************************
*** string ***
***********************************************************/

/***********************************************************
*** type ***
***********************************************************/

// This is the most basic type.
type Type struct {
	*Object

	// Type which this should extend
	Extends Reference `kego:"{\"default\": \"kego.io/system:object\"}"`

	// Is this type an interface?
	Interface Bool `kego:"{\"default\": false}"`

	// Array of interface types that this type should support
	Is []Reference

	// This is the native json type that represents this type. If omitted, default is object.
	Native String `kego:"{\"default\": \"object\"}"`

	// Link from rule types to the parent type. This should not be added to the json, it gets wired up when the json is unmarshaled.
	Parent *Type

	// Each field is listed with it's type
	Properties map[string]*Property

	// Embedded type that defines restriction rules for this type. By convention, the ID should be this type prefixed with the @ character.
	Rule *Type
}

func (o *Type) Base() *Object {
	return o.Object
}

func init() {

	json.RegisterType("kego.io/system:@array", reflect.TypeOf(&Array_rule{}))

	json.RegisterType("kego.io/system:@bool", reflect.TypeOf(&Bool_rule{}))

	json.RegisterType("kego.io/system:@import", reflect.TypeOf(&Import_rule{}))

	json.RegisterType("kego.io/system:@imports", reflect.TypeOf(&Imports_rule{}))

	json.RegisterType("kego.io/system:@map", reflect.TypeOf(&Map_rule{}))

	json.RegisterType("kego.io/system:@number", reflect.TypeOf(&Number_rule{}))

	json.RegisterType("kego.io/system:@object", reflect.TypeOf(&Object_rule{}))

	json.RegisterType("kego.io/system:@property", reflect.TypeOf(&Property_rule{}))

	json.RegisterType("kego.io/system:@reference", reflect.TypeOf(&Reference_rule{}))

	json.RegisterType("kego.io/system:@rule", reflect.TypeOf(&Rule_rule{}))

	json.RegisterType("kego.io/system:@string", reflect.TypeOf(&String_rule{}))

	json.RegisterType("kego.io/system:@type", reflect.TypeOf(&Type_rule{}))

	json.RegisterType("kego.io/system:array", reflect.TypeOf(&Array{}))

	json.RegisterType("kego.io/system:bool", reflect.TypeOf(&Bool{}))

	json.RegisterType("kego.io/system:import", reflect.TypeOf(&Import{}))

	json.RegisterType("kego.io/system:imports", reflect.TypeOf(&Imports{}))

	json.RegisterType("kego.io/system:map", reflect.TypeOf(&Map{}))

	json.RegisterType("kego.io/system:number", reflect.TypeOf(&Number{}))

	json.RegisterType("kego.io/system:object", reflect.TypeOf(&Object{}))

	json.RegisterType("kego.io/system:property", reflect.TypeOf(&Property{}))

	json.RegisterType("kego.io/system:reference", reflect.TypeOf(&Reference{}))

	json.RegisterType("kego.io/system:string", reflect.TypeOf(&String{}))

	json.RegisterType("kego.io/system:type", reflect.TypeOf(&Type{}))

}
