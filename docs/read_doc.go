package docs

// ReadDoc returns the Swagger specification in JSON format.
// This provides a simple, valid doc.json to avoid template execution errors.
func ReadDoc() string {
    return `{"swagger":"2.0","info":{"description":"Project and Task Management API","title":"Mini Jira Backend","version":"1.0"},"host":"localhost:8080","basePath":"/","paths":{}}`
}
