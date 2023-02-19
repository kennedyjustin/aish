# aish - An artificially intelligent shell

`aish` is a command line utility that helps with scripting tasks. [OpenAI's Codex model](https://openai.com/blog/openai-codex/) is used to generate shell code using natural language.

## Installation

```
go install github.com/kennedyjustin/aish@latest
```

## Setup

```
aish
```

## Usage

```
aish find all files ending in .txt containing 'foo' recursively
```

## Inspiration

- https://openai.com/blog/codex-apps
- https://platform.openai.com/codex-javascript-sandbox
- https://github.com/jjviana/witty
- https://github.com/tom-doerr/zsh_codex
- https://www.warp.dev/
- https://fig.io/

## TODO

- Make more accurate
- Remember previous prompts?
- Change usage to `aish -c <command>`, with just `aish` opening a new shell entirely
- Teach usage when running aish without additional args
- Use charm.sh libraries to prompt you if the completed text is correct
- contextually aware of you system, file system, etc
- Provide mechanism for giving aish cli documentation, so it can generate advanced scripts w/ model fine tuning, etc.
- Better documentation
