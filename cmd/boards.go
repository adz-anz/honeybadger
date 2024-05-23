package cmd

import (
	"encoding/json"
	"net/http"

	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"
)

type boardGraphSettings struct {
	HideMarkers       bool `json:"hide_markers,omitempty"`
	LogScale          bool `json:"log_scale,omitempty"`
	OmitMissingValues bool `json:"omit_missing_values,omitempty"`
	StackedGraphs     bool `json:"stacked_graphs,omitempty"`
	UTCXaxis          bool `json:"utc_xaxis,omitempty"`
	OverlaidCharts    bool `json:"overlaid_charts,omitempty"`
}

type boardQuery struct {
	Caption           string             `json:"caption,omitempty"`
	GraphSettings     boardGraphSettings `json:"graph_settings,omitempty"`
	QueryStyle        string             `json:"query_style,omitempty"`
	Dataset           string             `json:"dataset,omitempty"`
	QueryID           string             `json:"query_id,omitempty"`
	QueryAnnotationID string             `json:"query_annotation_id,omitempty"`
}

type boardLinks struct {
	BoardURL string `json:"board_url,omitempty"`
}

type board struct {
	Name         string       `json:"name,omitempty"`
	Description  string       `json:"description,omitempty"`
	Style        string       `json:"style,omitempty"`
	ColumnLayout string       `json:"column_layout,omitempty"`
	Queries      []boardQuery `json:"queries,omitempty"`
	Links        boardLinks   `json:"links,omitempty"`
	ID           string       `json:"id,omitempty"`
}

func newBoardsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "boards",
		Aliases: []string{"b"},
		Short: "Manage Boards",
		Long: "Boards are a place to pin and save useful queries and graphs you want to\n" +
			"retain for later reuse and reference.\n" +
			"\n" +
			"This command allows you to list, create, update, and delete Boards.",
	}

	cmd.AddCommand(newBoardsCreateCmd())
	cmd.AddCommand(newBoardsListCmd())
	cmd.AddCommand(newBoardsGetCmd())
	cmd.AddCommand(newBoardsUpdateCmd())
	cmd.AddCommand(newBoardsDeleteCmd())

	return cmd
}

// Create a Board
// https://docs.honeycomb.io/api/tag/Boards#operation/createBoard
func newBoardsCreateCmd() *cobra.Command {
	var (
		bName                                string
		bDescription                         string
		bColumnLayout                        string
		bQueryCaptions                       []string
		bQueryGraphSettingsHideMarkers       []bool
		bQueryGraphSettingsLogScale          []bool
		bQueryGraphSettingsOmitMissingValues []bool
		bQueryGraphSettingsStackedGraphs     []bool
		bQueryGraphSettingsUTCXAxis          []bool
		bQueryGraphSettingsOverlaidCharts    []bool
		bQueryStyle                          []string
		bQueryDataset                        []string
		bQueryId                             []string
		bQueryAnnotationId                   []string
		bQueryLength                         int
	)

	cmd := &cobra.Command{
		Use:   "create",
		Aliases: []string{"add", "new"},
		Short: "Create a Board comprised of one or more Queries.",
		Long:  "Create a Board comprised of one or more Queries.",
		Run: func(cmd *cobra.Command, args []string) {
			var b = board{
				Name:         bName,
				Description:  bDescription,
				ColumnLayout: bColumnLayout,
				Queries:      []boardQuery{},
			}

			bQueryLength = max(
				len(bQueryCaptions),
				len(bQueryGraphSettingsHideMarkers),
				len(bQueryGraphSettingsLogScale),
				len(bQueryGraphSettingsOmitMissingValues),
				len(bQueryGraphSettingsStackedGraphs),
				len(bQueryGraphSettingsUTCXAxis),
				len(bQueryGraphSettingsOverlaidCharts),
				len(bQueryStyle),
				len(bQueryDataset),
				len(bQueryId),
				len(bQueryAnnotationId),
			)

			for idx := 0; idx < bQueryLength; idx++ {
				var newBoardQuery boardQuery

				if len(bQueryCaptions) > idx {
					newBoardQuery.Caption = bQueryCaptions[idx]
				}
				if len(bQueryGraphSettingsHideMarkers) > idx {
					newBoardQuery.GraphSettings.HideMarkers = bQueryGraphSettingsHideMarkers[idx]
				}
				if len(bQueryGraphSettingsLogScale) > idx {
					newBoardQuery.GraphSettings.LogScale = bQueryGraphSettingsLogScale[idx]
				}
				if len(bQueryGraphSettingsOmitMissingValues) > idx {
					newBoardQuery.GraphSettings.OmitMissingValues = bQueryGraphSettingsOmitMissingValues[idx]
				}
				if len(bQueryGraphSettingsStackedGraphs) > idx {
					newBoardQuery.GraphSettings.StackedGraphs = bQueryGraphSettingsStackedGraphs[idx]
				}
				if len(bQueryGraphSettingsUTCXAxis) > idx {
					newBoardQuery.GraphSettings.UTCXaxis = bQueryGraphSettingsUTCXAxis[idx]
				}
				if len(bQueryGraphSettingsOverlaidCharts) > idx {
					newBoardQuery.GraphSettings.OverlaidCharts = bQueryGraphSettingsOverlaidCharts[idx]
				}
				if len(bQueryStyle) > idx {
					newBoardQuery.QueryStyle = bQueryStyle[idx]
				}
				if len(bQueryDataset) > idx {
					newBoardQuery.Dataset = bQueryDataset[idx]
				}
				if len(bQueryId) > idx {
					newBoardQuery.QueryID = bQueryId[idx]
				}
				if len(bQueryAnnotationId) > idx {
					newBoardQuery.QueryAnnotationID = bQueryAnnotationId[idx]
				}

				b.Queries = append(b.Queries, newBoardQuery)
			}
			var bodyMarshal, err = json.Marshal(b)
			if err != nil {
				log.WithFields(log.Fields{
					"_function": "newBoardsCreateCmd",
					"err": err,
					"board": b,
				}).Fatal("Error received when attempting to marshal a board.")
			}
			var p = Payload{
				Method: http.MethodPost,
				Path:   "/1/boards",
				Body:   bodyMarshal,
			}

			err = p.Execute()
			if err != nil {
				log.WithFields(log.Fields{
					"_function": "newBoardsCreateCmd",
					"err": err,
					"payload": p,
				}).Fatal("Error received when attempting to create a new board.")
			}
		},
	}

	cmd.Flags().StringVarP(&bName, "name", "n", "", "The name of the Board.")
	cmd.MarkFlagRequired("name")
	cmd.Flags().StringVarP(&bDescription, "description", "d", "",
		"A description of the Board.")
	cmd.Flags().StringVarP(&bColumnLayout, "column_layout", "c", "",
		"The number of columns to layout on the board.")
	cmd.Flags().StringArrayVar(&bQueryCaptions, "qc", nil,
		"Descriptive text to contextualize the value of the Query within the Board.")
	cmd.Flags().BoolSliceVar(&bQueryGraphSettingsHideMarkers, "qgshm", nil,
		"Hide markers on the graph.")
	cmd.Flags().BoolSliceVar(&bQueryGraphSettingsLogScale, "qgsls", nil,
		"Use a log scale, rather than a linear scale.")
	cmd.Flags().BoolSliceVar(&bQueryGraphSettingsOmitMissingValues, "qgsomv", nil,
		"Omit missing values from the graph.")
	cmd.Flags().BoolSliceVar(&bQueryGraphSettingsStackedGraphs, "qgssg", nil,
		"Display groups as stacked colored areas under their line graphs.")
	cmd.Flags().BoolSliceVar(&bQueryGraphSettingsUTCXAxis, "qgsuxa", nil,
		"Displays the X axis in Coordinated Universal Time, the time at 0° longitude.")
	cmd.Flags().BoolSliceVar(&bQueryGraphSettingsOverlaidCharts, "qgsoc", nil,
		"Combines any visualized AVG, MIN, MAX, and PERCENTILE clauses into a single chart.")
	cmd.Flags().StringArrayVar(&bQueryStyle, "qs", nil,
		"How the query should be displayed on the board. Enum: \"graph\" \"table\" \"combo\"")
	cmd.Flags().StringArrayVar(&bQueryDataset, "qd", nil,
		"The Dataset to Query. Required if using the deprecated query. Note: this field can take either name (\"My Dataset\") or slug (\"my_dataset\"); the response will always use the name.")
	cmd.Flags().StringArrayVar(&bQueryId, "qid", nil,
		"The ID of a Query object. Cannot be used with query. Query IDs can be retrieved from the UI or from the Query API.")
	cmd.Flags().StringArrayVar(&bQueryAnnotationId, "qaid", nil,
		"The ID of a Query Annotation that provides a name and description for the Query. The Query Annotation must apply to the query_id or query specified.")

	return cmd
}

// List All Boards
// https://docs.honeycomb.io/api/tag/Boards#operation/listBoards
func newBoardsListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "Retrieves a list of all non-secret Boards within an environment.",
		Long: "Retrieves a list of all non-secret Boards within an environment.\n" +
			"\n" +
			"Note: For Honeycomb Classic users, all boards within Classic will be returned.",
		Run: func(cmd *cobra.Command, args []string) {
			var p = Payload{
				Method: http.MethodGet,
				Path:   "/1/boards",
			}

			var err = p.Execute()
			if err != nil {
				log.WithFields(log.Fields{
					"_function": "newBoardsListCmd",
					"err": err,
					"payload": p,
				}).Fatal("Error received when attempting to list boards.")
			}
		},
	}

	return cmd
}

func newBoardsGetCmd() *cobra.Command {
	var (
		bId string
	)

	cmd := &cobra.Command{
		Use:     "get",
		Aliases: []string{},
		Short:   "Get a single Board by ID.",
		Long:    "Get a single Board by ID.",
		Run: func(cmd *cobra.Command, args []string) {
			var p = Payload{
				Method: http.MethodGet,
				Path:   "/1/boards/" + bId,
			}

			var err = p.Execute()
			if err != nil {
				log.WithFields(log.Fields{
					"_function": "newBoardsGetCmd",
					"err": err,
					"payload": p,
				}).Fatal("Error received when attempting to get a single board.")
			}
		},
	}

	cmd.Flags().StringVarP(&bId, "id", "i", "", "The unique identifier (ID) of a Board.")
	cmd.MarkFlagRequired("id")

	return cmd
}

func newBoardsUpdateCmd() *cobra.Command {
	var (
		bId                                  string
		bName                                string
		bDescription                         string
		bColumnLayout                        string
		bQueryCaptions                       []string
		bQueryGraphSettingsHideMarkers       []bool
		bQueryGraphSettingsLogScale          []bool
		bQueryGraphSettingsOmitMissingValues []bool
		bQueryGraphSettingsStackedGraphs     []bool
		bQueryGraphSettingsUTCXAxis          []bool
		bQueryGraphSettingsOverlaidCharts    []bool
		bQueryStyle                          []string
		bQueryDataset                        []string
		bQueryId                             []string
		bQueryAnnotationId                   []string
		bQueryLength                         int
	)

	cmd := &cobra.Command{
		Use:   "update",
		Aliases: []string{"up", "edit", "modify", "change", "set"},
		Short: "Update a Board by ID.",
		Long:  "Update a Board by ID.",
		Run: func(cmd *cobra.Command, args []string) {
			var b = board{
				Name:         bName,
				Description:  bDescription,
				ColumnLayout: bColumnLayout,
				Queries:      []boardQuery{},
			}

			bQueryLength = max(
				len(bQueryCaptions),
				len(bQueryGraphSettingsHideMarkers),
				len(bQueryGraphSettingsLogScale),
				len(bQueryGraphSettingsOmitMissingValues),
				len(bQueryGraphSettingsStackedGraphs),
				len(bQueryGraphSettingsUTCXAxis),
				len(bQueryGraphSettingsOverlaidCharts),
				len(bQueryStyle),
				len(bQueryDataset),
				len(bQueryId),
				len(bQueryAnnotationId),
			)

			for idx := 0; idx < bQueryLength; idx++ {
				var newBoardQuery boardQuery

				if len(bQueryCaptions) > idx {
					newBoardQuery.Caption = bQueryCaptions[idx]
				}
				if len(bQueryGraphSettingsHideMarkers) > idx {
					newBoardQuery.GraphSettings.HideMarkers = bQueryGraphSettingsHideMarkers[idx]
				}
				if len(bQueryGraphSettingsLogScale) > idx {
					newBoardQuery.GraphSettings.LogScale = bQueryGraphSettingsLogScale[idx]
				}
				if len(bQueryGraphSettingsOmitMissingValues) > idx {
					newBoardQuery.GraphSettings.OmitMissingValues = bQueryGraphSettingsOmitMissingValues[idx]
				}
				if len(bQueryGraphSettingsStackedGraphs) > idx {
					newBoardQuery.GraphSettings.StackedGraphs = bQueryGraphSettingsStackedGraphs[idx]
				}
				if len(bQueryGraphSettingsUTCXAxis) > idx {
					newBoardQuery.GraphSettings.UTCXaxis = bQueryGraphSettingsUTCXAxis[idx]
				}
				if len(bQueryGraphSettingsOverlaidCharts) > idx {
					newBoardQuery.GraphSettings.OverlaidCharts = bQueryGraphSettingsOverlaidCharts[idx]
				}
				if len(bQueryStyle) > idx {
					newBoardQuery.QueryStyle = bQueryStyle[idx]
				}
				if len(bQueryDataset) > idx {
					newBoardQuery.Dataset = bQueryDataset[idx]
				}
				if len(bQueryId) > idx {
					newBoardQuery.QueryID = bQueryId[idx]
				}
				if len(bQueryAnnotationId) > idx {
					newBoardQuery.QueryAnnotationID = bQueryAnnotationId[idx]
				}

				b.Queries = append(b.Queries, newBoardQuery)
			}
			var bodyMarshal, err = json.Marshal(b)
			if err != nil {
				log.WithFields(log.Fields{
					"_function": "newBoardsUpdateCmd",
					"err": err,
					"board": b,
				}).Fatal("Error received when attempting to marshal a board.")
			}
			var p = Payload{
				Method: http.MethodPut,
				Path:   "/1/boards/" + bId,
				Body:   bodyMarshal,
			}

			err = p.Execute()
			if err != nil {
				log.WithFields(log.Fields{
					"_function": "newBoardsUpdateCmd",
					"err": err,
					"payload": p,
				}).Fatal("Error received when attempting to update an existing board.")
			}
		},
	}

	cmd.Flags().StringVarP(&bId, "id", "i", "", "The unique identifier (ID) of a Board.")
	cmd.MarkFlagRequired("id")
	cmd.Flags().StringVarP(&bName, "name", "n", "", "The name of the Board.")
	cmd.MarkFlagRequired("name")
	cmd.Flags().StringVarP(&bDescription, "description", "d", "",
		"A description of the Board.")
	cmd.Flags().StringVarP(&bColumnLayout, "column_layout", "c", "",
		"The number of columns to layout on the board.")
	cmd.Flags().StringArrayVar(&bQueryCaptions, "qc", nil,
		"Descriptive text to contextualize the value of the Query within the Board.")
	cmd.Flags().BoolSliceVar(&bQueryGraphSettingsHideMarkers, "qgshm", nil,
		"Hide markers on the graph.")
	cmd.Flags().BoolSliceVar(&bQueryGraphSettingsLogScale, "qgsls", nil,
		"Use a log scale, rather than a linear scale.")
	cmd.Flags().BoolSliceVar(&bQueryGraphSettingsOmitMissingValues, "qgsomv", nil,
		"Omit missing values from the graph.")
	cmd.Flags().BoolSliceVar(&bQueryGraphSettingsStackedGraphs, "qgssg", nil,
		"Display groups as stacked colored areas under their line graphs.")
	cmd.Flags().BoolSliceVar(&bQueryGraphSettingsUTCXAxis, "qgsux", nil,
		"Displays the X axis in Coordinated Universal Time, the time at 0° longitude.")
	cmd.Flags().BoolSliceVar(&bQueryGraphSettingsOverlaidCharts, "qgsoc", nil,
		"Combines any visualized AVG, MIN, MAX, and PERCENTILE clauses into a single chart.")
	cmd.Flags().StringArrayVar(&bQueryStyle, "qs", nil,
		"How the query should be displayed on the board. Enum: \"graph\" \"table\" \"combo\"")
	cmd.Flags().StringArrayVar(&bQueryDataset, "qd", nil,
		"The Dataset to Query. Required if using the deprecated query. Note: this field can take either name (\"My Dataset\") or slug (\"my_dataset\"); the response will always use the name.")
	cmd.Flags().StringArrayVar(&bQueryId, "qid", nil,
		"The ID of a Query object. Cannot be used with query. Query IDs can be retrieved from the UI or from the Query API.")
	cmd.Flags().StringArrayVar(&bQueryAnnotationId, "qaid", nil,
		"The ID of a Query Annotation that provides a name and description for the Query. The Query Annotation must apply to the query_id or query specified.")

	return cmd
}

func newBoardsDeleteCmd() *cobra.Command {
	var (
		bId string
	)

	cmd := &cobra.Command{
		Use:     "delete",
		Aliases: []string{"rm", "remove", "del"},
		Short:   "Delete a single Board by ID.",
		Long:    "Delete a single Board by ID.",
		Run: func(cmd *cobra.Command, args []string) {
			var p = Payload{
				Method: http.MethodDelete,
				Path:   "/1/boards/" + bId,
			}

			var err = p.Execute()
			if err != nil {
				log.WithFields(log.Fields{
					"_function": "newBoardsDeleteCmd",
					"err": err,
					"payload": p,
				}).Fatal("Error received when attempting to delete an existing board.")
			}
		},
	}

	cmd.Flags().StringVarP(&bId, "id", "i", "", "The unique identifier (ID) of a Board.")
	cmd.MarkFlagRequired("id")

	return cmd
}
