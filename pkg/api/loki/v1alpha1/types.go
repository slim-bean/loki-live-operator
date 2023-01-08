package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:resource:path="grafanaagents"
// +kubebuilder:resource:singular="grafanaagent"
// +kubebuilder:resource:categories="agent-operator"

// LokiLive defines a Loki Live deployment.
type LokiLive struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec holds the specification of the desired behavior for the Grafana Agent
	// cluster.
	Spec LokiLiveSpec `json:"spec,omitempty"`
}

// LokiLiveSpec is a specification of the desired behavior of the Loki Live cluster
type LokiLiveSpec struct {
	// LogLevel controls the log level of the generated pods. Defaults to "info" if not set.
	LogLevel string `json:"logLevel,omitempty"`
	// LogFormat controls the logging format of the generated pods. Defaults to "logfmt" if not set.
	LogFormat string `json:"logFormat,omitempty"`
}
