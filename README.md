# Sesepuh Hub
CLI proxy to talk to Large Language Models (LLMs) like ChatGPT, Claude, and others

screenshots image
![sesepuh-hub](https://private-user-images.githubusercontent.com/11377023/429922735-83a204ac-8163-41b0-abc0-ce72b0ee2ea8.png?jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJnaXRodWIuY29tIiwiYXVkIjoicmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbSIsImtleSI6ImtleTUiLCJleHAiOjE3NDUxNTI0OTMsIm5iZiI6MTc0NTE1MjE5MywicGF0aCI6Ii8xMTM3NzAyMy80Mjk5MjI3MzUtODNhMjA0YWMtODE2My00MWIwLWFiYzAtY2U3MmIwZWUyZWE4LnBuZz9YLUFtei1BbGdvcml0aG09QVdTNC1ITUFDLVNIQTI1NiZYLUFtei1DcmVkZW50aWFsPUFLSUFWQ09EWUxTQTUzUFFLNFpBJTJGMjAyNTA0MjAlMkZ1cy1lYXN0LTElMkZzMyUyRmF3czRfcmVxdWVzdCZYLUFtei1EYXRlPTIwMjUwNDIwVDEyMjk1M1omWC1BbXotRXhwaXJlcz0zMDAmWC1BbXotU2lnbmF0dXJlPTkxZWY0YmJiYmI2ZTE2ZGUwZTgxNzUwMDE5MWQxZTVkOGE4NWZiOTg1MzRkNmUwMzM1OGJlNTNmNzg0MGMyZjgmWC1BbXotU2lnbmVkSGVhZGVycz1ob3N0In0.9pzTQBHWGzYNuJfzVIeR9iusXnpMDToAb94sWxvmS2s)

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

