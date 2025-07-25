package target_groups

import (
	"github.com/nttcom/eclcloud/v3"
	"github.com/nttcom/eclcloud/v3/pagination"
)

type commonResult struct {
	eclcloud.Result
}

// CreateResult represents the result of a Create operation.
// Call its Extract method to interpret it as a TargetGroup.
type CreateResult struct {
	commonResult
}

// ShowResult represents the result of a Show operation.
// Call its Extract method to interpret it as a TargetGroup.
type ShowResult struct {
	commonResult
}

// UpdateResult represents the result of a Update operation.
// Call its Extract method to interpret it as a TargetGroup.
type UpdateResult struct {
	commonResult
}

// DeleteResult represents the result of a Delete operation.
// Call its ExtractErr method to determine if the request succeeded or failed.
type DeleteResult struct {
	eclcloud.ErrResult
}

// CreateStagedResult represents the result of a CreateStaged operation.
// Call its Extract method to interpret it as a TargetGroup.
type CreateStagedResult struct {
	commonResult
}

// ShowStagedResult represents the result of a ShowStaged operation.
// Call its Extract method to interpret it as a TargetGroup.
type ShowStagedResult struct {
	commonResult
}

// UpdateStagedResult represents the result of a UpdateStaged operation.
// Call its Extract method to interpret it as a TargetGroup.
type UpdateStagedResult struct {
	commonResult
}

// CancelStagedResult represents the result of a CancelStaged operation.
// Call its ExtractErr method to determine if the request succeeded or failed.
type CancelStagedResult struct {
	eclcloud.ErrResult
}

// ConfigurationInResponse represents a configuration in a target group.
type ConfigurationInResponse struct {

	// - Members (real servers) of the target group
	Members []MemberInResponse `json:"members,omitempty"`
}

// MemberInResponse represents a member in a target group.
type MemberInResponse struct {

	// - IP address of the member (real server)
	IPAddress string `json:"ip_address"`

	// - Port number of the member (real server)
	Port int `json:"port"`

	// - Weight for the member (real server)
	// - If `policy.algorithm` is `"weighted-round-robin"` or `"weighted-least-connection"`, uses this parameter
	Weight int `json:"weight"`
}

// TargetGroup represents a target group.
type TargetGroup struct {

	// - ID of the target group
	ID string `json:"id"`

	// - Name of the target group
	Name string `json:"name"`

	// - Description of the target group
	Description string `json:"description"`

	// - Tags of the target group (JSON object format)
	Tags map[string]interface{} `json:"tags"`

	// - Configuration status of the target group
	//   - `"ACTIVE"`
	//     - There are no configurations of the target group that waiting to be applied
	//   - `"CREATE_STAGED"`
	//     - The target group has been added and waiting to be applied
	//   - `"UPDATE_STAGED"`
	//     - Changed configurations of the target group exists that waiting to be applied
	//   - `"DELETE_STAGED"`
	//     - The target group has been removed and waiting to be applied
	// - For detail, refer to the API reference appendix
	//     - https://sdpf.ntt.com/services/docs/managed-lb/service-descriptions/api_reference_appendix.html
	ConfigurationStatus string `json:"configuration_status"`

	// - Operation status of the load balancer which the target group belongs to
	//   - `"NONE"` :
	//     - There are no operations of the load balancer
	//     - The load balancer and related resources can be operated
	//   - `"PROCESSING"`
	//     - The latest operation of the load balancer is processing
	//     - The load balancer and related resources cannot be operated
	//   - `"COMPLETE"`
	//     - The latest operation of the load balancer has been succeeded
	//     - The load balancer and related resources can be operated
	//   - `"STUCK"`
	//     - The latest operation of the load balancer has been stopped
	//     - The operators will investigate the operation
	//     - The load balancer and related resources cannot be operated
	//   - `"ERROR"`
	//     - The latest operation of the load balancer has been failed
	//     - The operation was roll backed normally
	//     - The load balancer and related resources can be operated
	// - For detail, refer to the API reference appendix
	//     - https://sdpf.ntt.com/services/docs/managed-lb/service-descriptions/api_reference_appendix.html
	OperationStatus string `json:"operation_status"`

	// - ID of the load balancer which the target group belongs to
	LoadBalancerID string `json:"load_balancer_id"`

	// - ID of the owner tenant of the target group
	TenantID string `json:"tenant_id"`

	// - Members (real servers) of the target group
	Members []MemberInResponse `json:"members,omitempty"`

	// - Running configurations of the target group
	// - If `changes` is `true`, return object
	// - If current configuration does not exist, return `null`
	Current ConfigurationInResponse `json:"current,omitempty"`

	// - Added or changed configurations of the target group that waiting to be applied
	// - If `changes` is `true`, return object
	// - If staged configuration does not exist, return `null`
	Staged ConfigurationInResponse `json:"staged,omitempty"`
}

// ExtractInto interprets any commonResult as a target group, if possible.
func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "target_group")
}

// Extract is a function that accepts a result and extracts a TargetGroup resource.
func (r commonResult) Extract() (*TargetGroup, error) {
	var targetGroup TargetGroup

	err := r.ExtractInto(&targetGroup)

	return &targetGroup, err
}

// TargetGroupPage is the page returned by a pager when traversing over a collection of target group.
type TargetGroupPage struct {
	pagination.LinkedPageBase
}

// IsEmpty checks whether a TargetGroupPage struct is empty.
func (r TargetGroupPage) IsEmpty() (bool, error) {
	is, err := ExtractTargetGroups(r)

	return len(is) == 0, err
}

// ExtractTargetGroupsInto interprets the results of a single page from a List() call, producing a slice of target group entities.
func ExtractTargetGroupsInto(r pagination.Page, v interface{}) error {
	return r.(TargetGroupPage).Result.ExtractIntoSlicePtr(v, "target_groups")
}

// ExtractTargetGroups accepts a Page struct, specifically a NetworkPage struct, and extracts the elements into a slice of TargetGroup structs.
// In other words, a generic collection is mapped into a relevant slice.
func ExtractTargetGroups(r pagination.Page) ([]TargetGroup, error) {
	var s []TargetGroup

	err := ExtractTargetGroupsInto(r, &s)

	return s, err
}
