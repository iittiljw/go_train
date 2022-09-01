package lib

import (
	"context"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"os"
)

var listCmd = &cobra.Command{
	Use:          "list",
	Short:        "list pods ",
	Example:      "kubectl pods list [flags]",
	SilenceUsage: true,
	RunE: func(c *cobra.Command, args []string) error {
		ns, err := c.Flags().GetString("namespace")
		if err != nil {
			return err
		}
		if ns == "" {
			ns = "default"
		}
		
		list, err := client.CoreV1().Pods(ns).List(context.Background(), v1.ListOptions{LabelSelector: Labels, FieldSelector: Fields})
		if err != nil {
			return err
		}
		
		FilterListByJSON(list)
		table := tablewriter.NewWriter(os.Stdout)
		//设置头
		table.SetHeader(InitHeader(table))
		for _, pod := range list.Items {
			podRow := []string{pod.Name, pod.Namespace, pod.Status.PodIP,
				string(pod.Status.Phase)}
			if ShowLabels {
				podRow = append(podRow, Map2String(pod.Labels))
			}
			table.Append(podRow)
		}
		table.SetAutoWrapText(false)
		table.SetAutoFormatHeaders(true)
		table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
		table.SetAlignment(tablewriter.ALIGN_LEFT)
		table.SetCenterSeparator("")
		table.SetColumnSeparator("")
		table.SetRowSeparator("")
		table.SetHeaderLine(false)
		table.SetBorder(false)
		table.SetTablePadding("\t") // pad with tabs
		table.SetNoWhiteSpace(true)
		table.Render()
		return nil
	},
}
