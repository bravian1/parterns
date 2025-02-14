# Fractal
A WebAssembly-powered geometric pattern generator built with Go.

## Overview

Fractal is a creative coding project that generates mesmerizing geometric patterns using WebAssembly and Go. It currently specializes in creating random line patterns within triangular boundaries, with plans to expand into more complex fractal-based designs.

## Features

- Real-time pattern generation using WebAssembly
- Efficient Go-based rendering engine
- Triangle-based boundary system
- Random line pattern generation
- Custom WebAssembly file system integration
- Lightweight architecture for optimal performance

## Quick Start

1. Clone the repository:
    ```bash
    git clone https://github.com/bravian1/parterns.git
    cd parterns
    ```

2. Build the WebAssembly module:
    ```bash
    GOOS=js GOARCH=wasm go build -o main.wasm
    cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
    ```

3. Start the server:
    ```bash
    go run server.go
    ```

## Implementation

Include the WebAssembly executor in your HTML:
```html
<script src="wasm_exec.js"></script>
```

## Future Development

- Fractal-based pattern generation
- Multiple geometric shape support
- Color schemes and gradients
- Interactive pattern controls
- Enhanced rendering options

## Contributing

1. Fork the repository
2. Create your feature branch: `git checkout -b feature/amazing-feature`
3. Commit your changes: `git commit -m 'Add amazing feature'`
4. Push to the branch: `git push origin feature/amazing-feature`
5. Open a Pull Request

## License

Licensed under the BSD License. See [LICENSE](LICENSE) for details.

## Acknowledgements

- Go team for WebAssembly support
- Contributors and community members
