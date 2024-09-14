# Copypasta: Advanced Web Content Retrieval Tool

## Project Overview

Copypasta is an innovative command-line interface (CLI) tool developed as a software engineering internship project. It enhances the functionality of the `wget` command, providing a user-friendly interface for downloading web content with improved visual feedback and progress tracking. This project demonstrates proficiency in Go programming, concurrent processing, and creating efficient, user-centric software solutions.

## Key Features and Technical Highlights

### 1. Concurrent Processing with Goroutines
- **Implementation**: Utilized Go's goroutines for parallel processing of stdout and stderr streams.
- **Benefit**: Significantly improves performance by enabling simultaneous download and output handling.

### 2. Real-time Progress Tracking
- **Feature**: Custom-built `ProgressIndicator` struct for live download progress visualization.
- **Technical Detail**: Implements a thread-safe incrementing mechanism with a controlled update frequency to balance responsiveness and system resource usage.

### 3. Context-Aware Execution
- **Implementation**: Leveraged Go's `context` package to implement timeout functionality.
- **Advantage**: Enhances robustness by preventing indefinite hanging on slow or failed downloads.

### 4. Modular and Extensible Design
- **Structure**: Organized codebase with clear separation of concerns (main execution, progress tracking, output processing).
- **Benefit**: Facilitates easy maintenance and future feature additions.

### 5. Enhanced User Interface
- **Feature**: Colorized console output for improved readability and user experience.
- **Technical Aspect**: Integrated the `fatih/color` library, demonstrating ability to effectively incorporate external packages.

## Technical Implementation Details

### Installation and Dependencies

```bash
go get github.com/fatih/color
go build -o copypasta main.go
```

### Usage

```bash
./copypasta <URL>
```

### Code Structure

1. **Main Function**: Entry point of the program. Handles URL parsing, command execution, and error reporting.
2. **ProgressIndicator Struct**: 
   ```go
   type ProgressIndicator struct {
       total     int
       current   int
       lastPrint time.Time
   }
   ```
   Methods include `NewProgressIndicator()`, `Increment()`, and `Print()`.
3. **processOutput Function**: Handles formatting and coloring of the `wget` command output.

### Key Components

1. **URL Parsing and Validation**:
   ```go
   url, err := url.Parse(os.Args[1])
   ```

2. **Context with Timeout**:
   ```go
   ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
   ```

3. **Command Execution**:
   ```go
   cmd := exec.CommandContext(ctx, "wget", "--mirror", "--convert-links", "--adjust-extension", "--page-requisites", url.String())
   ```

4. **Goroutines for Output Handling**:
   ```go
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
   ```

## Problem-Solving Approach

1. **Requirement Analysis**: Identified limitations in existing web content retrieval tools and defined clear objectives for improvement.
2. **Design Phase**: Architected a solution that leverages Go's strengths in concurrent programming and system-level operations.
3. **Implementation**: Developed the core functionality with a focus on performance and user experience.
4. **Testing and Refinement**: Iterated on the initial implementation, fine-tuning the progress indicator and error handling mechanisms.

## Technical Skills Demonstrated

- **Go Programming**: Proficient use of Go's concurrency models, standard library, and third-party packages.
- **CLI Application Development**: Created an intuitive command-line tool with a focus on user experience.
- **System Integration**: Successfully interfaced with and enhanced existing system utilities (`wget`).
- **Error Handling and Robustness**: Implemented comprehensive error checking and timeout mechanisms.
- **Performance Optimization**: Balanced resource usage and responsiveness in progress tracking and output handling.

## Challenges and Solutions

1. **Challenge**: Accurate progress tracking for varied website structures.
   **Solution**: Implemented a flexible progress indicator based on "Downloaded:" messages, with plans for future enhancements.

2. **Challenge**: Balancing real-time updates with system resource usage.
   **Solution**: Implemented a time-based update mechanism in the `ProgressIndicator` to limit update frequency.

## Future Enhancements

- Implement multi-URL support for batch downloading.
- Develop a configuration file system for customizable `wget` options.
- Create a RESTful API wrapper to enable integration with web applications.
- Improve progress tracking accuracy for various website structures.

## Conclusion

The Copypasta project showcases the ability to develop robust, efficient software solutions. It demonstrates proficiency in Go programming, understanding of concurrent programming concepts, and commitment to creating user-centric applications. This project reflects readiness to contribute effectively to challenging software engineering tasks in a professional environment.