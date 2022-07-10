FROM btwiuse/arch

RUN curl -sL https://go.dev/dl/go1.19rc1.linux-amd64.tar.gz | tar -xzC /usr/local/

RUN cp /usr/local/go/bin/go /usr/bin/go

RUN pacman -Syu --noconfirm --overwrite='*' \
    docker clang android-ndk rustup deno pkgfile yarn npm wget htop dstat neofetch &&\
    yes | pacman -Scc

ENV CGO_ENABLED=0
ENV ANDROID_NDK_HOME=/opt/android-ndk
ENV PATH=/usr/local/go/bin:/go/bin:${PATH}:/opt/android-ndk/toolchains/llvm/prebuilt/linux-x86_64/bin

USER btwiuse

RUN go install -v golang.org/x/tools/gopls@latest

RUN echo 'set -o vi' >> $HOME/.bashrc
RUN echo 'export EDITOR=vim' >> $HOME/.bashrc
RUN echo 'export PATH=${HOME}/go/bin:${PATH}' >> $HOME/.bashrc
