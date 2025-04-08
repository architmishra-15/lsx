// logos.go

package main


var Icons = map[string]string {

	//  Folders
	"folder":        "\uf07b",
	"open_folder":   "\uf07c",
	"git_folder":    "\uf1d3",

	// Framework Specific
	"tailwind":                  "\ue8ba", // Tailwindcss
	"vue":                       "\ue6a0", // Vue
	"vue.config.json":           "\ued4a",
	"vite":                      "\ue8d7",
	"nextjs":                    "\ue83e",
	"svelte":                    "\ue8b7",
	"eslint":                    "\ue74b",
	"package.json":              "\uf487",

	// Programming Languages
	".asm":    "\ue6ab", // Assembly
	".go":     "\ue627", //	Go 
	"go.mod":  "\ue65e", // Go packages file
	".py":     "\ue73c", // Python
	".js":     "\ue781", // JavaScript
	".ts":     "\ue628", // TypeScript
	".jsx":    "\ued46", // React JS 
	".tsx":    "\ued46", // React TS
	".html":   "\ue736", // HTML
	".css":    "\ue749", // CSS 
	".json":   "\ue60b", // JSON 
	".md":     "\ue73e", // Markdown 
	".java":   "\ue738", // Java
	".c":      "\ue61e", // C
	".o":      "\uf013", // O files
	".cpp":    "\ue61d", // C++ 
	".cs":     "\uf81a", // C#
	".rb":     "\ue739", // Ruby
	".php":    "\ue73d", // PHP
	".rs":     "\ue7a8", // Rust 
	".swift":  "\ue755", // Swift
	".scala":  "\ue737", // Scala 
	".dart":   "\ue798", // Dart 
	".kt":     "\ue634", // Kotlin 
	".ex":     "\ue62d", // Elixir
	".exs":    "\ue62d", // Elixir
	".hs":     "\ue777", // Haskell
	".sh":     "\ue795", // Shell Script
	".pl":     "\ue769", // Perl
	".r":      "\ue76c", // R
	".coffee": "\ue7b1", // CoffeeScript
	".d":      "\ue7af", // D Language
	".m":      "\ue82a", // COBOL
	".mat":    "\ue82a", // COBOL (alternative extension)
	".ps1":    "\ue7d8", // PowerShell
	".jil":    "\ue80d", // Julia
	".lua":    "\ue826", // Lua
	".ml":     "\ue62b", // OCaml
	".f90":    "\ue7de", // Fortran
	".vim":    "\ue62b", // Vim
	".vimrc":  "\ue62b", // Vim
	".bash":   "\ue760", // Bash
	".bashrc": "\ue760", // Bash
	".zsh":    "\ue760", // Zsh
	".zshrc":  "\ue760", // Zsh
	".ads":    "\ue6b5", //Ada
	".cbl":    "\ue6b5", // Cobol
	".db":     "\uf01b", // SQL Database
	".sql":    "\uf1c0", // SQL Query Files
	".fs":     "\ue7a7", // F# 
	".fsi":    "\ue7a7", // F# (interface)
	".rkt":    "\ue7d4", // Racket
	".clj":    "\ue7d0", // Clojure
	".vb":     "\ufbe8", // Visual Basic
	".vba":    "\ufbe8", // Visual Basic for Applications
	
	// Markup and Config Files
	".xml":   "\ue62c",
	".yml":   "\ue60c",
	".yaml":  "\ue60c",
	".csv":   "\ue60d",
	".toml":  "\ue60e",
	".ini":   "\ue60e",
	".conf":  "\ue615",
	".tex":   "\ue69b", // LaTeX

	// Document types
	".txt":  "\uf15c", // Text
	".pdf":  "\uf1c1", // PDF
	".doc":  "\uf1c2", // Document 
	".docx": "\uf1c2", // MS Document 
	".xls":  "\uf1c3", // Excel 
	".xlsx": "\uf1c3", // Excel
	".ppt":  "\uf1c4", // PowerPoint
	".pptx": "\uf1c4", // PowerPoint
	".odt":   "\uf15c", // Open Document Text
	".ods":   "\uf1c3", // Open Document Spreadsheet
	".odp":   "\uf1c4", // Open Document Presentation
	".log":   "\uf4ed", // Log files
	".ipynb": "\ue80f", // Jupyter Notebook
	// Image types
	".jpg":  "\uf1c5", // 
	".jpeg": "\uf1c5", // 
	".png":  "\uf1c5", // 
	".gif":  "\uf1c5", // 
	".svg":  "\uf1c5", // 

	// Archives
	".zip":  "\uf1c6", // 
	".tar":  "\uf1c6", // 
	".gz":   "\uf1c6", // 
	".rar":  "\uf1c6", // 
	".7z":   "\uf1c6", // 

	// Multimedia - Video Files
	".mp4": "\uf03d", // MP4 Video
	".avi": "\uf03d", // AVI Video
	".mkv": "\uf03d", // MKV Video
	".mov": "\uf03d", // MOV Video
	".flv": "\uf03d", // FLV Video

	// eBooks
	".epub": "\uf02d",
	".mobi": "\uf02d",

	// git and other dotfiles
	".gitignore":       "\ue65d",
	".gitattributes":   "\ue65d",

	// Executable
	"executable": "\ueae8", // Binary
	"dockerfile": "\ue7b0", // Dockerfile
	"makefile":   "\ue70e", // makefile
	"cmake":      "\ue794", // CMakeFile

	// Default/unknown
	"default": "\uf15b", //
}


var Color = map[string]string {

	// Reset
	"reset": "\x1b[0m",

	// Styles
	"bold":          "\x1b[1m",
	"dim":           "\x1b[2m",
	"italic":        "\x1b[3m",
	"underline":     "\x1b[4m",
	"blink":         "\x1b[5m",
	"reverse":       "\x1b[7m",
	"hidden":        "\x1b[8m",
	"strikethrough": "\x1b[9m",

	// Foreground Colors
	"black":   "\x1b[30m",
	"red":     "\x1b[31m",
	"green":   "\x1b[32m",
	"yellow":  "\x1b[33m",
	"blue":    "\x1b[34m",
	"magenta": "\x1b[35m",
	"cyan":    "\x1b[36m",
	"white":   "\x1b[37m",

	// Bright Foreground Colors
	"bright_black":   "\x1b[90m",
	"bright_red":     "\x1b[91m",
	"bright_green":   "\x1b[92m",
	"bright_yellow":  "\x1b[93m",
	"bright_blue":    "\x1b[94m",
	"bright_magenta": "\x1b[95m",
	"bright_cyan":    "\x1b[96m",
	"bright_white":   "\x1b[97m",

	// Background Colors (prefixed with BG_)
	"BG_black":   "\x1b[40m",
	"BG_red":     "\x1b[41m",
	"BG_green":   "\x1b[42m",
	"BG_yellow":  "\x1b[43m",
	"BG_blue":    "\x1b[44m",
	"BG_magenta": "\x1b[45m",
	"BG_cyan":    "\x1b[46m",
	"BG_white":   "\x1b[47m",

	// Bright Background Colors
	"BG_bright_black":   "\x1b[100m",
	"BG_bright_red":     "\x1b[101m",
	"BG_bright_green":   "\x1b[102m",
	"BG_bright_yellow":  "\x1b[103m",
	"BG_bright_blue":    "\x1b[104m",
	"BG_bright_magenta": "\x1b[105m",
	"BG_bright_cyan":    "\x1b[106m",
	"BG_bright_white":   "\x1b[107m",

}
