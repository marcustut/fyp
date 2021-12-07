from transformers import pipeline
from transformers import BartForConditionalGeneration
from transformers import AutoTokenizer
from transformers.pipelines.base import Pipeline
import os

class TextSummarizer():
    '''Text summarizers using different models and tokenizers. Includes saving of pre-trained models into directories.'''

    summarizer: Pipeline
    model: BartForConditionalGeneration
    tokenizer: AutoTokenizer

    def __init__(self, checkpoint='sshleifer/distilbart-cnn-12-6') -> None:
        '''Initialises the TextSummarizer object with a summarizer, model and tokenizer.'''
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
        '''Checks whether directory path is empty.'''
        if os.path.exists(path) and not os.path.isfile(path):
            # Checking if the directory is empty or not
            if not os.listdir(path):
                return True
            else:
                return False
        else:
            os.makedirs(path)

    def summarize(self, chunks) -> dict:
        '''Summarizes text.'''
        return self.summarizer(chunks, return_text='True')