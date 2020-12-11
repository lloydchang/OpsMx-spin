/*
 * Spinnaker API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Url struct {
	UserInfo string `json:"userInfo,omitempty"`
	File string `json:"file,omitempty"`
	Query string `json:"query,omitempty"`
	Ref string `json:"ref,omitempty"`
	Path string `json:"path,omitempty"`
	DefaultPort int32 `json:"defaultPort,omitempty"`
	SerializedHashCode int32 `json:"serializedHashCode,omitempty"`
	DeserializedFields *UrlStreamHandler `json:"deserializedFields,omitempty"`
	Content *interface{} `json:"content,omitempty"`
	Protocol string `json:"protocol,omitempty"`
	Host string `json:"host,omitempty"`
	Authority string `json:"authority,omitempty"`
	Port int32 `json:"port,omitempty"`
}