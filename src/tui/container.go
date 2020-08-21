package tui

import (
	"os/exec"
	"strings"

	"github.com/rivo/tview"
)

const (
	containerIDIndex = 0
	namesIndex       = 1
	statusIndex      = 2
)

// container リストに表示するコンテナモデル
type container struct {
	id     string
	name   string
	status string
}

// containers 複数のContainerモデルを持つ構造体
type containers []container

// containerTuiWorker コンテナ一覧表示プロセスに関わる情報を保持する構造体
type containerTuiWorker struct {
	page *tview.Pages
	app  *tview.Application
}

// NewContainerTuiWorker コンテナ一覧表示プロセスに関わる情報を保持する構造体を返却する
func NewContainerTuiWorker(p *tview.Pages, a *tview.Application) *containerTuiWorker {
	return &containerTuiWorker{
		page: p,
		app:  a,
	}
}

// Handle コンテナ一覧画面に関する操作の流れを集約しているメソッド
func (ctw *containerTuiWorker) HandleContainer() {
	cs, err := getContainers()
	if err != nil {
		panic(err)
	}
	// コンテナ一覧を表示するためのリスト
	list := tview.NewList()
	for _, container := range cs {
		list.AddItem(container.name, container.status, 'a', func() {
			// コンテナが選択された時に出現するモーダル
			modal := getModal(container, ctw.page)
			ctw.page.AddPage("modal", modal, true, true)
		})
	}
	list.AddItem("Quit", "Press to exit", 'q', func() {
		ctw.page.RemovePage("c-list")
	})
	ctw.page.AddPage("c-list", list, true, true)
}

// ------------------ TUIに関するメソッド ------------------------

// getModal モーダルをアプリのページに表示する
func getModal(c container, p *tview.Pages) *tview.Modal {
	// モーダル
	modal := tview.NewModal().
		SetText("What do you want to next?").
		AddButtons([]string{"Start", "Stop", "Cancel"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Cancel" {
				p.RemovePage("modal")
				return
			}
			if buttonLabel == "Start" {
				if err := startContainer(c); err != nil {
					panic(err)
				}
				p.RemovePage("modal")

				// 起動完了モーダル
				completeModal := tview.NewModal().
					SetText("Completed to start the container!").
					AddButtons([]string{"Ok"}).
					SetDoneFunc(func(buttonIndex int, buttonLabel string) {
						if buttonLabel == "Ok" {
							p.RemovePage("completeModal")
						}
					})
				p.AddPage("completeModal", completeModal, true, true)
			}
			if buttonLabel == "Stop" {
				if err := stopContainer(c); err != nil {
					panic(err)
				}
				p.RemovePage("modal")

				// 消去完了モーダル
				completeModal := tview.NewModal().
					SetText("Completed to stop the container!").
					AddButtons([]string{"Ok"}).
					SetDoneFunc(func(buttonIndex int, buttonLabel string) {
						if buttonLabel == "Ok" {
							p.RemovePage("completeModal")
						}
					})
				p.AddPage("completeModal", completeModal, true, true)
			}
		})
	return modal
}

// ------------------ 外部コマンドによるコンテナ操作メソッド ------------------------

// getContainers 外部コマンドを実行し、Dockerコンテナを取得する
func getContainers() (containers, error) {

	cs := containers{}

	// 外部コマンドの実行よりコンテナ一覧を取得
	out, err := exec.Command("docker", "ps", "-a", "--format", "\"{{.ID}} {{.Names}} {{.Status}}\"").Output()
	if err != nil {
		return nil, err
	}

	// 標準出力より情報を取得し、containerモデルに情報を追加する
	slice := strings.Split(string(out), "\n")
	slice = slice[:len(slice)-1] // 改行区切りのため、末尾は空なので削除
	for _, c := range slice {
		container := container{}
		elements := strings.Split(c, " ")
		for i, e := range elements {
			if i == containerIDIndex {
				container.id = e
			}
			if i == namesIndex {
				container.name = e
			}
			if i == statusIndex {
				container.status = e
			}
		}
		cs = append(cs, container)
	}

	return cs, nil
}

// startContainer コンテナを起動する外部コマンドを実行する
func startContainer(c container) error {
	_, err := exec.Command("docker", "start", c.name).Output()
	if err != nil {
		return err
	}
	return nil
}

// stopContainer コンテナを停止する外部コマンドを実行する
func stopContainer(c container) error {
	_, err := exec.Command("docker", "stop", c.name).Output()
	if err != nil {
		return err
	}
	return nil
}
