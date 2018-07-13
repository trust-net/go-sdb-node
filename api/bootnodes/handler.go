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

var running map[string]string

func init() {
	running = make(map[string]string)
}

func GetBootnodes (r *http.Request) (api.ApiResponse, api.Error) {
	log.Printf("request to get all bootnodes")
	list := make([]bootNodeResponse,0)
	for id,url := range running {
		list = append(list, bootNodeResponse{ID: id, Url: url})
	}
	return list, nil
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
			url := "boot://" + req.ID + ".fake.url"
			running[req.ID] = url
			return bootNodeResponse{ID: req.ID, Url: url}, nil
		} else {
			// there is already a boot node for specified network ID
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
