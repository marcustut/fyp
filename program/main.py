from inputer import Inputer
from text_summarizer import TextSummarizer
from adapter import Adapter
# from mdutils.mdutils import MdUtils
from bs4 import BeautifulSoup
import requests
import transformers
transformers.logging.set_verbosity_debug()

# def get_article(url: str):
#     r = requests.get(url)
#     soup = BeautifulSoup(r.text, 'html.parser')
#     results = soup.find_all(['h1', 'p'])
#     text = [result.text for result in results]
#     article = ' '.join(text)
#     article_len = len(article.split(' '))
#     article = format_input(text=article)
#     return article, article_len

# def format_input(text: str):
#     text = text.replace('.', '.<eos>')
#     text = text.replace('!', '!<eos>')
#     text = text.replace('?', '?<eos>')
#     return text


# def get_sentences(article: str):
#     sentences = article.split('<eos>')

#     return sentences

# # Return chunked text to have less than 500 words
# def chunk_text(sentences: 'list[str]', max_chunk: int):
#     current_chunk = 0
#     chunks = []

#     for sentence in sentences:
#         if len(chunks) == current_chunk + 1:
#             # Check if the chunk is less than max_chunk
#             if len(chunks[current_chunk]) + len(sentence.split(' ')) <= max_chunk:
#                 chunks[current_chunk].extend(sentence.split(' '))
#             # Next chunk
#             else:
#                 current_chunk += 1
#                 chunks.append(sentence.split(' '))
#         else:
#             chunks.append(sentence.split(' '))

#     for chunk_id in range (len(chunks)):
#         chunks[chunk_id] = ' '.join(chunks[chunk_id])

#     return chunks

# # Add markdown annotation to summary
# def convert_markdown(results: dict):
#     md = "" # Markdown output to be fed into slide generator
#     txt = ""

#     for result in results:
#         for head, val in result.items():
#             if(len(txt) == 0): # If first item
#                 md += "# " + head + "\n\n- " + val
#                 txt += val
#             elif(len(txt) < 1000 and result != results[-1]): # If the number of characters on a slide exceeds 1000, create new slide
#                 md += "\n-" + val
#                 txt += val
#             else:
#                 md += "\n\n---\n\n"
#                 txt = ""

#     return md

# # Feed markdown into slides
# def generate_slides(results: dict):
#     txt = "" # Keep track of the current slide text length
#     md_file = MdUtils(file_name='slides_test.md')
#     md_file.create_md_file()

#     # Metadata
#     md_file.new_paragraph('---')
#     md_file.new_paragraph('theme: apple-basic')
#     md_file.new_paragraph('background: https://source.unsplash.com/collection/94734566/1920x1080')
#     md_file.new_paragraph('class: text-center')
#     md_file.new_paragraph('class: text-center')
#     md_file.new_paragraph('highlighter: shiki')
#     md_file.new_paragraph('lineNumbers: false')
#     md_file.new_paragraph('info: |\n## Slidev Starter Template\n Presentation slides for developers.\n\nLearn more at [Sli.dev](https://sli.dev)')
#     md_file.new_paragraph('drawings:\npersist: false')
#     md_file.new_paragraph('title: Welcome to Slidev')
#     md_file.new_paragraph('---')

#     # First slide
#     md_file.new_header(level=1, title='Welcome to Slidev')
#     md_file.new_paragraph('Presentation slides for developers')
#     md_file.new_paragraph('<div class="pt-12">' \
#   '<span @click="$slidev.nav.next" class="px-2 py-1 rounded cursor-pointer" hover="bg-white bg-opacity-10">' \
#     'Press Space for next page <carbon:arrow-right class="inline"/>' \
#   '</span>' \
# '</div>' \


#  ' <button @click="$slidev.nav.openInEditor()" title="Open in Editor" class="text-xl icon-btn opacity-50 !border-none !hover:text-white">' \
#     '<carbon:edit />' \
#   '</button>' \
#   '<a href="https://github.com/slidevjs/slidev" target="_blank" alt="GitHub"' \
#     'class="text-xl icon-btn opacity-50 !border-none !hover:text-white">' \
#     '<carbon-logo-github />' \
#   '</a>' \
# '</div>')

#     md_file.write('<!--' \
#     'The last comment block of each slide will be treated as slide notes. It will be visible and editable in Presenter Mode along with the slide. [Read more in the docs](https://sli.dev/guide/syntax.html#notes)' \
#     '-->')

#     md_file.new_paragraph('---')

#     for result in results:
#         for head, val in result.items():
#             if(len(txt) == 0):
#                 md_file.new_header(level=1, title=head)
#                 md_file.new_paragraph('- ' + val)
#                 txt += val
#             elif(len(txt) < 1000 and result != results[-1]): # If number of characters on current slide is within 1000 and it is not the last slide, create new slide
#                 md_file.new_paragraph('- ' + val)
#                 txt += val
#             else:
#                 md_file.new_paragraph('---')
#                 txt += val

#     print('Slides generated.')

# Convert dict summary object into string
def extract_sentence(results: dict):
    summary = ''
    for result in results:
        summary += ''.join(str(val[1:]) + "\n" for _, val in result.items())

    summary = summary.replace(' .', '.')
    summary = summary.replace(" !", "!")
    summary = summary.replace(" ?", "?")

    return summary

# Generate summary statistics
def generate_statistics(summary: str, words_before: int):
    words_after = len(summary.split(' '))
    reduced_by = (words_before - words_after) / words_before * 100

    print("Number of words in summary: " + str(words_after))
    print("Number of words in original article: " + str(words_before))
    print("Reduced by: " + str(round(reduced_by, 2)) + "%\n")
    print(summary)

# Main function
def main():
    print('---START OF PROGRAM---')

    inputer = Inputer(type='url')
    chunks, article_len = inputer.get_input(inp=input('URL: '))
    # article, article_len = get_article(url=input('URL: '))
    text_summarizer = TextSummarizer().mode # Abstractive summarisation
    adapter = Adapter()

    # sentences = get_sentences(article=article)
    # chunks = chunk_text(sentences=sentences, max_chunk=500)
    results = text_summarizer.summarize(chunks=chunks)
    adapter.convert_markdown(results=results)
    summary = extract_sentence(results=results)
    # generate_slides(results=results)
    generate_statistics(summary=summary, words_before=article_len)

if __name__ == "__main__":
    main()