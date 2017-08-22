package codegen

import (
	"path/filepath"
	"text/template"

	"goa.design/goa/codegen"
	httpdesign "goa.design/goa/http/design"
)

// ExampleCLI returns an example client tool main implementation.
func ExampleCLI(root *httpdesign.RootExpr) codegen.File {
	path := filepath.Join("cmd", codegen.SnakeCase(root.Design.API.Name)+"cli", "main.go")
	sections := func(genPkg string) []*codegen.Section {
		specs := []*codegen.ImportSpec{
			{Path: "context"},
			{Path: "encoding/json"},
			{Path: "flag"},
			{Path: "fmt"},
			{Path: "net/http"},
			{Path: "net/url"},
			{Path: "os"},
			{Path: "strings"},
			{Path: "time"},
			{Path: "goa.design/goa/http", Name: "goahttp"},
			{Path: genPkg + "/http/cli"},
		}
		s := []*codegen.Section{
			codegen.Header("", "main", specs),
			&codegen.Section{Template: mainCLITmpl, Data: root},
		}

		return s
	}

	return codegen.NewSource(path, sections)
}

var mainCLITmpl = template.Must(template.New("cli-main").Parse(mainCLIT))

// input: map[string]interface{}{"Services":[]ServiceData, "APIPkg": string}
const mainCLIT = `func main() {
	var (
		addr    = flag.String("url", "http://localhost:8080", "` + "`" + `URL` + "`" + ` to service host")
		verbose = flag.Bool("verbose", false, "Print request and response details")
		timeout = flag.Int("timeout", 30, "Maximum number of ` + "`" + `seconds` + "`" + ` to wait for response")
	)
	flag.Usage = usage
	flag.Parse()

	var (
		scheme string
		host   string
	)
	{
		u, err := url.Parse(*addr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "invalid URL %#v: %s", *addr, err)
			os.Exit(1)
		}
		scheme = u.Scheme
		host = u.Host
		if scheme == "" {
			scheme = "http"
		}
	}

	var (
		doer goahttp.Doer
	)
	{
		doer = &http.Client{Timeout: time.Duration(*timeout) * time.Second}
		if *verbose {
			doer = goahttp.NewDebugDoer(doer)
		}
	}

	endpoint, payload, err := cli.ParseEndpoint(
		scheme,
		host,
		doer,
		goahttp.RequestEncoder,
		goahttp.ResponseDecoder,
		*verbose,
	)
	if err != nil {
		if err == flag.ErrHelp {
			os.Exit(0)
		}
		fmt.Fprintln(os.Stderr, err.Error())
		fmt.Fprintln(os.Stderr, "run '"+os.Args[0]+" --help' for detailed usage.")
		os.Exit(1)
	}

	data, err := endpoint(context.Background(), payload)

	if *verbose {
		doer.(goahttp.DebugDoer).Fprint(os.Stderr)
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	if data != nil && !*verbose {
		m, _ := json.MarshalIndent(data, "", "    ")
		fmt.Println(string(m))
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, ` + "`" + `%s is a command line client for the {{ .Design.API.Name }} API.

Usage:
    %s [-url URL][-timeout SECONDS][-verbose] SERVICE ENDPOINT [flags]

    -url URL: specify service URL (http://localhost:8080)
    -timeout: Maximum number of seconds to wait for response (30)
    -debug:   print debug details (false)

Commands:
%s
Additional help:
    %s SERVICE [ENDPOINT] --help

Example:
%s
` + "`" + `, os.Args[0], os.Args[0], indent(cli.UsageCommands()), os.Args[0], indent(cli.UsageExamples()))
}

func indent(s string) string {
	if s == "" {
		return ""
	}
	return "    " + strings.Replace(s, "\n", "\n    ", -1)
}
`
