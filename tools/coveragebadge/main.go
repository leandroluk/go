// tools/coveragebadge/main.go
package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

func totalCoverage(profilePath string) (string, float64, error) {
	cmd := exec.Command("go", "tool", "cover", "-func="+profilePath)
	b, err := cmd.CombinedOutput()
	if err != nil {
		return "", 0, fmt.Errorf("go tool cover falhou: %v\n%s", err, string(b))
	}

	lines := strings.Split(strings.TrimSpace(string(b)), "\n")
	if len(lines) == 0 {
		return "", 0, fmt.Errorf("saida vazia do go tool cover")
	}

	last := lines[len(lines)-1]
	fields := strings.Fields(last)
	if len(fields) == 0 {
		return "", 0, fmt.Errorf("linha final inesperada: %q", last)
	}

	raw := fields[len(fields)-1]
	raw = strings.TrimSuffix(raw, "%")

	value, err := strconv.ParseFloat(raw, 64)
	if err != nil {
		return "", 0, fmt.Errorf("percentual invalido: %q", raw)
	}

	return raw, value, nil
}

func badgeColor(p float64) string {
	switch {
	case p >= 90:
		return "brightgreen"
	case p >= 80:
		return "green"
	case p >= 70:
		return "yellowgreen"
	case p >= 60:
		return "yellow"
	case p >= 50:
		return "orange"
	default:
		return "red"
	}
}

func download(url, filepath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("http %d: %s", resp.StatusCode, strings.TrimSpace(string(body)))
	}

	f, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	return err
}

func escape(s string) string {
	return strings.NewReplacer("-", "--", "_", "__", " ", "_").Replace(s)
}

func fail(err error) {
	fmt.Fprintln(os.Stderr, err.Error())
	os.Exit(1)
}

func main() {
	in := flag.String("in", "coverage.out", "")
	out := flag.String("out", "badges/coverage.svg", "")
	label := flag.String("label", "coverage", "")
	flag.Parse()

	percentText, percentValue, err := totalCoverage(*in)
	if err != nil {
		fail(err)
	}

	color := badgeColor(percentValue)

	if err := os.MkdirAll(filepath.Dir(*out), 0o755); err != nil {
		fail(err)
	}

	url := fmt.Sprintf("https://img.shields.io/badge/%s-%s%%25-%s.svg", escape(*label), escape(percentText), escape(color))
	if err := download(url, *out); err != nil {
		fail(err)
	}
}
