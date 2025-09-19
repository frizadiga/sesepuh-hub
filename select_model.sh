#!/usr/bin/env bash
# alias: `n/a`
# desc: flatten choice of model without worrying about vendor
# usage: fn_select_model.sh [args]

# set -x # uncomment to debug

__self_path_file=$(readlink -f "$0")
__self_path_dir=$(dirname "${__self_path_file}")

fn_select_model() {
  local model_name=$(yq -r '.models[].name' "${__self_path_dir}/enums.yml" | fzf) || exit 1

  # query once
  model_data=$(yq '.models[] | select(.name == "'"${model_name}"'")' "${__self_path_dir}/enums.yml")

  if [ -z "${model_data}" ]; then
    echo "Error: Model '${model_name}' not found in enums.yml" >&2
    exit 1
  fi

  echo '' >&2
  echo 'Selected:' >&2
  echo "$model_data" >&2
  echo '' >&2

  # Extract values
  local model=$(echo "${model_data}" | yq '.name')
  local vendor=$(echo "${model_data}" | yq '.vendor')

  echo "SESEPUH_HUB_MODEL=${model}" >&2
  echo "SESEPUH_HUB_VENDOR=${vendor}" >&2

  export SESEPUH_HUB_MODEL="${model}"
  export SESEPUH_HUB_VENDOR="${vendor}"
}

fn_select_model "$@"
