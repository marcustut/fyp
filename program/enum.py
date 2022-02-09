from enum import Enum

class SummarizeMode(str, Enum):
    Abs = 'abs'
    Ext = 'ext'

class SummarizeType(str, Enum):
    Txt = 'txt'
    Url = 'url'
    Pdf = 'pdf'