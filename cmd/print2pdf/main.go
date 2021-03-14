package main

import (
	"os"
	"fmt"
	"path"

	"github.com/spf13/cobra"
	"github.com/cravler/go-print2pdf/internal/app"
)

var version = "0.x-dev"

func main() {
	printToPDF := app.NewPrintToPDF()

	rootCmdName := path.Base(os.Args[0])
	rootCmd := app.NewRootCmd(rootCmdName + " <URL>", version, func(c *cobra.Command, args []string) error {
		timeout, _ := c.Flags().GetInt("timeout")

		printToPDF.Scale, _ = c.Flags().GetFloat64("scale")
		printToPDF.PaperWidth, _ = c.Flags().GetFloat64("paper-width")
		printToPDF.PaperHeight, _ = c.Flags().GetFloat64("paper-height")
		printToPDF.MarginTop, _ = c.Flags().GetFloat64("margin-top")
		printToPDF.MarginBottom, _ = c.Flags().GetFloat64("margin-bottom")
		printToPDF.MarginLeft, _ = c.Flags().GetFloat64("margin-left")
		printToPDF.MarginRight, _ = c.Flags().GetFloat64("margin-right")

		printToPDF.PageRanges, _ = c.Flags().GetString("page-ranges")
		printToPDF.HeaderTemplate, _ = c.Flags().GetString("header-template")
		printToPDF.FooterTemplate, _ = c.Flags().GetString("footer-template")

		printToPDF.Landscape, _ = c.Flags().GetBool("landscape")
		printToPDF.PrintBackground, _ = c.Flags().GetBool("print-background")
		printToPDF.PreferCSSPageSize, _ = c.Flags().GetBool("prefer-css-page-size")
		printToPDF.DisplayHeaderFooter, _ = c.Flags().GetBool("display-header-footer")
		printToPDF.IgnoreInvalidPageRanges, _ = c.Flags().GetBool("ignore-invalid-page-ranges")

		pdf, err := app.GeneratePDF(args[0], printToPDF, timeout)
		if err != nil {
			return err
		}

		fmt.Printf("%s", pdf)

		return nil
	})

	rootCmd.Flags().SortFlags = false

	rootCmd.Flags().Int("timeout", 0, "Process timeout in seconds")

	rootCmd.Flags().Float64("scale", printToPDF.Scale, "Scale of the webpage rendering")
	rootCmd.Flags().Float64("paper-width", printToPDF.PaperWidth, "Paper width in inches")
	rootCmd.Flags().Float64("paper-height", printToPDF.PaperHeight, "Paper height in inches")
	rootCmd.Flags().Float64("margin-top", printToPDF.MarginTop, "Top margin in inches")
	rootCmd.Flags().Float64("margin-bottom", printToPDF.MarginBottom, "Bottom margin in inches")
	rootCmd.Flags().Float64("margin-left", printToPDF.MarginLeft, "Left margin in inches")
	rootCmd.Flags().Float64("margin-right", printToPDF.MarginRight, "Right margin in inches")

	rootCmd.Flags().String("page-ranges", printToPDF.PageRanges, "Paper ranges to print, e.g., '1-5, 8, 11-13'")
	rootCmd.Flags().String("header-template", printToPDF.HeaderTemplate, "HTML template for the print header")
	rootCmd.Flags().String("footer-template", printToPDF.FooterTemplate, "HTML template for the print footer")

	rootCmd.Flags().BoolP("landscape", "l", printToPDF.Landscape, "Landscape paper orientation")
	rootCmd.Flags().BoolP("print-background", "b", printToPDF.PrintBackground, "Print background graphics")
	rootCmd.Flags().BoolP("prefer-css-page-size", "c", printToPDF.PreferCSSPageSize, "Prefer page size defined by CSS")
	rootCmd.Flags().BoolP("display-header-footer", "d", printToPDF.DisplayHeaderFooter, "Display header and footer")
	rootCmd.Flags().BoolP("ignore-invalid-page-ranges", "i", printToPDF.IgnoreInvalidPageRanges, "Silently ignore invalid but successfully parsed page ranges, such as '3-2'")

	app.ApplyDefaultFlags(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		rootCmd.Println(err)
		os.Exit(1)
	}
}