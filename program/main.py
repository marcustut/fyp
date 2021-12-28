from inputer import Inputer
from outputer import Outputer
from text_summarizer import TextSummarizer
from adapter import Adapter
import transformers
transformers.logging.set_verbosity_debug() # To check if the model is running

def main():
    print('---START OF PROGRAM---')

    inputer = Inputer(type='url')
    text_summarizer = TextSummarizer(mode='abs') # Abstractive summarisation
    adapter = Adapter()
    outputer = Outputer()

    chunks, article_len = inputer.get_input(inp='https://en.wikipedia.org/wiki/Rococo')
    results = text_summarizer.body.summarize(chunks=chunks)
    results = text_summarizer.title.summarize(results=results)
    adapter.convert_markdown(results=results)
    summary = outputer.get_output(results=results)
    outputer.generate_statistics(summary=summary, words_before=article_len)

if __name__ == "__main__":
    main()