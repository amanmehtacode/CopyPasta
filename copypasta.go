package main

import (
	"bufio"
	"context"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/fatih/color"
)

type ProgressIndicator struct {
	total     int
	current   int
	lastPrint time.Time
}

func NewProgressIndicator() *ProgressIndicator {
	return &ProgressIndicator{
		total:     100,
		current:   0,
		lastPrint: time.Now(),
	}
}

func (p *ProgressIndicator) Increment() {
	p.current++
	if time.Since(p.lastPrint) > 500*time.Millisecond {
		p.Print()
		p.lastPrint = time.Now()
	}
}

func (p *ProgressIndicator) Print() {
	percent := float64(p.current) / float64(p.total) * 100
	fmt.Printf("\rProgress: [%-20s] %.2f%%", strings.Repeat("=", int(percent/5)), percent)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: copypasta <URL>")
		os.Exit(1)
	}

	url, err := url.Parse(os.Args[1])
	if err != nil {
		fmt.Println("Invalid URL:", err)
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	cmd := exec.CommandContext(ctx, "wget", "--mirror", "--convert-links", "--adjust-extension", "--page-requisites", url.String())

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Error creating StdoutPipe:", err)
		os.Exit(1)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Println("Error creating StderrPipe:", err)
		os.Exit(1)
	}

	err = cmd.Start()
	if err != nil {
		fmt.Println("Error starting command:", err)
		os.Exit(1)
	}

	var wg sync.WaitGroup
	wg.Add(2)

	progress := NewProgressIndicator()

	go func() {
		defer wg.Done()
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			line := scanner.Text()
			processOutput(line, color.FgGreen)
			if strings.Contains(line, "Downloaded:") {
				progress.Increment()
			}
		}
	}()

	go func() {
		defer wg.Done()
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			processOutput(scanner.Text(), color.FgYellow)
		}
	}()

	wg.Wait()
	err = cmd.Wait()
	if err != nil {
		color.Red("\nCommand finished with error: %v", err)
	} else {
		color.Green("\nDownload completed successfully!")
	}
}

func processOutput(line string, textColor color.Attribute) {
	if strings.Contains(line, "Downloaded:") {
		parts := strings.Split(line, " ")
		if len(parts) >= 2 {
			color.Set(color.FgCyan)
			fmt.Print("Downloaded: ")
			color.Set(textColor)
			fmt.Println(parts[1])
			color.Unset()
			return
		}
	}

	if strings.Contains(line, "Saving to:") {
		parts := strings.SplitN(line, "Saving to: ", 2)
		if len(parts) == 2 {
			color.Set(color.FgMagenta)
			fmt.Print("Saving to: ")
			color.Set(textColor)
			fmt.Println(parts[1])
			color.Unset()
			return
		}
	}

	color.Set(textColor)
	fmt.Println(line)
	color.Unset()
}
