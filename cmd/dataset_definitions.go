package cmd

import (
	"encoding/json"
	"net/http"

	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"
)

type datasetDefinitionColumn struct {
	// The name of the Column or Derived Column to map to this Dataset
	// Definition Type. An empty string clears the mapping, potentially
	// reverting to a default mapping.
	Name string `json:"name,omitempty"`

	// Optional: column for regular columns and derived_column for derived
	// columns when setting Dataset Definitions. Honeycomb does not use this
	// field when updating Dataset definitions.
	ColumnType string `json:"column_type,omitempty"`
}

type datasetDefinition struct {
	// The unique identifier (ID) for each span.
	SpanID datasetDefinitionColumn `json:"span_id,omitempty"`

	// The ID of the trace this span belongs to.
	TraceID datasetDefinitionColumn `json:"trace_id,omitempty"`

	// The Parent Span ID - The ID of this span's parent span, the call
	// location the current span was called from.
	ParentID datasetDefinitionColumn `json:"parent_id,omitempty"`

	// The name of the function or method where the span was created.
	Name datasetDefinitionColumn `json:"name,omitempty"`

	// The name of the instrumented service.
	ServiceName datasetDefinitionColumn `json:"service_name,omitempty"`

	// Span Duration - How much time the span took, in milliseconds.
	DurationMs datasetDefinitionColumn `json:"duration_ms,omitempty"`

	// Metadata: Kind - The kind of Span. For example, client or server. The
	// use of this field to identify Span Events and Links is deprecated. Use
	// the field Metadata: Annotation Type.
	SpanKind datasetDefinitionColumn `json:"span_kind,omitempty"`

	// Metadata: Annotation Type - The type of span annotation. For example,
	// span_event or link. This lets Honeycomb visualize this type of event
	// differently in a trace. Do not use this field for other purposes.
	AnnotationType datasetDefinitionColumn `json:"annotation_type,omitempty"`

	// Metadata: Link Span ID - Links let you tie traces and spans to one
	// another. The Link Span ID lets you link to a different span (when used
	// with Link Trace ID).
	LinkSpanID datasetDefinitionColumn `json:"link_span_id,omitempty"`

	// Metadata: Link Trace ID - Links let you tie traces and spans to one
	// another. The Link Trace Id lets you link to a different trace or a
	// different span in the same trace (when used with Link Span ID).
	LinkTraceID datasetDefinitionColumn `json:"link_trace_id,omitempty"`

	// Use a Boolean or String to indicate error.
	Error datasetDefinitionColumn `json:"error,omitempty"`

	// Indicates the success, failure, or other status of a request.
	Status datasetDefinitionColumn `json:"status,omitempty"`

	// The HTTP URL or equivalent route processed by the request.
	Route datasetDefinitionColumn `json:"route,omitempty"`

	// The user making the request in the system.
	User datasetDefinitionColumn `json:"user,omitempty"`
}

func newDatasetDefinitionsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "dataset_definitions",
		Aliases: []string{"dd"},
		Short:   "Manage Dataset Definitions",
		Long: "Dataset definitions describe the fields with special meaning in the Dataset.\n" +
			"\n" +
			"Refer to the Dataset Definitions documentation for more information.\n" +
			"\n" +
			"Honeycomb automatically creates these Dataset definition fields when the Dataset is\n" +
			"created. Manual creation of Dataset definitions is not needed.",
	}

	cmd.AddCommand(
		newDatasetDefinitionsUpdateCmd(),
		newDatasetDefinitionsGetCmd(),
	)

	return cmd
}

// Set or Update Dataset Definitions
// https://docs.honeycomb.io/api/tag/Dataset-Definitions#operation/patchDatasetDefinitions
func newDatasetDefinitionsUpdateCmd() *cobra.Command {
	var (
		dSlug            string
		ddSpanIDName     string
		ddTraceIDName    string
		ddParentIDName   string
		ddName           string
		ddServiceName    string
		ddDurationMs     string
		ddSpanKind       string
		ddAnnotationType string
		ddLinkSpanID     string
		ddLinkTraceID    string
		ddError          string
		ddStatus         string
		ddRoute          string
		ddUser           string
	)
	cmd := &cobra.Command{
		Use:     "update",
		Aliases: []string{"up", "edit", "modify", "change", "set"},
		Short:   "Set or Update Dataset Definitions.",
		Long: "Set or update one or more definitions for a Dataset.\n" +
			"\n" +
			"Note: While the PATCH payload can include the column_type, Honeycomb does not use\n" +
			"this field when updating Dataset Definitions.",
		Run: func(cmd *cobra.Command, args []string) {
			var dd = datasetDefinition{
				SpanID: datasetDefinitionColumn{
					Name: ddSpanIDName,
				},
				TraceID: datasetDefinitionColumn{
					Name: ddTraceIDName,
				},
				ParentID: datasetDefinitionColumn{
					Name: ddParentIDName,
				},
				Name: datasetDefinitionColumn{
					Name: ddName,
				},
				ServiceName: datasetDefinitionColumn{
					Name: ddServiceName,
				},
				DurationMs: datasetDefinitionColumn{
					Name: ddDurationMs,
				},
				SpanKind: datasetDefinitionColumn{
					Name: ddSpanKind,
				},
				AnnotationType: datasetDefinitionColumn{
					Name: ddAnnotationType,
				},
				LinkSpanID: datasetDefinitionColumn{
					Name: ddLinkSpanID,
				},
				LinkTraceID: datasetDefinitionColumn{
					Name: ddLinkTraceID,
				},
				Error: datasetDefinitionColumn{
					Name: ddError,
				},
				Status: datasetDefinitionColumn{
					Name: ddStatus,
				},
				Route: datasetDefinitionColumn{
					Name: ddRoute,
				},
				User: datasetDefinitionColumn{
					Name: ddUser,
				},
			}
			var bodyMarshal, err = json.Marshal(dd)
			if err != nil {
				log.WithFields(log.Fields{
					"_function":          "newDatasetDefinitionsUpdateCmd",
					"err":                err,
					"dataset_definition": dd,
				}).Fatal("Error received when attempting to marshal a dataset definition.")
			}
			var p = payload{
				Method:   http.MethodPatch,
				Path:     "/1/dataset_definitions/" + dSlug,
				Body:     bodyMarshal,
				Response: &datasetDefinition{},
			}

			err = p.GetResponse(true)
			if err != nil {
				log.WithFields(log.Fields{
					"_function": "newDatasetDefinitionsUpdateCmd",
					"err":       err,
					"payload":   p,
				}).Fatal("Error received when attempting to update a dataset definition.")
			}
		},
	}

	cmd.Flags().StringVar(&dSlug, "slug", "",
		"The dataset slug.")
	cmd.MarkFlagRequired("slug")
	cmd.Flags().StringVar(&ddSpanIDName, "span-id", "",
		"The unique identifier (ID) for each span.")
	cmd.Flags().StringVar(&ddTraceIDName, "trace-id", "",
		"The ID of the trace this span belongs to.")
	cmd.Flags().StringVar(&ddParentIDName, "parent-id", "",
		"The ID of this span's parent span, the call location the current span was called from.")
	cmd.Flags().StringVar(&ddName, "name", "",
		"The name of the function or method where the span was created.")
	cmd.Flags().StringVar(&ddServiceName, "service-name", "",
		"The name of the instrumented service.")
	cmd.Flags().StringVar(&ddDurationMs, "duration-ms", "",
		"How much time the span took, in milliseconds.")
	cmd.Flags().StringVar(&ddSpanKind, "span-kind", "",
		"Metadata: Kind - The kind of Span. For example, client or server. The use of this"+
			"field to identify Span Events and Links is deprecated. Use the field Metadata: Annotation Type.")
	cmd.Flags().StringVar(&ddAnnotationType, "annotation-type", "",
		"Metadata: Annotation Type - The type of span annotation. For example, span_event or link."+
			"This lets Honeycomb visualize this type of event differently in a trace. Do not use this field for other purposes.")
	cmd.Flags().StringVar(&ddLinkSpanID, "link-span-id", "",
		"Metadata: Link Span ID - Links let you tie traces and spans to one another. The Link Span"+
			"ID lets you link to a different span (when used with Link Trace ID).")
	cmd.Flags().StringVar(&ddLinkTraceID, "link-trace-id", "",
		"Metadata: Link Trace ID - Links let you tie traces and spans to one another. The Link Trace"+
			"Id lets you link to a different trace or a different span in the same trace (when used with Link Span ID).")
	cmd.Flags().StringVar(&ddError, "error", "",
		"Use a Boolean or String to indicate error.")
	cmd.Flags().StringVar(&ddStatus, "status", "",
		"Indicates the success, failure, or other status of a request.")
	cmd.Flags().StringVar(&ddRoute, "route", "",
		"The HTTP URL or equivalent route processed by the request.")
	cmd.Flags().StringVar(&ddUser, "user", "",
		"The user making the request in the system.")

	return cmd
}

// Get All Dataset Definitions
// https://docs.honeycomb.io/api/tag/Dataset-Definitions#operation/listDatasetDefinitions
func newDatasetDefinitionsGetCmd() *cobra.Command {
	var (
		dSlug string
	)
	cmd := &cobra.Command{
		Use:     "get",
		Aliases: []string{"ls", "list"},
		Short:   "Get all Dataset Definitions.",
		Long: "Get all definitions for a Dataset.\n" +
			"\n" +
			"The response returns an object with a Dataset Definition for each set Dataset Definition type.",
		Run: func(cmd *cobra.Command, args []string) {
			var p = payload{
				Method:   http.MethodGet,
				Path:     "/1/dataset_definitions/" + dSlug,
				Response: &datasetDefinition{},
			}

			var err = p.GetResponse(true)
			if err != nil {
				log.WithFields(log.Fields{
					"_function": "newDatasetDefinitionsGetCmd",
					"err":       err,
					"payload":   p,
				}).Fatal("Error received when attempting to get a dataset definition.")
			}
		},
	}

	cmd.Flags().StringVar(&dSlug, "slug", "",
		"The dataset slug.")
	cmd.MarkFlagRequired("slug")

	return cmd
}
