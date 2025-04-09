# datepick

Minimal date-picker terminal UI.

![Demo: selecting a date with `datepick`.](https://vhs.charm.sh/vhs-4lG0HocX3uyEmfMR5MupSx.gif)

Standalone version of the `date` date-picker command https://github.com/charmbracelet/gum/pull/683 (stale) would add to `gum`. I'll archive this tool if/when that gets merged.

## Installation

```bash
go install github.com/lukasschwab/datepick/cmd/datepick@latest
```

## Usage

Pick a date, starting from the current date by default:

```bash
datepick
```

Or starting from an initial date of your choosing:

```bash
datepick --value 2023-11-28
```

For detailed usage instructions:

```bash
datepick --help
```

<!--
<img src="https://vhs.charm.sh/vhs-2Gkiemx0ALZZBmODcSrg2I.gif" width="600" alt="Running gum date with a prompt specified" />
-->
