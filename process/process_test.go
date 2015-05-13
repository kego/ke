package process

import (
	"io/ioutil"
	"testing"

	"kego.io/system"
	"kego.io/uerr"

	"os"

	"path/filepath"

	"github.com/stretchr/testify/assert"
)

func TestScan(t *testing.T) {

	test := `{
	"description": "a",
	"type": "system:type",
	"id": "b",
	"properties": {
		"c": {
			"type": "system:property",
			"item": {
				"type": "system:@string"
			}
		}
	}
}`

	d, err := ioutil.TempDir("", "")
	assert.NoError(t, err)
	defer os.Remove(d)

	f := filepath.Join(d, "a.json")
	err = ioutil.WriteFile(f, []byte(test), 0644)
	assert.NoError(t, err)
	defer os.Remove(f)

	err = Scan(d, "d.e/f", map[string]string{})
	assert.NoError(t, err)
	ty, ok := system.GetType("d.e/f:b")
	assert.True(t, ok)
	assert.Equal(t, "a", ty.Description)

}

func TestScan_rule(t *testing.T) {

	test := `{
	"description": "a",
	"type": "system:type",
	"id": "b",
	"properties": {
		"c": {
			"type": "system:property",
			"item": {
				"type": "system:@string"
			}
		}
	},
	"rule": {
		"description": "d",
		"type": "system:type",
		"id": "@b",
		"is": ["system:rule"],
		"properties": {
			"e": {
				"description": "f",
				"type": "system:property",
				"item": {
					"type": "system:@string"
				}
			}
		}
	}
}`

	d, err := ioutil.TempDir("", "")
	assert.NoError(t, err)
	defer os.Remove(d)

	f := filepath.Join(d, "a.json")
	err = ioutil.WriteFile(f, []byte(test), 0644)
	assert.NoError(t, err)
	defer os.Remove(f)

	err = Scan(d, "d.e/f", map[string]string{})
	assert.NoError(t, err)
	ty, ok := system.GetType("d.e/f:b")
	assert.True(t, ok)
	assert.Equal(t, "@b", ty.Rule.Id)

}

func TestScan_errors(t *testing.T) {

	err := Scan("/this-folder-doesnt-exist", "", map[string]string{})
	uerr.Assert(t, err, "XHHQSAVCKK")

	d, err := ioutil.TempDir("", "")
	assert.NoError(t, err)
	defer os.Remove(d)

	f := filepath.Join(d, "a.json")
	err = ioutil.WriteFile(f, []byte("foo"), 0644)
	assert.NoError(t, err)
	defer os.Remove(f)

	err = Scan(d, "d.e/f", map[string]string{})
	uerr.Assert(t, err, "XHHQSAVCKK")
	uerr.Stack(t, err, "DHTURNTIXE")

}
