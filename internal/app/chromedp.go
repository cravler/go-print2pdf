package app

import (
	"time"
	"context"

	"github.com/chromedp/chromedp"
	"github.com/chromedp/cdproto/page"
)

func NewPrintToPDF() *page.PrintToPDFParams {
	return &page.PrintToPDFParams{
		// Scale of the webpage rendering.
		Scale: 1,
		// Paper width in inches.
		PaperWidth: 8.5,
		// Paper height in inches.
		PaperHeight: 11,
		// Top margin in inches, 0.4 inches = ~1cm.
		MarginTop: 0,
		// Bottom margin in inches, 0.4 inches = ~1cm.
		MarginBottom: 0,
		// Left margin in inches, 0.4 inches = ~1cm.
		MarginLeft: 0,
		// Right margin in inches, 0.4 inches = ~1cm.
		MarginRight: 0,
		// Paper ranges to print, e.g., '1-5, 8, 11-13'.
		// Defaults to the empty string, which means print all pages.
		PageRanges: "",
		// HTML template for the print header.
		// Should be valid HTML markup with following classes used to inject printing values into them:
		//  - date: formatted print date
		//  - title: document title
		//  - url: document location
		//  - pageNumber: current page number
		//  - totalPages: total pages in the document
		// For example, <span class=title></span> would generate span containing the title.
		HeaderTemplate: "",
		// HTML template for the print footer.
		// Should use the same format as the headerTemplate.
		FooterTemplate: "",
		// Paper orientation.
		Landscape: false,
		// Print background graphics.
		PrintBackground: false,
		// Whether or not to prefer page size as defined by css.
		// Defaults to false, in which case the content will be scaled to fit the paper size.
		PreferCSSPageSize: false,
		// Display header and footer.
		DisplayHeaderFooter: false,
		// Whether to silently ignore invalid but successfully parsed page ranges, such as '3-2'.
		IgnoreInvalidPageRanges: false,
	}
}

func GeneratePDF(url string, printToPDF *page.PrintToPDFParams, timeout int) ([]byte, error) {
	opts := append(
		chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("ignore-certificate-errors", true),
	)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if timeout > 0 {
		ctx, cancel = context.WithTimeout(ctx, time.Second * time.Duration(timeout))
		defer cancel()
	}

	ctx, cancel = chromedp.NewExecAllocator(ctx, opts...)
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	var pdf []byte
	if err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			pdf, _, err = printToPDF.Do(ctx)
			return err
		}),
	); err != nil {
		return pdf, err
	}
	return pdf, nil
}