#!/bin/bash

ROOT_DIR=$(dirname $(dirname $(realpath -m ${0})))

loop_unix(){
  cd ${ROOT_DIR}/bin
  ls -1d {android,darwin,linux,*bsd}/* | while read dir; do
    pushd $dir
    [[ -f k0s ]] && chmod +x k0s
    compressed="${OLDPWD}/k0s-${dir////-}.tar.gz"
    tar cz * > "$compressed"
    popd
  done
}

loop_windows(){
  cd ${ROOT_DIR}/bin
  ls -1d windows/* | while read dir; do
    pushd $dir
    [[ -f k0s.exe ]] && chmod +x k0s.exe
    compressed="${OLDPWD}/k0s-${dir////-}.zip"
    zip - * > "$compressed"
    popd
  done
}

main(){
  loop_unix
  loop_windows
}

main
