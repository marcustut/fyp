from http.client import HTTPException
from flask import Flask
from flask_restful import Resource, Api, reqparse
from enum import SummarizeMode
import ast

from program.enum import SummarizeMode
from inputer import Inputer
from outputer import Outputer
from text_summarizer import TextSummarizer
from adapter import Adapter
import transformers
transformers.logging.set_verbosity_debug() # To check if the model is running

app = Flask(__name__)
api = Api(app)

class Summarize(Resource):
    def post(self):
        parser = reqparse.RequestParser()

        parser.add_argument('mode', required=True)
        parser.add_argument('url', required=False)
        parser.add_argument('url', required=False)

        args = parser.parse_args()

        try:
            summarize_mode = SummarizeMode(args['mode'])
        except ValueError:
            raise HTTPException(status_code=404, detail=f"Mode '{args['mode']}' is invalid.")

    pass

api.add_resource(Summarize, '/summarize')

if __name__ == '__main__':
    app.run()