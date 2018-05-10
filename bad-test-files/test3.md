---
title: "Test Doc with broken YAML code block"
---

Hi this is a random paragraph. Should pass.

```yaml
valid:
  yaml: true
```

This is more stuff. Should fail due to a tab.

```yaml
invalid:
	yaml: false
```
Even more stuff. Should fail due to a single-space indent.

```yaml
invalid:
 yaml: false
```
