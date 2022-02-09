from http.client import HTTPException
from django.forms import ValidationError
from flask import Flask
from flask_restful import Resource, Api, reqparse
from django.core.validators import URLValidator
from grpc import StatusCode
import requests
import ast

# from program.enum import SummarizeMode, SummarizeType
from inputer import Inputer
from outputer import Outputer
from text_summarizer import TextSummarizer
from adapter import Adapter
import transformers
transformers.logging.set_verbosity_debug() # To check if the model is running
import json

app = Flask(__name__)
api = Api(app)

class Summarize(Resource):
    def get(self):
        return {"message": "Connection successful"}, 200
        
    def post(self):
        parser = reqparse.RequestParser()

        parser.add_argument('mode', required=True, choices=['abs', 'ext'])
        parser.add_argument('type', required=True, choices=['txt', 'url', 'pdf'])
        parser.add_argument('input', required=True)
        parser.add_argument('maxChunk', default=500, required=False, type=int, choices=range(50, 500, 50))
        parser.add_argument('maxCharPerSlide', default=500, required=False, type=int, choices=range(100, 500, 50)) # max_len

        # TODO: Add themes argument

        args = parser.parse_args()

        # Validate url
        if args['type'] == 'url':
            validateURL = URLValidator()

            # Check if input looks like an URL
            try:
                validateURL(args['input'])
            except ValidationError:
                raise HTTPException(StatusCode=400, detail=f"Input is an invalid URL.")

            inputer = Inputer(type=args['type'])

            # Get URL

            try:
                chunks, article_len = inputer.get_input(inp=args['input'])
            except requests.ConnectionError as ex:
                raise HTTPException(StatusCode=404, detail=f"URL does not exist. {ex} ")

            # Summarize
            try:
                text_summarizer = TextSummarizer(mode=args['mode'])
                results = text_summarizer.body.summarize(chunks=chunks)
                results = text_summarizer.title.summarize(results=results)
            except Exception as ex:
                raise HTTPException(StatusCode=500, detail=f"{ex}")

            # Convert to markdown
            try:
                adapter = Adapter()
                adapter.convert_markdown(results=results)
            except Exception as ex:
                raise HTTPException(StatusCode=500, detail=f"{ex}")

            # Output statistics
            outputer = Outputer()
            summary = outputer.get_output(results=results)
            words_after, words_before, reduced = outputer.generate_statistics(summary=summary, words_before=article_len)

            # Return markdown string/file via JSON
            response = {
                "summary": adapter.md,
                "wordsAfter": words_after,
                "wordsBefore": words_before,
                "reducedByPercentage": reduced
            }

            return json.dump(response), 200

        # try:
        #     summarize_mode = SummarizeMode(args['mode'])
        # except ValueError:
        #     raise HTTPException(status_code=404, detail=f"Mode '{args['mode']}' is invalid.")

        # try:
        #     summarize_type = SummarizeType(args['type'])
        # except ValueError:
        #     raise HTTPException(status_code=404, detail=f"Type '{args['type']}' is invalid.")

        # Validate maxChunk

    pass

api.add_resource(Summarize, '/summarize')

if __name__ == '__main__':
    app.run()