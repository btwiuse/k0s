FROM btwiuse/k0s AS k0s

FROM btwiuse/arch

RUN curl -sL https://go.dev/dl/go1.19rc2.linux-amd64.tar.gz | tar -xzC /usr/local/ && cp /usr/local/go/bin/go /usr/bin/go

RUN pacman -Syu --noconfirm --overwrite='*' \
    android-ndk &&\
    yes | pacman -Scc

RUN pacman -Syu --noconfirm --overwrite='*' \
    docker clang rustup deno pkgfile yarn npm wget htop dstat neofetch cmake gopls staticcheck delve go-tools &&\
    yes | pacman -Scc

RUN pkgfile -u

ENV CGO_ENABLED=0 \
    ANDROID_NDK_HOME=/opt/android-ndk \
    PATH=/usr/local/go/bin:/go/bin:${PATH}:/opt/android-ndk/toolchains/llvm/prebuilt/linux-x86_64/bin

USER btwiuse

# RUN go install -v golang.org/x/tools/gopls@latest
# RUN go install -v honnef.co/go/tools/cmd/staticcheck@latest
# RUN go install -v github.com/go-delve/delve/cmd/dlv@latest
# RUN go install -v golang.org/dl/gotip@latest && gotip download
# RUN go install -v github.com/btwiuse/baize@latest

RUN echo 'set -o vi' >> $HOME/.bashrc &&\
    echo 'export EDITOR=vim' >> $HOME/.bashrc &&\
    echo 'export PATH=${HOME}/go/bin:${PATH}' >> $HOME/.bashrc

COPY --from=k0s /usr/bin/k0s /usr/bin/k0s