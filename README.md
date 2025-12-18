<p align="center">
  <img width="600" alt="BrowserWing" src="https://raw.githubusercontent.com/browserwing/browserwing/main/docs/assets/banner.svg">
</p>

<p align="center">
  <img alt="Go" src="https://img.shields.io/badge/Go-1.21%2B-00ADD8?logo=go&logoColor=white" />
  <img alt="React" src="https://img.shields.io/badge/React-18-61DAFB?logo=react&logoColor=white" />
  <img alt="TypeScript" src="https://img.shields.io/badge/TypeScript-5-3178C6?logo=typescript&logoColor=white" />
  <img alt="Vite" src="https://img.shields.io/badge/Vite-5-646CFF?logo=vite&logoColor=white" />
  <img alt="pnpm" src="https://img.shields.io/badge/pnpm-9-F69220?logo=pnpm&logoColor=white" />
  <img alt="MCP" src="https://img.shields.io/badge/MCP-Model%20Context%20Protocol-7B61FF" />
</p>

<p align="center">
  English · <a href="./README.zh-CN.md">简体中文</a> · <a href="./README.ja.md">日本語</a> · <a href="./README.es.md">Español</a> · <a href="./README.pt.md">Português</a>
</p>

<p align="center"><a href="https://browserwing.com">browserwing.com</a></p>


https://github.com/user-attachments/assets/7018126f-01c8-468f-a30d-3ca36f769876


## Highlights

- Browser management: launch, control, and persist sessions with cookies
- Script recording: capture clicks, inputs, navigations, and edit visually
- MCP integration: convert scripts to Model Context Protocol commands
- LLM support: configure OpenAI, Claude, DeepSeek and more for extraction

## Requirements

- Google Chrome or Chromium installed and accessible in your environment.

## Quick Start

### Option A — Download Release (recommended)

- Download the prebuilt binary for your OS from [Releases](https://github.com/browserwing/browserwing/releases).
- Run the app and open `http://localhost:8080`.

```bash
# Linux/macOS
chmod +x ./browserwing
./browserwing --port 8080

# Windows (PowerShell)
./browserwing.exe --port 8080
```

### Option B — Build from Source

```bash
# Install deps (Go + pnpm required)
make install

# Build integrated binary (frontend embedded)
make build-embedded
./build/browserwing --port 8080

# Or build all targets and packages
make build-all
make package
```

## Why BrowserWing

- Active ecosystem and rich scripts help you finish tasks fast
- Token-efficient design with fast performance
- Automates complex web tasks with reliable, replayable scripts
- Bridges recorded actions with LLM workflows via MCP commands
- Keeps sessions stable with cookie and storage management
- Designed for data extraction, RPA, testing, and agent-driven browsing

## Architecture

```mermaid
flowchart LR
  User((You))
  Frontend[React + Vite]
  Backend[Go API]
  Browser[Recorder / Player / Manager]
  MCP[MCP Server]
  LLMs[OpenAI, Claude, DeepSeek, etc.]

  User --> Frontend
  Frontend --> Backend
  Backend --> Browser
  Backend --> MCP
  MCP --> LLMs
```

## Usage Guide

1. Manage browsers: start instances, configure profiles, handle cookies
2. Record scripts: capture steps and save for replay/editing
3. Convert to MCP: expose scripts as MCP tools/commands
4. Call with LLMs: let models orchestrate browser automation via MCP

## Contributing

- Issues and PRs are welcome. Please include clear steps to reproduce or a concise rationale.
- For feature ideas, open a discussion with use cases and expected outcomes.

## Community

Discord: [https://discord.gg/BkqcApRj](https://discord.gg/BkqcApRj)
twitter: [https://x.com/chg80333](https://x.com/chg80333)

## Acknowledgements

- Inspired by modern browser automation, agentic workflows, and MCP.

## License

- MIT License. See `LICENSE`.

## Disclaimer

- Do not use for illegal purposes or to violate site terms.
- Intended for personal learning and legitimate automation only.
