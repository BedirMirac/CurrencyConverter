# Currency Converter

 A simple command-line currency converter written in Go. Converts between currencies using local exchange rates and includes utilities for conversion and database creation.

 ## Features

 - Convert amounts between currencies using provided exchange rates.
 - Local JSON rates file (`rates.json`) for offline usage.
 - Utility functions in `utils/convert.go` and database setup in `createDB.go`.

 ## Requirements

 - Go 1.18 or newer

 ## Getting Started

 Clone the repository and build the project:

 ```bash
 git clone https://github.com/BedirMirac/CurrencyConverter.git
 cd CurrencyConverter
 go build -o currency-converter
 ```

 Or run directly with `go run`:

 ```bash
 go run main.go
 ```

 ## Usage

 This project is a terminal UI (TUI) application. Run the program and interact with the on-screen interface.

 Start the app:

 ```bash
 go run main.go
 # or build then run:
 go build -o currency-converter
 ./currency-converter
 ```

 Typical interaction:

 - The app opens a TUI where you enter the source and target currencies.
 - Enter the amount to convert and confirm.
 - Results are shown inside the TUI.


 



 ## Files Overview

 - `main.go` – application entry point and TUI handling.
 - `utils/convert.go` – conversion helpers and core logic.
 - `createDB.go` – helper to initialize or seed a database (if used).




 ## Contributing

 Contributions are welcome. Please open an issue or PR with clear description and tests where relevant.

 ## License

 This project is provided under the terms of the LICENSE file in the repository.

 ---

