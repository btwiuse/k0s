FROM btwiuse/arch

RUN pacman -Syu --noconfirm --overwrite='*' docker clang android-ndk rustup deno pkgfile yarn npm wget htop dstat neofetch && yes | pacman -Scc

RUN curl -sL https://go.dev/dl/go1.19rc1.linux-amd64.tar.gz | tar -xzC /usr/local/

RUN cp /usr/local/go/bin/go /usr/bin/go

USER btwiuse

RUN echo 'export CGO_ENABLED=0' >> $HOME/.bashrc
RUN echo 'export EDITOR=vim' >> $HOME/.bashrc
RUN echo '. /etc/profile.d/android-ndk.sh' >> $HOME/.bashrc
RUN echo 'export PATH=/usr/local/go/bin:${HOME}/go/bin:${ANDROID_NDK}/toolchains/llvm/prebuilt/linux-x86_64/bin:${PATH}' >> $HOME/.bashrc
