package cmd

import (
	"encoding/json"
	"net/http"

	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"
)

type boardGraphSettings struct {
	HideMarkers       bool `json:"hide_markers"`
	LogScale          bool `json:"log_scale"`
	OmitMissingValues bool `json:"omit_missing_values"`
	StackedGraphs     bool `json:"stacked_graphs"`
	UTCXAxis          bool `json:"utc_xaxis"`
	OverlaidCharts    bool `json:"overlaid_charts"`
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
		Use:     "boards",
		Aliases: []string{"b"},
		Short:   "Manage Boards",
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
	cmd.AddCommand(newBoardsAddQueryCmd())
	// cmd.AddCommand(newBoardsUpdateQueryCmd())
	// cmd.AddCommand(newBoardsDeleteQueryCmd())

	return cmd
}

// Add a query to a Board
// CUSTOM
func newBoardsAddQueryCmd() *cobra.Command {
	var (
		bId                                  string
		bQueryCaption                        string
		bQueryGraphSettingsHideMarkers       bool
		bQueryGraphSettingsLogScale          bool
		bQueryGraphSettingsOmitMissingValues bool
		bQueryGraphSettingsStackedGraphs     bool
		bQueryGraphSettingsUTCXAxis          bool
		bQueryGraphSettingsOverlaidCharts    bool
		bQueryStyle                          string
		bQueryDataset                        string
		bQueryId                             string
		bQueryAnnotationId                   string
	)

	cmd := &cobra.Command{
		Use:     "add_query",
		Aliases: []string{"aq"},
		Short:   "Add a Query to a Board.",
		Long:    "Add a Query to a Board.",
		Run: func(cmd *cobra.Command, args []string) {
			// Get the board first, so we can append a new query to it.
			var pGet = Payload{
				Method:   http.MethodGet,
				Path:     "/1/boards/" + bId,
				Response: &board{},
			}

			var err = pGet.GetResponse()
			if err != nil {
				log.WithFields(log.Fields{
					"_function": "newBoardsAddQueryCmd",
					"err":       err,
					"payload":   pGet,
				}).Fatal("Error received when attempting to get the board to update.")
			}

			// Update the returned board with the new query.
			pGet.Response.(*board).Queries = append(pGet.Response.(*board).Queries, boardQuery{
				Caption: bQueryCaption,
				GraphSettings: boardGraphSettings{
					HideMarkers:       bQueryGraphSettingsHideMarkers,
					LogScale:          bQueryGraphSettingsLogScale,
					OmitMissingValues: bQueryGraphSettingsOmitMissingValues,
					StackedGraphs:     bQueryGraphSettingsStackedGraphs,
					UTCXAxis:          bQueryGraphSettingsUTCXAxis,
					OverlaidCharts:    bQueryGraphSettingsOverlaidCharts,
				},
				QueryStyle:        bQueryStyle,
				Dataset:           bQueryDataset,
				QueryID:           bQueryId,
				QueryAnnotationID: bQueryAnnotationId,
			})

			bodyMarshal, err := json.Marshal(pGet.Response)
			if err != nil {
				log.WithFields(log.Fields{
					"_function": "newBoardsAddQueryCmd",
					"err":       err,
					"board":     pGet.Response,
				}).Fatal("Error received when attempting to marshal a board.")
			}
			var pPut = Payload{
				Method:   http.MethodPut,
				Path:     "/1/boards/" + bId,
				Body:     bodyMarshal,
				Response: &board{},
			}

			err = pPut.PrintResponse()
			if err != nil {
				log.WithFields(log.Fields{
					"_function": "newBoardsAddQueryCmd",
					"err":       err,
					"payload":   pPut,
				}).Fatal("Error received when attempting to add a new query to a board.")
			}
		},
	}

	cmd.Flags().StringVarP(&bId, "id", "i", "",
		"The unique identifier (ID) of a Board.")
	cmd.MarkFlagRequired("id")
	cmd.Flags().StringVarP(&bQueryCaption, "caption", "c", "",
		"Descriptive text to contextualize the value of the Query within the Board.")
	cmd.Flags().BoolVarP(&bQueryGraphSettingsHideMarkers, "hide_mdarkers", "H", false,
		"Hide markers on the graph.")
	cmd.Flags().BoolVarP(&bQueryGraphSettingsLogScale, "log_scale", "L", false,
		"Use a log scale, rather than a linear scale.")
	cmd.Flags().BoolVarP(&bQueryGraphSettingsOmitMissingValues, "omit_missing", "O", false,
		"Omit missing values from the graph.")
	cmd.Flags().BoolVarP(&bQueryGraphSettingsStackedGraphs, "stacked_graphs", "S", false,
		"Display groups as stacked colored areas under their line graphs.")
	cmd.Flags().BoolVarP(&bQueryGraphSettingsUTCXAxis, "utx_axis", "U", false,
		"Displays the X axis in Coordinated Universal Time, the time at 0° longitude.")
	cmd.Flags().BoolVarP(&bQueryGraphSettingsOverlaidCharts, "overlaid_charts", "V", false,
		"Combines any visualized AVG, MIN, MAX, and PERCENTILE clauses into a single chart.")
	cmd.Flags().StringVarP(&bQueryStyle, "style", "s", "",
		"How the query should be displayed on the board. Enum: \"graph\" \"table\" \"combo\"")
	cmd.Flags().StringVarP(&bQueryDataset, "dataset", "d", "",
		"The Dataset to Query. Required if using the deprecated query. Note: this field can take either name (\"My Dataset\") or slug (\"my_dataset\"); the response will always use the name.")
	cmd.Flags().StringVarP(&bQueryId, "query_id", "q", "",
		"The ID of a Query object. Cannot be used with query. Query IDs can be retrieved from the UI or from the Query API.")
	cmd.Flags().StringVarP(&bQueryAnnotationId, "annotation_id", "a", "",
		"The ID of a Query Annotation that provides a name and description for the Query. The Query Annotation must apply to the query_id or query specified.")

	return cmd
}

// Create a Board
// https://docs.honeycomb.io/api/tag/Boards#operation/createBoard
func newBoardsCreateCmd() *cobra.Command {
	var (
		bName         string
		bDescription  string
		bColumnLayout string
	)

	cmd := &cobra.Command{
		Use:     "create",
		Aliases: []string{"add", "new"},
		Short:   "Create a Board.",
		Long:    "Create a Board without any Queries - these can be added after creation.",
		Run: func(cmd *cobra.Command, args []string) {
			var b = board{
				Name:         bName,
				Description:  bDescription,
				ColumnLayout: bColumnLayout,
			}

			var bodyMarshal, err = json.Marshal(b)
			if err != nil {
				log.WithFields(log.Fields{
					"_function": "newBoardsCreateCmd",
					"err":       err,
					"board":     b,
				}).Fatal("Error received when attempting to marshal a board.")
			}
			var p = Payload{
				Method:   http.MethodPost,
				Path:     "/1/boards",
				Body:     bodyMarshal,
				Response: &board{},
			}

			err = p.PrintResponse()
			if err != nil {
				log.WithFields(log.Fields{
					"_function": "newBoardsCreateCmd",
					"err":       err,
					"payload":   p,
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
				Method:   http.MethodGet,
				Path:     "/1/boards",
				Response: &[]board{},
			}

			var err = p.PrintResponse()
			if err != nil {
				log.WithFields(log.Fields{
					"_function": "newBoardsListCmd",
					"err":       err,
					"payload":   p,
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
				Method:   http.MethodGet,
				Path:     "/1/boards/" + bId,
				Response: &board{},
			}

			var err = p.PrintResponse()
			if err != nil {
				log.WithFields(log.Fields{
					"_function": "newBoardsGetCmd",
					"err":       err,
					"payload":   p,
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
	)

	cmd := &cobra.Command{
		Use:     "update",
		Aliases: []string{"up", "edit", "modify", "change", "set"},
		Short:   "Update a Board by ID.",
		Long:    "Update a Board by ID, leaving existing queries as-is.",
		Run: func(cmd *cobra.Command, args []string) {
			// Get the board first, so we can append a new query to it.
			var pGet = Payload{
				Method:   http.MethodGet,
				Path:     "/1/boards/" + bId,
				Response: &board{},
			}

			var err = pGet.GetResponse()
			if err != nil {
				log.WithFields(log.Fields{
					"_function": "newBoardsUpdateCmd",
					"err":       err,
					"payload":   pGet,
				}).Fatal("Error received when attempting to get the board to update.")
			}

			var b = board{
				Name:         bName,
				Description:  bDescription,
				ColumnLayout: bColumnLayout,
				Queries:      pGet.Response.(*board).Queries,
			}

			bodyMarshal, err := json.Marshal(b)
			if err != nil {
				log.WithFields(log.Fields{
					"_function": "newBoardsUpdateCmd",
					"err":       err,
					"board":     b,
				}).Fatal("Error received when attempting to marshal a board.")
			}
			var pPut = Payload{
				Method:   http.MethodPut,
				Path:     "/1/boards/" + bId,
				Body:     bodyMarshal,
				Response: &board{},
			}

			err = pPut.PrintResponse()
			if err != nil {
				log.WithFields(log.Fields{
					"_function": "newBoardsUpdateCmd",
					"err":       err,
					"payload":   pPut,
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
				Method:   http.MethodDelete,
				Path:     "/1/boards/" + bId,
				Response: nil,
			}

			var err = p.PrintResponse()
			if err != nil {
				log.WithFields(log.Fields{
					"_function": "newBoardsDeleteCmd",
					"err":       err,
					"payload":   p,
				}).Fatal("Error received when attempting to delete an existing board.")
			}
		},
	}

	cmd.Flags().StringVarP(&bId, "id", "i", "", "The unique identifier (ID) of a Board.")
	cmd.MarkFlagRequired("id")

	return cmd
}
