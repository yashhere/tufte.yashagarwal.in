VERSION=`curl -s https://api.github.com/repos/gohugoio/hugo/releases/latest | grep -oP '"tag_name": "\K(.*)(?=")'`
VERSION="${VERSION//v}"
curl -L -s https://github.com/gohugoio/hugo/releases/download/v${VERSION}/hugo_${VERSION}_Linux-64bit.tar.gz | tar -xz && mv hugo /usr/local/bin/hugo