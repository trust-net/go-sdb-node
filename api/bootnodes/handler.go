package bootnodes

import (
	"net/http"
	"encoding/json"
	"log"
	"github.com/trust-net/go-sdb-node/api"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/p2p/discover"
)

type bootNodeRequest struct {
	ID 		string 			`json:"id,omitempty"`
}

type bootNodeResponse struct {
	ID 		string 			`json:"id,omitempty"`
	Url 		string 			`json:"url,omitempty"`
}

type BootNodesHandler struct {
	addr string
	running map[string]*discover.Table
}

func NewBootNodesHandler(addr string) *BootNodesHandler {
	return &BootNodesHandler{
		addr: addr,
		running: make(map[string]*discover.Table),
	}
}

func (h *BootNodesHandler) GetBootnodes(r *http.Request) (api.ApiResponse, api.Error) {
	log.Printf("request to get all bootnodes")
	list := make([]bootNodeResponse,0)
	for id,node := range h.running {
		list = append(list, bootNodeResponse{ID: id, Url: node.Self().String()})
	}
	return list, nil
}

func (h *BootNodesHandler) PostBootnodesStart(r *http.Request) (api.ApiResponse, api.Error) {
	var req bootNodeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || len(req.ID) == 0 {
		log.Printf("ERROR: Failed to parse request body '%s'", err)
		return nil, api.ApiError(api.ERR_BAD_REQUEST,"bad request")
	} else {
		log.Printf("%s -- request to start bootnode for network: %s", r.Host, req.ID)
		if _, found := h.running[req.ID]; !found {
			// start a new boot node
			if node, err := h.startNewNode(); err == nil {
				h.running[req.ID] = node
				return bootNodeResponse{ID: req.ID, Url: node.Self().String()}, nil
			} else {
				log.Printf("ERROR: Failed to create new boot node '%s'", err)
				return nil, api.ApiError(api.ERR_FAILED, "failed to create boot node")
			}
		} else {
			// there is already a boot node for specified network ID
			return nil, api.ApiError(api.ERR_CONFLICT, "boot node already exists")		
		}
	}
}

func (h *BootNodesHandler) startNewNode() (*discover.Table, error) {
	if nodeKey, err := crypto.GenerateKey(); err != nil {
		return nil, err
	} else {
		if table, err := discover.ListenUDP(nodeKey, h.addr + ":0", nil, "", nil); err != nil {
			return nil, err
		} else {
			return table, nil
		}
	}
}

func (h *BootNodesHandler) stopRunningNode(table *discover.Table) {
	if table != nil {
		table.Close()
	}
}

func (h *BootNodesHandler) PostBootnodesStop(r *http.Request) (api.ApiResponse, api.Error) {
	var req bootNodeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || len(req.ID) == 0 {
		log.Printf("ERROR: Failed to parse request body '%s'", err)
		return nil, api.ApiError(api.ERR_BAD_REQUEST,"bad request")
	} else {
		log.Printf("request to stop bootnode for network: %s", req.ID)
		if node, found := h.running[req.ID]; found {
			// stop running node
			h.stopRunningNode(node)
			delete(h.running, req.ID)
			return nil, nil
		} else {
			// there was no already running boot node
			return nil, api.ApiError(api.ERR_NOT_FOUND, "boot node does not exists")		
		}
	}
}
