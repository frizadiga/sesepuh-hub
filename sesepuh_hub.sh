#!/usr/bin/env bash
# alias: 'n/a'
# desc: fn_sesepuh_hub description.
# usage: fn_sesepuh_hub.sh [args]
# flags: @WIP:1 @TODO:0 @FIXME:0 @BUG:0 @OPTIMIZE:0 @REFACTOR:0 @DEPRECATED:0

set -e -u -o pipefail
# set -x # uncomment to debug

declare -r __self_path_file=$(readlink -f "$0")
declare -r __self_path_dir=$(dirname "${__self_path_file}")

# source "${TOOLS_DIR}/ansi-utils.sh"
# source "$HOME/Documents/tools/aliases"

# check if script run directly or indirect
# if [ "${0}" = "${BASH_SOURCE}" ]; then
#   echo "Script is being run directly"
# else
#   echo "Script is being sourced"
# fi

fn_sesepuh_hub() {
  # echo $(red $(bold 'fn_sesepuh_hub'))

  "${__self_path_dir}/sesepuh-hub" --prompt "$*"
}

fn_sesepuh_hub "$*"
