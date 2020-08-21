package tui

import (
	"os/exec"
	"regexp"
	"strings"

	"github.com/gdamore/tcell"

	"github.com/rivo/tview"
)

// language 言語情報を持つ構造体
type language struct {
	name string
	ver  string
}

// languages 複数の言語情報の構造体を持つ
type languages []language

// dictionaryTuiWorker ずかん操作のTUI画面に関するプロパティを持つ構造体
type dictionaryTuiWorker struct {
	page *tview.Pages
	app  *tview.Application
}

// NewDictionaryTuiWorker ずかん操作のTUI画面に関するプロパティを持つ構造体を生成する
func NewDictionaryTuiWorker(p *tview.Pages, a *tview.Application) *dictionaryTuiWorker {
	return &dictionaryTuiWorker{
		page: p,
		app:  a,
	}
}

func (dtw *dictionaryTuiWorker) HandleDictionary() {
	ls := languages{}
	targetLang := []string{"Java", "Go", "Ruby", "PHP"}
	for _, t := range targetLang {
		l := language{}
		l.name = t
		ls = append(ls, l)
	}
	ls.getVersion()
	ls.createTable(dtw)
}

func (ls languages) getVersion() {
	for _, lang := range ls {
		if lang.name == "Java" {
			lang.ver = extractVersion("java", "-version")
		} else if lang.name == "Go" {
			lang.ver = extractVersion("go", "version")
		} else if lang.name == "Ruby" {
			lang.ver = extractVersion("ruby", "--version")
		} else if lang.name == "PHP" {
			lang.ver = extractVersion("php", "--version")
		}
	}
}

func (ls languages) createTable(dtw *dictionaryTuiWorker) {
	table := tview.NewTable().SetBorders(true)
	cols, rows := 2, 5
	word := 0
	for c := 0; c < cols; c++ {
		for r := 0; r < rows; r++ {
			color := tcell.ColorWhite
			//if c < 1 || r < 1 {
			//	color = tcell.ColorYellow
			//}
			if c < 1 {
				table.SetCell(r, c,
					tview.NewTableCell(ls[word].name).
						SetTextColor(color).
						SetAlign(tview.AlignCenter))
				word = (word + 1) % len(ls)
			} else {
				table.SetCell(r, c,
					tview.NewTableCell(ls[word].ver).
						SetTextColor(color).
						SetAlign(tview.AlignCenter))
				word = (word + 1) % len(ls)
			}
		}
	}
	dtw.page.AddPage("ver-list", table, true, true)
}

func extractVersion(lang string, fmt string) string {
	out, err := exec.Command(lang, fmt).Output()
	if err != nil {
		panic(err)
	}

	slice := strings.Split(string(out), "\n") //改行区切りで出力を取得
	vstr := slice[0]
	r := regexp.MustCompile("(^.*?)[0-9]+\\.[0-9]+")
	v := r.FindString(vstr)
	return v
}
