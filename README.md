<center><h1>LazyAI💤</h1></center>

LazyAI is a Go TUI (Terminal User Interface) application that brings AI assistance right to your clipboard. It features a dashboard to manage history, AI model, and prompts. Users can simply copy text, and LazyAI automatically sends it to the AI, generates a response, and copies it to the clipboard for easy pasting.

## Table of contents

- [Table of contents](#table-of-contents)
- [Demonstration](#demonstration)
- [Commands](#commands)
- [Screenshots](#screenshots)
- [Features](#features)
  - [✨ **Seamless Clipboard Integration**](#-seamless-clipboard-integration)
  - [🎛️ **Intuitive Dashboard**](#️-intuitive-dashboard)
  - [🚀 **Automatic Text Processing**](#-automatic-text-processing)
  - [⚡ **Instant AI Responses**](#-instant-ai-responses)
  - [🔧 **Flexible Configuration**](#-flexible-configuration)
  - [🖥️ **Terminal-based User Interface**](#️-terminal-based-user-interface)
- [Installation](#installation)
- [Advanced Options](#advanced-options)
  - [Detached mode](#detached-mode)
  - [Provide default prompt](#provide-default-prompt)
- [Contributions](#contributions)
- [Contributors](#contributors)
- [Acknowledgements](#acknowledgements)

## Demonstration
<!-- video here -->

https://github.com/user-attachments/assets/2fddd89d-6f78-452a-b5e0-271a21a13e3a

## Commands

Most of the commands mentioned below, you can access by running

```
lazyAi -help
```

To start application

```
lazyAi
```

To run in detached mode

```
lazyAi -d
```

Run by providing default prompt (works with and without detached mode)

```
lazyAi -p "my default prompt"
```

## Screenshots

![Screenshot_16-07-2024_170814](https://github.com/user-attachments/assets/a1be0381-a590-4a7d-903c-5221234981ed)

## Features

#### ✨ **Seamless Clipboard Integration**

- Effortlessly interact with AI by simply copying text
- No need to switch between applications or manually paste content

#### 🎛️ **Intuitive Dashboard**

- Manage your interaction history at a glance
- Easily switch between AI models for different tasks
- Customize and save prompts for quick access

#### 🚀 **Automatic Text Processing**

- LazyAI detects when you've copied text and springs into action
- No manual triggering required - it's always ready to assist

#### ⚡ **Instant AI Responses**

- Generated content is immediately copied to your clipboard
- Paste AI-generated responses instantly, anywhere you need them

#### 🔧 **Flexible Configuration**

- Run in normal or detached mode to suit your workflow
- Set default prompts for specialized tasks or projects

#### 🖥️ **Terminal-based User Interface**

- Lightweight and fast, perfect for power users
- No need for a graphical environment - use it right in your terminal

## Installation

Install LazyAI using the to our releases page: [LazyAI Releases](link)

## Advanced Options

### Detached mode

LazyAI can run in detached mode, which means it will run in the background and automatically start when you copy text. To run in detached mode, use the `-d` flag:

### Provide default prompt

You can provide default prompt, using `-p` flag, while executing the main command, this becomes helpful when you are running *cli* in *detached mode*.

## Contributions

We welcome and appreciate contributions from the community! Whether it's bug fixes, feature enhancements, documentation improvements, or any other valuable input, your contributions help make LazyAI better for everyone.

Before making a contribution, please read our [CONTRIBUTING.md](CONTRIBUTING.md) file for guidelines on how to submit your changes and the pull request process.

We look forward to your pull requests and thank you for your interest in improving LazyAI!

## Contributors

To be made with [Contrib rocks](https://contrib.rocks/)

## Acknowledgements

LazyAI wouldn't be possible without these amazing open-source projects:

- [clipboard by atotto](https://github.com/atotto/clipboard)
- [glamour by charmbracelet](https://github.com/charmbracelet/glamour)
- [tcell by gdamore](https://github.com/gdamore/tcell)
- [systray by getlantern](https://github.com/getlantern/systray)
- [tview by rivo](https://github.com/rivo/tview)
