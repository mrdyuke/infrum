package generator

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/mrdyuke/infrum/domain"
)

func TestNewGenerator(t *testing.T) {
	g, err := NewGenerator()
	if err != nil {
		t.Fatalf("NewGenerator() failed: %v", err)
	}
	if g == nil {
		t.Fatal("NewGenerator() returned nil")
	}
	if g.AppName != "MyApp" {
		t.Errorf("expected default AppName %q, got %q", "MyApp", g.AppName)
	}
	if g.TargetDir == "" {
		t.Error("expected non-empty TargetDir")
	}
	if len(g.LibList) != 0 {
		t.Errorf("expected empty LibList, got %d items", len(g.LibList))
	}
}

func TestGenerate_EmptyAppName(t *testing.T) {
	g := &Generator{
		AppName:   "",
		TargetDir: t.TempDir(),
		LibList: domain.LibraryList{
			{LibName: "test", LibPath: domain.LibraryPaths{}},
		},
	}

	err := g.Generate()
	if err == nil {
		t.Fatal("expected error for empty app name, got nil")
	}
	if !strings.Contains(err.Error(), "empty app name") {
		t.Errorf("expected 'empty app name' error, got: %v", err)
	}
}

func TestGenerate_EmptyLibList(t *testing.T) {
	g := &Generator{
		AppName:   "testapp",
		TargetDir: t.TempDir(),
		LibList:   domain.LibraryList{},
	}

	err := g.Generate()
	if err == nil {
		t.Fatal("expected error for empty LibList, got nil")
	}
	if !strings.Contains(err.Error(), "no libraries selected") {
		t.Errorf("expected 'no libraries selected' error, got: %v", err)
	}
}

func TestGenerate_WhitespaceAppName(t *testing.T) {
	g := &Generator{
		AppName:   "   ",
		TargetDir: t.TempDir(),
		LibList: domain.LibraryList{
			{LibName: "test", LibPath: domain.LibraryPaths{}},
		},
	}

	err := g.Generate()
	if err == nil {
		t.Fatal("expected error for whitespace-only app name, got nil")
	}
	if !strings.Contains(err.Error(), "empty app name") {
		t.Errorf("expected 'empty app name' error, got: %v", err)
	}
}

func TestGenerate_Success(t *testing.T) {
	tmpDir := t.TempDir()

	libs := domain.LibraryList{
		{
			LibName: "stdlib",
			LibPath: domain.LibraryPaths{
				"/": domain.LibraryFile{
					"main.go": `package main

import "fmt"

func main() {
    fmt.Println("{{.AppName}}")
}
`,
				},
			},
		},
	}

	g := &Generator{
		AppName:   "testy",
		TargetDir: tmpDir,
		LibList:   libs,
	}

	err := g.Generate()
	if err != nil {
		t.Fatalf("Generate() failed: %v", err)
	}

	// Verify the project directory was created
	projectDir := filepath.Join(tmpDir, "testy")
	if _, err := os.Stat(projectDir); os.IsNotExist(err) {
		t.Fatal("project directory was not created")
	}

	// Verify main.go was created with rendered template
	mainGo := filepath.Join(projectDir, "main.go")
	data, err := os.ReadFile(mainGo)
	if err != nil {
		t.Fatalf("failed to read main.go: %v", err)
	}
	if !strings.Contains(string(data), `fmt.Println("testy")`) {
		t.Errorf("main.go does not contain rendered app name. Got:\n%s", string(data))
	}

	// Verify go.mod was created
	goMod := filepath.Join(projectDir, "go.mod")
	if _, err := os.Stat(goMod); os.IsNotExist(err) {
		t.Fatal("go.mod was not created")
	}
	goModData, _ := os.ReadFile(goMod)
	if !strings.Contains(string(goModData), "module testy") {
		t.Errorf("go.mod does not contain correct module name. Got:\n%s", string(goModData))
	}

	// Verify the generated project compiles
	cmd := exec.Command("go", "build", "-o", "/dev/null", "./...")
	cmd.Dir = projectDir
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("generated project failed to build:\n%s\nError: %v", string(out), err)
	}
}

func TestGenerate_TemplateRendering(t *testing.T) {
	tmpDir := t.TempDir()

	libs := domain.LibraryList{
		{
			LibName: "config",
			LibPath: domain.LibraryPaths{
				"/": domain.LibraryFile{
					"app.conf": `name={{.AppName}}
`,
				},
				"/internal/": domain.LibraryFile{
					"handler.go": `package internal

const App = "{{.AppName}}"
`,
				},
			},
		},
	}

	g := &Generator{
		AppName:   "myapp",
		TargetDir: tmpDir,
		LibList:   libs,
	}

	if err := g.Generate(); err != nil {
		t.Fatalf("Generate() failed: %v", err)
	}

	projectDir := filepath.Join(tmpDir, "myapp")

	// Check root file
	data, err := os.ReadFile(filepath.Join(projectDir, "app.conf"))
	if err != nil {
		t.Fatalf("failed to read app.conf: %v", err)
	}
	if string(data) != "name=myapp\n" {
		t.Errorf("app.conf content wrong. Got: %q", string(data))
	}

	// Check file in subdirectory
	data, err = os.ReadFile(filepath.Join(projectDir, "internal", "handler.go"))
	if err != nil {
		t.Fatalf("failed to read internal/handler.go: %v", err)
	}
	if !strings.Contains(string(data), `const App = "myapp"`) {
		t.Errorf("handler.go not rendered correctly. Got: %s", string(data))
	}
}

func TestGenerate_MultipleLibraries(t *testing.T) {
	tmpDir := t.TempDir()

	libs := domain.LibraryList{
		{
			LibName: "core",
			LibPath: domain.LibraryPaths{
				"/": domain.LibraryFile{
					"core.txt": "core-{{.AppName}}",
				},
			},
		},
		{
			LibName: "extra",
			LibPath: domain.LibraryPaths{
				"/extras/": domain.LibraryFile{
					"extra.txt": "extra-{{.AppName}}",
				},
			},
		},
	}

	g := &Generator{
		AppName:   "multi",
		TargetDir: tmpDir,
		LibList:   libs,
	}

	if err := g.Generate(); err != nil {
		t.Fatalf("Generate() failed: %v", err)
	}

	projectDir := filepath.Join(tmpDir, "multi")

	// Check root file from first lib
	data, err := os.ReadFile(filepath.Join(projectDir, "core.txt"))
	if err != nil {
		t.Fatalf("failed to read core.txt: %v", err)
	}
	if string(data) != "core-multi" {
		t.Errorf("core.txt content wrong. Got: %q", string(data))
	}

	// Check file in subdirectory from second lib
	data, err = os.ReadFile(filepath.Join(projectDir, "extras", "extra.txt"))
	if err != nil {
		t.Fatalf("failed to read extras/extra.txt: %v", err)
	}
	if string(data) != "extra-multi" {
		t.Errorf("extra.txt content wrong. Got: %q", string(data))
	}
}

func TestGenerate_InvalidTemplate(t *testing.T) {
	tmpDir := t.TempDir()

	libs := domain.LibraryList{
		{
			LibName: "bad",
			LibPath: domain.LibraryPaths{
				"/": domain.LibraryFile{
					"bad.tmpl": "hello {{.NonExistentField}",
				},
			},
		},
	}

	g := &Generator{
		AppName:   "test",
		TargetDir: tmpDir,
		LibList:   libs,
	}

	err := g.Generate()
	if err == nil {
		t.Fatal("expected error for invalid template, got nil")
	}
	if !strings.Contains(err.Error(), "failed to parse template") {
		t.Errorf("expected template parse error, got: %v", err)
	}
}
