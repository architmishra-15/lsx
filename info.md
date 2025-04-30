# lsx - Modern ls Command with File Icons

A modern remake of the Unix `ls` command written in Go, featuring colorized outputs and file type icons.

## Project Structure

The project is organized into multiple files, each with a specific responsibility:

```
lsx/
├── main.go           # Main program logic and entry point
├── flags.go          # Command-line flag handling
├── print_format.go   # Multi-column display formatting
├── list_view.go      # Unix-like long format display
├── file_utils.go     # Utility functions for file info formatting
└── show_help.go      # Help information display
```

## Functionality Overview

The `lsx` command supports the following flags:

- `-l`: Long format showing detailed file information
- `-a`: Show all files including hidden files (dotfiles)
- `-d` or `--directory`: List directories themselves, not their contents
- `-h` or `--human-readable`: Show file sizes in human-readable format (KB, MB, GB)
- `-v` or `--version`: Show version information
- `--help`: Display help information
- Combined flags like `-la` are also supported

## Detailed File Documentation

### main.go

The main entry point of the application that orchestrates the execution flow.

**Functions**:
- `main()`: Program entry point
- `processPath(pattern string, flags Flags)`: Handles glob pattern processing
- `printFilesWithExtension(dirPath string, ext string, flags Flags)`: Lists files matching an extension
- `printFile(filePath string, flags Flags)`: Displays information for a single file
- `printDirectoryContents(dirPath string, flags Flags)`: Lists contents of a directory

**External Calls**:
- Calls `handle_flag()` from flags.go
- Calls `PrintFilesInColumns()` or `PrintListView()` based on flags

### flags.go

Handles command-line flag parsing and processing.

**Types**:
- `Flags`: Structure holding flag states
  ```go
  type Flags struct {
    LongFormat    bool // -l flag
    AllFiles      bool // -a flag
    DirectoryOnly bool // -d or --directory flag
    HumanReadable bool // -h or --human-readable flag
    Help          bool // --help flag
    Version       bool // -v or --version flag
  }
  ```

**Functions**:
- `parseFlags(args []string) (Flags, []string)`: Parses command line arguments into a Flags structure
- `handle_flag(args []string) (Flags, []string)`: Processes flags and handles help/version flags
- `show_version()`: Displays version information 

**External Calls**:
- Calls `showHelp()` from show_help.go

### print_format.go

Handles the standard multi-column formatted display of files.

**Functions**:
- `isConfigFile(name string) string`: Identifies framework configuration files
- `PrintFilesInColumns(files []os.FileInfo, numColumns int, flags ...Flags)`: Displays files in columns with icons

**External Calls**:
- Calls `PrintListView()` from list_view.go when long format flag is set
- Calls `getFileTypeColorAndIcon()` function (defined elsewhere)
- Uses `Color` and `Icons` maps (defined elsewhere)

### list_view.go

Implements Unix-like long format listings (like `ls -l`).

**Functions**:
- `PrintListView(files []os.FileInfo, flags Flags)`: Main function for ls -l style display
- `calculateTotalBlocks(files []os.FileInfo) int64`: Calculates disk usage in 1K blocks
- `findMaxSizeLength(files []os.FileInfo, humanReadable bool) int`: Helps with output alignment
- `getOwnerAndGroup(file os.FileInfo) (string, string)`: Gets owner/group info
- `getLinkCount(file os.FileInfo) int`: Gets number of hard links
- `printFileListView(file os.FileInfo, maxSizeLen int, humanReadable bool)`: Formats and prints a single file entry

**External Calls**:
- Calls `FormatFileSize()` from file_utils.go
- Calls `FormatPermissions()` from file_utils.go
- Calls `FormatModTime()` from file_utils.go
- Calls `getFileTypeColorAndIcon()` function
- Uses `Color` map

### file_utils.go

Provides utility functions for file information formatting.

**Functions**:
- `FormatFileSize(size int64, humanReadable bool) string`: Formats file size (optionally in human-readable form)
- `FormatPermissions(info os.FileInfo) string`: Formats file permissions as rwx-style string
- `formatPermissionBits(mode os.FileMode, r, w, x os.FileMode) string`: Helper for permission formatting
- `FormatModTime(modTime time.Time) string`: Formats modification time

**Constants**:
- Size units: KB, MB, GB, TB

### show_help.go

Contains the help display functionality.

**Functions**:
- `showHelp()`: Displays usage information and available options

## Color and Icon System

The application uses ANSI color codes to provide colorized output based on file types:
- Directories: Blue
- Executables: Green
- Images: Magenta
- Videos: Cyan
- Configuration files: Yellow
- Regular files: White

Each file type also has an associated icon, enhancing visual recognition.

## Usage Examples

```bash
# List files in the current directory
lsx

# List all files including hidden ones
lsx -a

# List files in long format
lsx -l

# Long format with human-readable sizes
lsx -lh

# List all files in long format
lsx -la

# List the directory itself, not its contents
lsx -d .
```

## Building and Running

To build the project:

```bash
go build -o lsx.exe
```

To run without building:

```bash
go run .
```

## Detailed Per-File Documentation

### main.go

The main entry point and controller for the lsx application.

#### Functions

- `main()`
  - Description: Entry point of the program that parses arguments and handles execution flow
  - Inputs: None (uses os.Args internally)
  - Outputs: None
  - Calls: handle_flag(), printDirectoryContents(), processPath()

- `processPath(pattern string, flags Flags)`
  - Description: Processes patterns like *.extension or specific file paths
  - Inputs: pattern string, flags Flags
  - Outputs: None (prints to console)
  - Calls: printFilesWithExtension(), printFile()
  
- `printFilesWithExtension(dirPath string, ext string, flags Flags)`
  - Description: Lists files that match a specific extension
  - Inputs: dirPath string, ext string, flags Flags
  - Outputs: None (prints to console)
  - Calls: PrintFilesInColumns()

- `printFile(filePath string, flags Flags)`
  - Description: Shows information for a single file
  - Inputs: filePath string, flags Flags
  - Outputs: None (prints to console)
  - Calls: PrintFilesInColumns()

- `printDirectoryContents(dirPath string, flags Flags)`
  - Description: Lists contents of a directory based on flag settings
  - Inputs: dirPath string, flags Flags
  - Outputs: None (prints to console)
  - Calls: PrintFilesInColumns()

### flags.go

Handles command-line flag parsing and processing.

#### Types

- `Flags`
  - Description: Structure that stores the state of each command-line flag
  - Fields:
    - LongFormat (bool): -l flag
    - AllFiles (bool): -a flag
    - DirectoryOnly (bool): -d/--directory flag
    - HumanReadable (bool): -h/--human-readable flag
    - Help (bool): --help flag
    - Version (bool): -v/--version flag

#### Functions

- `parseFlags(args []string) (Flags, []string)`
  - Description: Parses command line arguments into Flags structure and remaining args
  - Inputs: args []string (command-line arguments)
  - Outputs: (Flags, []string) - parsed flags and remaining arguments
  - Special Feature: Handles combined flags like -la

- `handle_flag(args []string) (Flags, []string)`
  - Description: Processes flags and executes help/version commands if needed
  - Inputs: args []string (command-line arguments)
  - Outputs: (Flags, []string) - processed flags and remaining arguments
  - Calls: showHelp(), show_version()

- `show_version()`
  - Description: Displays version and copyright information
  - Inputs: None
  - Outputs: None (prints to console)

### print_format.go

Handles the formatted display of files in columns with icons and colors.

#### Functions

- `isConfigFile(name string) string`
  - Description: Identifies framework configuration files to use specific icons
  - Inputs: name string (filename)
  - Outputs: string (framework identifier or empty string)

- `PrintFilesInColumns(files []os.FileInfo, numColumns int, flags ...Flags)`
  - Description: Main display function that shows files in customizable columns with icons
  - Inputs: files []os.FileInfo, numColumns int, flags ...Flags
  - Outputs: None (prints to console)
  - Calls: PrintListView() if long format flag is set
  - Special Feature: Handles file name overflow for long names

### list_view.go

Implements Unix-style long format listing similar to ls -l command.

#### Functions

- `PrintListView(files []os.FileInfo, flags Flags)`
  - Description: Main function for Unix-like long format display
  - Inputs: files []os.FileInfo, flags Flags
  - Outputs: None (prints to console)
  - Calls: calculateTotalBlocks(), findMaxSizeLength(), printFileListView()

- `calculateTotalBlocks(files []os.FileInfo) int64`
  - Description: Calculates total disk usage in 1K blocks (Unix compatibility)
  - Inputs: files []os.FileInfo
  - Outputs: int64 (total blocks)

- `findMaxSizeLength(files []os.FileInfo, humanReadable bool) int`
  - Description: Finds max size string length for output alignment
  - Inputs: files []os.FileInfo, humanReadable bool
  - Outputs: int (max size string length)
  - Calls: FormatFileSize()

- `getOwnerAndGroup(file os.FileInfo) (string, string)`
  - Description: Gets owner and group names (simplified for Windows)
  - Inputs: file os.FileInfo
  - Outputs: string, string (owner and group names)

- `getLinkCount(file os.FileInfo) int`
  - Description: Gets number of hard links (simulated on Windows)
  - Inputs: file os.FileInfo
  - Outputs: int (link count)

- `printFileListView(file os.FileInfo, maxSizeLen int, humanReadable bool)`
  - Description: Formats and prints a single file entry in Unix-style format
  - Inputs: file os.FileInfo, maxSizeLen int, humanReadable bool
  - Outputs: None (prints to console)
  - Calls: FormatFileSize(), FormatModTime(), getFileTypeColorAndIcon()

### file_utils.go

Provides utility functions for formatting file information.

#### Functions

- `FormatFileSize(size int64, humanReadable bool) string`
  - Description: Formats file size, optionally in human-readable form (KB, MB, GB)
  - Inputs: size int64, humanReadable bool
  - Outputs: string (formatted size)

- `FormatPermissions(info os.FileInfo) string`
  - Description: Formats file permissions in Unix-style rwx format
  - Inputs: info os.FileInfo
  - Outputs: string (formatted permissions like -rw-r--r--)
  - Calls: formatPermissionBits()

- `formatPermissionBits(mode os.FileMode, r, w, x os.FileMode) string`
  - Description: Helper function to format permission bits for user/group/other
  - Inputs: mode os.FileMode, r, w, x os.FileMode
  - Outputs: string (rwx or r-- etc.)

- `FormatModTime(modTime time.Time) string`
  - Description: Formats file modification time in Unix-style format
  - Inputs: modTime time.Time
  - Outputs: string (formatted date and time)

### show_help.go

Displays help information when requested.

#### Functions

- `showHelp()`
  - Description: Displays program usage and available options
  - Inputs: None
  - Outputs: None (prints to console)
