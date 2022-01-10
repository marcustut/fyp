# Use pytorch GPU base image
FROM gcr.io/cloud-aiplatform/training/pytorch-gpu.1-10

# set working directory
WORKDIR /app

# Install required packages
RUN pip install google-cloud-storage transformers datasets tqdm cloudml-hypertune pandas

# Copies the trainer code to the docker image.
COPY ../trainer/__init__.py /app/trainer/__init__.py
COPY ../trainer/train.py /app/trainer/train.py

# Set up the entry point to invoke the trainer.
ENTRYPOINT ["python", "-m", "trainer.train"]