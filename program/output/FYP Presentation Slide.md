---
theme: seriph
background: https://source.unsplash.com/collection/94734566/1920x1080
download: true
class: text-center
highlighter: shiki
lineNumbers: false
info: |
  ## Slidev Starter Template
  Presentation slides for developers.

  Learn more at [Sli.dev](https://sli.dev)
drawings:
  persist: false
title: Welcome to Slidev
---

# SliGen

AI Text Summarizer & Slides Generator

<div class="absolute bottom-1/4 left-1/2 transform -translate-x-1/2 -translate-y-1/2 text-xs text-true-gray-600">
Presented by <strong>Ling Li Ya</strong> & <strong>Lee Kai Yang</strong>
</div>

<div class="abs-br m-6 flex gap-2">
  <a href="https://github.com/marcustut/fyp" target="_blank" alt="GitHub"
    class="text-xl icon-btn opacity-50 !border-none !hover:text-white">
    <carbon-logo-github />
  </a>
</div>

<!--
Notes
-->

---

# Problem

<div class="flex items-center justify-center">
<img src="https://i.ibb.co/n6L5FY4/image-2022-03-09-201124.png" alt="image-2022-03-09-201124" border="0">
</div>

- Students and lecturers need to create slides from long assignment reports and lecture textbooks.

---

# Solution

<div class="flex items-center justify-center">
<img src="https://i.ibb.co/NFv3xSs/image-2022-03-09-201845.png" alt="image-2022-03-09-201845" border="0">
</div>

- AI slide generator: text summarization and slide generation.

---

# Extractive Summarization

<div class="flex items-center justify-center">
<img src="https://i.ibb.co/dGjwzLP/image-2022-03-09-205221.png" alt="image-2022-03-09-205221" border="0">
</div>

- It selects the most important sentences from the document and concatenates them to form a summary.

---

# Abstractive Summarization

<div class="flex items-center justify-center">
<img src="https://i.ibb.co/zr7jM7y/image-2022-03-09-205644.png" alt="image-2022-03-09-205644" border="0">
</div>

- The summarizer will not only understand the input document but also generate its own summary from its understanding.

---

# Summarization Algorithms

<div class="flex items-center justify-center">
<img src="https://i.ibb.co/JCxrvFk/image-2022-03-09-210011.png" alt="image-2022-03-09-210011" border="0">
</div>

- These three models are based on an architecture known as the transformer architecture.

- BERT and BART are used to generate summaries for the slide body, whereas T5 is used to generate the slide title.

---

<div class="flex items-center justify-center">
<div style="width:100%;height:0px;position:relative;padding-bottom:56.338%;"><iframe src="https://streamable.com/e/6eot77?autoplay=1&nocontrols=1" frameborder="0" width="100%" height="100%" allowfullscreen allow="autoplay" style="width:100%;height:100%;position:absolute;left:0px;top:0px;overflow:hidden;"></iframe></div>
</div>

---

# System Flowchart

<div class="pt-4 pb-6 flex items-center justify-center">
<img src="https://i.ibb.co/WDyH7Zy/Untitled-Diagram-drawio.png" alt="Untitled-Diagram-drawio" border="0">
</div>

# Markdown Rule

1. "# " for title
2. "- " for bullet list item
3. "\n\n" for new line after terminating punctuation such as ", . !"
4. "\n\n---\n\n" for new slide every 150, 250 or 500 characters
