package commands

import (
	"bytes"
	"github.com/microsoft/abstrakt/internal/chartservice"
	"github.com/microsoft/abstrakt/internal/composeservice"
	"github.com/microsoft/abstrakt/internal/tools/logger"
	"github.com/spf13/cobra"
	"path"
	"strings"
)

var templateType string
var constellationFilePath string
var mapsFilePath string
var outputPath string

var composeCmd = &cobra.Command{
	Use:   "compose",
	Short: "Compose a package into requested template type",
	Long: `Compose is for composing a package based on mapsFilePath and constellationFilePath and template (default value is helm).

Example: abstrakt compose -t [template type] -f [constellationFilePath] -m [mapsFilePath] -o [outputPath]`,

	Run: func(cmd *cobra.Command, args []string) {

		if templateType == "helm" || templateType == "" {
			service := composeservice.NewComposeService()
			_ = service.LoadFromFile(constellationFilePath, mapsFilePath)
			chart, err := service.Compose("Output", outputPath)
			if err != nil {
				logger.Errorf("Could not compose: %v", err)
				return
			}

			err = chartservice.SaveChartToDir(chart, outputPath)

			if err != nil {
				logger.Errorf("There was an error saving the chart: %v", err)
				return
			}

			logger.Infof("Chart was saved to: %v", outputPath)

			var out *bytes.Buffer
			out, err = chartservice.BuildChart(path.Join(outputPath, "Output"))

			if err != nil {
				logger.Errorf("There was an error saving the chart: %v", err)
				return
			}

			logger.PrintBuffer(out, true)

		} else {
			logger.Errorf("Template type: %v is not known", templateType)
			return
		}

		logger.Debugf("args: %v", strings.Join(args, " "))
		logger.Debugf("template: %v", templateType)
		logger.Debugf("constellationFilePath: %v", constellationFilePath)
		logger.Debugf("mapsFilePath: %v", mapsFilePath)
		logger.Debugf("outputPath: %v", outputPath)
	},
}

func init() {
	composeCmd.Flags().StringVarP(&constellationFilePath, "constellationFilePath", "f", "", "constellation file path")
	_ = composeCmd.MarkFlagRequired("constellationFilePath")
	composeCmd.Flags().StringVarP(&mapsFilePath, "mapsFilePath", "m", "", "maps file path")
	_ = composeCmd.MarkFlagRequired("mapsFilePath")
	composeCmd.Flags().StringVarP(&outputPath, "outputPath", "o", "", "destination directory")
	_ = composeCmd.MarkFlagRequired("outputPath")
	composeCmd.Flags().StringVarP(&templateType, "template type", "t", "helm", "output template type")
}
