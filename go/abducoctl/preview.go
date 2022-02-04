package abducoctl

import (
	"fmt"
	"log"
	"strings"

	"github.com/k0kubun/pp"
	fuzzyfinder "github.com/ktr0731/go-fuzzyfinder"
)

type Preview struct {
	Name      string
	AlbumName string
	Artist    string
}

func Previews() []Preview {
	p := []Preview{}
	l, _ := List()
	for _, i := range l {
		p = append(p, Preview{
			Name:      i.Session,
			Artist:    i.Session,
			AlbumName: i.Session,
		})
	}
	return p
}

func DoPreview() {
	p, _ := List()
	idx, err := fuzzyfinder.FindMulti(
		p,
		func(i int) string {
			c := `$`
			if p[i].Username == `root` {
				c = `#`
			}
			s := fmt.Sprintf(`%s@%s <%d> %s %s`,
				p[i].Username,
				p[i].Started,
        p[i].PID,
				c,
				strings.Join(p[i].Executables[1:len(p[i].Executables)], ` `),
			)
			return s
		},
		fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}
			return fmt.Sprintf("Session: %s (%s) | PID %d | User %s",
				p[i].Session,
				p[i].PID,
				p[i].Username)
		}))
	if err != nil {
		log.Fatal(err)
	}
	pp.Printf("selected: %v\n", p[idx[0]])
	//for _, id := range idx {
	//		fmt.Println(p[id])
	//	}
	/*
	   if Exists(answers.Session) {
	       Connect(answers.Session)
	   }
	*/
}
