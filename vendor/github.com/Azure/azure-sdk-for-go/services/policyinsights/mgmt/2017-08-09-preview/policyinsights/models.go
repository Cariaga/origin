package policyinsights

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
)

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/policyinsights/mgmt/2017-08-09-preview/policyinsights instead.
// PolicyStatesResource enumerates the values for policy states resource.
type PolicyStatesResource string

const (
	// Default ...
	Default PolicyStatesResource = "default"
	// Latest ...
	Latest PolicyStatesResource = "latest"
)

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/policyinsights/mgmt/2017-08-09-preview/policyinsights instead.
// PossiblePolicyStatesResourceValues returns an array of possible values for the PolicyStatesResource const type.
func PossiblePolicyStatesResourceValues() []PolicyStatesResource {
	return []PolicyStatesResource{Default, Latest}
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/policyinsights/mgmt/2017-08-09-preview/policyinsights instead.
// Column column definition.
type Column struct {
	// Ordinal - Ordinal value of the column in a record.
	Ordinal *int32 `json:"ordinal,omitempty"`
	// Name - Name of the column.
	Name *string `json:"name,omitempty"`
	// DataType - Data type of the column.
	DataType *string `json:"dataType,omitempty"`
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/policyinsights/mgmt/2017-08-09-preview/policyinsights instead.
// Operation operation definition.
type Operation struct {
	// Name - Operation name.
	Name *string `json:"name,omitempty"`
	// Display - Display metadata associated with the operation.
	Display *OperationDisplay `json:"display,omitempty"`
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/policyinsights/mgmt/2017-08-09-preview/policyinsights instead.
// OperationDisplay display metadata associated with the operation.
type OperationDisplay struct {
	// Provider - Resource provider name.
	Provider *string `json:"provider,omitempty"`
	// Resource - Resource name on which the operation is performed.
	Resource *string `json:"resource,omitempty"`
	// Operation - Operation name.
	Operation *string `json:"operation,omitempty"`
	// Description - Operation description.
	Description *string `json:"description,omitempty"`
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/policyinsights/mgmt/2017-08-09-preview/policyinsights instead.
// OperationsListResults list of available operations.
type OperationsListResults struct {
	autorest.Response `json:"-"`
	// Value - List of available operations.
	Value *[]Operation `json:"value,omitempty"`
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/policyinsights/mgmt/2017-08-09-preview/policyinsights instead.
// PolicyEventsQueryResults query results.
type PolicyEventsQueryResults struct {
	autorest.Response `json:"-"`
	// Value - Query results.
	Value *[]PolicyEventsQueryResultsTable `json:"value,omitempty"`
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/policyinsights/mgmt/2017-08-09-preview/policyinsights instead.
// PolicyEventsQueryResultsTable query results table.
type PolicyEventsQueryResultsTable struct {
	// Metadata - Metadata about the query results.
	Metadata *PolicyEventsQueryResultsTableMetadata `json:"metadata,omitempty"`
	// Columns - List of columns included in query results.
	Columns *PolicyEventsQueryResultsTableColumns `json:"columns,omitempty"`
	// Rows - Query result rows, each representing a policy event record.
	Rows *[][]interface{} `json:"rows,omitempty"`
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/policyinsights/mgmt/2017-08-09-preview/policyinsights instead.
// PolicyEventsQueryResultsTableColumns list of columns included in query results.
type PolicyEventsQueryResultsTableColumns struct {
	// AdditionalProperties - Unmatched properties from the message are deserialized this collection
	AdditionalProperties map[string]interface{} `json:""`
	// Timestamp - Timestamp for the policy event record.
	Timestamp *Column `json:"Timestamp,omitempty"`
	// ResourceID - Resource ID.
	ResourceID *Column `json:"ResourceId,omitempty"`
	// PolicyAssignmentID - Policy assignment ID.
	PolicyAssignmentID *Column `json:"PolicyAssignmentId,omitempty"`
	// PolicyDefinitionID - Policy definition ID.
	PolicyDefinitionID *Column `json:"PolicyDefinitionId,omitempty"`
	// EffectiveParameters - Effective parameters for the policy assignment.
	EffectiveParameters *Column `json:"EffectiveParameters,omitempty"`
	// IsCompliant - Flag which states whether the resource is compliant against the policy assignment it was evaluated against.
	IsCompliant *Column `json:"IsCompliant,omitempty"`
	// SubscriptionID - Subscription ID.
	SubscriptionID *Column `json:"SubscriptionId,omitempty"`
	// ResourceType - Resource type.
	ResourceType *Column `json:"ResourceType,omitempty"`
	// ResourceLocation - Resource location.
	ResourceLocation *Column `json:"ResourceLocation,omitempty"`
	// ResourceGroup - Resource group name.
	ResourceGroup *Column `json:"ResourceGroup,omitempty"`
	// ResourceTags - List of resource tags.
	ResourceTags *Column `json:"ResourceTags,omitempty"`
	// PolicyAssignmentName - Policy assignment name.
	PolicyAssignmentName *Column `json:"PolicyAssignmentName,omitempty"`
	// PolicyAssignmentOwner - Policy assignment owner.
	PolicyAssignmentOwner *Column `json:"PolicyAssignmentOwner,omitempty"`
	// PolicyAssignmentParameters - Policy assignment parameters.
	PolicyAssignmentParameters *Column `json:"PolicyAssignmentParameters,omitempty"`
	// PolicyAssignmentScope - Policy assignment scope.
	PolicyAssignmentScope *Column `json:"PolicyAssignmentScope,omitempty"`
	// PolicyDefinitionName - Policy definition name.
	PolicyDefinitionName *Column `json:"PolicyDefinitionName,omitempty"`
	// PolicyDefinitionAction - Policy definition action, i.e. effect.
	PolicyDefinitionAction *Column `json:"PolicyDefinitionAction,omitempty"`
	// PolicyDefinitionCategory - Policy definition category.
	PolicyDefinitionCategory *Column `json:"PolicyDefinitionCategory,omitempty"`
	// PolicySetDefinitionID - Policy set definition ID, if the policy assignment is for a policy set.
	PolicySetDefinitionID *Column `json:"PolicySetDefinitionId,omitempty"`
	// PolicySetDefinitionName - Policy set definition name, if the policy assignment is for a policy set.
	PolicySetDefinitionName *Column `json:"PolicySetDefinitionName,omitempty"`
	// PolicySetDefinitionOwner - Policy set definition owner, if the policy assignment is for a policy set.
	PolicySetDefinitionOwner *Column `json:"PolicySetDefinitionOwner,omitempty"`
	// PolicySetDefinitionCategory - Policy set definition category, if the policy assignment is for a policy set.
	PolicySetDefinitionCategory *Column `json:"PolicySetDefinitionCategory,omitempty"`
	// PolicySetDefinitionParameters - Policy set definition parameters, if the policy assignment is for a policy set.
	PolicySetDefinitionParameters *Column `json:"PolicySetDefinitionParameters,omitempty"`
	// ManagementGroupIds - Comma seperated list of management group IDs, which represent the hierarchy of the management groups the resource is under.
	ManagementGroupIds *Column `json:"ManagementGroupIds,omitempty"`
	// PolicyDefinitionReferenceID - Reference ID for the policy definition inside the policy set, if the policy assignment is for a policy set.
	PolicyDefinitionReferenceID *Column `json:"PolicyDefinitionReferenceId,omitempty"`
	// TenantID - Tenant ID for the policy event record.
	TenantID *Column `json:"TenantId,omitempty"`
	// PrincipalOid - Principal object ID for the user who initiated the resource operation that triggered the policy event.
	PrincipalOid *Column `json:"PrincipalOid,omitempty"`
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/policyinsights/mgmt/2017-08-09-preview/policyinsights instead.
// MarshalJSON is the custom marshaler for PolicyEventsQueryResultsTableColumns.
func (peqrt PolicyEventsQueryResultsTableColumns) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if peqrt.Timestamp != nil {
		objectMap["Timestamp"] = peqrt.Timestamp
	}
	if peqrt.ResourceID != nil {
		objectMap["ResourceId"] = peqrt.ResourceID
	}
	if peqrt.PolicyAssignmentID != nil {
		objectMap["PolicyAssignmentId"] = peqrt.PolicyAssignmentID
	}
	if peqrt.PolicyDefinitionID != nil {
		objectMap["PolicyDefinitionId"] = peqrt.PolicyDefinitionID
	}
	if peqrt.EffectiveParameters != nil {
		objectMap["EffectiveParameters"] = peqrt.EffectiveParameters
	}
	if peqrt.IsCompliant != nil {
		objectMap["IsCompliant"] = peqrt.IsCompliant
	}
	if peqrt.SubscriptionID != nil {
		objectMap["SubscriptionId"] = peqrt.SubscriptionID
	}
	if peqrt.ResourceType != nil {
		objectMap["ResourceType"] = peqrt.ResourceType
	}
	if peqrt.ResourceLocation != nil {
		objectMap["ResourceLocation"] = peqrt.ResourceLocation
	}
	if peqrt.ResourceGroup != nil {
		objectMap["ResourceGroup"] = peqrt.ResourceGroup
	}
	if peqrt.ResourceTags != nil {
		objectMap["ResourceTags"] = peqrt.ResourceTags
	}
	if peqrt.PolicyAssignmentName != nil {
		objectMap["PolicyAssignmentName"] = peqrt.PolicyAssignmentName
	}
	if peqrt.PolicyAssignmentOwner != nil {
		objectMap["PolicyAssignmentOwner"] = peqrt.PolicyAssignmentOwner
	}
	if peqrt.PolicyAssignmentParameters != nil {
		objectMap["PolicyAssignmentParameters"] = peqrt.PolicyAssignmentParameters
	}
	if peqrt.PolicyAssignmentScope != nil {
		objectMap["PolicyAssignmentScope"] = peqrt.PolicyAssignmentScope
	}
	if peqrt.PolicyDefinitionName != nil {
		objectMap["PolicyDefinitionName"] = peqrt.PolicyDefinitionName
	}
	if peqrt.PolicyDefinitionAction != nil {
		objectMap["PolicyDefinitionAction"] = peqrt.PolicyDefinitionAction
	}
	if peqrt.PolicyDefinitionCategory != nil {
		objectMap["PolicyDefinitionCategory"] = peqrt.PolicyDefinitionCategory
	}
	if peqrt.PolicySetDefinitionID != nil {
		objectMap["PolicySetDefinitionId"] = peqrt.PolicySetDefinitionID
	}
	if peqrt.PolicySetDefinitionName != nil {
		objectMap["PolicySetDefinitionName"] = peqrt.PolicySetDefinitionName
	}
	if peqrt.PolicySetDefinitionOwner != nil {
		objectMap["PolicySetDefinitionOwner"] = peqrt.PolicySetDefinitionOwner
	}
	if peqrt.PolicySetDefinitionCategory != nil {
		objectMap["PolicySetDefinitionCategory"] = peqrt.PolicySetDefinitionCategory
	}
	if peqrt.PolicySetDefinitionParameters != nil {
		objectMap["PolicySetDefinitionParameters"] = peqrt.PolicySetDefinitionParameters
	}
	if peqrt.ManagementGroupIds != nil {
		objectMap["ManagementGroupIds"] = peqrt.ManagementGroupIds
	}
	if peqrt.PolicyDefinitionReferenceID != nil {
		objectMap["PolicyDefinitionReferenceId"] = peqrt.PolicyDefinitionReferenceID
	}
	if peqrt.TenantID != nil {
		objectMap["TenantId"] = peqrt.TenantID
	}
	if peqrt.PrincipalOid != nil {
		objectMap["PrincipalOid"] = peqrt.PrincipalOid
	}
	for k, v := range peqrt.AdditionalProperties {
		objectMap[k] = v
	}
	return json.Marshal(objectMap)
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/policyinsights/mgmt/2017-08-09-preview/policyinsights instead.
// PolicyEventsQueryResultsTableMetadata metadata about the query results.
type PolicyEventsQueryResultsTableMetadata struct {
	// GeneratedQuery - Internal query generated. Used for diagnostics purposes.
	GeneratedQuery *string `json:"generatedQuery,omitempty"`
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/policyinsights/mgmt/2017-08-09-preview/policyinsights instead.
// PolicyStatesQueryResults query results.
type PolicyStatesQueryResults struct {
	autorest.Response `json:"-"`
	// Value - Query results.
	Value *[]PolicyStatesQueryResultsTable `json:"value,omitempty"`
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/policyinsights/mgmt/2017-08-09-preview/policyinsights instead.
// PolicyStatesQueryResultsTable query results table.
type PolicyStatesQueryResultsTable struct {
	// Metadata - Metadata about the query results.
	Metadata *PolicyStatesQueryResultsTableMetadata `json:"metadata,omitempty"`
	// Columns - List of columns included in query results.
	Columns *PolicyStatesQueryResultsTableColumns `json:"columns,omitempty"`
	// Rows - Query result rows, each representing a policy state record.
	Rows *[][]interface{} `json:"rows,omitempty"`
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/policyinsights/mgmt/2017-08-09-preview/policyinsights instead.
// PolicyStatesQueryResultsTableColumns list of columns included in query results.
type PolicyStatesQueryResultsTableColumns struct {
	// AdditionalProperties - Unmatched properties from the message are deserialized this collection
	AdditionalProperties map[string]interface{} `json:""`
	// Timestamp - Timestamp for the policy state record.
	Timestamp *Column `json:"Timestamp,omitempty"`
	// ResourceID - Resource ID.
	ResourceID *Column `json:"ResourceId,omitempty"`
	// PolicyAssignmentID - Policy assignment ID.
	PolicyAssignmentID *Column `json:"PolicyAssignmentId,omitempty"`
	// PolicyDefinitionID - Policy definition ID.
	PolicyDefinitionID *Column `json:"PolicyDefinitionId,omitempty"`
	// EffectiveParameters - Effective parameters for the policy assignment.
	EffectiveParameters *Column `json:"EffectiveParameters,omitempty"`
	// IsCompliant - Flag which states whether the resource is compliant against the policy assignment it was evaluated against.
	IsCompliant *Column `json:"IsCompliant,omitempty"`
	// SubscriptionID - Subscription ID.
	SubscriptionID *Column `json:"SubscriptionId,omitempty"`
	// ResourceType - Resource type.
	ResourceType *Column `json:"ResourceType,omitempty"`
	// ResourceLocation - Resource location.
	ResourceLocation *Column `json:"ResourceLocation,omitempty"`
	// ResourceGroup - Resource group name.
	ResourceGroup *Column `json:"ResourceGroup,omitempty"`
	// ResourceTags - List of resource tags.
	ResourceTags *Column `json:"ResourceTags,omitempty"`
	// PolicyAssignmentName - Policy assignment name.
	PolicyAssignmentName *Column `json:"PolicyAssignmentName,omitempty"`
	// PolicyAssignmentOwner - Policy assignment owner.
	PolicyAssignmentOwner *Column `json:"PolicyAssignmentOwner,omitempty"`
	// PolicyAssignmentParameters - Policy assignment parameters.
	PolicyAssignmentParameters *Column `json:"PolicyAssignmentParameters,omitempty"`
	// PolicyAssignmentScope - Policy assignment scope.
	PolicyAssignmentScope *Column `json:"PolicyAssignmentScope,omitempty"`
	// PolicyDefinitionName - Policy definition name.
	PolicyDefinitionName *Column `json:"PolicyDefinitionName,omitempty"`
	// PolicyDefinitionAction - Policy definition action, i.e. effect.
	PolicyDefinitionAction *Column `json:"PolicyDefinitionAction,omitempty"`
	// PolicyDefinitionCategory - Policy definition category.
	PolicyDefinitionCategory *Column `json:"PolicyDefinitionCategory,omitempty"`
	// PolicySetDefinitionID - Policy set definition ID, if the policy assignment is for a policy set.
	PolicySetDefinitionID *Column `json:"PolicySetDefinitionId,omitempty"`
	// PolicySetDefinitionName - Policy set definition name, if the policy assignment is for a policy set.
	PolicySetDefinitionName *Column `json:"PolicySetDefinitionName,omitempty"`
	// PolicySetDefinitionOwner - Policy set definition owner, if the policy assignment is for a policy set.
	PolicySetDefinitionOwner *Column `json:"PolicySetDefinitionOwner,omitempty"`
	// PolicySetDefinitionCategory - Policy set definition category, if the policy assignment is for a policy set.
	PolicySetDefinitionCategory *Column `json:"PolicySetDefinitionCategory,omitempty"`
	// PolicySetDefinitionParameters - Policy set definition parameters, if the policy assignment is for a policy set.
	PolicySetDefinitionParameters *Column `json:"PolicySetDefinitionParameters,omitempty"`
	// ManagementGroupIds - Comma seperated list of management group IDs, which represent the hierarchy of the management groups the resource is under.
	ManagementGroupIds *Column `json:"ManagementGroupIds,omitempty"`
	// PolicyDefinitionReferenceID - Reference ID for the policy definition inside the policy set, if the policy assignment is for a policy set.
	PolicyDefinitionReferenceID *Column `json:"PolicyDefinitionReferenceId,omitempty"`
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/policyinsights/mgmt/2017-08-09-preview/policyinsights instead.
// MarshalJSON is the custom marshaler for PolicyStatesQueryResultsTableColumns.
func (psqrt PolicyStatesQueryResultsTableColumns) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if psqrt.Timestamp != nil {
		objectMap["Timestamp"] = psqrt.Timestamp
	}
	if psqrt.ResourceID != nil {
		objectMap["ResourceId"] = psqrt.ResourceID
	}
	if psqrt.PolicyAssignmentID != nil {
		objectMap["PolicyAssignmentId"] = psqrt.PolicyAssignmentID
	}
	if psqrt.PolicyDefinitionID != nil {
		objectMap["PolicyDefinitionId"] = psqrt.PolicyDefinitionID
	}
	if psqrt.EffectiveParameters != nil {
		objectMap["EffectiveParameters"] = psqrt.EffectiveParameters
	}
	if psqrt.IsCompliant != nil {
		objectMap["IsCompliant"] = psqrt.IsCompliant
	}
	if psqrt.SubscriptionID != nil {
		objectMap["SubscriptionId"] = psqrt.SubscriptionID
	}
	if psqrt.ResourceType != nil {
		objectMap["ResourceType"] = psqrt.ResourceType
	}
	if psqrt.ResourceLocation != nil {
		objectMap["ResourceLocation"] = psqrt.ResourceLocation
	}
	if psqrt.ResourceGroup != nil {
		objectMap["ResourceGroup"] = psqrt.ResourceGroup
	}
	if psqrt.ResourceTags != nil {
		objectMap["ResourceTags"] = psqrt.ResourceTags
	}
	if psqrt.PolicyAssignmentName != nil {
		objectMap["PolicyAssignmentName"] = psqrt.PolicyAssignmentName
	}
	if psqrt.PolicyAssignmentOwner != nil {
		objectMap["PolicyAssignmentOwner"] = psqrt.PolicyAssignmentOwner
	}
	if psqrt.PolicyAssignmentParameters != nil {
		objectMap["PolicyAssignmentParameters"] = psqrt.PolicyAssignmentParameters
	}
	if psqrt.PolicyAssignmentScope != nil {
		objectMap["PolicyAssignmentScope"] = psqrt.PolicyAssignmentScope
	}
	if psqrt.PolicyDefinitionName != nil {
		objectMap["PolicyDefinitionName"] = psqrt.PolicyDefinitionName
	}
	if psqrt.PolicyDefinitionAction != nil {
		objectMap["PolicyDefinitionAction"] = psqrt.PolicyDefinitionAction
	}
	if psqrt.PolicyDefinitionCategory != nil {
		objectMap["PolicyDefinitionCategory"] = psqrt.PolicyDefinitionCategory
	}
	if psqrt.PolicySetDefinitionID != nil {
		objectMap["PolicySetDefinitionId"] = psqrt.PolicySetDefinitionID
	}
	if psqrt.PolicySetDefinitionName != nil {
		objectMap["PolicySetDefinitionName"] = psqrt.PolicySetDefinitionName
	}
	if psqrt.PolicySetDefinitionOwner != nil {
		objectMap["PolicySetDefinitionOwner"] = psqrt.PolicySetDefinitionOwner
	}
	if psqrt.PolicySetDefinitionCategory != nil {
		objectMap["PolicySetDefinitionCategory"] = psqrt.PolicySetDefinitionCategory
	}
	if psqrt.PolicySetDefinitionParameters != nil {
		objectMap["PolicySetDefinitionParameters"] = psqrt.PolicySetDefinitionParameters
	}
	if psqrt.ManagementGroupIds != nil {
		objectMap["ManagementGroupIds"] = psqrt.ManagementGroupIds
	}
	if psqrt.PolicyDefinitionReferenceID != nil {
		objectMap["PolicyDefinitionReferenceId"] = psqrt.PolicyDefinitionReferenceID
	}
	for k, v := range psqrt.AdditionalProperties {
		objectMap[k] = v
	}
	return json.Marshal(objectMap)
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/policyinsights/mgmt/2017-08-09-preview/policyinsights instead.
// PolicyStatesQueryResultsTableMetadata metadata about the query results.
type PolicyStatesQueryResultsTableMetadata struct {
	// GeneratedQuery - Internal query generated. Used for diagnostics purposes.
	GeneratedQuery *string `json:"generatedQuery,omitempty"`
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/policyinsights/mgmt/2017-08-09-preview/policyinsights instead.
// QueryFailure error response.
type QueryFailure struct {
	// Error - Error definition.
	Error *QueryFailureError `json:"error,omitempty"`
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/policyinsights/mgmt/2017-08-09-preview/policyinsights instead.
// QueryFailureError error definition.
type QueryFailureError struct {
	// Code - Service specific error code which serves as the substatus for the HTTP error code.
	Code *string `json:"code,omitempty"`
	// Message - Description of the error.
	Message *string `json:"message,omitempty"`
}
