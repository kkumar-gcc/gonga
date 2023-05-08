# Gonga - A Social Media API Framework

Gonga is a social media API framework written in Go that provides developers with the tools they need to build their own social media platforms. With Gonga, you can easily add user authentication, social connections, and much more to your application.

## Features

- User authentication
- Social connections (friends, followers, etc.)
- Post and comment creation
- Direct messaging
- Push notifications


## Getting Started

To get started with Gonga, simply clone the repository and run the following command:

```sh
go run main.go serve
```
This will start the Gonga server on port 8080. You can then start making requests to the API.

## Gonga CLI Commands

Gonga also includes CLI commands that can be used to manage your application. Here are the available commands:

- `migrate`- Run database migrations.
- `seed` - Seed the database with initial data.
- `serve` - Start the Gonga server.
- `test` - Run the test suite.

To run a command, use the following syntax:

```sh
go run main.go <command>
```

For example, to run the migrations command:

```sh
go run main.go migrate
```

## Contributing

If you are interested in contributing to Gonga, please read our [contribution guidelines]("https://github.com/kkumar-gcc/gonga/blob/main/CONTRIBUTING.md") before getting started. We welcome all contributions, big or small!

## License

Gonga is licensed under the [MIT License]("").