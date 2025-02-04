package main

import (
	"bytes"

	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
	"github.com/spf13/cobra"
	"github.com/ugent-library/alma-go"
	"github.com/ugent-library/marc"
)

var getBibParams = alma.GetBibParams{}

var getItemsParams = alma.GetHoldingItemsParams{}

func init() {
	rootCmd.AddCommand(getBibCmd)
	getBibCmd.Flags().StringVar(&getBibParams.View, "view", "full", `"full" or "brief"`)
	getBibCmd.Flags().StringSliceVar(&getBibParams.Expand, "expand", nil, "")

	getBibCmd.AddCommand(getBibRecordCmd)

	getBibCmd.AddCommand(getHoldingsCmd)

	getBibCmd.AddCommand(getHoldingCmd)

	getHoldingCmd.AddCommand(getHoldingItemsCmd)
	getHoldingItemsCmd.Flags().IntVar(&getItemsParams.Limit, "limit", 10, "")
	getHoldingItemsCmd.Flags().IntVar(&getItemsParams.Offset, "offset", 0, "")
	getHoldingItemsCmd.Flags().StringVar(&getItemsParams.Expand, "expand", "", "")
	getHoldingItemsCmd.Flags().StringVar(&getItemsParams.UserID, "user-id", "", "")
	getHoldingItemsCmd.Flags().StringVar(&getItemsParams.CurrentLibrary, "current-library", "", "")
	getHoldingItemsCmd.Flags().StringVar(&getItemsParams.CurrentLocation, "current-location", "", "")
	getHoldingItemsCmd.Flags().StringVarP(&getItemsParams.Query, "query", "q", "", "")
	getHoldingItemsCmd.Flags().StringVar(&getItemsParams.OrderBy, "order-by", "", "")
	getHoldingItemsCmd.Flags().StringVar(&getItemsParams.Direction, "direction", "", "")
	getHoldingItemsCmd.Flags().StringVar(&getItemsParams.CreateDateFrom, "create-date-from", "", "")
	getHoldingItemsCmd.Flags().StringVar(&getItemsParams.CreateDateTo, "create-date-to", "", "")
	getHoldingItemsCmd.Flags().StringVar(&getItemsParams.ModifyDateFrom, "modify-date-from", "", "")
	getHoldingItemsCmd.Flags().StringVar(&getItemsParams.ModifyDateTo, "modify-date-to", "", "")
	getHoldingItemsCmd.Flags().StringVar(&getItemsParams.ReceiveDateFrom, "receive-date-from", "", "")
	getHoldingItemsCmd.Flags().StringVar(&getItemsParams.ReceiveDateTo, "receive-date-to", "", "")
	getHoldingItemsCmd.Flags().StringVar(&getItemsParams.ExpectedReceiveDateFrom, "expected-receive-date-from", "", "")
	getHoldingItemsCmd.Flags().StringVar(&getItemsParams.ExpectedReceiveDateTo, "expected-receive-date-to", "", "")
	getHoldingItemsCmd.Flags().StringVar(&getItemsParams.View, "view", "brief", `"full" or "brief"`)
}

var getBibCmd = &cobra.Command{
	Use:   "bib [mms-id]",
	Short: "Get bib",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		resBody, err := almaClient.RawGetBib(cmd.Context(), args[0], getBibParams)
		if err != nil {
			return err
		}

		return writeJSON(cmd, resBody)
	},
}

var getBibRecordCmd = &cobra.Command{
	Use:   "record [mms-id]",
	Short: "Get bib record only",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		resData, err := almaClient.GetBib(cmd.Context(), args[0], alma.GetBibParams{})
		if err != nil {
			return err
		}

		if prettify {
			dec := marc.NewMARCXMLDecoder(bytes.NewBuffer([]byte(resData.Record())))
			rec, err := dec.Decode()
			if err != nil {
				return err
			}

			colTag := "tag"
			colInd1 := "ind1"
			colInd2 := "ind2"
			colVal := "val"

			sfStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#88f"))

			cols := []table.Column{
				table.NewColumn(colTag, "Tag", 5).WithStyle(
					lipgloss.NewStyle().
						Foreground(lipgloss.Color("#88f")),
				),
				table.NewColumn(colInd1, "Ind 1", 5),
				table.NewColumn(colInd2, "Ind 2", 5),
				table.NewColumn(colVal, "Value", 65),
			}

			rows := []table.Row{
				table.NewRow(table.RowData{
					colTag: "LDR",
					colVal: rec.Leader,
				}),
			}
			for _, field := range rec.ControlFields {
				rows = append(rows, table.NewRow(table.RowData{
					colTag: field.Tag,
					colVal: field.Value,
				}))
			}
			for _, field := range rec.DataFields {
				val := ""
				for _, sf := range field.SubFields {
					val += sfStyle.Render(sf.Code) + sf.Value
				}
				rows = append(rows, table.NewRow(table.RowData{
					colTag:  field.Tag,
					colInd1: field.Ind1,
					colInd2: field.Ind2,
					colVal:  val,
				}))
			}

			t := table.New(cols).
				WithRows(rows).
				HeaderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("10")).Bold(true)).
				Focused(true).
				WithBaseStyle(
					lipgloss.NewStyle().
						BorderForeground(lipgloss.Color("#a38")).
						Align(lipgloss.Left),
				).
				WithMultiline(true)

			return writeString(cmd, t.View())
		}

		return writeString(cmd, resData.Record())
	},
}

var getHoldingsCmd = &cobra.Command{
	Use:   "holdings [mms-id]",
	Short: "Get bib holdings",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		resBody, err := almaClient.RawGetHoldings(cmd.Context(), args[0])
		if err != nil {
			return err
		}

		return writeJSON(cmd, resBody)
	},
}

var getHoldingCmd = &cobra.Command{
	Use:   "holding [mms-id] [holding-id]",
	Short: "Get bib holding",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		resBody, err := almaClient.RawGetHolding(cmd.Context(), args[0], args[1])
		if err != nil {
			return err
		}

		return writeJSON(cmd, resBody)
	},
}

var getHoldingItemsCmd = &cobra.Command{
	Use:   "items [mms-id] [holding-id]",
	Short: "Get bib holding items",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		resBody, err := almaClient.RawGetHoldingItems(cmd.Context(), args[0], args[1], getItemsParams)
		if err != nil {
			return err
		}

		return writeJSON(cmd, resBody)
	},
}
