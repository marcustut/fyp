from http.client import HTTPException
from django.forms import ValidationError
from flask import Flask, send_from_directory
import flask.scaffold
flask.helpers._endpoint_from_view_func = flask.scaffold._endpoint_from_view_func
from flask_restful import Resource, Api, reqparse
from django.core.validators import URLValidator
import requests
import ast

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
                file_name = adapter.convert_markdown(results=results)
            except Exception as ex:
                raise HTTPException(StatusCode=500, detail=f"{ex}")

            # Output statistics
            outputer = Outputer()
            summary = outputer.get_output(results=results)
            words_after, words_before, reduced = outputer.generate_statistics(summary=summary, words_before=article_len)

            # Return markdown string/file via JSON
            response = {
                "fileName": file_name,
                "wordsAfter": words_after,
                "wordsBefore": words_before,
                "reducedByPercentage": reduced
            }

            return json.dumps(response), 200

    pass

api.add_resource(Summarize, '/summarize')

@app.route("/output/<path:path>", methods=['GET'])
def send_static_files(path):
    return send_from_directory('output', path)

if __name__ == '__main__':
    app.run()