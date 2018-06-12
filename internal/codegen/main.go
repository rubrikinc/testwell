package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path"
	"sort"
	"strings"
	"text/template"
)

var (
	flagPkg  = flag.String("pkg", "assert", "The output package.")
	flagTmpl = flag.String(
		"tmpl",
		"../internal/assert.tmpl",
		"The template source.",
	)
)

// We split the string here because we do not want phabricator to thing this
// very file is generated.
var header = "// " + "@generated\n\n"

// For the convenience in `assert.tmpl`. We offer functions returning sets of
// types. Using the Go template terminology, they can be pipelined together to
// form bigger sets (unions). For example, {{ PrimitiveTypes | ErrorType }}.
// For this to work, we need those functions to take 0 or 1 arguments. The
// closest we can achieve in Go is to take 0 or N arguments (variadics).
// Because we do not want duplicates, and we want a stable output across runs,
// we pass our set trough the function `uniqAndOrdered` before returning it.

func numericalTypes(accs ...types) types {
	accs = append(accs, strings2types([]string{
		"int", "uint",
		"int8", "uint8",
		"int16", "uint16",
		"int32", "uint32",
		"int64", "uint64",
		"float32", "float64",
		"complex64", "complex128",
		"byte",
	}))
	return uniqAndOrdered(accs...)
}

func primitiveTypes(accs ...types) types {
	accs = append(accs, numericalTypes(accs...))
	accs = append(accs, strings2types([]string{
		"uintptr", "rune", "bool", "string",
	}))
	return uniqAndOrdered(accs...)
}

type tmplData struct {
	Package string
}

type typeInfo struct {
	N string
	T string
}

func (t *typeInfo) String() string {
	return t.T
}

type types []typeInfo

func (ts types) Len() int { return len(ts) }
func (ts types) Swap(i, j int) {
	ts[i], ts[j] = ts[j], ts[i]
}
func (ts types) Less(i, j int) bool {
	return strings.Compare(ts[i].N, ts[j].N) < 0
}

func uniqAndOrdered(accs ...types) types {
	s := map[string]typeInfo{}
	for _, acc := range accs {
		for _, e := range acc {
			s[e.T] = e
		}
	}
	r := make(types, 0, len(s))
	for _, t := range s {
		r = append(r, t)
	}
	sort.Sort(r)
	return r
}

func strings2types(ss []string) types {
	r := make(types, 0, len(ss))
	for _, s := range ss {
		r = append(r, typeInfo{T: s, N: strings.Title(s)})
	}
	return r
}

func main() {
	flag.Parse()
	tmpl := template.New("")
	tmpl = tmpl.Funcs(template.FuncMap{
		"NumericalTypes": numericalTypes,
		"PrimitiveTypes": primitiveTypes,
		"ErrorType": func(accs ...types) types {
			accs = append(accs, strings2types([]string{"error"}))
			return uniqAndOrdered(accs...)
		},
		"GenericType": func(accs ...types) types {
			accs = append(accs, types{typeInfo{N: "", T: "interface{}"}})
			return uniqAndOrdered(accs...)
		},
		"Capitalize": strings.Title,
	})

	outputFilename := fmt.Sprintf("%s.go", *flagPkg)
	fmt.Printf("Generating %s from %s\n", outputFilename, *flagTmpl)
	tmpl = template.Must(tmpl.ParseFiles(*flagTmpl))

	outputFile, err := os.OpenFile(
		outputFilename,
		os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
		0644,
	)
	if err != nil {
		panic(err)
	}

	success := false
	defer func() {
		if !success {
			os.Remove(outputFile.Name())
		}
	}()

	gofmtcmd := exec.Command("goimports")
	gofmtcmd.Stdout = outputFile
	gofmtcmd.Stderr = os.Stderr

	output, err := gofmtcmd.StdinPipe()
	if err != nil {
		panic(err)
	}

	if err = gofmtcmd.Start(); err != nil {
		panic(err)
	}

	fmt.Fprint(output, header)

	data := tmplData{
		Package: *flagPkg,
	}

	if err = tmpl.ExecuteTemplate(output, path.Base(*flagTmpl), data); err != nil {
		panic(err)
	}

	if err = output.Close(); err != nil {
		panic(err)
	}

	if err = gofmtcmd.Wait(); err != nil {
		panic(err)
	}

	success = true
}
