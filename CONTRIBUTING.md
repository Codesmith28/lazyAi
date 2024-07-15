# Contributing to LazyAI

First off, thank you for considering contributing to LazyAI! It's people like you that make LazyAI such a great tool.

LazyAI is a Go TUI (Terminal User Interface) application that brings AI assistance right to your clipboard. We welcome contributions that help improve its functionality, performance, and user experience.

## Submitting Issues

Before submitting a new issue, please search to see if someone has filed a similar issue before. If there is a similar issue, you can add your comments to the existing issue.

When submitting a new issue, please provide:

1. A clear and descriptive title
2. A detailed description of the issue
3. Steps to reproduce the problem
4. Expected behavior
5. Actual behavior
   
## Contributing Code

Here's how you can contribute code to LazyAI:

1. Fork the repository on GitHub.
2. Clone your fork locally:
 ```
 git clone https://github.com/Codesmith28/lazyAi.git
 ```
3. Redirect to main directory
```
cd lazyAi
```
4. Run the application:
- with UI
  ```
  go run main.go
  ```
- in detached mode
  ```
  go run main.go -d
  ```
*Note*: A log of your last state will be created in `lazyai.log` in the location where your executable is present. This can be helpful for debugging.

5. Create a new branch for your feature or bug fix:
```
git checkout -b feature-or-fix-name
```
6. Make your changes and commit them with a clear commit message.
7. Push your changes to your fork on GitHub:
```
git push origin feature-or-fix-name
```
8. Thats it! You can create a PR now when all necessary changes are made.
