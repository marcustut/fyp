#!/bin/bash

# install nvm
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.1/install.sh | bash
echo 'export NVM_DIR="$([ -z "${XDG_CONFIG_HOME-}" ] && printf %s "${HOME}/.nvm" || printf %s "${XDG_CONFIG_HOME}/nvm")"' >> ~/.bashrc
echo '[ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh" # This loads nvm' >> ~/.bashrc
source ~/.bashrc
nvm install node

# allow node to run on port 80
sudo setcap cap_net_bind_service=+ep `readlink -f \`which node\``

# init npm project
mkdir slidev && cd slidev && npm init -y

# install dependencies
npm install @slidev/cli@0.28.9 @slidev/theme-default@0.21.2 @slidev/theme-apple-basic@0.20.0 @slidev/theme-seriph@0.21.2 @slidev/theme-bricks@0.0.2 @slidev/theme-shibainu@0.0.2

# copy .md from bucket
aws s3 cp s3://${BUCKET_NAME}/slides/${USER_ID}/${SLIDE_ID}.md slides.md
#aws s3 cp s3://ai-text-summarizer/slides/USER_01FVCVRPPZ13BNZ5HNCGWCC89W/SLID_01FWDJKJ5NWKEBRR13Z3Q69X33.md slides.md

# run slidev
npx slidev --port 80 --remote
