package main

import (
	"embed"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

//go:embed dict.tsv
var dictFile embed.FS

func main() {
	data, err := dictFile.ReadFile("dict.tsv")
	if err != nil {
		fmt.Printf("Could not read dict.tsv: %v\n", err)
		return
	}

	rawLines := strings.Split(strings.TrimSpace(string(data)), "\n")
	var lines []string
	for _, line := range rawLines {
		if strings.TrimSpace(line) != "" {
			lines = append(lines, line)
		}
	}

	if len(lines) == 0 {
		fmt.Println("dict.tsv is empty.")
		return
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	found := false
	for i := 0; i < 10; i++ {
		pick := lines[r.Intn(len(lines))]
		parts := strings.Split(pick, "\t")

		if len(parts) >= 4 {
			renderKuhgesagt(parts[0], parts[1], parts[2], parts[3])
			found = true
			break
		}
	}

}

func renderKuhgesagt(word, def, de, en string) {
	content := []string{
		fmt.Sprintf("WORT: %s", word),
		fmt.Sprintf("DEF:  %s", def),
		"",
	}
	content = append(content, wrap("DE: "+de, 55)...)
	content = append(content, wrap("EN: "+en, 55)...)

	width := 0
	for _, l := range content {
		if len(l) > width {
			width = len(l)
		}
	}
	if width < 40 {
		width = 40
	}

	fmt.Println("  " + strings.Repeat("-", width+2))
	for _, l := range content {
		fmt.Printf("< %s%s >\n", l, strings.Repeat(" ", width-len(l)))
	}
	fmt.Println("  " + strings.Repeat("-", width+2))

	fmt.Println("        \\   ^__^")
	fmt.Println("         \\  (oo)\\_______")
	fmt.Println("            (__)\\       )\\/\\")
	fmt.Println("                ||----w |")
	fmt.Println("                ||     ||")
}

func wrap(text string, limit int) []string {
	var result []string
	words := strings.Fields(text)
	if len(words) == 0 {
		return result
	}

	curr := words[0]
	for _, w := range words[1:] {
		if len(curr)+1+len(w) > limit {
			result = append(result, curr)
			curr = w
		} else {
			curr += " " + w
		}
	}
	result = append(result, curr)
	return result
}
