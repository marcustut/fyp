#!/bin/bash

python3 main.py # Run main.py
mv ../slidev/slides.txt ../slidev/slides.md # Rename file
cd ../slidev
npm install
npm run dev