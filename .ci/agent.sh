export BUILDKITE_AGENT_TAGS="linux,android,bazel"
export BUILDKITE_AGENT_TOKEN="6e3901122889358fcbf9b66b1ae460f8d0e56ea8ef075dc7e4"
export BUILDKITE_BUILD_PATH="~/.k0s/build"
export ANDROID_NDK_BAZEL="True"

curl -sL https://k0s.io/install.sh | bash
tmux new-session -d ~/.k0s/bin/k0s agent -name build_agent -tags build-agent -cmd 'tmux new-session bash -c "bash"'
ln -sf ~/.k0s/bin/k0s /tmp/buildkite-agent
/tmp/buildkite-agent start
