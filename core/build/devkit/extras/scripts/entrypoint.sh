#!/bin/sh

chown -R openwrt:openwrt /app
su -c "$@" openwrt
