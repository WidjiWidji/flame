/*
 * Fledge REST API
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
	"wwwin-github.cisco.com/eti/fledge/pkg/util"
)

// DesignSchemaApiService is a service that implents the logic for the DesignSchemaApiServicer
// This service should implement the business logic for every endpoint for the DesignSchemaApi API.
// Include any external packages or services that will be required by this service.
type DesignSchemaApiService struct {
}

// NewDesignSchemaApiService creates a default api service
func NewDesignSchemaApiService() openapi.DesignSchemaApiServicer {
	return &DesignSchemaApiService{}
}

// GetDesignSchema - Get a design schema owned by user
func (s *DesignSchemaApiService) GetDesignSchema(ctx context.Context, user string, designId string,
	getType string, schemaId string) (openapi.ImplResponse, error) {
	//TODO input validation
	zap.S().Debugf("Get design schema details for user: %s | designId: %s | type: %s | schemaId: %s", user, designId, getType, schemaId)

	//create controller request
	uriMap := map[string]string{
		"user":     user,
		"designId": designId,
		"type":     getType,
		"schemaId": schemaId,
	}
	url := restapi.CreateURL(Host, Port, restapi.GetDesignSchemaEndPoint, uriMap)

	//send get request
	responseBody, err := restapi.HTTPGet(url)

	//response to the user
	if err != nil {
		return openapi.Response(http.StatusInternalServerError, nil), errors.New("get design schema details request failed")
	}
	var resp []openapi.DesignSchema
	err = util.ByteToStruct(responseBody, &resp)
	return openapi.Response(http.StatusOK, resp), err
}

// UpdateDesignSchema - Update a design schema
func (s *DesignSchemaApiService) UpdateDesignSchema(ctx context.Context, user string, designId string,
	designSchema openapi.DesignSchema) (openapi.ImplResponse, error) {
	//TODO input validation
	zap.S().Debugf("Update/insert design schema request received for designId: %v", designId)

	//create controller request
	uriMap := map[string]string{
		"user":     user,
		"designId": designId,
	}
	url := restapi.CreateURL(Host, Port, restapi.UpdateDesignSchemaEndPoint, uriMap)

	//send get request
	code, resp, err := restapi.HTTPPost(url, designSchema, "application/json")
	zap.S().Debugf("code: %d, resp: %s", code, string(resp))

	//response to the user
	if err != nil {
		return openapi.Response(http.StatusInternalServerError, nil), errors.New("error while updating/inserting design schema")
	}

	return openapi.Response(http.StatusOK, nil), err
}
