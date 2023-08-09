package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/quill-sql/quill-go"
)

const OrganizationIDContextKey = "ORGANIZATION_ID"

func main() {
	// TODO: ADD TO DOCS: client := quill.NewClient(quill.ClientParams{PrivateKey: os.Getenv("QUILL_PRIVATE_KEY"), DatabaseConnectionString: os.Getenv("POSTGRES_READ")})
	client := quill.NewClient(quill.QuillClientParams{PrivateKey: "pk_cb3c270a0eb908b06e04e89fd9436f8f3d8c1a85a38276239ae32cbe6e0d98d5", DatabaseConnectionString: "postgres://postgres:SeeQuill99**@db.fnaxkqsjsnisbokmyqtb.supabase.co:6543/postgres"})

	// Add an endpoint
	http.HandleFunc("/quill", func(w http.ResponseWriter, r *http.Request) {
		// fetch organizationID from your existing auth middleware
		// organizationID, _ := r.Context().Value(OrganizationIDContextKey).(string)
		var organizationID = "2"

		// Convert json body.metadata to RequestMetadata
		body := &quill.RequestBody{}
		err := json.NewDecoder(r.Body).Decode(body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		result, err := client.Query(organizationID, body.Metadata)
		if err != nil {
			fmt.Fprint(w, err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	})
	// var organizationID = "2"
	// quillClient := quill.NewClient(quill.QuillClientParams{PrivateKey: os.Getenv("QUILL_PRIVATE_KEY"), DatabaseConnectionString: os.Getenv("POSTGRES_READ")})
	// http.HandleFunc("/quill", func(w http.ResponseWriter, r *http.Request) {
	// 	// assuming organizationID can be fetched via auth middleware
	// 	organizationID, _ := r.Context().Value(OrganizationIDContextKey).(string)
	// 	result, err := quillClient.Query(organizationID, r.Body)
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// 	w.Header().Set("Content-Type", "application/json")
	// 	json.NewEncoder(w).Encode(result)
	// })

	// Start the server
	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
