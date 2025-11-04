package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/wasuppu/smu"
)

type NavItem struct {
	Name string
	URL  string
}

var (
	navItems []NavItem
)

const (
	templateFile = "templates/template.html"
	outputDir    = "./static"
	homeFile     = "./content/home.md"
	projectsFile = "./content/projects.md"
	contactFile  = "./content/contact.md"
	postsDir     = "./content/posts"
	imgsDir      = "./content/imgs"
	assetDir     = "./templates/asset"
	cssPath      = "css/style.css"
)

func main() {
	serverMode := flag.Bool("s", false, "start HTTP server")
	port := flag.Int("p", 3000, "server port")
	flag.Parse()

	tpl, err := template.ParseFiles(templateFile)

	if err != nil {
		log.Fatal(err)
	}

	createDirIfNotExist(outputDir)

	copyDir(imgsDir, filepath.Join(outputDir, filepath.Base(imgsDir)))
	copyDirContents(assetDir, outputDir)

	log.Println("Start Rendering")
	addNavItems(homeFile, projectsFile, postsDir, contactFile)
	must(generatePageFromFile(tpl, homeFile))
	must(generatePageFromFile(tpl, projectsFile))
	must(generatePageFromFile(tpl, contactFile))
	must(generatePosts(tpl, postsDir))
	must(generatePage(tpl, "index.html", map[string]any{
		"CSS":      cssPath,
		"Title":    "space",
		"Body":     "",
		"NavItems": navItems,
	}))
	must(generatePage(tpl, "404.html", map[string]any{
		"CSS":      cssPath,
		"Title":    "Not Found",
		"Body":     "<h1>404 - Page Not Found</h1><p>The requested page could not be found.</p>",
		"NavItems": navItems,
	}))
	log.Println("End Rendering")

	if *serverMode {
		startServer(*port)
	}
}

func startServer(port int) {
	staticDir, err := filepath.Abs(outputDir)
	if err != nil {
		log.Fatalf("Failed to get the static directory path: %v", err)
	}

	fileServer := http.FileServer(http.Dir(staticDir))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.ServeFile(w, r, filepath.Join(staticDir, "index.html"))
			return
		}

		// Check if the requested file exists
		filePath := filepath.Join(staticDir, r.URL.Path)
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			// File not found, serve 404 page
			http.ServeFile(w, r, filepath.Join(staticDir, "404.html"))
			return
		}

		fileServer.ServeHTTP(w, r)
	})

	addr := fmt.Sprintf(":%d", port)
	log.Printf("Started server on: http://localhost%s\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func generatePosts(tpl *template.Template, dirpath string) error {
	entries, err := os.ReadDir(dirpath)
	if err != nil {
		return err
	}

	var subpageNavItems []NavItem
	if len(entries) > 0 {
		createDirIfNotExist(filepath.Join(outputDir, "posts"))

		for _, item := range navItems {
			subpageNavItems = append(subpageNavItems, NavItem{
				Name: item.Name,
				URL:  "../" + strings.TrimPrefix(item.URL, "./"),
			})
		}
	}

	var indexItems []string
	for _, entry := range entries {
		mdFile := filepath.Join(postsDir, entry.Name())
		mdContent, err := os.ReadFile(mdFile)
		if err != nil {
			return err
		}

		htmlContent := string(smu.Process(mdContent))
		title := extractTitle(htmlContent)
		m := map[string]any{
			"CSS":      "../" + cssPath,
			"Title":    title,
			"Body":     htmlContent,
			"NavItems": subpageNavItems,
		}

		name := purename(mdFile)
		htmlFilename := filepath.Join("posts", name+".html")
		if err := generatePage(tpl, htmlFilename, m); err != nil {
			log.Printf("Warning: Failed to write %s: %v\n", htmlFilename, err)
			continue
		}

		log.Printf("Rendered post:%s to %s\n", name, htmlFilename)
		indexItems = append(indexItems, fmt.Sprintf(`<p><a href="%s">%s</a></p>`, filepath.Join("posts", name+".html"), title))
	}

	m := map[string]any{
		"CSS":      cssPath,
		"Title":    "posts",
		"Body":     "<h1>Posts</h1>" + strings.Join(indexItems, ""),
		"NavItems": navItems,
	}

	return generatePage(tpl, "posts.html", m)
}

func generatePageFromFile(tpl *template.Template, filename string) error {
	htmlContent, err := md2html(filename)
	if err != nil {
		return err
	}
	name := purename(filename)
	m := map[string]any{
		"CSS":      cssPath,
		"Title":    name,
		"Body":     string(htmlContent),
		"NavItems": navItems,
	}
	return generatePage(tpl, name+".html", m)
}

func generatePage(tpl *template.Template, filename string, m map[string]any) error {
	var tplbuffer bytes.Buffer
	tpl.Execute(&tplbuffer, m)
	return os.WriteFile(filepath.Join(outputDir, filename), tplbuffer.Bytes(), 0644)
}

func md2html(filename string) ([]byte, error) {
	mdContent, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	htmlContent := smu.Process(mdContent)
	return htmlContent, nil
}

func extractTitle(text string) string {
	if h1Start := strings.Index(text, "<h1>"); h1Start != -1 {
		h1End := strings.Index(text[h1Start:], "</h1>")
		if h1End != -1 {
			title := text[h1Start+4 : h1Start+h1End]
			return strings.TrimSpace(title)
		}
	}
	return ""
}

func purename(filename string) string {
	basename := filepath.Base(filename)
	ext := filepath.Ext(basename)
	return strings.TrimSuffix(basename, ext)
}

func addNavItems(ss ...string) {
	for _, s := range ss {
		name := purename(s)
		navItems = append(navItems, NavItem{
			Name: strings.ToUpper(name[:1]) + name[1:],
			URL:  fmt.Sprintf("./%s.html", name),
		})
	}
}

func createDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.Mkdir(dir, 0755); err != nil {
			log.Printf("Failed to create directory %s: %v\n", dir, err)
		}
	}
}

func copyDirContents(src, dst string) error {
	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			if err = copyDir(srcPath, dstPath); err != nil {
				return err
			}
		} else {
			if err = copyFile(srcPath, dstPath); err != nil {
				return err
			}
		}
	}

	return nil
}

func copyDir(src, dst string) error {
	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}

	if err = os.MkdirAll(dst, srcInfo.Mode()); err != nil {
		return err
	}

	return copyDirContents(src, dst)
}

func copyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	srcInfo, err := srcFile.Stat()
	if err != nil {
		return err
	}

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	if _, err = io.Copy(dstFile, srcFile); err != nil {
		return err
	}

	return os.Chmod(dst, srcInfo.Mode())
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
