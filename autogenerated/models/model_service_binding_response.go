/*
 * Open Service Broker API
 *
 * The Open Service Broker API defines an HTTP(S) interface between Platforms and Service Brokers.
 *
 * API version: master - might contain changes that are not yet released
 * Contact: open-service-broker-api@googlegroups.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type ServiceBindingResponse struct {

	Metadata ServiceBindingMetadata `json:"metadata,omitempty"`

	Credentials map[string]interface{} `json:"credentials,omitempty"`

	SyslogDrainUrl string `json:"syslog_drain_url,omitempty"`

	RouteServiceUrl string `json:"route_service_url,omitempty"`

	VolumeMounts []ServiceBindingVolumeMount `json:"volume_mounts,omitempty"`

	Endpoints []ServiceBindingEndpoint `json:"endpoints,omitempty"`
}
