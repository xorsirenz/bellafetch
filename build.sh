#!/bin/sh
PACKAGE_NAME="bellafetch"
PACKAGE_VERSION="2.0.0"
SOURCE_FILE="bellafetch.go"

TARGETS=(
    "linux/amd64"
    )

if [ $(whoami) = "root" ]; then
  :
else
  echo -e "\033[31mError: \033[0mNot running as root!"
  exit 1
fi

for target in "${TARGETS[@]}"; do
    goos_arch="${target%:*}"
    IFS='/' read -r goos goarch <<< "$goos_arch"
    GOOS="$goos"
    GOARCH="$goarch"
    
    echo "Building $FILE_NAME for $GOOS/$GOARCH..."
    
    FILE_NAME="${PACKAGE_NAME}-${PACKAGE_VERSION}-${GOOS}-${GOARCH}"

    go build -o "${FILE_NAME}" cmd/"${SOURCE_FILE}"
    tar -czvf "${FILE_NAME}.tar.gz" "${FILE_NAME}"

    if cp ${FILE_NAME} /usr/local/bin/${PACKAGE_NAME}; then
        rm "${FILE_NAME}"
        echo "✅ Build successful: ${FILE_NAME}"
    else
        echo "❌ Build failed!"
        exit 1
    fi
done
