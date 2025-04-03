# Sesepuh Hub
CLI proxy to talk to Large Language Models (LLMs) like ChatGPT, Claude, and others

## Installation
```bash
git clone <this-repo>
cd sesepuh-hub
make install
make build
# use the binary
./sesepuh-hub --prompt "Hello, how are you?"
```
or for some direct usage convinience call it via `sesepuh-hub.sh` command
```bash
sesepuh-hub.sh "Hello, how are you?"
```

## What is Sesepuh Hub do?
- **Multi-LLM Support**: Connect to multiple LLMs like ChatGPT, Claude, and others. so you always have same way to talk to them
- **Support Token Streaming**: You can set env `SESEPUH_NEED_STREAM=1` to get the response token by token

Still on early development, but you can try it out and give feedback

