FROM btwiuse/arch

RUN pacman -Syu --noconfirm --overwrite='*' go

ENV PATH="${PATH}:/asdf/.asdf/shims:/asdf/.asdf/bin"

RUN git clone --depth 1 https://github.com/asdf-vm/asdf.git $HOME/.asdf && \
    echo '. $HOME/.asdf/asdf.sh' >> $HOME/.bashrc && \
    echo '. $HOME/.asdf/asdf.sh' >> $HOME/.profile

ENV ANDROID_BAZEL_NDK=True

ENV GOWORK=off

ENV PATH=${PATH}:/go/bin