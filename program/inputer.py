from bs4 import BeautifulSoup
import requests

class Inputer():
    '''Handles everything regarding the input text, including obtaining, preprocessing and formatting of text.
    `type`: the input type selected.
    `valid_types`: valid input types
    `max_chunk`: the maximum length of a text chunk.
    '''

    type: str
    valid_types = {'url', 'txt', 'pdf'}
    max_chunk: int

    def __init__(self, type: str, max_chunk=500) -> None:
        '''Initialises the Inputer object.'''
        if(self.__contains__(type)):
            self.type = type
        else:
            raise Exception('Invalid input type given. Valid types are', self.valid_type)
        self.max_chunk = max_chunk
        pass

    def __contains__(self, type: str) -> bool:
        '''Checks if input type is valid.'''
        if(type in self.valid_types):
            return True
        else:
            return False

    def get_input(self, inp: str) -> 'list':
        '''Gets input based on the input type.'''
        if(self.type == "url"):
            article, article_len = self.__get_article(url=inp)
        sentences = self.__get_sentences(article=article)
        chunks = self.__chunk_text(sentences=sentences)

        # TODO: Extract from PDF
        

        return chunks, article_len

    def __get_article(self, url: str) -> 'list[str]':
        '''Gets article from URL.'''
        r = requests.get(url)
        soup = BeautifulSoup(r.text, 'html.parser')
        results = soup.find_all(['h1', 'p'])
        text = [result.text for result in results]
        article = ' '.join(text)
        article_len = len(article.split())
        article = self.__add_tokens(text=article)
        return article, article_len

    def __add_tokens(self, text: str) -> str:
        '''Adds tokens to text for easier processing.'''
        text = text.replace('.', '.<eos>')
        text = text.replace('!', '!<eos>')
        text = text.replace('?', '?<eos>')
        return text


    def __get_sentences(self, article: str) -> 'list[str]':
        '''Gets individual sentences for text chunking.'''
        sentences = article.split('<eos>')

        return sentences

    def __chunk_text(self, sentences: 'list[str]') -> 'list[str]':
        '''Chunks text for each chunk to be less than the max length.'''
        current_chunk = 0
        chunks = []

        for sentence in sentences:
            if len(chunks) == current_chunk + 1:
                # Check if the chunk is less than max_chunk
                if len(chunks[current_chunk]) + len(sentence.split()) <= self.max_chunk:
                    chunks[current_chunk].extend(sentence.split())
                # Next chunk
                else:
                    current_chunk += 1
                    chunks.append(sentence.split())
            else:
                chunks.append(sentence.split())

        for chunk_id in range (len(chunks)):
            chunks[chunk_id] = ' '.join(chunks[chunk_id])

        return chunks