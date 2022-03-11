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

    # text_summarizer_abs = ''
    # text_summarizer_ext = ''

    def post(self):
        parser = reqparse.RequestParser()

        parser.add_argument('mode', required=True, choices=['abs', 'ext'])
        parser.add_argument('type', required=True, choices=['txt', 'url', 'pdf'])
        parser.add_argument('input', required=True)
        parser.add_argument('maxChunk', default=500, required=False, type=int, choices=range(50, 501, 50))
        parser.add_argument('maxCharPerSlide', default=500, required=False, type=int, choices=range(100, 501, 50)) # max_len
        parser.add_argument('theme', default='apple-basic', required=False, choices=['apple-basic', 'seriph', 'default', 'bricks', 'shibainu'])
        parser.add_argument('outputName', required=True)

        args = parser.parse_args()

        # Validate url
        if args['type'] == 'url':
            validateURL = URLValidator()

            # Check if input looks like an URL
            try:
                validateURL(args['input'])
            except ValidationError:
                raise HTTPException(StatusCode=400, detail=f"Input is an invalid URL.")

            inputer = Inputer(type=args['type'], max_chunk=args['maxChunk'])

            # Get URL
            try:
                chunks, article_len = inputer.get_input(inp=args['input'])
            except requests.ConnectionError as ex:
                raise HTTPException(StatusCode=404, detail=f"URL does not exist. {ex} ")

        elif args['type'] == 'pdf':

            # Check if path exists
            if os.path.isfile(f"./uploads/{args['input']}"):
                # Extract text
                inputer = Inputer(type=args['type'], max_chunk=args['maxChunk'])
                try:
                    chunks, article_len = inputer.get_input(inp="./uploads/" + args['input'])
                except Exception as ex:
                    raise HTTPException(StatusCode=422, detail=f"Failed to extract text from PDF file. {ex}")
            else:
                raise HTTPException(StatusCode=404, detail="File does not exist.")

        elif args['type'] == 'txt':
            # Check if path exists
            if os.path.isfile(f"./uploads/{args['input']}"):
                inputer = Inputer(type=args['type'], max_chunk=args['maxChunk'])
                # Read TXT file
                try:
                    chunks, article_len = inputer.get_input(inp="./uploads/" + args['input'])
                except Exception as ex:
                    raise HTTPException(StatusCode=422, detail=f"Failed to read content from TXT file. {ex}")
            else:
                raise HTTPException(StatusCode=404, detail="File does not exist.")

        # Choose summarizer
        # if args['mode'] == 'abs':
        #     text_summarizer = self.text_summarizer_abs
        # else:
        #     text_summarizer = self.text_summarizer_ext

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
            adapter.convert_markdown(file_name=args['outputName'], results=results)
        except Exception as ex:
            raise HTTPException(StatusCode=500, detail=f"{ex}")

        # Output statistics
        outputer = Outputer()
        summary = outputer.get_output(results=results)
        words_after, words_before, reduced = outputer.generate_statistics(words_before=article_len)

        # Return markdown string/file via JSON
        response = {
            "fileName": args['outputName'],
            "wordsAfter": words_after,
            "wordsBefore": words_before,
            "reducedByPercentage": reduced
        }

        return response, 200

    # def __init__(self) -> None:
    #     try:
    #         self.text_summarizer_abs = TextSummarizer(mode='abs')
    #         self.text_summarizer_ext = TextSummarizer(mode='ext')
    #     except Exception as ex:
    #         raise HTTPException(StatusCode=500, detail=f"Failed to initialise summarizer. {ex}")

    pass

api.add_resource(Summarize, '/summarize')

@app.route("/output/<path:path>", methods=['GET'])
def send_static_files(path):
    return send_from_directory('output', path)

def allowed_file(filename):
    return '.' in filename and filename.rsplit('.', 1)[1].lower() in ALLOWED_EXTENSIONS

@app.route('/estimate', methods=['POST'])
def count_word():
    type_ = request.args.get('type')
    input_ = request.args.get('input')

    if type_ == 'url':
        validateURL = URLValidator()

        # Check if input looks like an URL
        try:
            validateURL(input_)
        except ValidationError:
            raise HTTPException(StatusCode=400, detail=f"Input is an invalid URL.")

        inputer = Inputer(type=type_)

        # Get URL
        try:
            _, article_len = inputer.get_input(inp=input_)

        except requests.ConnectionError as ex:
            raise HTTPException(StatusCode=404, detail=f"URL does not exist. {ex} ")

    elif type_ == 'pdf':
        # Check if path exists
        if os.path.isfile(f"./uploads/{input_}"):
            # Extract text
            inputer = Inputer(type=type_)
            try:
                _, article_len = inputer.get_input(inp="./uploads/" + input_)
            except Exception as ex:
                raise HTTPException(StatusCode=422, detail=f"Failed to extract text from PDF file. {ex}")
        else:
            raise HTTPException(StatusCode=404, detail="File does not exist.")

    elif type_ == 'txt':
        # Check if path exists
        if os.path.isfile(f"./uploads/{input_}"):
            inputer = Inputer(type=type_)
            # Read TXT file
            try:
                _, article_len = inputer.get_input(inp="./uploads/" + input_)
            except Exception as ex:
                raise HTTPException(StatusCode=422, detail=f"Failed to read content from TXT file. {ex}")
        else:
            raise HTTPException(StatusCode=404, detail="File does not exist.")

    else:
        raise HTTPException(StatusCode=403, detail="File type not supported.")

    # 8 second per word
    raw = round(article_len * 0.15)
    mins = int((raw - raw % 60) / 60)
    secs = raw % 60

    return {"articleLength": article_len, "minutes": mins, "seconds": secs}, 200

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

# with app.app_context():
#     print("Run before app.run()")
#     summarize = Summarize()

if __name__ == '__main__':
    app.run()
