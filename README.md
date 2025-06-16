# Sesepuh Hub
CLI proxy to talk to Large Language Models (LLMs) like ChatGPT, Claude, Google (Gemini), and others

![screenshot](./.assets/0.png)

Sesepuh Hub is a tool designed to interact with various AI models, including those from OpenAI, Anthropic, Google, and more. You can utilize different AI models for generating content based on user-defined prompts.

## What is Sesepuh Hub do?
- **Multi-LLM support**: Connect to multiple major LLMs like OpenAI, Claude, Gemini, and others. so you always have same way to talk to them
- **Support token streaming**: You can set env `SESEPUH_NEED_STREAM=1` to get the response token by token
- **Act as a standard interface in stdout to connect to LLMs**: You can use `sesepuh-hub` as a standard interface to connect to LLMs, allowing integration with your own tools

## Features

- Support for multiple AI vendors:
  - OpenAI
  - Anthropic
  - Google
  - Mistral
  - Ollama
  - XAI
- Synchronization and streaming modes for interactions with AI models.
- Easy prompt handling through command line flags or environment variables.

## Getting Started

### Prerequisites

- Go version 1.24.0 or later
- Environment variables for API keys and models set up:
  - `OPENAI_API_KEY`
  - `ANTHROPIC_API_KEY`
  - `GEMINI_API_KEY`
  - `MISTRAL_API_KEY`
  - `XAI_API_KEY`
  - `SESEPUH_HUB_MODEL` (optional)
  - `SESEPUH_HUB_VENDOR` (choose your vendor)

### Installation

```bash
git clone https://github.com/frizadiga/sesepuh-hub.git
cd sesepuh-hub
make install
make build
```

run the binary with a prompt:
```bash
# run the binary
./sesepuh-hub --prompt "eli5 general relativity?"
```

or for some direct usage convinience call it via `sesepuh-hub.sh` command
```bash
sesepuh-hub.sh "explain needle in a haystack llm?"
```

### Select Model and Vendor

Sesepuh Hub provides an interactive model selection tool that lets you choose from all available models without worrying about specific vendors. This tool automatically sets the required environment variables for your selected model.

#### Interactive Model Selection

Use the interactive model selector to browse and choose from all available models:

```bash
# Using Make target
make select_model

# Or source the script directly
source ./select_model.sh
```

This will:
1. Display a searchable list of all available models using `fzf`
2. Show model details (name, vendor, description, max tokens, etc.)
3. Automatically set the environment variables:
   - `SESEPUH_HUB_MODEL` - The selected model name
   - `SESEPUH_HUB_VENDOR` - The corresponding vendor

#### Example Usage

```bash
# Select a model interactively
$ make select_model

# The tool will show available models in fzf:
# > gpt-4
#   gpt-3.5-turbo
#   claude-3-sonnet
#   gemini-2.0-flash
#   ...

# After selection, environment variables are automatically set:
Selected:
name: gemini-2.0-flash
vendor: google
description: "Latest Gemini model with improved performance"
max_tokens: 1048576

SESEPUH_HUB_MODEL: gemini-2.0-flash
SESEPUH_HUB_VENDOR: google

# Now you can use sesepuh-hub with the selected model
$ ./sesepuh-hub --prompt "Hello, how are you?"
```

#### Manual Environment Setup

If you prefer to set the environment variables manually:

```bash
# Set your preferred model and vendor
export SESEPUH_HUB_MODEL="gpt-4"
export SESEPUH_HUB_VENDOR="openai"

# Or for other models:
export SESEPUH_HUB_MODEL="claude-3-sonnet"
export SESEPUH_HUB_VENDOR="anthropic"

export SESEPUH_HUB_MODEL="gemini-2.0-flash"
export SESEPUH_HUB_VENDOR="google"
```

#### Available Models

The model selection tool reads from `enums.yml` which contains all supported models with their specifications:

- **OpenAI**: gpt-3.5-turbo, gpt-4, gpt-4-turbo, gpt-4o, etc.
- **Anthropic**: claude-3-haiku, claude-3-sonnet, claude-3-opus, etc.
- **Google**: gemini-1.5-pro, gemini-2.0-flash, etc.
- **Mistral**: mistral-small, mistral-medium, mistral-large, etc.
- **XAI**: grok-beta, grok-vision-beta, etc.
- **Ollama**: Local models via Ollama

#### Tips

- Use arrow keys or type to search/filter models in the interactive selector
- Press `Esc` to cancel the selection
- The environment variables persist in your current shell session
- You can check current settings with: `echo $SESEPUH_HUB_MODEL $SESEPUH_HUB_VENDOR`

### Development Usage

To run the application with a prompt, use the command:

```bash
make dev PROMPT="your prompt here"
```

You can also run with specific vendors:

```bash
make ollama PROMPT="your prompt here"
make openai PROMPT="your prompt here"
make gemini PROMPT="your prompt here"
make mistral PROMPT="your prompt here"
make claude PROMPT="your prompt here"
make xai PROMPT="your prompt here"
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contributing

If you'd like to contribute, please fork the repository and submit a pull request. Contributions are welcome!
