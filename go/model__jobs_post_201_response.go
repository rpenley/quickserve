/*
 * Batch Processing Service API
 *
 * An API for managing batch processing jobs.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type JobsPost201Response struct {

	// Unique identifier of the job.
	JobId string `json:"jobId,omitempty"`

	// Current status of the job.
	Status string `json:"status,omitempty"`
}
