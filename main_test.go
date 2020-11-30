package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_migrate(t *testing.T) {
	os.Setenv("TEST_PATTERN", "*")
	testDataDir := "test_data"
	dir := t.TempDir()

	expected, err := ioutil.ReadFile(filepath.Join(testDataDir, "v1_to_v4beta1_expected"))
	assert.NoError(t, err)

	testData, err := ioutil.ReadFile(filepath.Join(testDataDir, "v1_to_v4beta1"))
	assert.NoError(t, err)

	err = ioutil.WriteFile(filepath.Join(dir, "v1_to_v4beta1"), testData, 0666)
	assert.NoError(t, err)

	err = migrate(dir)
	assert.NoError(t, err)

	result, err := ioutil.ReadFile(filepath.Join(dir, "v1_to_v4beta1"))
	assert.NoError(t, err)

	assert.True(t, reflect.DeepEqual(result, expected), fmt.Sprintf("migrated file does not match: %v, want: %v", string(result), string(expected)))

}

func Test_migrateRequirementsFile(t *testing.T) {

	testDataDir := "test_data"
	dir := t.TempDir()

	expected, err := ioutil.ReadFile(filepath.Join(testDataDir, "jx-requirements.yml-expected"))
	assert.NoError(t, err)

	testData, err := ioutil.ReadFile(filepath.Join(testDataDir, "jx-requirements.yml"))
	assert.NoError(t, err)

	err = ioutil.WriteFile(filepath.Join(dir, "jx-requirements.yml"), testData, 0666)
	assert.NoError(t, err)

	err = migrate(dir)
	assert.NoError(t, err)

	result, err := ioutil.ReadFile(filepath.Join(dir, "jx-requirements.yml"))
	assert.NoError(t, err)

	assert.True(t, reflect.DeepEqual(result, expected), fmt.Sprintf("migrated file does not match: %v, want: %v", string(result), string(expected)))

}
