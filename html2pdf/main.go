package main

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/html2pdf", func(writer http.ResponseWriter, request *http.Request) {
		request.ParseForm()
		targetUrl := request.FormValue("url")
		err := ChromedpPrintPdf(targetUrl, "./file.pdf")
		if err != nil {
			fmt.Println(err)
			return
		}
	})

	http.ListenAndServe(":8081", nil)
	ch := make(chan int)
	<-ch

}

func ChromedpPrintPdf(url string, to string) error {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var buf []byte
	tests := []struct {
		sel    string
		by     chromedp.QueryOption
		width  int64
		height int64
	}{
		{`.titlefold`, chromedp.ByQueryAll, 239, 239},
	}

	err := chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.ActionFunc(func(ctx context.Context) error {
			for _, test := range tests {
				var nodeList []*cdp.Node
				err := chromedp.Run(ctx, chromedp.Nodes(test.sel,&nodeList,test.by))
				if err != nil {
					return fmt.Errorf("chromedp Run failed,err:%+v", err)

				}
				for i, node := range nodeList{
					model,err := dom.GetBoxModel().WithNodeID(node.NodeID).Do(ctx)
					if err != nil {
						 fmt.Println("chromedp GetBoxModel failed,err", err)
						 continue
					}
					log.Printf("第%d个测试，长:%d，宽:%d\n", i, model.Width, model.Height)
				}

			}
			return nil

		}),
		chromedp.ActionFunc(func(ctx context.Context) error {
			err := chromedp.Run(ctx, chromedp.Tasks{
				chromedp.Navigate(url),
				chromedp.ActionFunc(func(ctx context.Context) error {
					var err error
					buf, _, err = page.PrintToPDF().WithPrintBackground(true).
						Do(ctx)
					return err
				}),
			})
			if err != nil {
				return fmt.Errorf("chromedp Run failed,err:%+v", err)
			}
			if err := ioutil.WriteFile(to, buf, 0644); err != nil {
				return fmt.Errorf("write to file failed,err:%+v", err)
			}
			return nil
		}),
	})
	if err != nil {
		return err
	}



	return nil
}
