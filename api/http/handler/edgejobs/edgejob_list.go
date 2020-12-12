package edgejobs

import (
	"net/http"

	httperror "github.com/portainer/libhttp/error"
	"github.com/portainer/libhttp/response"
)

// GET request on /api/edge_jobs
func (handler *Handler) edgeJobList(w http.ResponseWriter, r *http.Request) *httperror.HandlerError {
	edgeJobs, err := handler.DataStore.EdgeJob().EdgeJobs()
	if err != nil {
		return &httperror.HandlerError{http.StatusInternalServerError, "Unable to retrieve Edge jobs from the database", err}
	}

	return response.JSON(w, edgeJobs)
}
