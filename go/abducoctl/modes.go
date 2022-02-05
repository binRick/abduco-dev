package abducoctl

type Mode struct {
	Name    string
	Handler func(string) func()
}

func init() {
	//	pp.Println(modes)
	//	os.Exit(1)

}

func Modes() []Mode {
	return modes
}

func GetMode(n string) Mode {
	for _, m := range Modes() {
		if m.Name == n {
			return m
		}
	}
	return Mode{}
}

//NewSessionNameString())

var modes = []Mode{}

/*
		Mode{"json", func() bool {
			fmt.Fprintf(os.Stdout, "%s\n", fmt.Sprintf(`%s`, JSON()))
			return true
		}},
		Mode{"path", func() bool {
			fmt.Fprintf(os.Stdout, "%s\n", fmt.Sprintf(`%s`, Path()))
			return true
		}},
		Mode{"names", func() bool {
			fmt.Fprintf(os.Stdout, "%s\n", fmt.Sprintf(`%s`, Names()))
			return true
		}},
		Mode{"pids", func() bool {
			fmt.Fprintf(os.Stdout, "%s\n", fmt.Sprintf(`%s`, PIDs()))
			return true
		}},
}*/
