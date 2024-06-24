#!/usr/bin/env sh

VERSION=V1.0.0

# Determine system type
SYSTEM_TYPE=$(uname -s)

# Determine architecture type
ARCH_TYPE=$(uname -m)

# Check if the script is being run with root privileges
if [ "$(id -u)" -ne 0 ]; then
    SUDO="sudo"
fi

# Create root directory
$SUDO mkdir -p /opt/godem

# Construct the download URL based on system type and architecture type
DOWNLOAD_URL="https://github.com/Luna-CY/dem/releases/download/${VERSION}"

case "$SYSTEM_TYPE" in
    Linux)
        case "$ARCH_TYPE" in
            x86_64)
                DOWNLOAD_URL="${DOWNLOAD_URL}/linux-amd64-${VERSION}.tar.gz"
                ;;
            arm64)
                DOWNLOAD_URL="${DOWNLOAD_URL}/linux-arm64-${VERSION}.tar.gz"
                ;;
        esac
        ;;
    Darwin)
        case "$ARCH_TYPE" in
            x86_64)
                DOWNLOAD_URL="${DOWNLOAD_URL}/darwin-amd64-${VERSION}.tar.gz"
                ;;
            arm64)
                DOWNLOAD_URL="${DOWNLOAD_URL}/darwin-arm64-${VERSION}.tar.gz"
                ;;
        esac
        ;;
esac

# Download and extract the file
curl -L -o /opt/godem/godem-${VERSION}.tar.gz "$DOWNLOAD_URL"
$SUDO mkdir -p /usr/local/bin
$SUDO tar zxf /opt/godem/godem-${VERSION}.tar.gz -C /usr/local/bin
