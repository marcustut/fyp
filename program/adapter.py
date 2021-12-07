import os

class Adapter():
    '''Adapting plain string texts into markdown format specifically for Slidev.
    `slide_len`: number of characters in the slide body on the current slide.
    `max_len`: maximum allowed number of characters in the slide body on one slide.
    `md`: markdown output to be fed into the slide generator.
    `metadata`: Slidev metadata that defines themes and settings.
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
            '---\n\n',
            'theme: apple-basic\n',
            'background: https://source.unsplash.com/collection/94734566/1920x1080\n',
            'class: text-center\n',
            'highlighter: shiki\n',
            'lineNumbers: false\n',
            '''
            info: |
                ## Slidev Starter Template
                Presentation slides for developers.

                Learn more at [Sli.dev](https://sli.dev)''',
            '''
            drawings:
                persist: false
            ''',
            'title: Welcome to Slide\n',
            '\n---\n\n',
            '# Welcome to Slidev\n',
            'Presentation slides for developers\n',
            '''
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
            ''',
            '''
            <!--
                The last comment block of each slide will be treated as slide notes. It will be visible and editable in Presenter Mode along with the slide. [Read more in the docs](https://sli.dev/guide/syntax.html#notes)
                -->
            '''
        ]
        pass

    def convert_markdown(self, results: dict):
        '''Converts summarised results into markdown.'''

        # Add first slide
        self.md += self.__add_slide()

        try:
            for result in results:
                for head, val in result.items():
                    if(self.slide_len == 0): # If first item
                        self.md += self.__add_header(head)
                        self.md += self.__add_ulist(val)
                        self.__update_length(val)
                    elif(self.slide_len < 1000 and result != results[-1]): # If the number of characters on a slide exceeds 1000, create new slide
                        self.__add_ulist(val)
                        self.__update_length(val)
                    else:
                        self.md += self.__add_slide()
                        self.__reset_length()
        except:
            raise Exception('Markdown conversion error')

        try:
            self.__create_file()
        except:
            raise Exception('Markdown file creation error')

    def __update_length(self, val: str) -> None:
        '''Counts the total number of characters in the current slide.'''
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
        First creates a text file for writing of strings, then renames the `.txt` file to `.md`.
        '''

        with open('slides_test.txt', 'w') as f:
            f.writelines(self.metadata)
            f.write(self.md)

        f.close()

        os.rename('slides_test.txt', 'slides_test.md')