package types

// Request type for querying all resources
type QueryAllResourceRequest struct {
    Owner string `json:"owner"`
    Name  string `json:"name"`
}

// Response type for querying all resources
type QueryAllResourceResponse struct {
    Resources []Resource `json:"resources"`
}

