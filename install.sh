#!/usr/bin/env bash
set -e

function get_arch() {
    a=$(uname -m)
    case ${a} in
    "x86_64" | "amd64")
        echo "amd64"
        ;;
    "i386" | "i486" | "i586")
        echo "386"
        ;;
    "aarch64" | "arm64")
        echo "arm64"
        ;;
    "armv6l" | "armv7l")
        echo "not support"
	;;
    "s390x")
        echo "not support"
        ;;
    *)
        echo ${NIL}
        ;;
    esac
}

function get_os() {
     echo $(uname -s | awk '{print tolower($0)}')
}

main() {
    local release="0.3.2"
    local os=$(get_os)
    local arch=$(get_arch)
    local dest_file="${HOME}/k_${release}_${os}_${arch}.tar.gz"
    local url="https://github.com/zaunist/k/releases/download/v${release}/k_${release}_${os}_${arch}.tar.gz"

    echo "[1/3] Downloading ${url}"
    rm -f "${dest_file}"
    if [ -x "$(command -v wget)" ]; then
        wget -q -P "${HOME}" "${url}"
    else
        curl -s -S -L -o "${dest_file}" "${url}"
    fi

    echo "[2/3] Install k to the ${HOME}/.k/bin"
    mkdir -p "${HOME}/.k/bin"
    tar -xz -f "${dest_file}" -C "${HOME}/.k/bin"
    chmod +x "${HOME}/.k/bin/k"

    echo "[3/3] Set environment variables"
    if [ -x "$(command -v bash)" ]; then
        cat >>${HOME}/.bashrc <<-'EOF'
		# ===== set k environment variables =====
		export PATH="${HOME}/.k/bin:$PATH"
		EOF
    fi

    if [ -x "$(command -v zsh)" ]; then
        cat >>${HOME}/.zshrc <<-'EOF'
		# ===== set k environment variables =====
		export PATH="${HOME}/.k/bin:$PATH"
		EOF
    fi

    echo "Reopen your terminal for use K,if you want to use it immediately,"
    echo "try to execute the follow command in your terminal"
    echo "source .bashrc"
    echo "source .zshrc"
    exit 0
}

main
