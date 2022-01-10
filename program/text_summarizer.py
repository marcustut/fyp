from transformers import pipeline
from transformers import BartForConditionalGeneration, AutoModelForSeq2SeqLM
from transformers import AutoTokenizer
from transformers.pipelines.base import Pipeline
from summarizer import Summarizer
import os
import pickle


class TextSummarizer():
    '''Text summarizers using different models and tokenizers. Includes saving of pre-trained models into directories.
    `valid_modes`: the valid modes of summarisation, including 'abs' and 'ext' for abstractive and extractive summarisation respectively.
    `body`: summarizer for the slide body.
    `title`: summarizer for the slide title.
    `summarizer`: the summarizer pipeline.
    `model`: the weights of the model.
    `tokenizer`: the model tokenizer.
    '''

    valid_modes = {'abs', 'ext'}
    body: 'TextSummarizer'
    title: 'Title'
    summarizer: any
    model: any
    tokenizer: AutoTokenizer

    def __init__(self, mode: str) -> None:
        '''Initialises the TextSummarizer object with the summarization mode and the title summarizer.'''
        if(self.__contains__(mode) == False):
            raise Exception(
                'Invalid summarizer mode given. Valid mdoes are', self.valid_modes)

        if(mode == 'abs'):
            self.body = Abstractive()
        elif(mode == 'ext'):
            self.body = Extractive()

        self.title = Title()
        pass

    def __contains__(self, mode: str) -> bool:
        '''Checks if input type is valid.'''
        if(mode in self.valid_modes):
            return True
        else:
            return False

class Extractive(TextSummarizer):
    '''The mode of summarization that can be chosen by the user. This is extractive summarization where a summary is pieced together from the original sentences in the document.'''

    def __init__(self) -> None:
        self.__create_summarizer(path='../models/extractive/')
        self.__save_model(path='../models/extractive/', model=self.summarizer)
        pass

    def __create_summarizer(self, path: str):
        # If file is not empty
        if(self.__is_empty(path) == False):
            self.summarizer = pickle.load(open(path + 'bert-ext.pkl', 'rb'))
        else:
            self.summarizer = Summarizer()

    def __save_model(self, path: str, model) -> None:
        '''Saves the model and tokenizer to a directory.'''
        if(self.__is_empty(path)):
            pickle.dump(model, open(path + 'bert-ext.pkl', 'wb'))

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

    def summarize(self, chunks: 'list[str]', ratio=0.05) -> 'list[dict]':
        results = []
        for chunk in chunks:
            dict_ = {'summary_text': ' ' + self.summarizer(chunk, ratio=ratio)}
            results.append(dict_)
        return results

class Abstractive(TextSummarizer):
    '''The mode of summarization that can be chosen by the user. This is abstractive summarization where new sentences are generated from the original document.
    '''

    def __init__(self, checkpoint='sshleifer/distilbart-cnn-12-6') -> None:
        self.__create_summarizer(checkpoint=checkpoint)
        self.__save_model(path=checkpoint, model=self.model,
                          tokenizer=self.tokenizer)
        pass

    def __create_summarizer(self, checkpoint: str, ) -> Pipeline:
        '''Creates a summarizer from loading a model and tokenizer.'''
        checkpoint = 'sshleifer/distilbart-cnn-12-6'
        model = BartForConditionalGeneration.from_pretrained(checkpoint)
        tokenizer = AutoTokenizer.from_pretrained(checkpoint)

        self.summarizer, self.model, self.tokenizer = pipeline(
            'summarization', model=model, tokenizer=tokenizer), model, tokenizer

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
        '''Summarizes text and returns body summary'''
        results = self.summarizer(chunks, return_text='True')
        print(results)
        return results


class Title(TextSummarizer):
    '''An internal class that will be used to generate slide titles, and combining the title with body results.
    '''

    def __init__(self, file='../models/title-generator-t5-arxiv-16-4.pkl') -> None:
        '''Initialises the title summarizer.'''
        self.__create_summarizer(file=file)
        # self.__save_model(path=file, model=self.model, tokenizer=self.tokenizer)
        pass

    # Local title generator
    def __create_summarizer(self, file: str):
        '''Creates the summarizer.'''
        self.model = pickle.load(open(file, 'rb'))
        self.summarizer = self.model.predict
        pass

    # PEGASUS SUMMARIZER
    # def __init__(self, checkpoint='sshleifer/distill-pegasus-xsum-16-4') -> None:
    #     '''Initialises the title summarizer.'''
    #     self.__create_summarizer(checkpoint=checkpoint)
    #     self.__save_model(path=checkpoint, model=self.model,
    #                       tokenizer=self.tokenizer)
    #     pass

    # def __create_summarizer(self, checkpoint: str):
    #     '''Creates the summarizer.'''
    #     model = AutoModelForSeq2SeqLM.from_pretrained(checkpoint)
    #     tokenizer = AutoTokenizer.from_pretrained(checkpoint)

    #     self.summarizer, self.model, self.tokenizer = pipeline(
    #         'summarization', model=model, tokenizer=tokenizer), model, tokenizer

    # def __save_model(self, path: str, model, tokenizer) -> None:
    #     '''Saves the model and tokenizer to a directory.'''
    #     if(self.__is_empty('../models/' + path)):
    #         model.save_pretrained(save_directory='../models/' + path)

    #     if(self.__is_empty('../tokenizers/' + path)):
    #         tokenizer.save_pretrained(save_directory='../tokenizers/' + path)

    # def __is_empty(self, path: str) -> bool:
    #     '''Checks whether directory path is empty and creates one if it does not exist.'''
    #     if os.path.exists(path) and not os.path.isfile(path):
    #         # Checking if the directory is empty or not
    #         if not os.listdir(path):
    #             return True
    #         else:
    #             return False
    #     else:
    #         os.makedirs(path)

    def __transpose_dict(self, dict_: dict) -> dict:
        '''Transposes the keys and values of the dictionary object. Based on the assumption that all keys and values are unique.'''
        return {y:x for x, y in dict_.items()}

    def summarize(self, results: 'list[dict]') -> 'list[dict]':
        '''Summarizes from the result body to give a title, then combines these pairs together.'''
        print('Title Summarizer')
        new_results = []
        for dict_ in results:
            # Now the body is the key and the title is the value
            dict_ = self.__transpose_dict(dict_)
            for body in dict_:
                # Summarize the given text
                dict_[body] = self.summarizer(body) # Returns a list
                pass
            dict_ = self.__transpose_dict(dict_)
            new_results.append(dict_)
            print('dict_:', dict_)

        return new_results

    # def summarize(self, results: 'list[dict]') -> 'list[dict]':
    #     '''Summarizes from the result body to give a title, then combines these pairs together.'''
    #     print('Title Summarizer')
    #     new_results = []
    #     for dict_ in results:
    #         # Now the body is the key and the title is the value
    #         dict_ = self.__transpose_dict(dict_)
    #         for body in dict_:
    #             # Summarize the given text
    #             title = self.summarizer(body, return_text='True', max_length=20)
    #             dict_[body] = title[0]['summary_text'] # Only one dict returned in an array
    #             pass
    #         dict_ = self.__transpose_dict(dict_)
    #         new_results.append(dict_)
    #         print('dict_:', dict_)

    #     return new_results
