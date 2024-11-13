# Tree Utility

This utility displays a directory tree, listing both directories and files when the `-f` flag is provided, along with file sizes.

## Usage

To run the utility, use the following command:

```bash
go run main.go <path> [-f]
```

Example output:

```bash
$ go run main.go . -f
├───main.go (1881b)
├───main_test.go (1318b)
└───testdata
	├───project
	│	├───file.txt (19b)
	│	└───gopher.png (70372b)
	├───static
	│	├───css
	│	│	└───body.css (28b)
	│	├───html
	│	│	└───index.html (57b)
	│	└───js
	│		└───site.js (10b)
	├───zline
	│	└───empty.txt (empty)
	└───zzfile.txt (empty)

$ go run main.go .
└───testdata
	├───project
	├───static
	│	├───css
	│	├───html
	│	└───js
	└───zline
```

