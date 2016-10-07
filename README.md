# orca

Orca is an opinionated tool for running your local development environment in [Docker](https://docker.com) containers. It lets you specify the desired setup along your project files in the `.orca` directory. It has a `prepare` feature which lets you separate the Docker image build into two steps. This allows you to use an ssh-agent for downloading private repositores (i.e. private Rubygems).

## Usage

```
# manage (global) dependency services
orca [--global] up|down|run|prepare|...

# manage the app with docker-compose
orca app up|down|run|...

# build an application image
orca build [--force-rebuild]
```

## Contributing
1. Fork it!
2. Create your feature branch: `git checkout -b my-new-feature`
3. Commit your changes: `git commit -am 'Add some feature'`
4. Push to the branch: `git push origin my-new-feature`
5. Submit a pull request! :)

## History

- v0.0.1 (2016-10-07): initial version

## License

This project is licensed under the MIT License. See LICENSE for details.
