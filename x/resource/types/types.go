package types

import (
    "fmt"
)

// Resource defines the structure for the resource object
type Resource struct {
    Name        string `json:"name"`
    Description string `json:"description"`
    Owner       string `json:"owner"`
}

func (r *Resource) Reset() {
    *r = Resource{}
}

func (r *Resource) String() string {
    return fmt.Sprintf("Resource{Name: %s, Description: %s, Owner: %s}", r.Name, r.Description, r.Owner)
}

func (r *Resource) ProtoMessage() {}

var KeyPrefixResource = []byte("resource/")

