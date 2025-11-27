# updep

> **Note**: This project is currently under active development. Features and APIs may change.

Interactive TUI for updating JavaScript dependencies. Works with npm, yarn, pnpm, and bun.

## Installation

### Using npx (no installation needed)

```bash
# npm
npx updep

# Yarn
yarn dlx updep

# pnpm
pnpm dlx updep

# Bun
bunx updep
```

### Install globally

```bash
npm install -g updep
```

Then run from any project:

```bash
updep
```

## Usage

Navigate to your JavaScript project directory and run:

```bash
# npm
npx updep

# Yarn
yarn dlx updep

# pnpm
pnpm dlx updep

# Bun
bunx updep
```

### Keybindings

- `↑/↓` or `j/k` - Navigate between packages
- `Space` - Toggle package selection
- `w` - Select wanted version
- `l` - Select latest version
- `Enter` - Update selected packages
- `?` - Show full help footer
- `q` or `Ctrl+C` - Quit

## Requirements

- Node.js 18 or higher
- npm, yarn, pnpm, or bun installed
- A JavaScript project with a `package.json` file

## License

MIT - see [LICENSE](https://github.com/snikoletopoulos/updep/blob/main/LICENSE)

## Links

- [GitHub Repository](https://github.com/snikoletopoulos/updep)
- [Report Issues](https://github.com/snikoletopoulos/updep/issues)
- [Discussions](https://github.com/snikoletopoulos/updep/discussions)
