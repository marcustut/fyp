from http.client import HTTPException
from django.forms import ValidationError
from flask import request, send_from_directory
from flask_cors import CORS
import flask.scaffold
flask.helpers._endpoint_from_view_func = flask.scaffold._endpoint_from_view_func
from flask_restful import Resource, Api, reqparse
from app import app
from django.core.validators import URLValidator
from werkzeug.utils import secure_filename
import requests
import os

from inputer import Inputer
from outputer import Outputer
from text_summarizer import TextSummarizer
from adapter import Adapter
import transformers
transformers.logging.set_verbosity_debug() # To check if the model is running

cors = CORS(app, origins=["*"])
api = Api(app)

ALLOWED_EXTENSIONS = set(['txt', 'pdf'])

class Summarize(Resource):

    def post(self):
        parser = reqparse.RequestParser()

        parser.add_argument('mode', required=True, choices=['abs', 'ext'])
        parser.add_argument('type', required=True, choices=['txt', 'url', 'pdf'])
        parser.add_argument('input', required=True)
        parser.add_argument('maxChunk', default=500, required=False, type=int, choices=range(50, 500, 50))
        parser.add_argument('maxCharPerSlide', default=500, required=False, type=int, choices=range(100, 500, 50)) # max_len
        parser.add_argument('maxCharPerSlide', default=500, required=False, type=int, choices=range(100, 500, 50)) # max_len
        # TODO: Add themes argument
        parser.add_argument('theme', default='apple-basic', required=False, type=int, choices=['apple-basic', 'seriph', 'default']) # max_len

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
                adapter = Adapter(theme=args['theme'])
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

            return response, 200

    pass

api.add_resource(Summarize, '/summarize')

@app.route("/output/<path:path>", methods=['GET'])
def send_static_files(path):
    return send_from_directory('output', path)

def allowed_file(filename):
    return '.' in filename and filename.rsplit('.', 1)[1].lower() in ALLOWED_EXTENSIONS

@app.route('/uploads', methods=['POST'])
def upload_file():
    # Check if request has file
    if 'file' not in request.files:
        return {'message': 'No file part in the request'}, 400

    file = request.files['file']
    if file.filename == '':
        return {'message': 'No file selected for uploading'}, 400

    if file and allowed_file(file.filename):
        filename = secure_filename(file.filename)
        file.save(os.path.join(app.config['UPLOAD_FOLDER'], filename))
        return {'message': 'File successfully uploaded'}, 201

    else:
        return {'message': 'Allowed file types are txt and pdf'}, 400

if __name__ == '__main__':
    app.run(debug=True) # Debug mode, auto reload on change