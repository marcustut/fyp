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
npm install @slidev/cli @slidev/theme-default @slidev/theme-apple-basic @slidev/theme-seriph @slidev/theme-bricks @slidev/theme-shibainu

# copy .md from bucket
aws s3 cp s3://${BUCKET_NAME}/slides/${USER_ID}/${SLIDE_ID}.md slides.md

# run slidev
npx slidev --port 80 --remote
