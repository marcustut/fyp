class Adapter():
    '''Adapting plain string texts into markdown format specifically for Slidev.
    `slide_len`: number of characters in the slide body on the current slide.
    `max_len`: maximum allowed number of characters in the slide body on one slide.
    `md`: markdown output to be fed into the slide generator.
    `metadata`: Slidev metadata that defines themes and settings, which acts like a template.
    '''

    slide_len: int
    max_len: int
    md: str
    metadata: 'list[str]'

    def __init__(self, max_len=1000) -> None:
        '''Initialises an adapter'''
        self.slide_len = 0
        self.max_len = max_len
        self.md = ''
        self.metadata = [
'''---
theme: apple-basic
background: https://source.unsplash.com/collection/94734566/1920x1080
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

# Welcome to Slidev

Presentation slides for developers

<div class="pt-12">
  <span @click="$slidev.nav.next" class="px-2 py-1 rounded cursor-pointer" hover="bg-white bg-opacity-10">
    Press Space for next page <carbon:arrow-right class="inline"/>
  </span>
</div>

<div class="abs-br m-6 flex gap-2">
  <button @click="$slidev.nav.openInEditor()" title="Open in Editor" class="text-xl icon-btn opacity-50 !border-none !hover:text-white">
    <carbon:edit />
  </button>
  <a href="https://github.com/slidevjs/slidev" target="_blank" alt="GitHub"
    class="text-xl icon-btn opacity-50 !border-none !hover:text-white">
    <carbon-logo-github />
  </a>
</div>

<!--
The last comment block of each slide will be treated as slide notes. It will be visible and editable in Presenter Mode along with the slide. [Read more in the docs](https://sli.dev/guide/syntax.html#notes)
-->'''
]
        pass

    def convert_markdown(self, results: dict):
        '''Converts summarised results into markdown.'''

        # Add first slide
        self.md += self.__add_slide()

        try:
            for result in results:
                for head, val in result.items():
                    if(self.slide_len > self.max_len):# If the number of characters on a slide exceeds max_len, create new slide
                        self.md += self.__add_slide()
                        self.__reset_length()

                    if(self.slide_len == 0): # If first item in slide
                        self.md += self.__add_header(head)

                    self.md += self.__add_ulist(val)
                    self.__update_length(val)
        except:
            raise Exception('Markdown conversion error')

        try:
            self.__create_file()
        except:
            raise Exception('Markdown file creation error')

    def __update_length(self, val: str) -> None:
        '''Updates the total number of characters in the current slide.'''
        self.slide_len += len(val)

    def __reset_length(self) -> None:
        '''Resets the text length for a new slide.'''
        self.slide_len = 0

    def __add_header(self, head: str) -> str:
        '''Adds a slide header.'''
        return "# " + head + "\n"

    def __add_ulist(self, val: str) -> str:
        '''Adds an unnumbered list item.'''
        return "\n-" + val

    def __add_slide(self) -> str:
        '''Adds a slide.'''
        return "\n\n---\n\n"

    def __create_file(self) -> None:
        '''
        Creates a Markdown file with a Slidev template and writes the target Markdown string.
        First creates a text file for writing of strings, for further renaming of the `.txt` file to `.md`.
        '''

        with open('../slidev/slides.txt', 'w') as f:
            f.writelines(self.metadata)
            f.write(self.md)

        f.close()

        # For some reason, the .md file will go blank if the below command is executed throughout the main.py program. Bash shell script is used instead.
        # os.rename('slides_test.txt', 'slides_test.md')