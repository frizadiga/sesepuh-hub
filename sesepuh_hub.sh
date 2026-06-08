#!/usr/bin/env bash
# alias: sesepuh_hub
# desc: Sesepuh Hub CLI wrapper for interacting with various LLMs (OpenAI, Anthropic, Google, etc.)
# usage: sesepuh_hub.sh "your prompt here"

set -e -u -o pipefail
# set -x # uncomment to debug

declare -r __self_path_file=$(readlink -f "$0")
declare -r __self_path_dir=$(dirname "${__self_path_file}")

fn_sesepuh_hub() {
  "${__self_path_dir}/sesepuh-hub" --prompt "$*"
}

fn_sesepuh_hub "$*"
