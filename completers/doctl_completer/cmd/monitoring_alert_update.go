package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var monitoring_alert_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an alert policy",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(monitoring_alert_updateCmd).Standalone()
	monitoring_alert_updateCmd.Flags().String("compare", "", "The comparator of the alert policy. Either `GreaterThan` or `LessThan`")
	monitoring_alert_updateCmd.Flags().String("description", "", "A description of the alert policy.")
	monitoring_alert_updateCmd.Flags().StringSlice("emails", []string{}, "Emails to send alerts to.")
	monitoring_alert_updateCmd.Flags().Bool("enabled", true, "Whether the alert policy is enabled.")
	monitoring_alert_updateCmd.Flags().StringSlice("entities", []string{}, "Entities to apply the alert against. (e.g. a droplet ID for a droplet alert policy)")
	monitoring_alert_updateCmd.Flags().StringSlice("slack-channels", []string{}, "Slack channels to send alerts to.")
	monitoring_alert_updateCmd.Flags().StringSlice("slack-urls", []string{}, "Slack URLs to send alerts to.")
	monitoring_alert_updateCmd.Flags().StringSlice("tags", []string{}, "Tags to apply the alert against.")
	monitoring_alert_updateCmd.Flags().String("type", "", "The type of alert policy.")
	monitoring_alert_updateCmd.Flags().Int("value", 0, "The value of the alert policy to compare against.")
	monitoring_alert_updateCmd.Flags().String("window", "5m", "The window to apply the alert policy conditions against.")
	monitoring_alertCmd.AddCommand(monitoring_alert_updateCmd)
}