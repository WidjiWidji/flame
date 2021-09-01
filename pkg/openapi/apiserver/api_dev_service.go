/*
 * Job REST API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package apiserver

import (
	"context"
	"errors"
	"net/http"

	"go.uber.org/zap"

	"wwwin-github.cisco.com/eti/fledge/pkg/openapi"
	"wwwin-github.cisco.com/eti/fledge/pkg/restapi"
)

// DevApiService is a service that implents the logic for the DevApiServicer
// This service should implement the business logic for every endpoint for the DevApi API.
// Include any external packages or services that will be required by this service.
type DevApiService struct {
}

// NewDevApiService creates a default api service
func NewDevApiService() openapi.DevApiServicer {
	return &DevApiService{}
}

// JobNodes - Nodes information for the job
func (s *DevApiService) JobNodes(ctx context.Context, user string, jobNodes openapi.JobNodes) (openapi.ImplResponse, error) {
	zap.S().Debugf("Add nodes for user: %s | designId: %v", user, jobNodes.DesignId)

	//create controller request
	uriMap := map[string]string{
		"user": user,
	}
	url := restapi.CreateURL(Host, Port, restapi.JobNodesEndPoint, uriMap)

	//send get request
	_, _, err := restapi.HTTPPost(url, jobNodes, "application/json")

	//response to the user
	if err != nil {
		return openapi.Response(http.StatusInternalServerError, nil), errors.New("add nodes for job request failed")
	}
	return openapi.Response(http.StatusOK, nil), err
}

// UpdateJobNodes - Update or add new nodes information for the job
func (s *DevApiService) UpdateJobNodes(ctx context.Context, user string, jobNodes openapi.JobNodes) (openapi.ImplResponse, error) {
	zap.S().Debugf("Update/add new nodes for user: %s | designId: %v", user, jobNodes.DesignId)

	//create controller request
	uriMap := map[string]string{
		"user": user,
	}
	url := restapi.CreateURL(Host, Port, restapi.JobNodesEndPoint, uriMap)

	//send get request
	_, _, err := restapi.HTTPPut(url, jobNodes, "application/json")

	//response to the user
	if err != nil {
		return openapi.Response(http.StatusInternalServerError, nil), errors.New("update/add nodes for job request failed")
	}
	return openapi.Response(http.StatusOK, nil), err
}
