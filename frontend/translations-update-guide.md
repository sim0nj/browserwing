# Dashboard 翻译更新指南

## 需要更新的语言

简体中文 (zh-CN) 已更新 ✅  
其他语言需要手动更新 ⚠️

---

## 繁体中文 (zh-TW) 更新

在 `translations.ts` 中找到 `'zh-TW'` 部分的 dashboard，替换为：

```typescript
'dashboard.hero.title': '瀏覽器自動化平台',
'dashboard.hero.subtitle': '原生支援 MCP & Skills 協議，讓 AI 工具輕鬆控制瀏覽器',
'dashboard.quickstart.title': '快速整合',
'dashboard.quickstart.subtitle': '選擇一種方式開始使用',
'dashboard.quickstart.mcp.title': 'MCP Server',
'dashboard.quickstart.mcp.subtitle': '適用於所有支援 MCP 的 AI 工具',
'dashboard.quickstart.mcp.copy': '複製配置',
'dashboard.quickstart.mcp.copied': '已複製',
'dashboard.quickstart.mcp.copyButton': '複製配置',
'dashboard.quickstart.mcp.instruction': '貼到支援 MCP 的 AI 工具配置檔案中',
'dashboard.quickstart.skill.title': 'Skills 檔案',
'dashboard.quickstart.skill.subtitle': '適用於所有支援 Skills 的 AI 工具',
'dashboard.quickstart.skill.description': '26+ 瀏覽器操作 API，完整的自動化能力',
'dashboard.quickstart.skill.downloadButton': '下載 Skills 檔案',
'dashboard.quickstart.skill.downloading': '下載中...',
'dashboard.quickstart.skill.instruction': '匯入到支援 Skills 協議的 AI 工具中',
'dashboard.steps.title': '三步開始使用',
'dashboard.steps.step': '步驟',
'dashboard.steps.step1.title': '選擇整合方式',
'dashboard.steps.step1.desc': '複製 MCP 配置或下載 Skills 檔案',
'dashboard.steps.step2.title': '匯入 AI 工具',
'dashboard.steps.step2.desc': '將配置匯入支援 MCP/Skills 的 AI 工具',
'dashboard.steps.step3.title': '開始使用',
'dashboard.steps.step3.desc': '透過自然語言指令控制瀏覽器',
'dashboard.advanced.title': '進階功能',
'dashboard.advanced.subtitle': '需要更多自訂？探索進階特性',
'dashboard.advanced.browser.title': '錄製腳本',
'dashboard.advanced.browser.desc': '在瀏覽器中錄製操作，生成可重放的自動化腳本',
'dashboard.advanced.scripts.title': '管理腳本',
'dashboard.advanced.scripts.desc': '匯入匯出腳本，轉換為 Skills 檔案供 AI 使用',
'dashboard.advanced.config.title': 'LLM 配置',
'dashboard.advanced.config.desc': '配置 LLM 模型進行智慧資料提取',
'dashboard.advanced.learnMore': '了解更多',
```

---

## 英文 (en) 更新

在 `translations.ts` 中找到 `'en'` 部分的 dashboard，替换为：

```typescript
'dashboard.hero.title': 'Browser Automation Platform',
'dashboard.hero.subtitle': 'Native support for MCP & Skills protocols, enabling AI tools to control browsers effortlessly',
'dashboard.quickstart.title': 'Quick Integration',
'dashboard.quickstart.subtitle': 'Choose a method to get started',
'dashboard.quickstart.mcp.title': 'MCP Server',
'dashboard.quickstart.mcp.subtitle': 'Compatible with all MCP-enabled AI tools',
'dashboard.quickstart.mcp.copy': 'Copy Config',
'dashboard.quickstart.mcp.copied': 'Copied',
'dashboard.quickstart.mcp.copyButton': 'Copy Config',
'dashboard.quickstart.mcp.instruction': 'Paste into your MCP-enabled AI tool configuration',
'dashboard.quickstart.skill.title': 'Skills File',
'dashboard.quickstart.skill.subtitle': 'Compatible with all Skills-enabled AI tools',
'dashboard.quickstart.skill.description': '26+ browser automation APIs with complete capabilities',
'dashboard.quickstart.skill.downloadButton': 'Download Skills File',
'dashboard.quickstart.skill.downloading': 'Downloading...',
'dashboard.quickstart.skill.instruction': 'Import into your Skills-enabled AI tool',
'dashboard.steps.title': 'Get Started in Three Steps',
'dashboard.steps.step': 'Step',
'dashboard.steps.step1.title': 'Choose Integration',
'dashboard.steps.step1.desc': 'Copy MCP config or download Skills file',
'dashboard.steps.step2.title': 'Import to AI Tool',
'dashboard.steps.step2.desc': 'Import config into your MCP/Skills-enabled AI tool',
'dashboard.steps.step3.title': 'Start Using',
'dashboard.steps.step3.desc': 'Control browser with natural language commands',
'dashboard.advanced.title': 'Advanced Features',
'dashboard.advanced.subtitle': 'Need more customization? Explore advanced capabilities',
'dashboard.advanced.browser.title': 'Record Scripts',
'dashboard.advanced.browser.desc': 'Record browser operations and generate replayable automation scripts',
'dashboard.advanced.scripts.title': 'Manage Scripts',
'dashboard.advanced.scripts.desc': 'Import/export scripts and convert to Skills files for AI',
'dashboard.advanced.config.title': 'LLM Configuration',
'dashboard.advanced.config.desc': 'Configure LLM models for intelligent data extraction',
'dashboard.advanced.learnMore': 'Learn More',
```

---

## 西班牙语 (es) 更新

在 `translations.ts` 中找到 `'es'` 部分的 dashboard，替换为：

```typescript
'dashboard.hero.title': 'Plataforma de Automatización del Navegador',
'dashboard.hero.subtitle': 'Soporte nativo para los protocolos MCP y Skills, permitiendo que las herramientas de IA controlen los navegadores sin esfuerzo',
'dashboard.quickstart.title': 'Integración Rápida',
'dashboard.quickstart.subtitle': 'Elige un método para comenzar',
'dashboard.quickstart.mcp.title': 'Servidor MCP',
'dashboard.quickstart.mcp.subtitle': 'Compatible con todas las herramientas de IA habilitadas para MCP',
'dashboard.quickstart.mcp.copy': 'Copiar Config',
'dashboard.quickstart.mcp.copied': 'Copiado',
'dashboard.quickstart.mcp.copyButton': 'Copiar Config',
'dashboard.quickstart.mcp.instruction': 'Pegar en la configuración de tu herramienta de IA compatible con MCP',
'dashboard.quickstart.skill.title': 'Archivo Skills',
'dashboard.quickstart.skill.subtitle': 'Compatible con todas las herramientas de IA habilitadas para Skills',
'dashboard.quickstart.skill.description': 'Más de 26 API de automatización del navegador con capacidades completas',
'dashboard.quickstart.skill.downloadButton': 'Descargar Archivo Skills',
'dashboard.quickstart.skill.downloading': 'Descargando...',
'dashboard.quickstart.skill.instruction': 'Importar en tu herramienta de IA compatible con Skills',
'dashboard.steps.title': 'Comienza en Tres Pasos',
'dashboard.steps.step': 'Paso',
'dashboard.steps.step1.title': 'Elegir Integración',
'dashboard.steps.step1.desc': 'Copiar config MCP o descargar archivo Skills',
'dashboard.steps.step2.title': 'Importar a Herramienta IA',
'dashboard.steps.step2.desc': 'Importar config en tu herramienta de IA compatible con MCP/Skills',
'dashboard.steps.step3.title': 'Comenzar a Usar',
'dashboard.steps.step3.desc': 'Controlar el navegador con comandos de lenguaje natural',
'dashboard.advanced.title': 'Funciones Avanzadas',
'dashboard.advanced.subtitle': '¿Necesitas más personalización? Explora capacidades avanzadas',
'dashboard.advanced.browser.title': 'Grabar Scripts',
'dashboard.advanced.browser.desc': 'Grabar operaciones del navegador y generar scripts de automatización reproducibles',
'dashboard.advanced.scripts.title': 'Gestionar Scripts',
'dashboard.advanced.scripts.desc': 'Importar/exportar scripts y convertir a archivos Skills para IA',
'dashboard.advanced.config.title': 'Configuración LLM',
'dashboard.advanced.config.desc': 'Configurar modelos LLM para extracción inteligente de datos',
'dashboard.advanced.learnMore': 'Más Información',
```

---

## 日语 (ja) 更新

在 `translations.ts` 中找到 `'ja'` 部分的 dashboard，替换为：

```typescript
'dashboard.hero.title': 'ブラウザ自動化プラットフォーム',
'dashboard.hero.subtitle': 'MCP & Skills プロトコルをネイティブサポート、AIツールが簡単にブラウザを制御',
'dashboard.quickstart.title': 'クイック統合',
'dashboard.quickstart.subtitle': '方法を選んで始める',
'dashboard.quickstart.mcp.title': 'MCPサーバー',
'dashboard.quickstart.mcp.subtitle': 'すべてのMCP対応AIツールと互換性あり',
'dashboard.quickstart.mcp.copy': '設定をコピー',
'dashboard.quickstart.mcp.copied': 'コピー済み',
'dashboard.quickstart.mcp.copyButton': '設定をコピー',
'dashboard.quickstart.mcp.instruction': 'MCP対応AIツールの設定ファイルに貼り付け',
'dashboard.quickstart.skill.title': 'Skillsファイル',
'dashboard.quickstart.skill.subtitle': 'すべてのSkills対応AIツールと互換性あり',
'dashboard.quickstart.skill.description': '26以上のブラウザ操作API、完全な自動化機能',
'dashboard.quickstart.skill.downloadButton': 'Skillsファイルをダウンロード',
'dashboard.quickstart.skill.downloading': 'ダウンロード中...',
'dashboard.quickstart.skill.instruction': 'Skills対応AIツールにインポート',
'dashboard.steps.title': '3ステップで開始',
'dashboard.steps.step': 'ステップ',
'dashboard.steps.step1.title': '統合方法を選択',
'dashboard.steps.step1.desc': 'MCP設定をコピーまたはSkillsファイルをダウンロード',
'dashboard.steps.step2.title': 'AIツールにインポート',
'dashboard.steps.step2.desc': 'MCP/Skills対応AIツールに設定をインポート',
'dashboard.steps.step3.title': '使用開始',
'dashboard.steps.step3.desc': '自然言語コマンドでブラウザを制御',
'dashboard.advanced.title': '高度な機能',
'dashboard.advanced.subtitle': 'さらにカスタマイズが必要？高度な機能を探索',
'dashboard.advanced.browser.title': 'スクリプト録画',
'dashboard.advanced.browser.desc': 'ブラウザ操作を録画し、再生可能な自動化スクリプトを生成',
'dashboard.advanced.scripts.title': 'スクリプト管理',
'dashboard.advanced.scripts.desc': 'スクリプトのインポート/エクスポート、AIのためにSkillsファイルに変換',
'dashboard.advanced.config.title': 'LLM設定',
'dashboard.advanced.config.desc': 'インテリジェントなデータ抽出のためにLLMモデルを設定',
'dashboard.advanced.learnMore': '詳細を見る',
```

---

## 快速更新方法

由于翻译文件太大，建议：

1. 打开 `frontend/src/i18n/translations.ts`
2. 搜索对应语言的 `'dashboard.hero.title'`
3. 删除从 `'dashboard.hero.title'` 到 `'dashboard.advanced.learnMore'` 的所有行
4. 粘贴上面对应语言的新翻译

---

## 已完成

- ✅ 简体中文 (zh-CN) - 已在 translations.ts 中更新
- ⚠️  繁体中文 (zh-TW) - 需手动更新
- ⚠️  英文 (en) - 需手动更新
- ⚠️  西班牙语 (es) - 需手动更新（可选）
- ⚠️  日语 (ja) - 需手动更新（可选）
