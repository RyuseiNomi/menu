package src

import (
	"log"

	t "github.com/RyuseiNomi/menu/src/tui"
	"github.com/rivo/tview"
)

func Handle() {
	app := tview.NewApplication()
	page := tview.NewPages()
	list := tview.NewList().
		AddItem("ずかん", "", 'd', nil).
		AddItem("コンテナ", "", 'c', func() {
			ctw := t.NewContainerTuiWorker(page, app)
			ctw.Handle()
		}).
		AddItem("アプリ", "", 'a', nil).
		AddItem("レポート", "", 'r', nil).
		AddItem("せってい", "", 's', nil).
		AddItem("とじる", "", 'q', func() {
			app.Stop()
			log.Println("aaa")
		})
	page.AddPage("list", list, true, true)
	if err := app.SetRoot(page, true).Run(); err != nil {
		panic(err)
	}
}
