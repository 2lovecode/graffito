package modg

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"os"
	"os/exec"
	"strings"
)

type graph struct {
	pool *HashMap
	root *Pkg
}

const rootName = "graffito"

func newGraph() *graph {
	o := new(graph)
	o.pool = NewHashMap()

	rootPkg := NewPkg(rootName, "")
	o.pool.Add(rootPkg)

	o.root = rootPkg
	return o
}

func (g *graph) add(before *Pkg, after *Pkg) {
	if before == nil || after == nil {
		return
	}

	if n := g.pool.Find(before.ID()); n != nil {
		before = n
	} else {
		g.pool.Add(before)
	}

	if m := g.pool.Find(after.ID()); m != nil {
		after = m
	} else {
		g.pool.Add(after)
	}
	//fmt.Println(before, after)
	before.AddDep(after.ID())
	after.AddRDep(before.ID())
}

func (g *graph) Print() {
	if g.root == nil {
		fmt.Println("空")
	}

	// 遍历
	deep := 0
	g.print(g.root.ID(), deep, false)
}

func (g *graph) print(id PkgID, deep int, final bool) {
	// 安全代码
	//if deep > 50 {
	//	return
	//}

	if o := g.pool.Find(id); o != nil {
		if !o.IsVisit {
			prefix := ""

			if deep > 0 {
				//dependsLen := len(o.depends)?
				suffix := "   ├─"
				if final {
					suffix = "   └─"
				}

				prefix = strings.Repeat("   │", deep-1) + suffix
			}

			if o.version != "" {
				fmt.Printf("  %s %s [%s]\n", prefix, o.name, o.version)
			} else {
				fmt.Printf("%s %s\n", prefix, o.name)
			}
			o.IsVisit = true
			//fmt.Printf("xxxxx %d deep %d\n", len(o.depends), deep)
			depLen := len(o.depends)
			if depLen > 0 {
				for kk, eachD := range o.depends {
					nextDeep := deep + 1
					isFinal := false
					if kk >= (depLen - 1) {
						isFinal = true
					}
					g.print(eachD, nextDeep, isFinal)
				}
			}
		}
	}

	return
}

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use: "modg",
		Run: func(cmd *cobra.Command, args []string) {
			pwd, err := os.Getwd()
			goroot := os.Getenv("GOROOT")
			isTest := false
			if len(args) > 0 && args[0] == "test" {
				isTest = true
			}
			fmt.Println(isTest)
			if err == nil {
				if goroot == "" {
					goroot = pwd
				}
				fmt.Println("****************")
				fmt.Println("Current Dir:", pwd)
				fmt.Println("GoRoot:", goroot)
				fmt.Println("****************")

				goBin := fmt.Sprintf("%s/bin/go", goroot)

				modCmd := exec.Command(goBin, "mod", "graph")
				res, e := modCmd.Output()

				if e == nil {
					var bufReader *bufio.Reader
					if isTest {
						fp, fErr := os.Open(pwd + "/tools/modg/test.txt")
						if fErr == nil {
							bufReader = bufio.NewReader(fp)
						} else {
							fmt.Println(fErr)
						}
					} else {
						reader := bytes.NewReader(res)
						//
						bufReader = bufio.NewReader(reader)
					}

					if bufReader == nil {
						fmt.Println("错误")
						os.Exit(0)
					}
					gph := newGraph()

					for {
						line, readErr := bufReader.ReadString('\n')
						line = strings.TrimSpace(line)
						if readErr != nil && readErr != io.EOF {
							// TODO notice message
							break
						}
						if readErr == io.EOF {
							break
						}
						lineSlice := strings.Split(line, " ")
						if len(lineSlice) == 2 {
							beforeSlice := strings.Split(lineSlice[0], "@")
							afterSlice := strings.Split(lineSlice[1], "@")
							if len(beforeSlice) == 1 && beforeSlice[0] == rootName {
								beforeSlice = append(beforeSlice, "")
							}
							if len(beforeSlice) == 2 && len(afterSlice) == 2 {
								gph.add(NewPkg(beforeSlice[0], beforeSlice[1]), NewPkg(afterSlice[0], afterSlice[1]))
							}
						}
					}
					gph.Print()
				}
			}
		},
	}
}

func BuildGraph() {

}
