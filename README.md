# templator

## Add dependency

```
go get github.com/ssongin/templator
```

## YAML Configuration

This project uses a YAML file to pre compile files using templates. Below is the structure and examples for your configuration.

### Joiners

Joins multiple file or folders into one using template

#### Structure

```yaml
joiners:
  - template: <path to template file>
    destination: <output directory>
    join:
      - generate: <output file name>
        source:
          - <file path, glob or directory>
          - ...
```

#### Fields

- **template**: Path to the template file used for joining.
- **destination**: Directory where the generated files will be placed.
- **join**: List of join operations.
  - **generate**: Name of the output file to generate.
  - **source**: List of source files, directories, or glob patterns to include.

#### Examples

##### Example 1: Explicit Files

```yaml
joiners:
  - template: ./resources/template/util/filejoin.tmpl
    destination: ./resources/temp
    join:
      - generate: source.temp
        source:
          - ./resources/js/global/constants.js
          - ./resources/js/htmx/base/htmx-base.js
          - ./resources/css/global/global.css
          - ./resources/css/global/constants.css
          - ./resources/html/index.html
```

##### Example 2: Using Glob Patterns and Directories

```yaml
joiners:
  - template: ./resources/template/util/filejoin.tmpl
    destination: ./resources/temp
    join:
      - generate: source.temp
        source:
          - ./resources/js/**/*.js      # All JS files recursively
          - ./resources/css/global/     # All files in this directory (non-recursive)
          - ./resources/html/index.html # Single file
```

#### Notes

- You can use glob patterns (e.g., `**/*.js`) to include multiple files.
- If you specify a directory, all files in that directory (non-recursive) will be included.
- All paths should be relative to the project root or the location of the YAML file.
- Path injection is prevented: only files within allowed directories