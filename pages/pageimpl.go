package pages

import (
	"bufio"
	"fmt"
	"gamewebserver/handleerror"
	"io"
	"os"
	"strings"
	"sync"
)

type Pager struct {
	pages map[string]string
	mutex sync.Mutex
}

func CreatePager() *Pager {
	p := new(Pager)
	p.pages = make(map[string]string, 0)
	p.findAllPages()
	return p
}

func (p *Pager) GetPage(nameFile string) string {
	if _, ok := p.pages[nameFile]; ok {
		return p.pages[nameFile]
	}
	panic("no page")
}

func (p *Pager) findAllPages() {
	// path, err := os.Executable()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(path)
	path := `/home/zhenya/tests/go/game/webserver/html`
	dirs, err := os.ReadDir(path)
	handleerror.HandleErrorWithPanic(err)
	var wg sync.WaitGroup
	for _, d := range dirs {
		if d.IsDir() {
			continue
		}
		if !strings.Contains(d.Name(), "html") {
			continue
		}
		wg.Add(1)
		go p.readFile(&wg, path+string(os.PathSeparator)+d.Name(), d.Name())
	}

	wg.Wait()

	for _, html := range p.pages {
		fmt.Println(html)
	}
}

func (p *Pager) readFile(wg *sync.WaitGroup, path, namePage string) {
	defer wg.Done()
	file, err := os.Open(path)
	handleerror.HandleErrorWithPanic(err)

	reader := bufio.NewReader(file)
	var builder strings.Builder
	for {
		line, _, err := reader.ReadLine()
		if err != nil && err == io.EOF {
			break
		}
		builder.WriteString(string(line))
	}

	parts := strings.Split(namePage, ".")
	defer p.mutex.Unlock()
	p.mutex.Lock()
	p.pages[parts[0]] = builder.String()
}
