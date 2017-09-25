package main

const (
	DefaultPort = 8080
	DefaultRoot = "/"
)

type Application struct {
	Title          string
	Port           int
	Root           string
	MainController home.HomeController
}

func (a *Application) Init(title string) {
	a.Title = title
	a.Port = DefaultPort
	a.Root = DefaultRoot
	a.MainController.Init("Home", a.Root)
}

func (hc *HomeController) printOnPage(body string) {
	http.HandleFunc(hc.Base.SubRoot, (func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, body)
	}))
}

func (hc *HomeController) Index(appTitle string) {
	pageTitle := hc.Base.Name + " - " + appTitle

	hc.printOnPage("<html>" +
		"<head>" +
		"<title>" + pageTitle + "</title>" +
		"</head>" +
		"<body>" +
		"<h1>Hello!</h1>" +
		"</body>" +
		"</html>")
}

func main() {
	var app Application
	app.Init("Project Seabass")
	app.MainController.Index(app.Title)
	psb.RunServer(app.Port)
}
