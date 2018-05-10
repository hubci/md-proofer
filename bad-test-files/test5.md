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

This is a second error, only one space and a tab.

```yaml
invalid:
	yaml: false
invalid2:
 yaml: false
 ```
