package utils

import (
	"github.com/quill-sql/quill-go/pkg/assets"
)

// Function to get a type name by Oid
func ConvertTypeToPostgres(dataTypeID int) string {
	for _, typ := range assets.PGTypes {
		if dataTypeID == typ.OID {
			return typ.Typname
		}
	}
	return "varchar"
}
