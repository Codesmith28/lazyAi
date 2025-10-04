#!/bin/bash
set -e # exit if any command fails

echo "========================================"
echo " 🚀 Starting full build for lazyAi"
echo " Root directory: $(pwd)"
echo "========================================"
echo

echo "[1/2] 🪟 Building Windows version..."
(
    cd windows
    echo "   → Entered $(pwd)"
    ./build.sh
    echo "   ✓ Windows build completed"
)
echo

echo "[2/2] 🐧 Building Linux version..."
(
    cd linux
    echo "   → Entered $(pwd)"
    ./build.sh
    echo "   ✓ Linux build completed"
)
echo

echo "========================================"
echo " ✅ All builds finished successfully!"
echo " Windows binary: lazyAi/windows/lazyAi.exe"
echo " Linux binary:   lazyAi/linux/lazyAi"
echo "========================================"
