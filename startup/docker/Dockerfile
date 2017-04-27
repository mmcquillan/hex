# Example Hex Docker File
FROM hexbotio/hex:beta
RUN apt-get update --fix-missing \
  && apt-get install -y build-essential \
  && git -C /opt clone https://github.com/tj/n \
  && make install -C /opt/n \
  && n 5.3 --quiet \
  && npm install -g azure-cli \
ADD hex.json /etc/hex.json
