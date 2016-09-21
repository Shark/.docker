config_find_dir() {
  declare cur="${1:-$(pwd)}"
  if [[ -d "$cur/.docker" ]]; then
    pushd . > /dev/null
    cd "$cur/.docker"
    echo "$(pwd)"
    popd > /dev/null
    return 0
  fi

  pushd . > /dev/null
  cd "$cur"
  local abspath="$(pwd)"
  popd > /dev/null

  if [[ $abspath != '/' ]]; then
    config_find_dir "$cur/.."
  else
    return 1
  fi
}
