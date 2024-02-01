+++
title = "Http Middlewares"
description = "This section explains the structure of the SDK (Sofareware Development Kit) API."
date = 2021-05-01T08:00:00+00:00
updated = 2021-05-01T08:00:00+00:00
draft = false
weight = 30
sort_by = "weight"
template = "docs/page.html"

[extra]
lead = 'This section explains the structure of the SDK (Sofareware Development Kit) API.'
toc = true
top = false
+++

## Overview

The SDK API is the primary method of extending the functionality of the system. It provides access to manipulate system accounts, network devices, theme configuration, user sessions and payment methods. Each plugin is provided with an instance of [`PluginApi`](../plugin-api), the root interface of our SDK's API.
