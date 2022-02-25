# TODO: Remove blank page

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

    def __init__(self, theme='apple-basic', max_len=500) -> None:
        '''Initialises an adapter'''
        self.slide_len = 0
        self.max_len = max_len
        self.md = ''
        self.metadata = [
f'''---
theme: {theme}
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

    def convert_markdown(self, file_name: str, results: dict) -> str:
        '''Converts summarised results into markdown.'''

        # Add first slide
        self.md += self.__add_newline()
        self.md += self.__add_slide()

        try:
            for result in results:
                for head, val in result.items():
                    val = self.__strip_whitespaces(summary=val)
                    # If the number of characters on a slide exceeds max_len, create new slide
                    if(self.slide_len > self.max_len):
                        self.md += self.__add_newline()
                        self.md += self.__add_slide()
                        self.__reset_length()

                    self.md += self.__add_newline()
                    self.md += self.__add_header(head)
                    self.md += self.__add_ulist(val)
                    self.__update_length(val)
        except:
            raise Exception('Markdown conversion error')

        try:
            self.md += '\n' # Add this to prevent client error
            file_name = self.__create_file(file_name)
        except:
            raise Exception('Markdown file creation error')

        return file_name

    def __strip_whitespaces(self, summary: str) -> str:
        '''Strips additional whitespaces in front of punctuations.'''
        summary = summary.replace(' .', '.')
        summary = summary.replace(" !", "!")
        summary = summary.replace(" ?", "?")

        return summary

    def __update_length(self, val: str) -> None:
        '''Updates the total number of characters in the current slide.'''
        self.slide_len += len(val)

    def __reset_length(self) -> None:
        '''Resets the text length for a new slide.'''
        self.slide_len = 0

    def __add_header(self, head: str) -> str:
        '''Adds a slide header.'''
        return '# ' + head + '\n'

    def __add_ulist(self, val: str) -> str:
        '''Adds an unnumbered list item.'''
        return '\n-' + val

    def __add_slide(self) -> str:
        '''Adds a slide.'''
        return '---'

    def __add_newline(self) -> str:
        return '\n\n'

    def __create_file(self, file_name: str) -> str:
        '''
        Creates a Markdown file with a Slidev template and writes the target Markdown string.
        First creates a text file for writing of strings, for further renaming of the `.txt` file to `.md`.
        '''

        # file_name = "slides.txt"

        # with open('../slidev/slides.txt', 'w', encoding='utf-8') as f:
        with open('./output/' + file_name, 'w', encoding='utf-8') as f:
            f.writelines(self.metadata)
            f.write(self.md)

        f.close()

        return file_name

        # For some reason, the .md file will go blank if the below command is executed throughout the main.py program. Bash shell script is used instead.
        # os.rename('slides_test.txt', 'slides_test.md')