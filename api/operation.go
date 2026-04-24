package api

import (
	"fmt"
	"strings"
)

type Operation string

const (
	OperationGetItems      Operation = "GetItems"
	OperationGetVariations Operation = "GetVariations"
	OperationSearchItems   Operation = "SearchItems"
)

type ErrInvalidResource struct {
	Resource  Resource
	Operation Operation
}

func (e ErrInvalidResource) Error() string {
	return fmt.Sprintf("invalid resource %q for operation %q", e.Resource, e.Operation)
}

func (o Operation) Path() string {
	name := string(o)
	return strings.ToLower(name[:1]) + name[1:]
}

func (o Operation) Validate(resources []Resource) error {
	for _, resource := range resources {
		validOperations := resourceOperationsMap[resource]
		if !existsInOperations(o, validOperations) {
			return ErrInvalidResource{Resource: resource, Operation: o}
		}
	}

	return nil
}
