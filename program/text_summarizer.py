from transformers import pipeline
from transformers import BartForConditionalGeneration, AutoModel
from transformers import AutoTokenizer
from transformers.pipelines.base import Pipeline
import os

class TextSummarizer():
    '''Text summarizers using different models and tokenizers. Includes saving of pre-trained models into directories.
    `valid_modes`: the valid modes of summarisation, including 'abs' and 'ext' for abstractive and extractive summarisation respectively.
    '''

    valid_modes = {'abs', 'ext'}

    def __init__(self, mode: str) -> None:
        '''Initialises the TextSummarizer object with the summarization mode and the title summarizer.'''
        if(self.__contains__(mode) == False):
            raise Exception('Invalid summarizer mode given. Valid mdoes are', self.valid_modes)

        self.__new__(mode)
        self.Title()

        pass

    def __new__(self, mode: str) -> 'TextSummarizer':
        if(mode == 'abs'):
            return self.Abstractive()
        pass

    def __contains__(self, mode: str) -> bool:
        '''Checks if input type is valid.'''
        if(mode in self.valid_modes):
            return True
        else:
            return False

    class Abstractive():
        '''The mode of summarization that can be chosen by the user.
        `summarizer`: the summarizer pipeline.
        `model`: the weights of the model.
        `tokenizer`: the model tokenizer.
        '''
        summarizer: Pipeline
        model: AutoModel
        tokenizer: AutoTokenizer

        def __init__(self, checkpoint='sshleifer/distilbart-cnn-12-6') -> None:
            self.summarizer, self.model, self.tokenizer = self.__create_summarizer(checkpoint=checkpoint)
            self.__save_model(path=checkpoint, model=self.model, tokenizer=self.tokenizer)
            pass

        def __create_summarizer(self, checkpoint: str, ) -> Pipeline:
            '''Creates a summarizer from loading a model and tokenizer.'''
            checkpoint = 'sshleifer/distilbart-cnn-12-6'
            model = BartForConditionalGeneration.from_pretrained(checkpoint)
            tokenizer = AutoTokenizer.from_pretrained(checkpoint)

            return pipeline('summarization', model=model, tokenizer=tokenizer), model, tokenizer

        def __save_model(self, path: str, model, tokenizer) -> None:
            '''Saves the model and tokenizer to a directory.'''
            if(self.__is_empty('../models/' + path)):
                model.save_pretrained(save_directory='../models/' + path)

            if(self.__is_empty('../tokenizers/' + path)):
                tokenizer.save_pretrained(save_directory='../tokenizers/' + path)

        def __is_empty(self, path: str) -> bool:
            '''Checks whether directory path is empty and creates one if it does not exist.'''
            if os.path.exists(path) and not os.path.isfile(path):
                # Checking if the directory is empty or not
                if not os.listdir(path):
                    return True
                else:
                    return False
            else:
                os.makedirs(path)

        def summarize(self, chunks: 'list[str]') -> 'list[dict]':
            '''Summarizes text.'''
            results = self.summarizer(chunks, return_text='True')
            return results

    class Title():
        '''An internal class that will be used to generate slide titles, and combining the title with body results.
        `title_summarizer`: the summarizer for the title.
        '''

        title_summarizer: any

        def __init__(self) -> None:
            '''Initialises the title summarizer.'''
            self.title_summarizer = self.__create_summarizer()
            pass

        def __create_summarizer(self):
            '''Creates the summarizer.'''
            pass

        def summarize(self, results: 'list[dict]') -> 'list[dict]':
            '''Summarizes from the result body to give a title, then combines these pairs together.'''
            for dict_ in results:
                for title, body in dict_:
                    pass

            return results
