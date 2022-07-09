FROM btwiuse/arch

RUN pacman -Syu --noconfirm --overwrite='*' go gopls docker clang android-ndk rustup deno pkgfile yarn npm wget

RUN curl -sL https://go.dev/dl/go1.19rc1.linux-amd64.tar.gz | tar -xzC /usr/local/

# asdf install golang 1.19rc1

USER btwiuse

RUN export NIGHTLY=nightly ;\
    rustup install $NIGHTLY ;\
    rustup target add wasm32-unknown-unknown --toolchain $NIGHTLY ;\
    rustup target add wasm32-wasi --toolchain $NIGHTLY ;\
    rustup default $NIGHTLY

RUN git clone --depth 1 https://github.com/asdf-vm/asdf.git $HOME/.asdf && \
    echo '. /home/btwiuse/.asdf/asdf.sh' >> $HOME/.bashrc

RUN echo 'export CGO_ENABLED=0' >> $HOME/.bashrc
RUN echo 'export EDITOR=vim' >> $HOME/.bashrc
RUN echo '. /etc/profile.d/android-ndk.sh' >> $HOME/.bashrc
RUN echo 'export PATH=/usr/local/go/bin:${HOME}/go/bin:${ANDROID_NDK}/toolchains/llvm/prebuilt/linux-x86_64/bin:${PATH}' >> $HOME/.bashrc
