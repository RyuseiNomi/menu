package tui

import (
	"os/exec"
	"regexp"

	"github.com/rivo/tview"
)

type reportTuiWorker struct {
	page *tview.Pages
	app  *tview.Application
}

func NewReportTuiWorker(p *tview.Pages, a *tview.Application) *reportTuiWorker {
	return &reportTuiWorker{
		page: p,
		app:  a,
	}
}

func (rtw *reportTuiWorker) HandleReport() {
	showModal(rtw)
}

func showModal(rtw *reportTuiWorker) {
	modal := tview.NewModal().
		SetText("現在の差分をStashします。よろしいですか？").
		AddButtons([]string{"Ok", "Cancel"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Ok" {
				rtw.app.Stop()
				// shellの実行
				err := exec.Command("sh", "../shell/report.sh").Run()
				if err != nil {
					panic(err)
				}
			} else if buttonLabel == "Cancel" {
				rtw.page.RemovePage("modal")
			}
		})
	rtw.page.AddPage("modal", modal, true, true)
}

func isUnderGitDirectory(rtw *reportTuiWorker) {
	out, err := exec.Command("git", "status").Output()
	if err != nil {
		panic(err)
	}
	r := regexp.MustCompile("/fatal/")
	v := r.FindString(string(out))
	if v != "" {
		errModal := tview.NewModal().
			SetText("git管理下のディレクトリに移動してください。").
			AddButtons([]string{"Ok"}).
			SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				if buttonLabel == "Ok" {
					rtw.page.RemovePage("errModal")
				}
			})
		rtw.page.AddPage("errModal", errModal, true, true)
	}
}
