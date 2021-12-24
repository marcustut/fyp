class Outputer():
    '''Handles everything regarding the output string text, including capitalising and generating summary statistics.
    `results`: the summary text returned from a TextSummarizer object for post-processing.
    `words_before`: number of words in the original text before summarising.
    `words_after`: number of words in the summary text after summarising.
    '''
    results: any
    words_before: int
    words_after: int

    def __init__(self) -> None:
        '''Initialise the Outputer object.'''
        pass

    def get_output(self, results: any) -> str:
        '''Gets the summary text in plain string format.'''
        self.results = results
        if(isinstance(self.results, list)):
            summary = self.__list_to_string()
        self.words_after = len(summary.split())

        return summary

    def __list_to_string(self) -> str:
        '''Converts a list of dict summary objects into string.'''
        summary = ''
        for result in self.results:
            summary += ''.join(str(val[1:]) + "\n" for _, val in result.items())

        self.__strip_whitespaces(summary=summary)
        return summary

    def __strip_whitespaces(self, summary: str) -> str:
        '''Strips additional whitespaces in front of punctuations.'''
        summary = summary.replace(' .', '.')
        summary = summary.replace(" !", "!")
        summary = summary.replace(" ?", "?")

        return summary

    def generate_statistics(self, summary: str, words_before: int) -> None:
        '''Generates summary statistics.'''
        self.words_before = words_before
        reduced_by = (self.words_before - self.words_after) / self.words_before * 100

        print("Number of words in summary: " + str(self.words_after))
        print("Number of words in original article: " + str(self.words_before))
        print("Reduced by: " + str(round(reduced_by, 2)) + "%\n")
        print(summary)