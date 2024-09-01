package espnffgo

import (
	"encoding/json"
	"fmt"
	"testing"
)

// PrettyPrint JSON response
func prettyPrint(t testing.TB, data interface{}) {
	t.Helper()
	printData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		t.Errorf("Error pretty printing JSON data")
	}
	fmt.Println(string(printData))
}
