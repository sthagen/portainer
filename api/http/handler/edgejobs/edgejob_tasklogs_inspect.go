package edgejobs

import (
	"net/http"
	"strconv"

	httperror "github.com/portainer/libhttp/error"
	"github.com/portainer/libhttp/request"
	"github.com/portainer/libhttp/response"
)

type fileResponse struct {
	FileContent string `json:"FileContent"`
}

// GET request on /api/edge_jobs/:id/tasks/:taskID/logs
func (handler *Handler) edgeJobTaskLogsInspect(w http.ResponseWriter, r *http.Request) *httperror.HandlerError {
	edgeJobID, err := request.RetrieveNumericRouteVariableValue(r, "id")
	if err != nil {
		return &httperror.HandlerError{http.StatusBadRequest, "Invalid Edge job identifier route variable", err}
	}

	taskID, err := request.RetrieveNumericRouteVariableValue(r, "taskID")
	if err != nil {
		return &httperror.HandlerError{http.StatusBadRequest, "Invalid Task identifier route variable", err}
	}

	logFileContent, err := handler.FileService.GetEdgeJobTaskLogFileContent(strconv.Itoa(edgeJobID), strconv.Itoa(taskID))
	if err != nil {
		return &httperror.HandlerError{http.StatusInternalServerError, "Unable to retrieve log file from disk", err}
	}

	return response.JSON(w, &fileResponse{FileContent: string(logFileContent)})
}

// fmt.Sprintf("/tmp/edge_jobs/%s/logs_%s", edgeJobID, taskID)
