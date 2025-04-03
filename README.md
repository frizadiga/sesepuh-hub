# Sesepuh Hub
CLI proxy to talk to Large Language Models (LLMs) like ChatGPT, Claude, and others

screenshots image
![sesepuh-hub](https://private-user-images.githubusercontent.com/11377023/429922735-83a204ac-8163-41b0-abc0-ce72b0ee2ea8.png?jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJnaXRodWIuY29tIiwiYXVkIjoicmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbSIsImtleSI6ImtleTUiLCJleHAiOjE3NDM2Nzg3MjQsIm5iZiI6MTc0MzY3ODQyNCwicGF0aCI6Ii8xMTM3NzAyMy80Mjk5MjI3MzUtODNhMjA0YWMtODE2My00MWIwLWFiYzAtY2U3MmIwZWUyZWE4LnBuZz9YLUFtei1BbGdvcml0aG09QVdTNC1ITUFDLVNIQTI1NiZYLUFtei1DcmVkZW50aWFsPUFLSUFWQ09EWUxTQTUzUFFLNFpBJTJGMjAyNTA0MDMlMkZ1cy1lYXN0LTElMkZzMyUyRmF3czRfcmVxdWVzdCZYLUFtei1EYXRlPTIwMjUwNDAzVDExMDcwNFomWC1BbXotRXhwaXJlcz0zMDAmWC1BbXotU2lnbmF0dXJlPWUyOWZkMDEzOTZkYWFmN2JlZjlhMWZiM2JjY2UyMmIxZWMxYmJjMDIwMTMyYTZmNTczODk5NGM5YzZmMmEzMmImWC1BbXotU2lnbmVkSGVhZGVycz1ob3N0In0.zFt_gMHq1d36blLDp-CydwoJ7sh-9K20gendPP9xa0E)

## Installation
```bash
git clone https://github.com/frizadiga/sesepuh-hub.git
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

