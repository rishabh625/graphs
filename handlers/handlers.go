package handlers

import (
	"encoding/json"
	"graphs/datastruct"
	"net/http"
)

//GraphInput Takes Graphs and vertex weight as input
func GraphInput(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "POST":
		var reqobj datastruct.InputGraph
		var resp datastruct.APIResponse
		if err := json.NewDecoder(r.Body).Decode(&reqobj); err != nil {
			//resp = &datastruct.APIResponse{}
			w.WriteHeader(http.StatusBadRequest)
		} else {
			graph := datastruct.CreateGraph(reqobj)
			resp = datastruct.GetShortestPath(reqobj.From, reqobj.To, graph)
			w.WriteHeader(http.StatusOK)
			byteresp, _ := json.Marshal(resp)
			w.Write(byteresp)
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		response := datastruct.APIResponse{}
		byteresp, _ := json.Marshal(response)
		w.Write(byteresp)
	}
}
