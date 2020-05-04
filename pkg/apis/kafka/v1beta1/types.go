package v1beta1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// KafkaUserSpec is the spec for kafkauser resource
type KafkaUserSpec struct {
	Authentication Authentication `json:"authentication,omitempty"`
	Authorization  Authorization  `json:"authorization,omitempty"`
	//  +optional
	Quotas Quotas `json:quotas,omitempty`
}

// KafkaUserStatus defines the observed state of kafkauser custom resource
type KafkaUserStatus struct {
	ObservedGeneration *int32       `json:"observedGeneration,omitempty"`
	Conditions         []Conditions `json:"conditions,omitempty"`
	Username           string       `json:username,omitempty`
	Secret             string       `json:secret,omitempty`
}

// Authentication defines authentication
type Authentication struct {
	Type string `json:"type,omitempty"`
}

// Authorization defines authorization
type Authorization struct {
	Acls []Acls `json:"acls,omitempty"`
	Type string `json:"type,omitempty"`
}

// Acls defines acls
type Acls struct {
	Host        string      `json:"host,omitempty"`
	Operation   string      `json:"operation,omitempty"`
	AclResource AclResource `json:"resource,omitempty"`
	Type        string      `json:"type,omitempty"`
}

// AclResource defines resource
type AclResource struct {
	Name        string `json:"name,omitempty"`
	PatternType string `json:"patternType,omitempty"`
	Type        string `json:"type,omitempty"`
}

// Quotas defines quotas
type Quotas struct {
	ConsumerByteRate  *int32 `json:"consumerByteRate,omitempty"`
	ProducerByteRate  *int32 `json:"producerByteRate,omitempty"`
	RequestPercentage *int32 `json:"requestPercentage,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KafkaUser describes kafkauser resource type
type KafkaUser struct {
	metav1.TypeMeta `json:",inline"`

	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec KafkaUserSpec `json:"spec"`

	// +optional
	Status KafkaUserStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KafkaUserList is a list of KafkaUser resources
type KafkaUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KafkaUser `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KafkaTopic describes a KafkaTopic resource
type KafkaTopic struct {
	metav1.TypeMeta `json:",inline"`

	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec KafkaTopicSpec `json:"spec"`

	// +optional
	Status KafkaTopicStatus `json:"status,omitempty"`
}

// KafkaTopicSpec is the spec for a kafkatopic resource
type KafkaTopicSpec struct {
	Partitions *int32            `json:"partitions,omitempty"`
	Replicas   *int32            `json:"replicas,omitempty"`
	Config     map[string]string `json:"config,omitempty"`
	TopicName  string            `json:"topicName,omitempty"`
}

// Condition
type Conditions struct {
	Type               string             `json:"type,omitempty"`
	Status             v1.ConditionStatus `json:"status,omitempty"`
	LastTransitionTime metav1.Time        `json:"lastTransition,omitempty"`
	Reason             string             `json:"reason,omitempty"`
	Message            string             `json:"message,omitempty"`
}

// KafkaTopicStatus defines the observed state of kafkatopic custom resource
type KafkaTopicStatus struct {
	ObservedGeneration *int32       `json:"observedGeneration,omitempty"`
	Conditions         []Conditions `json:"conditions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KafkaTopicList is a list of KafkaTopic resources
type KafkaTopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KafkaTopic `json:"items"`
}
