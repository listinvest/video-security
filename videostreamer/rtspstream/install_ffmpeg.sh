#set -ex

export LD_LIBRARY_PATH=$HOME/compiled/lib/
export PATH="$HOME/compiled/bin":$PATH
export PKG_CONFIG_PATH=$HOME/compiled/lib/pkgconfig

if [ ! -e "$HOME/nasm" ]; then
  # sudo apt-get -y install asciidoc xmlto # this fails :(
  git clone http://repo.or.cz/nasm.git "$HOME/nasm"
  cd "$HOME/nasm"
  ./autogen.sh
  ./configure --prefix="$HOME/compiled"
  make
  make install || echo "Installing docs fails but should be OK otherwise"
fi

if [ ! -e "$HOME/x264" ]; then
  git clone http://git.videolan.org/git/x264.git "$HOME/x264"
  cd "$HOME/x264"
   #git master as of this writing
   #git checkout 7d0ff22e8c96de126be9d3de4952edd6d1b75a8c
  ./configure --prefix="$HOME/compiled" --enable-pic --enable-static 
  make
  make install-lib-static
fi

if [ ! -e "$HOME/ffmpeg/libavcodec/libavcodec.a" ]; then
  git clone https://git.ffmpeg.org/ffmpeg.git "$HOME/ffmpeg" || echo "FFmpeg dir already exists"
  cd "$HOME/ffmpeg"
  ./configure --prefix="$HOME/compiled" --enable-libx264 --enable-gnutls --enable-gpl --enable-static --enable-shared
  make
  make install
fi

#gedit ~/.bashrc
#export GOPATH=$HOME/projects/
#export GOROOT=$HOME/go
#export PATH=$PATH:$GOROOT/bin:$HOME/compiled/bin/
#export LD_LIBRARY_PATH=$HOME/compiled/lib/
#export PKG_CONFIG_PATH=$PKG_CONFIG_PATH:$HOME/compiled/lib/pkgconfig
#source ~/.bashrc