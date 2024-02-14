#!/bin/bash
echo "thienhang - Print mount point"
echo "--------------------"
echo "✅ Mount Points:"
mount -v | grep "^/" | awk '{print $3}'
echo

echo "❌ Unmount Points:"
awk '{print $2}' /etc/fstab | grep "^/" | while read -r line; do
    grep -qs "$line" /proc/mounts || echo "$line"
done

echo ""
check_mount_and_space() {
    local mount_point="$1"

    # Check if the given mount point is mounted
    if grep -qs "$mount_point" /proc/mounts; then
        echo "Mount point '$mount_point' is mounted."
        echo "Free space on '$mount_point':"
        df -h --output=avail "$mount_point" | tail -n 1
    else
        echo "Error: Mount point '$mount_point' is not mounted."
    fi
}

# Usage: ./check_mount_and_space.sh <Mount Point>
if [ -z "$1" ]; then
    echo "Usage: $0 <Mount Point>"
    exit 1
fi

check_mount_and_space "$1"
