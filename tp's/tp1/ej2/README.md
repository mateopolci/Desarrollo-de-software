# 1. Aplicación de formularios (o ventanas)

Deberá contar con, por lo menos, dos botones que realicen dos tareas diferentes. Éstas pueden ser, matemáticas, lógicas o de cadenas de caracteres. No podrán coincidir con las del ejercicio anterior.

## General requirements

- [Golang](https://go.dev/doc/install) 1.20 or higher.
- NPM ([Node](https://nodejs.org/en/download/prebuilt-installer/current) 15 or higher).
- [Wails](https://wails.io/docs/gettingstarted/installation).

### Windows requirements

- [WebView2](https://developer.microsoft.com/en-us/microsoft-edge/webview2/?form=MA13LH)

### Linux requirements

- Linux requires the standard gcc build tools plus libgtk3 and libwebkit. Rather than list a ton of commands for different distros, Wails can try to determine what the installation commands are for your specific distribution. Run `wails doctor` after installation to be shown how to install the dependencies. If your distro/package manager is not supported, please consult the Add Linux Distro guide.

### MacOs requirements

- Wails requires that the xcode command line tools are installed. This can be done by running `xcode-select --install`.

## How to run

1) Clone the repo.
2) Enter ej2 directory with: `cd ej2`
3) Run `go mod tidy`.
4) Enter frontend directory with: `cd frontend`.
5) Install proyect dependencies: `npm install`.
6) Go back to ej2: `cd ../`.
    - Live development window: Run `wails dev`.
    - Build an executable: Run `wails build` (Generated in ./build/bin).
