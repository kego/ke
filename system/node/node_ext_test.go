package node_test

import (
	"testing"

	"reflect"

	"github.com/davelondon/ktest/assert"
	"github.com/davelondon/ktest/require"
	"golang.org/x/net/context"
	"kego.io/json"
	"kego.io/process/parser"
	"kego.io/system"
	"kego.io/system/node"
	"kego.io/tests"
	"kego.io/tests/data"
)

func TestUnmarshal(t *testing.T) {

	cb := tests.Context("kego.io/tests/data").Jauto().Sauto(parser.Parse)

	s := `{
	"type": "multi",
	"js": "a"
}`

	n, err := node.Unmarshal(cb.Ctx(), []byte(s))
	require.NoError(t, err)
	m, ok := n.Value.(*data.Multi)
	require.True(t, ok)
	assert.Equal(t, "a", m.Js)

}

func TestNode_Unpack(t *testing.T) {

	cb := tests.Context("kego.io/tests/data").Jauto().Sauto(parser.Parse)

	s := `{
	"type": "multi",
	"js": "a"
}`

	n := node.NewNode()
	require.NoError(t, n.Unpack(cb.Ctx(), json.PackString(s)))
	m, ok := n.Value.(*data.Multi)
	require.True(t, ok)
	assert.Equal(t, "a", m.Js)

}

func TestNode_Unpack3(t *testing.T) {

	cb := tests.Context("kego.io/tests/data").Jauto().Sauto(parser.Parse)

	s := `{
	"type": "multi",
	"a": "b"
}`

	n := node.NewNode()
	err := n.Unpack(cb.Ctx(), json.PackString(s))
	assert.HasError(t, err, "SRANLETJRS")

}

func setup(t *testing.T) (*tests.ContextBuilder, *node.Node) {

	cb := tests.Context("kego.io/tests/data").Jauto().Sauto(parser.Parse)

	m := `"type": "multi",
"js": "js1",
"ss": "ss1",
"sr": "sr1",
"jn": 1.1,
"sn": 1.2,
"jb": true,
"sb": false,
"i": { "type": "facea", "a": "ia" },
"ajs": [ "ajs0", "ajs1", "ajs2", "ajs3" ],
"ass": [ "ass0", "ass1", "ass2", "ass3" ],
"ajn": [ 2.1, 2.2, 2.3, 2.4 ],
"asn": [ 3.1, 3.2, 3.3, 3.4 ],
"ajb": [ true, false, false, false ],
"asb": [ false, true, false, false ],
"mjs": { "a": "mjsa", "b": "mjsb" },
"mss": { "a": "mssa", "b": "mssb" },
"mjn": { "a": 4.1, "b": 4.2 },
"msn": { "a": 5.1, "b": 5.2 },
"mjb": { "a": true, "b": false },
"msb": { "a": false, "b": true },
"anri": [ "anri0", { "type": "facea", "a": "anri1" }, { "type": "multi", "ss": "anri2" } ],
"mnri": { "a": "mnria", "b": { "type": "facea", "a": "mnrib" }, "c": { "type": "multi", "ss": "mnric" } }
`
	mm := m + `,
"m": { "type": "multi" },
"am": [ { "type": "multi" }, { "type": "multi" } ],
"mm": { "a": { "type": "multi" }, "b": { "type": "multi" } }`

	s := `{
	` + m + `,
	"m": {` + mm + `},
	"am": [ {` + mm + `}, {` + mm + `} ],
	"mm": { "a": {` + mm + `}, "b": {` + mm + `} }
}`
	n := node.NewNode()
	require.NoError(t, n.Unpack(cb.Ctx(), json.PackString(s)))
	return cb, n
}

func empty(t *testing.T) (*tests.ContextBuilder, *node.Node) {

	cb := tests.Context("kego.io/tests/data").Jauto().Sauto(parser.Parse)

	s := `{
	"type": "multi",
	"m": { "type": "multi" },
	"am": [ { "type": "multi" }, { "type": "multi" } ],
	"mm": { "a": { "type": "multi" }, "b": { "type": "multi" } }
}`
	n := node.NewNode()
	require.NoError(t, n.Unpack(cb.Ctx(), json.PackString(s)))
	return cb, n
}

func multi(t *testing.T, n *node.Node, m *data.Multi, test func(t *testing.T, n *node.Node, m *data.Multi)) {
	test(t, n.Map["mm"].Map["b"], m.Mm["b"])
	test(t, n.Map["mm"].Map["a"], m.Mm["a"])
	test(t, n.Map["am"].Array[1], m.Am[1])
	test(t, n.Map["am"].Array[0], m.Am[0])
	test(t, n.Map["m"], m.M)
	test(t, n, m)
}

func TestNode_Unpack2(t *testing.T) {

	cb, n := setup(t)

	test := func(t *testing.T, n *node.Node, m *data.Multi) {
		assert.Equal(t, "js1", m.Js)
		assert.Equal(t, "ss1", m.Ss.Value())
		assert.Equal(t, "sr1", m.Sr.Name)
		assert.Equal(t, 1.1, m.Jn)
		assert.Equal(t, 1.2, m.Sn.Value())
		assert.Equal(t, true, m.Jb)
		assert.Equal(t, false, m.Sb.Value())
		assert.Equal(t, "ia", m.I.Face())
		assert.Equal(t, "ajs0", m.Ajs[0])
		assert.Equal(t, "ajs1", m.Ajs[1])
		assert.Equal(t, "ass0", m.Ass[0].Value())
		assert.Equal(t, "ass1", m.Ass[1].Value())
		assert.Equal(t, 2.1, m.Ajn[0])
		assert.Equal(t, 2.2, m.Ajn[1])
		assert.Equal(t, 3.1, m.Asn[0].Value())
		assert.Equal(t, 3.2, m.Asn[1].Value())
		assert.Equal(t, true, m.Ajb[0])
		assert.Equal(t, false, m.Ajb[1])
		assert.Equal(t, false, m.Asb[0].Value())
		assert.Equal(t, true, m.Asb[1].Value())
		assert.Equal(t, "mjsa", m.Mjs["a"])
		assert.Equal(t, "mjsb", m.Mjs["b"])
		assert.Equal(t, "mssa", m.Mss["a"].Value())
		assert.Equal(t, "mssb", m.Mss["b"].Value())
		assert.Equal(t, 4.1, m.Mjn["a"])
		assert.Equal(t, 4.2, m.Mjn["b"])
		assert.Equal(t, 5.1, m.Msn["a"].Value())
		assert.Equal(t, 5.2, m.Msn["b"].Value())
		assert.Equal(t, true, m.Mjb["a"])
		assert.Equal(t, false, m.Mjb["b"])
		assert.Equal(t, false, m.Msb["a"].Value())
		assert.Equal(t, true, m.Msb["b"].Value())
		assert.Equal(t, "anri0", m.Anri[0].GetString(cb.Ctx()).Value())
		assert.Equal(t, "anri1", m.Anri[1].GetString(cb.Ctx()).Value())
		assert.Equal(t, "anri2", m.Anri[2].GetString(cb.Ctx()).Value())
		assert.Equal(t, "mnria", m.Mnri["a"].GetString(cb.Ctx()).Value())
		assert.Equal(t, "mnrib", m.Mnri["b"].GetString(cb.Ctx()).Value())
		assert.Equal(t, "mnric", m.Mnri["c"].GetString(cb.Ctx()).Value())
	}

	multi(t, n, n.Value.(*data.Multi), test)

}

func TestNode_SetValueString(t *testing.T) {

	cb, n := setup(t)

	test := func(t *testing.T, n *node.Node, m *data.Multi) {
		n.Map["js"].SetValueString(cb.Ctx(), "aa")
		assert.Equal(t, "aa", m.Js)

		n.Map["ss"].SetValueString(cb.Ctx(), "bb")
		assert.Equal(t, "bb", m.Ss.Value())

		n.Map["sr"].SetValueString(cb.Ctx(), "cc")
		assert.Equal(t, "cc", m.Sr.Name)

		n.Map["i"].Map["a"].SetValueString(cb.Ctx(), "dd")
		assert.Equal(t, "dd", m.I.Face())

		n.Map["ajs"].Array[0].SetValueString(cb.Ctx(), "ee")
		assert.Equal(t, "ee", m.Ajs[0])

		n.Map["ajs"].Array[1].SetValueString(cb.Ctx(), "ff")
		assert.Equal(t, "ff", m.Ajs[1])

		n.Map["mjs"].Map["a"].SetValueString(cb.Ctx(), "gg")
		assert.Equal(t, "gg", m.Mjs["a"])

		n.Map["mjs"].Map["b"].SetValueString(cb.Ctx(), "hh")
		assert.Equal(t, "hh", m.Mjs["b"])
	}

	multi(t, n, n.Value.(*data.Multi), test)

}

func TestNode_SetValueNumber(t *testing.T) {

	cb, n := setup(t)

	test := func(t *testing.T, n *node.Node, m *data.Multi) {
		n.Map["jn"].SetValueNumber(cb.Ctx(), 11.0)
		assert.Equal(t, 11.0, m.Jn)

		n.Map["sn"].SetValueNumber(cb.Ctx(), 12.0)
		assert.Equal(t, 12.0, m.Sn.Value())

		n.Map["ajn"].Array[0].SetValueNumber(cb.Ctx(), 13.0)
		assert.Equal(t, 13.0, m.Ajn[0])

		n.Map["ajn"].Array[1].SetValueNumber(cb.Ctx(), 14.0)
		assert.Equal(t, 14.0, m.Ajn[1])

		n.Map["mjn"].Map["a"].SetValueNumber(cb.Ctx(), 15.0)
		assert.Equal(t, 15.0, m.Mjn["a"])

		n.Map["mjn"].Map["b"].SetValueNumber(cb.Ctx(), 16.0)
		assert.Equal(t, 16.0, m.Mjn["b"])
	}

	multi(t, n, n.Value.(*data.Multi), test)
}

func TestNode_SetValueBool(t *testing.T) {

	cb, n := setup(t)

	test := func(t *testing.T, n *node.Node, m *data.Multi) {
		n.Map["jb"].SetValueBool(cb.Ctx(), true)
		assert.Equal(t, true, m.Jb)

		n.Map["sb"].SetValueBool(cb.Ctx(), true)
		assert.Equal(t, true, m.Sb.Value())

		n.Map["ajb"].Array[0].SetValueBool(cb.Ctx(), true)
		assert.Equal(t, true, m.Ajb[0])

		n.Map["ajb"].Array[1].SetValueBool(cb.Ctx(), true)
		assert.Equal(t, true, m.Ajb[1])

		n.Map["mjb"].Map["a"].SetValueBool(cb.Ctx(), true)
		assert.Equal(t, true, m.Mjb["a"])

		n.Map["mjb"].Map["b"].SetValueBool(cb.Ctx(), true)
		assert.Equal(t, true, m.Mjb["b"])
	}

	multi(t, n, n.Value.(*data.Multi), test)
}

func TestNode_DeleteObjectChild(t *testing.T) {
	cb, n := setup(t)

	test := func(t *testing.T, n *node.Node, m *data.Multi) {
		assert.NoError(t, n.DeleteObjectChild(cb.Ctx(), "js"))
		assert.True(t, n.Map["js"].Missing)
		assert.Equal(t, "", m.Js)
		assert.NoError(t, n.DeleteObjectChild(cb.Ctx(), "ss"))
		assert.Nil(t, m.Ss)
		assert.NoError(t, n.DeleteObjectChild(cb.Ctx(), "sr"))
		assert.Nil(t, m.Sr)
		assert.NoError(t, n.DeleteObjectChild(cb.Ctx(), "jn"))
		assert.Equal(t, 0.0, m.Jn)
		assert.NoError(t, n.DeleteObjectChild(cb.Ctx(), "sn"))
		assert.Nil(t, m.Sn)
		assert.NoError(t, n.DeleteObjectChild(cb.Ctx(), "jb"))
		assert.Equal(t, false, m.Jb)
		assert.NoError(t, n.DeleteObjectChild(cb.Ctx(), "sb"))
		assert.Nil(t, m.Sb)
		assert.NoError(t, n.DeleteObjectChild(cb.Ctx(), "i"))
		assert.Nil(t, m.I)
		assert.NoError(t, n.DeleteObjectChild(cb.Ctx(), "ajs"))
		assert.Nil(t, m.Ajs)
		assert.NoError(t, n.DeleteObjectChild(cb.Ctx(), "mjs"))
		assert.Nil(t, m.Mjs)

	}
	multi(t, n, n.Value.(*data.Multi), test)
}

func TestNode_DeleteObjectChild2(t *testing.T) {
	cb, n := setup(t)
	err := n.Map["mjs"].DeleteObjectChild(cb.Ctx(), "a")
	assert.IsError(t, err, "BMUSITINTC")
	err = n.Map["ajs"].DeleteObjectChild(cb.Ctx(), "a")
	assert.IsError(t, err, "BMUSITINTC")
}

func TestNode_DeleteMapChild(t *testing.T) {
	_, n := setup(t)
	test := func(t *testing.T, n *node.Node, m *data.Multi) {
		assert.NoError(t, n.Map["mjs"].DeleteMapChild("a"))
		assert.Equal(t, 1, len(n.Map["mjs"].Map))
		assert.Equal(t, 1, len(m.Mjs))
		assert.NoError(t, n.Map["mjs"].DeleteMapChild("b"))
		assert.Equal(t, 0, len(n.Map["mjs"].Map))
		assert.Equal(t, 0, len(m.Mjs))
		assert.NoError(t, n.Map["mm"].DeleteMapChild("b"))
		assert.Equal(t, 1, len(n.Map["mm"].Map))
		assert.Equal(t, 1, len(m.Mm))
	}
	multi(t, n, n.Value.(*data.Multi), test)
}

func TestNode_DeleteMapChild2(t *testing.T) {
	_, n := setup(t)
	err := n.DeleteMapChild("js")
	assert.IsError(t, err, "ACRGPCPPFK")
	err = n.Map["ajs"].DeleteMapChild("a")
	assert.IsError(t, err, "ACRGPCPPFK")
}

func TestNode_DeleteArrayChild(t *testing.T) {
	_, n := setup(t)
	test := func(t *testing.T, n *node.Node, m *data.Multi) {
		assert.NoError(t, n.Map["ajs"].DeleteArrayChild(1))
		assert.Equal(t, 3, len(n.Map["ajs"].Array))
		assert.Equal(t, 0, n.Map["ajs"].Array[0].Index)
		assert.Equal(t, 1, n.Map["ajs"].Array[1].Index)
		assert.Equal(t, 2, n.Map["ajs"].Array[2].Index)
		assert.Equal(t, []string{"ajs0", "ajs2", "ajs3"}, m.Ajs)

		assert.NoError(t, n.Map["ajs"].DeleteArrayChild(2))
		assert.Equal(t, 2, len(n.Map["ajs"].Array))
		assert.Equal(t, 0, n.Map["ajs"].Array[0].Index)
		assert.Equal(t, 1, n.Map["ajs"].Array[1].Index)
		assert.Equal(t, []string{"ajs0", "ajs2"}, m.Ajs)

		assert.NoError(t, n.Map["ajs"].DeleteArrayChild(0))
		assert.Equal(t, 1, len(n.Map["ajs"].Array))
		assert.Equal(t, 0, n.Map["ajs"].Array[0].Index)
		assert.Equal(t, []string{"ajs2"}, m.Ajs)

		assert.NoError(t, n.Map["ajs"].DeleteArrayChild(0))
		assert.Equal(t, 0, len(n.Map["ajs"].Array))
		assert.Equal(t, []string{}, m.Ajs)
	}
	multi(t, n, n.Value.(*data.Multi), test)
}

func TestNode_DeleteArrayChild2(t *testing.T) {
	_, n := setup(t)
	err := n.DeleteArrayChild(0)
	assert.IsError(t, err, "NFVEWWCSMV")
	err = n.Map["mjs"].DeleteArrayChild(0)
	assert.IsError(t, err, "NFVEWWCSMV")
}

func TestNode_ReorderArrayChild(t *testing.T) {
	_, n := setup(t)
	test := func(t *testing.T, n *node.Node, m *data.Multi) {
		assert.NoError(t, n.Map["ajs"].ReorderArrayChild(0, 1))
		assert.Equal(t, 4, len(n.Map["ajs"].Array))
		assert.Equal(t, 0, n.Map["ajs"].Array[0].Index)
		assert.Equal(t, 1, n.Map["ajs"].Array[1].Index)
		assert.Equal(t, 2, n.Map["ajs"].Array[2].Index)
		assert.Equal(t, 3, n.Map["ajs"].Array[3].Index)
		assert.Equal(t, "ajs1", n.Map["ajs"].Array[0].ValueString)
		assert.Equal(t, "ajs0", n.Map["ajs"].Array[1].ValueString)
		assert.Equal(t, "ajs2", n.Map["ajs"].Array[2].ValueString)
		assert.Equal(t, "ajs3", n.Map["ajs"].Array[3].ValueString)
		assert.Equal(t, []string{"ajs1", "ajs0", "ajs2", "ajs3"}, m.Ajs)
	}
	multi(t, n, n.Value.(*data.Multi), test)
}

func TestNode_ReorderArrayChild2(t *testing.T) {
	_, n := setup(t)
	err := n.ReorderArrayChild(0, 1)
	assert.IsError(t, err, "MHEXGBUQOL")
	err = n.Map["mjs"].ReorderArrayChild(0, 1)
	assert.IsError(t, err, "MHEXGBUQOL")
}

func newNonEmptyNode() *node.Node {
	n := node.NewNode()
	n.Parent = node.NewNode()
	n.Array = []*node.Node{node.NewNode(), node.NewNode()}
	n.Map = map[string]*node.Node{"a": node.NewNode(), "b": node.NewNode()}
	n.Key = "a"
	n.Index = 1
	n.Origin = system.NewReference("a", "b")
	n.ValueString = "a"
	n.ValueNumber = 2.0
	n.ValueBool = true
	n.Value = "a"
	n.Val = reflect.ValueOf("a")
	n.Null = false
	n.Missing = false
	n.Rule = &system.RuleWrapper{}
	n.Type = &system.Type{}
	n.JsonType = json.J_OBJECT
	return n
}

func TestNode_InitialiseRoot(t *testing.T) {
	n := newNonEmptyNode()
	n.InitialiseRoot()
	assert.Nil(t, n.Parent)
	assert.Equal(t, 0, len(n.Array))
	assert.Equal(t, 0, len(n.Map))
	assert.Equal(t, "", n.Key)
	assert.Equal(t, -1, n.Index)
	assert.Nil(t, n.Origin)
	assert.Equal(t, "", n.ValueString)
	assert.Equal(t, 0.0, n.ValueNumber)
	assert.Equal(t, false, n.ValueBool)
	assert.Nil(t, n.Value)
	assert.Equal(t, reflect.Value{}, n.Val)
	assert.True(t, n.Null)
	assert.True(t, n.Missing)
	assert.Nil(t, n.Rule)
	assert.Nil(t, n.Type)
	assert.Equal(t, json.J_NULL, n.JsonType)
}

func TestNode_InitialiseArrayChild(t *testing.T) {
	cb, n := setup(t)
	test := func(t *testing.T, n *node.Node, m *data.Multi) {
		c := node.NewNode()
		assert.NoError(t, c.AddToArray(cb.Ctx(), n.Map["ajs"], 4, false))
		assert.Equal(t, 5, len(n.Map["ajs"].Array))
		assert.True(t, n.Map["ajs"].Array[4].Null)
		assert.False(t, n.Map["ajs"].Array[4].Missing)
		assert.Equal(t, c, n.Map["ajs"].Array[4])
		assert.Equal(t, []string{"ajs0", "ajs1", "ajs2", "ajs3", ""}, m.Ajs)

		c1 := node.NewNode()
		assert.NoError(t, c1.AddToArray(cb.Ctx(), n.Map["ass"], 4, false))
		assert.Equal(t, 5, len(m.Ass))
		assert.Nil(t, m.Ass[4])

		c2 := node.NewNode()
		assert.NoError(t, c2.AddToArray(cb.Ctx(), n.Map["ajs"], 0, false))
		assert.Equal(t, 5, len(n.Map["ajs"].Array))
		assert.Equal(t, []string{"", "ajs1", "ajs2", "ajs3", ""}, m.Ajs)

		c3 := node.NewNode()
		assert.NoError(t, c3.AddToArray(cb.Ctx(), n.Map["ass"], 0, false))
		assert.Equal(t, 5, len(m.Ass))
		assert.Nil(t, m.Ass[0])

		c4 := node.NewNode()
		assert.NoError(t, c4.AddToArray(cb.Ctx(), n.Map["am"], 2, false))
		assert.Equal(t, 3, len(m.Am))
		assert.Nil(t, m.Am[2])
	}
	multi(t, n, n.Value.(*data.Multi), test)

}

func TestNode_InitialiseArrayChild2(t *testing.T) {
	cb, n := setup(t)
	c := node.NewNode()
	err := c.AddToArray(cb.Ctx(), n.Map["ajs"], 5, false)
	assert.IsError(t, err, "GHJIDXABLL")

	// kludge the Val to be empty
	n.Map["ajs"].Val.Set(reflect.ValueOf([]string{}))
	err = c.AddToArray(cb.Ctx(), n.Map["ajs"], 1, false)
	assert.IsError(t, err, "YFKMXFUPHY")

}

func TestNode_InitialiseMapChild(t *testing.T) {
	cb, n := setup(t)
	test := func(t *testing.T, n *node.Node, m *data.Multi) {

		c := node.NewNode()
		assert.NoError(t, c.AddToMap(cb.Ctx(), n.Map["mjs"], "c", false))
		assert.Equal(t, 3, len(n.Map["mjs"].Map))
		assert.True(t, n.Map["mjs"].Map["c"].Null)
		assert.False(t, n.Map["mjs"].Map["c"].Missing)
		assert.Equal(t, c, n.Map["mjs"].Map["c"])
		assert.Equal(t, "", m.Mjs["c"])

		c1 := node.NewNode()
		assert.NoError(t, c1.AddToMap(cb.Ctx(), n.Map["mss"], "c", false))
		assert.Equal(t, 3, len(n.Map["mss"].Map))
		assert.Nil(t, m.Mss["c"])

		c2 := node.NewNode()
		assert.NoError(t, c2.AddToMap(cb.Ctx(), n.Map["mm"], "c", false))
		assert.Equal(t, 3, len(m.Mm))
		assert.Nil(t, m.Mm["c"])
	}
	multi(t, n, n.Value.(*data.Multi), test)

}

func TestNode_SetValueZero(t *testing.T) {
	cb, n := setup(t)
	test := func(t *testing.T, n *node.Node, m *data.Multi) {

		sstring, ok := system.GetTypeFromCache(cb.Ctx(), "kego.io/system", "string")
		assert.True(t, ok)
		facea, ok := system.GetTypeFromCache(cb.Ctx(), "kego.io/tests/data", "facea")
		assert.True(t, ok)
		faceb, ok := system.GetTypeFromCache(cb.Ctx(), "kego.io/tests/data", "faceb")
		assert.True(t, ok)

		c1 := node.NewNode()
		assert.NoError(t, c1.AddToMap(cb.Ctx(), n.Map["mss"], "c", false))
		assert.NoError(t, c1.SetValueZero(cb.Ctx(), true, nil))
		assert.False(t, n.Map["mss"].Map["c"].Missing)
		assert.True(t, n.Map["mss"].Map["c"].Null)
		assert.Nil(t, m.Mss["c"])

		c1a := node.NewNode()
		assert.NoError(t, c1a.AddToMap(cb.Ctx(), n.Map["mss"], "d", false))
		assert.NoError(t, c1a.SetValueZero(cb.Ctx(), false, nil))
		assert.False(t, n.Map["mss"].Map["d"].Missing)
		assert.False(t, n.Map["mss"].Map["d"].Null)
		assert.NotNil(t, m.Mss["d"])
		assert.Equal(t, "", m.Mss["d"].Value())

		c2 := node.NewNode()
		assert.NoError(t, c2.AddToMap(cb.Ctx(), n.Map["mm"], "c", false))
		assert.NoError(t, c2.SetValueZero(cb.Ctx(), true, nil))
		assert.Nil(t, m.Mm["c"])

		c2a := node.NewNode()
		assert.NoError(t, c2a.AddToMap(cb.Ctx(), n.Map["mm"], "d", false))
		assert.NoError(t, c2a.SetValueZero(cb.Ctx(), false, nil))
		assert.NotNil(t, m.Mm["d"])
		assert.Equal(t, "", m.Mm["d"].Js)

		assert.NoError(t, n.Map["i"].SetValueZero(cb.Ctx(), true, faceb))
		assert.Nil(t, m.I)
		assert.IsType(t, &data.Faceb{}, m.I)
		assert.NoError(t, n.Map["nri"].SetValueZero(cb.Ctx(), false, facea))
		assert.NotNil(t, m.Nri)
		assert.IsType(t, &data.Facea{}, m.Nri)

		c3 := node.NewNode()
		assert.NoError(t, c3.AddToArray(cb.Ctx(), n.Map["ass"], 4, false))
		assert.NoError(t, c3.SetValueZero(cb.Ctx(), true, nil))
		assert.Nil(t, m.Ass[4])

		c3a := node.NewNode()
		assert.NoError(t, c3a.AddToArray(cb.Ctx(), n.Map["ass"], 5, false))
		assert.NoError(t, c3a.SetValueZero(cb.Ctx(), false, nil))
		assert.NotNil(t, m.Ass[5])

		c4 := node.NewNode()
		assert.NoError(t, c4.AddToArray(cb.Ctx(), n.Map["am"], 2, false))
		assert.NoError(t, c4.SetValueZero(cb.Ctx(), true, nil))
		assert.Nil(t, m.Am[2])

		c4a := node.NewNode()
		assert.NoError(t, c4a.AddToArray(cb.Ctx(), n.Map["am"], 3, false))
		assert.NoError(t, c4a.SetValueZero(cb.Ctx(), false, nil))
		assert.NotNil(t, m.Am[3])

		c5 := node.NewNode()
		assert.NoError(t, c5.AddToArray(cb.Ctx(), n.Map["anri"], 3, false))
		assert.NoError(t, c5.SetValueZero(cb.Ctx(), true, sstring))
		assert.Nil(t, m.Anri[3])
		assert.IsType(t, system.NewString(""), m.Anri[3])

		c5a := node.NewNode()
		assert.NoError(t, c5a.AddToArray(cb.Ctx(), n.Map["anri"], 4, false))
		assert.NoError(t, c5a.SetValueZero(cb.Ctx(), false, sstring))
		assert.NotNil(t, m.Anri[4])
		assert.IsType(t, system.NewString(""), m.Anri[4])

		c6 := node.NewNode()
		assert.NoError(t, c6.AddToArray(cb.Ctx(), n.Map["anri"], 5, false))
		assert.NoError(t, c6.SetValueZero(cb.Ctx(), true, facea))
		assert.Nil(t, m.Anri[5])
		assert.IsType(t, &data.Facea{}, m.Anri[5])

		c6a := node.NewNode()
		assert.NoError(t, c6a.AddToArray(cb.Ctx(), n.Map["anri"], 6, false))
		assert.NoError(t, c6a.SetValueZero(cb.Ctx(), false, facea))
		assert.NotNil(t, m.Anri[6])
		assert.IsType(t, &data.Facea{}, m.Anri[6])

		c7 := node.NewNode()
		assert.NoError(t, c7.AddToMap(cb.Ctx(), n.Map["mnri"], "d", false))
		assert.NoError(t, c7.SetValueZero(cb.Ctx(), true, sstring))
		assert.Nil(t, m.Mnri["d"])
		assert.IsType(t, system.NewString(""), m.Mnri["d"])

		c7a := node.NewNode()
		assert.NoError(t, c7a.AddToMap(cb.Ctx(), n.Map["mnri"], "e", false))
		assert.NoError(t, c7a.SetValueZero(cb.Ctx(), false, sstring))
		assert.NotNil(t, m.Mnri["e"])
		assert.IsType(t, system.NewString(""), m.Mnri["e"])

		c8 := node.NewNode()
		assert.NoError(t, c8.AddToMap(cb.Ctx(), n.Map["mnri"], "f", false))
		assert.NoError(t, c8.SetValueZero(cb.Ctx(), true, facea))
		assert.Nil(t, m.Mnri["f"])
		assert.IsType(t, &data.Facea{}, m.Mnri["f"])

		c8a := node.NewNode()
		assert.NoError(t, c8a.AddToMap(cb.Ctx(), n.Map["mnri"], "g", false))
		assert.NoError(t, c8a.SetValueZero(cb.Ctx(), false, facea))
		assert.NotNil(t, m.Mnri["g"])
		assert.IsType(t, &data.Facea{}, m.Mnri["g"])
	}
	multi(t, n, n.Value.(*data.Multi), test)

}

func TestNode_SetValueZero2(t *testing.T) {
	cb, n := empty(t)
	test := func(t *testing.T, n *node.Node, m *data.Multi) {
		sstring, ok := system.GetTypeFromCache(cb.Ctx(), "kego.io/system", "string")
		assert.True(t, ok)
		snumber, ok := system.GetTypeFromCache(cb.Ctx(), "kego.io/system", "number")
		assert.True(t, ok)

		assert.NoError(t, n.Map["ss"].SetValueZero(cb.Ctx(), true, sstring))
		assert.Nil(t, m.Ss)

		assert.NoError(t, n.Map["sn"].SetValueZero(cb.Ctx(), false, snumber))
		assert.NotNil(t, m.Sn)

		assert.NoError(t, n.Map["mnri"].SetValueZero(cb.Ctx(), true, nil))
		assert.Nil(t, m.Mnri)

		assert.NoError(t, n.Map["mi"].SetValueZero(cb.Ctx(), false, nil))
		assert.NotNil(t, m.Mi)
		assert.Equal(t, 0, len(m.Mi))

		assert.NoError(t, n.Map["anri"].SetValueZero(cb.Ctx(), true, nil))
		assert.Nil(t, m.Anri)

		assert.NoError(t, n.Map["ai"].SetValueZero(cb.Ctx(), false, nil))
		assert.NotNil(t, m.Ai)
		assert.Equal(t, 0, len(m.Ai))

	}
	multi(t, n, n.Value.(*data.Multi), test)

}

func TestNode_SetValueZero3(t *testing.T) {
	cb, n := empty(t)
	f, ok := system.GetTypeFromCache(cb.Ctx(), "kego.io/tests/data", "face")
	assert.True(t, ok)
	err := n.SetValueZero(cb.Ctx(), true, f)
	assert.HasError(t, err, "VHOSYBMDQL")
}

func TestNode_SetValueUnpack(t *testing.T) {
	cb, n := empty(t)
	err := n.SetValueUnpack(cb.Ctx(), json.PackString(`"a"`))
	assert.HasError(t, err, "VEPLUIJXSN")
}

type lab struct{}

func (l *lab) Label(ctx context.Context) string {
	return "b"
}

func TestNode_Label(t *testing.T) {
	n := node.NewNode()
	n.Parent = node.NewNode()
	n.Key = ""
	n.Value = ""
	n.Index = -1

	assert.Equal(t, "(?)", n.Label(context.Background()))

	n.Index = 1
	assert.Equal(t, "1", n.Label(context.Background()))

	n.Value = &lab{}
	assert.Equal(t, "b", n.Label(context.Background()))

	n.Key = "a"
	assert.Equal(t, "a", n.Label(context.Background()))

	n.Parent = nil
	assert.Equal(t, "root", n.Label(context.Background()))

	n = nil
	assert.Equal(t, "(nil)", n.Label(context.Background()))

}

func TestNode_Path(t *testing.T) {
	r := node.NewNode()
	a := node.NewNode()
	b := node.NewNode()
	a.Key = "a"
	b.Key = "b"
	a.Parent = r
	b.Parent = a
	assert.Equal(t, "root/a/b", b.Path())
}

func TestNode_Root(t *testing.T) {
	r := node.NewNode()
	a := node.NewNode()
	b := node.NewNode()
	a.Parent = r
	b.Parent = a
	assert.Equal(t, r, a.Root())
	assert.Equal(t, r, b.Root())
	n := (*node.Node)(nil)
	assert.Nil(t, n.Root())
}

func TestNode_DisplayType(t *testing.T) {
	cb, n := setup(t)

	dt, err := n.DisplayType(cb.Ctx())
	assert.NoError(t, err)
	assert.Equal(t, "multi", dt)

	dt, err = n.Map["ai"].DisplayType(cb.Ctx())
	assert.NoError(t, err)
	assert.Equal(t, "null", dt)

	dt, err = n.Map["ajs"].DisplayType(cb.Ctx())
	assert.NoError(t, err)
	assert.Equal(t, "[]json:string", dt)
}