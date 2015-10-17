package system

import (
	"reflect"

	"kego.io/json"
	"kego.io/kerr"
)

// Enforcer is a rule with properties that need to be enforced against data.
type Enforcer interface {
	Enforce(data interface{}, path string, aliases map[string]string) (bool, string, error)
}

type CollectionRule interface {
	GetItemsRule() RuleInterface
}

func init() {
	type dummyRule struct {
		*Object
		*Rule
		Default interface{}
	}
	json.RegisterInterface(reflect.TypeOf((*RuleInterface)(nil)).Elem(), reflect.TypeOf(&dummyRule{}))
}

type RuleHolder struct {
	Rule       RuleInterface
	RuleType   *Type
	ParentType *Type
}

func (r *RuleHolder) GetReflectType() (reflect.Type, error) {

	if r.Rule.GetRule().Interface {
		typ, _, ok := json.GetTypeInterface(r.ParentType.Id.Package, r.ParentType.Id.Name)
		if !ok {
			return nil, kerr.New("QGUVEUTXAN", nil, "Type interface for %s not found", r.ParentType.Id.Value())
		}
		return typ, nil
	}

	switch r.ParentType.Native.Value {
	case "object", "number", "bool", "string":
		typ, _, ok := json.GetType(r.ParentType.Id.Package, r.ParentType.Id.Name)
		if !ok {
			return nil, kerr.New("DLAJJPJDPL", nil, "Type %s not found", r.ParentType.Id.Value())
		}
		if r.ParentType.Native.Value != "object" && typ.Kind() == reflect.Ptr {
			return typ.Elem(), nil
		}
		return typ, nil
	case "array", "map":
		c, ok := r.Rule.(CollectionRule)
		if !ok {
			return nil, kerr.New("GSYSHQOWNH", nil, "Map types must have rule that implements CollectionRule")
		}
		itemsRule := c.GetItemsRule()
		items, err := NewRuleHolder(itemsRule)
		if err != nil {
			return nil, kerr.New("EDEMPPVUNW", err, "NewRuleHolder")
		}
		itemsType, err := items.GetReflectType()
		if err != nil {
			return nil, kerr.New("LMKEHHWHKL", err, "GetReflectType")
		}
		if r.ParentType.Native.Value == "map" {
			return reflect.MapOf(reflect.TypeOf(""), itemsType), nil
		}
		return reflect.SliceOf(itemsType), nil
	default:
		return nil, kerr.New("VDEORSSUWA", nil, "Unknown native %s", r.ParentType.Native.Value)
	}
}

func NewMinimalRuleHolder(t *Type) *RuleHolder {
	return &RuleHolder{Rule: nil, RuleType: nil, ParentType: t}
}
func NewRuleHolder(r RuleInterface) (*RuleHolder, error) {
	rt, pt, err := ruleTypes(r)
	if err != nil {
		return nil, kerr.New("VRCWUGOTMA", err, "ruleTypes")
	}
	return &RuleHolder{Rule: r, RuleType: rt, ParentType: pt}, nil
}

func ruleTypes(r RuleInterface) (ruleType *Type, parentType *Type, err error) {
	ruleReference, err := ruleTypeReference(r)
	if err != nil {
		return nil, nil, kerr.New("BNEKIFYDDL", err, "ruleTypeReference")
	}
	rt, ok := ruleReference.GetType()
	if !ok {
		return nil, nil, kerr.New("PFGWISOHRR", nil, "ruleReference.GetType: type %v not found", ruleReference.Value())
	}
	typeReference := ruleReference.ChangeToType()
	pt, ok := typeReference.GetType()
	if !ok {
		return nil, nil, kerr.New("KYCTDXKFYR", nil, "typeReference.GetType: type %v not found", typeReference.Value())
	}
	return rt, pt, nil
}

func ruleTypeReference(r RuleInterface) (*Reference, error) {
	ob, ok := r.(ObjectInterface)
	if !ok {
		return nil, kerr.New("VKFNPJDNVB", nil, "r does not implement ObjectInterface")
	}
	return &ob.GetObject().Type, nil

}

// ItemsRule returns Items rule for a collection Rule.
func (r *RuleHolder) ItemsRule() (*RuleHolder, error) {
	if !r.ParentType.IsNativeCollection() {
		return nil, kerr.New("VPAGXSTQHM", nil, "parentType %s is not a collection", r.ParentType.Id)
	}
	items, _, ok, err := RuleFieldByReflection(r.Rule, "Items")
	if err != nil {
		return nil, kerr.New("LIDXIQYGJD", err, "RuleFieldByReflection")
	}
	if !ok {
		return nil, kerr.New("VYTHGJTSNJ", nil, "RuleFieldByReflection could not find Items field")
	}
	rule, ok := items.(RuleInterface)
	if !ok {
		return nil, kerr.New("DIFVRMVWMC", nil, "items is not a rule")
	}
	rh, err := NewRuleHolder(rule)
	if err != nil {
		return nil, kerr.New("FGYMQPNBQJ", err, "NewRuleHolder")
	}
	return rh, nil
}

func RuleFieldByReflection(object interface{}, name string) (value interface{}, pointer interface{}, ok bool, err error) {
	v := reflect.ValueOf(object)
	if v.Kind() != reflect.Ptr {
		return nil, nil, false, kerr.New("QOYMWPXWUO", nil, "val.Kind (%s) is not a Ptr: %v", name, v.Kind())
	}
	if v.Elem().Kind() != reflect.Struct {
		return nil, nil, false, kerr.New("IGOUOBGXAN", nil, "val.Elem().Kind (%s) is not a Struct: %v", name, v.Elem().Kind())
	}
	value, pointer, _, found, zero, err := GetObjectField(v.Elem(), name)

	// zero => !ok
	return value, pointer, found && !zero, err
}

func GetMapMember(v reflect.Value, name string) (object interface{}, pointer interface{}, value reflect.Value, found bool, zero bool, err error) {
	for v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	member := v.MapIndex(reflect.ValueOf(name))
	return returnValue(member)
}
func GetArrayMember(v reflect.Value, index int) (object interface{}, pointer interface{}, value reflect.Value, found bool, zero bool, err error) {
	for v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	member := v.Index(index)
	return returnValue(member)
}
func GetObjectField(v reflect.Value, name string) (object interface{}, pointer interface{}, value reflect.Value, found bool, zero bool, err error) {
	for v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	field := v.FieldByName(name)
	return returnValue(field)
}
func returnValue(v reflect.Value) (object interface{}, pointer interface{}, value reflect.Value, found bool, zero bool, err error) {
	value = v
	empty := reflect.Value{}
	if v == empty {
		// Field does not exist
		return
	}
	if v.Kind() == reflect.Ptr {
		// If it's a pointer we should only return not found if
		// it's nil:
		if v.IsNil() {
			zero = true
			return
		}
	} else if v.Kind() == reflect.Map || v.Kind() == reflect.Slice {
		if v.Len() == 0 {
			zero = true
		}
	} else {
		// If it's not a pointer, we return not found if it's an
		// zero value
		nilValue := reflect.Zero(v.Type())
		if v.Interface() == nilValue.Interface() {
			zero = true
		}
	}
	found = true
	object = v.Interface()
	// This prevents **Foo being returned for pointer when value is already *Foo
	if v.Kind() == reflect.Ptr {
		pointer = v.Interface()
	} else if v.CanAddr() {
		pointer = v.Addr().Interface()
	}
	return
}

// HoldsDisplayType returns the string to display when communicating to
// the end user what type this rule holds.
func (r *RuleHolder) HoldsDisplayType(path string, aliases map[string]string) (string, error) {
	str, err := r.ParentType.Id.ValueContext(path, aliases)
	if err != nil {
		return "", kerr.New("OPIFCOHGWI", err, "ValueContext")
	}
	if r.Rule.GetRule().Interface || r.ParentType.Interface {
		str += "*"
	}
	return str, nil
}
