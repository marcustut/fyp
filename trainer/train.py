# Load Data
import json

data = '../dataset/arxiv-metadata-oai-snapshot.json'

def get_metadata():
    with open(data, 'r') as f:
        for line in f:
            yield line

metadata = get_metadata()

for paper in metadata:
    paper_dict = json.loads(paper)
    print('Title: {}\n\nAbstract: {}\nRef: {}'.format(paper_dict.get('title'), paper_dict.get('abstract'), paper_dict.get('journal-ref')))
    break

titles = []
abstracts = []
years = []
metadata = get_metadata()

for paper in metadata:
    paper_dict = json.loads(paper)
    ref = paper_dict.get('journal-ref')
    try:
        year = int(ref[-4:])
        if 2016 < year < 2021:
            years.append(year)
            titles.append(paper_dict.get('title'))
            abstracts.append(paper_dict.get('abstract'))
    except:
        pass

# Convert into DF
import pandas as pd

papers = pd.DataFrame({
    'title': titles,
    'abstract': abstracts,
    'year': years
})

# Splitting data using the 80:20 training-testing ratio
eval_df = papers.sample(frac=0.2, random_state=673)
train_df = papers.drop(eval_df.index)

from datasets import Dataset

train_data = Dataset.from_pandas(train_df)
eval_data = Dataset.from_pandas(eval_df)

train_data = train_data.remove_columns('__index_level_0__')
eval_data = eval_data.remove_columns('__index_level_0__')

from transformers import AutoTokenizer

path = 'sshleifer/distilbart-cnn-12-6'

tokenizer = AutoTokenizer.from_pretrained(path)

train_dataset = train_data.map(lambda x: tokenizer(train_data['abstract'], padding='max_length', truncation=True), remove_columns=train_data.column_names, batched=True)
eval_dataset = eval_data.map(lambda x: tokenizer(eval_data['abstract'], padding='max_length', truncation=True), remove_columns=eval_data.column_names, batched=True)

from transformers import AutoModelForSeq2SeqLM

model = AutoModelForSeq2SeqLM.from_pretrained(path)

from transformers import TrainingArguments

training_args = TrainingArguments(
    output_dir='trainer' + path,
    per_device_train_batch_size=8,
    per_device_eval_batch_size=8,
    num_train_epochs=4
)

from transformers import Trainer

trainer = Trainer(
    model=model,
    args=training_args,
    train_dataset=train_dataset,
    eval_dataset=eval_dataset
)

from torch import cuda
cuda.empty_cache()