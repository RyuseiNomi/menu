package src

import (
	"os/exec"

	t "github.com/RyuseiNomi/menu/src/tui"
	"github.com/rivo/tview"
)

// Handle 画面に表示するリストの描画と各イベントの呼び出しを行う
func Handle() {
	app := tview.NewApplication()
	page := tview.NewPages()
	list := tview.NewList().
		//AddItem("ずかん", "", 'd', func() {
		//	dtw := t.NewDictionaryTuiWorker(page, app)
		//	dtw.HandleDictionary()
		//}).
		AddItem("コンテナ", "", 'c', func() {
			ctw := t.NewContainerTuiWorker(page, app)
			ctw.HandleContainer()
		}).
		AddItem("アプリ", "", 'a', func() {
			app.Stop()
			_, err := exec.Command("open", "/Applications").Output()
			if err != nil {
				panic(err)
			}
		}).
		AddItem("ユーザ", "", 's', func() {
			// TODO シェル等を用いてユーザ名を取得する
			app.Stop()
			_, err := exec.Command("open", "/System/Library/PreferencePanes/Accounts.prefPane/").Output()
			if err != nil {
				panic(err)
			}
		}).
		AddItem("レポート", "", 'r', func() {
			rtw := t.NewReportTuiWorker(page, app)
			rtw.HandleReport()
		}).
		AddItem("とじる", "", 'q', func() {
			app.Stop()
		})
	page.AddPage("list", list, true, true)
	if err := app.SetRoot(page, true).Run(); err != nil {
		panic(err)
	}
}
