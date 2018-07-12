package bootnodes

import (
	"net/http"
	"encoding/json"
	"log"
	"github.com/trust-net/go-sdb-node/api"
)

type bootNodeRequest struct {
	ID 		string 			`json:"id,omitempty"`
}

type bootNodeResponse struct {
	ID 		string 			`json:"id,omitempty"`
	Url 		string 			`json:"url,omitempty"`
}

var running map[string]struct{}

func init() {
	running = make(map[string]struct{})
}

func PostBootnodesStart (r *http.Request) (api.ApiResponse, api.Error) {
	var req bootNodeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || len(req.ID) == 0 {
		log.Printf("ERROR: Failed to parse request body '%s'", err)
		return nil, api.ApiError(api.ERR_BAD_REQUEST,"bad request")
	} else {
		log.Printf("request to start bootnode for network: %s", req.ID)
		if _, found := running[req.ID]; !found {
			// one fake success response
			running[req.ID] = struct{}{}
			return bootNodeResponse{ID: req.ID, Url: "fake url: " + req.ID}, nil
		} else {
			// there is already a boot node running
			return nil, api.ApiError(api.ERR_CONFLICT, "boot node already exists")		
		}
	}
}

func PostBootnodesStop (r *http.Request) (api.ApiResponse, api.Error) {
	var req bootNodeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || len(req.ID) == 0 {
		log.Printf("ERROR: Failed to parse request body '%s'", err)
		return nil, api.ApiError(api.ERR_BAD_REQUEST,"bad request")
	} else {
		log.Printf("request to stop bootnode for network: %s", req.ID)
		if _, found := running[req.ID]; found {
			// one fake success response
			delete(running, req.ID)
			return nil, nil
		} else {
			// there was no already running boot node
			return nil, api.ApiError(api.ERR_NOT_FOUND, "boot node does not exists")		
		}
	}
}
