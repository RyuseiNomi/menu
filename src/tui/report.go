package tui

import (
	"bufio"
	"fmt"
	"os"
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
		SetText("これまでの 作業内容を stashに 書き残しますか？").
		AddButtons([]string{"はい", "いいえ"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "はい" {
				// git stash saveの実行
				rtw.app.Stop()
				stash()
			} else if buttonLabel == "いいえ" {
				rtw.page.RemovePage("modal")
			}
		})
	rtw.page.AddPage("modal", modal, true, true)
}

func stash() error {
	fmt.Println("以下の差分をStashします。")
	fmt.Println("-----------------------\n ")
	out, err := exec.Command("git", "status").Output()
	if err != nil {
		return err
	}
	r := regexp.MustCompile(`modified:  .*`)
	files := r.FindAllStringSubmatch(string(out), -1)
	for i := range files {
		fmt.Printf("\x1b[31m%s\x1b[0m \n", files[i])
	}
	fmt.Println("\n-----------------------")
	fmt.Println("stashに対する説明文を入力してください。")
	fmt.Print("説明:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	out, err = exec.Command("git", "stash", "save", text).Output()
	if err != nil {
		panic(err)
	}
	fmt.Print("stashが完了しました。")
	return nil
}

func isUnderGitDirectory(rtw *reportTuiWorker) {
	out, err := exec.Command("git", "status").Output()
	if err != nil {
		panic(err)
	}
	r := regexp.MustCompile("fatal")
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
