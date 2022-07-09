FROM btwiuse/arch

RUN pacman -Syu --noconfirm --overwrite='*' go gopls docker clang android-ndk

RUN curl -sL https://go.dev/dl/go1.19rc1.linux-amd64.tar.gz | tar -xzC /usr/local/

# asdf install golang 1.19rc1

USER btwiuse

RUN echo 'export CGO_ENABLED=0' >> $HOME/.bashrc
RUN echo 'export EDITOR=vim' >> $HOME/.bashrc
RUN echo 'export PATH=/usr/local/go/bin:${HOME}/go/bin:${PATH}' >> $HOME/.bashrc
RUN echo '. /etc/profile.d/android-ndk.sh' >> $HOME/.bashrc

RUN git clone --depth 1 https://github.com/asdf-vm/asdf.git $HOME/.asdf && \
    echo '. $HOME/.asdf/asdf.sh' >> $HOME/.bashrc
