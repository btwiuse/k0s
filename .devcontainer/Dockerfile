FROM btwiuse/arch

RUN pacman -Syu --noconfirm --overwrite='*' go gopls docker

ENV PATH="${PATH}:/asdf/.asdf/shims:/asdf/.asdf/bin"

RUN git clone --depth 1 https://github.com/asdf-vm/asdf.git $HOME/.asdf && \
    echo '. $HOME/.asdf/asdf.sh' >> $HOME/.bashrc && \
    echo '. $HOME/.asdf/asdf.sh' >> $HOME/.profile

ENV ANDROID_BAZEL_NDK=True

ENV PATH=/usr/local/go/bin:go/bin:/home/btwiuse/bin/go/bin:${PATH}

ENV CGO_ENABLED=0

RUN curl -sL https://go.dev/dl/go1.19rc1.linux-amd64.tar.gz | tar -xzC /usr/local/

# asdf install golang 1.19rc1
