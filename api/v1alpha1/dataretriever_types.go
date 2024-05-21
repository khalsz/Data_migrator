/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	batchv1 "k8s.io/api/batch/v1"

	corev1 "k8s.io/api/core/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// DataRetrieverSpec defines the desired state of DataRetriever
type DataRetrieverSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	//+kubebuilder:validation:MinLength=0
	// The schedule of the dataretriever application in Cron Format,
	// check https://cloud.google.com/scheduler/docs/configuring/cron-job-schedules#:~:text=A%20schedule%20is%20defined%20using,See%20more%20code%20actions..
	Schedule string `json:"schedule"`

	//+kubebuilder:validation:Minimum=0
	//optional start deadline in seconds. If deadline is missed, the job will be counted as failed
	// +optional
	StartDeadlineSecs *int64 `json:"startDeadlineSeconds,omitempty"`

	// Specifies how to treat concurrent executions of a Job.
	// Valid values are:
	// - "Allow" (default): allows CronJobs to run concurrently;
	// - "Forbid": forbids concurrent runs, skipping next run if previous run hasn't finished yet;
	// - "Replace": cancels currently running job and replaces it with a new one
	// +optional
	ConcurrencyPolicy ConcurrencyPolicy `json:"concurrencyPolicy,omitempty"`

	//+kubebuilder:validation:Minimum=0
	// Specifies the number of retries before marking this job failed.
	// Defaults to 6
	// +optional
	BackoffLimit int `json:"backoffLimit,omitempty"`

	// Name of the ConfigMap for GuestbookSpec's configuration
	// +kubebuilder:validation:MaxLength=15
	// +kubebuilder:validation:MinLength=1
	ConfigMapName string `json:"configMapName"`

	// Specifies the job that will be created when executing a CronJob.
	JobTemplate batchv1.JobTemplateSpec `json:"jobTemplate"`

	//+kubebuilder:validation:Minimum=0
	// optional maximum active duration of a job relative to its start time (in seconds).
	// system will terminate job if active time exceed this. It is reset if job is
	// suspended and reset when job resumes.
	ActiveDeadlineSeconds int64 `json:"activeDeadlineSeconds,omitempty"`
}

// DataRetrieverStatus defines the observed state of DataRetriever
type DataRetrieverStatus struct {
	// A list of pointers to currently running jobs.
	Active []corev1.ObjectReference `json:"active"`
	// Information when was the last time the job was successfully scheduled.
	LastScheduleTime *metav1.Time `json:"lastScheduleTime,omitempty"`
	// Information when was the last time the job successfully completed.
	LastSuccessfulTime *metav1.Time `json:"lastSuccessfulTime,omitempty"`
}

// ConcurrencyPolicy describes how the job will be handled.
// Only one of the following concurrent policies may be specified.
// If none of the following policies is specified, the default one
// is AllowConcurrent.
// +kubebuilder:validation:Enum=Allow;Forbid;Replace
type ConcurrencyPolicy string

const (
	// AllowConcurrent allows CronJobs to run concurrently.
	AllowConcurrent ConcurrencyPolicy = "Allow"

	// ForbidConcurrent forbids concurrent runs, skipping next run if previous
	// hasn't finished yet.
	ForbidConcurrent ConcurrencyPolicy = "Forbid"

	// ReplaceConcurrent cancels currently running job and replaces it with a new one.
	ReplaceConcurrent ConcurrencyPolicy = "Replace"
)

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// DataRetriever is the Schema for the dataretrievers API
type DataRetriever struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DataRetrieverSpec   `json:"spec,omitempty"`
	Status DataRetrieverStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// DataRetrieverList contains a list of DataRetriever
type DataRetrieverList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataRetriever `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataRetriever{}, &DataRetrieverList{})
}
