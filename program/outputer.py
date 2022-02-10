from typing import Tuple


class Outputer():
    '''Generates summary statistics regarding the output text.
    `results`: the summary text returned from a TextSummarizer object for statistical analysis.
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
        summary = self.__list_to_string()
        self.words_after = len(summary.split())

        return summary

    def __list_to_string(self) -> str:
        '''Converts a list of dict summary objects into string.'''
        summary = ''
        for result in self.results:
            summary += ''.join(str(val[1:]) + "\n" for _, val in result.items())

        return summary

    def generate_statistics(self, summary: str, words_before: int) -> Tuple[int, int, float]:
        '''Generates summary statistics. Returns `summary word count`, `original article word count`, `reduced by percentage`.'''
        self.words_before = words_before
        reduced_by = (self.words_before - self.words_after) / self.words_before * 100

        return self.words_after, self.words_before, round(reduced_by, 2)

        # print("Number of words in summary: " + str(self.words_after))
        # print("Number of words in original article: " + str(self.words_before))
        # print("Reduced by: " + str(round(reduced_by, 2)) + "%\n")
        # print(summary)