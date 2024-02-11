// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package main

import "encoding/json"
import "fmt"

type RunnableSchemaJson struct {
	// In corresponds to the JSON schema field "in".
	In *string `json:"in,omitempty" yaml:"in,omitempty" mapstructure:"in,omitempty"`

	// Name corresponds to the JSON schema field "name".
	Name *string `json:"name,omitempty" yaml:"name,omitempty" mapstructure:"name,omitempty"`

	// Notify corresponds to the JSON schema field "notify".
	Notify RunnableSchemaJsonNotify `json:"notify,omitempty" yaml:"notify,omitempty" mapstructure:"notify,omitempty"`

	// Out corresponds to the JSON schema field "out".
	Out *string `json:"out,omitempty" yaml:"out,omitempty" mapstructure:"out,omitempty"`

	// Post corresponds to the JSON schema field "post".
	Post *string `json:"post,omitempty" yaml:"post,omitempty" mapstructure:"post,omitempty"`

	// Pre corresponds to the JSON schema field "pre".
	Pre *string `json:"pre,omitempty" yaml:"pre,omitempty" mapstructure:"pre,omitempty"`

	// Run corresponds to the JSON schema field "run".
	Run string `json:"run" yaml:"run" mapstructure:"run"`
}

type RunnableSchemaJsonNotify map[string]interface{}

// UnmarshalJSON implements json.Unmarshaler.
func (j *RunnableSchemaJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["run"]; !ok || v == nil {
		return fmt.Errorf("field run in RunnableSchemaJson: required")
	}
	type Plain RunnableSchemaJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = RunnableSchemaJson(plain)
	return nil
}