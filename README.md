# Pokedex

A lightweight, terminal-based Pokémon exploration tool built with Go.
This project provides a simple CLI interface for discovering regions, exploring areas, catching Pokémon, and managing your personal Pokédex — all through clean, intuitive commands.

Whether you want a fun command-line toy or a hands-on Go project to deepen your backend skills, this Pokedex is fast, extensible, and easy to use.

## 🌟 Features

* **Interactive Exploration** — Traverse Pokémon regions using `map` and `mapb` to move forward and backward through routes.
* **Discover Local Pokémon** — Use `explore` to see which Pokémon inhabit the current area.
* **Catch Pokémon** — Try your luck using the `catch` command.
* **View Your Collection** — Access your caught Pokémon through the `pokedex` command.
* **Inspect Pokémon Details** — Display stats, types, and abilities with `inspect`.

Built to be simple, educational, and easy to extend if you want to add more commands or features later.

## 🎯 Motivation

Many CLI projects are either overly simplistic or require heavy frameworks.
This Pokedex strikes the balance:

* **Minimal & Fast** — No external dependencies beyond Go.
* **Perfect for Learning Go** — Great for practicing HTTP requests, JSON parsing, CLI design, and state management.
* **Extensible** — Easily expand commands or add new functionality.
* **Fun to Use** — Turns the Pokémon API into an interactive tool.

Whether you're learning Go or just want a fun terminal app, this project gives you everything you need — and nothing you don’t.

## 🚀 Installation & Usage

### Requirements

* **Go 1.21+**

No external packages needed — this project uses only the Go standard library.

### Build the Project

From the root of the repository:

```
go build
```

This creates an executable named `pokedex`.

### Run the Program

```
./pokedex
```

You’ll enter an interactive shell where you can run all available commands.

## 🕹️ Commands

Inside the Pokedex shell, you can use:

* `help` — Show all available commands
* `exit` — Quit the program
* `map` — View the next page of areas
* `mapb` — View the previous page of areas
* `explore <area>` — See which Pokémon live in an area
* `catch <pokemon>` — Attempt to catch a Pokémon
* `inspect <pokemon>` — View details for a caught Pokémon
* `pokedex` — List all Pokémon you’ve captured

## 🤝 Contributing

Clone the repository:

```
git clone https://github.com/USERNAME/pokedex
cd pokedex
```

Build and run the project:

```
go build
./pokedex
```

If you'd like to contribute, please fork the repository and submit a pull request to the `main` branch.
