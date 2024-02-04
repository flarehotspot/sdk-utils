#!/usr/bin/env node

const prod = process.env.NODE_ENV !== 'development';
const isDevkit = process.env.DEVKIT_BUILD;

const ldflags = prod ? `-ldflags="-s -w"` : '';
const trimpath = prod ? `-trimpath` : '';
const tags = prod ? (isDevkit ? `-tags="dev"` : '') : `-tags="mono dev"`;

module.exports = `${tags} ${ldflags} ${trimpath}`;
