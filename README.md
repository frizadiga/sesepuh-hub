# Sesepuh Hub
CLI proxy to talk to Large Language Models (LLMs) like ChatGPT, Claude, Google (Gemini), and others

![screenshot](./.assets/0.png)

## Installation
```bash
git clone https://github.com/frizadiga/sesepuh-hub.git
cd sesepuh-hub
make install
make build
# run the binary
./sesepuh-hub --prompt "eli5 general relativity?"
```
or for some direct usage convinience call it via `sesepuh-hub.sh` command
```bash
sesepuh-hub.sh "explain needle in a haystack llm?"
```

## What is Sesepuh Hub do?
- **Multi-LLM support**: Connect to multiple LLMs like ChatGPT, Claude, and others. so you always have same way to talk to them
- **Support token streaming**: You can set env `SESEPUH_NEED_STREAM=1` to get the response token by token
- **Act as standard interface to connect to LLMs**: You can use `sesepuh-hub` as a standard interface to connect to LLMs, so you can use it in your own tools

Still on early development, but you can try it out and give feedback

