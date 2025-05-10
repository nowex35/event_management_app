FROM apiaryio/dredd:latest

WORKDIR /dredd

# 必要なら追加のツールや設定ファイルをインストール・コピー
COPY openapi.yaml /dredd/openapi.yaml
# COPY dredd.yml /dredd/dredd.yml

# デフォルトコマンド（必要に応じて変更）
CMD ["dredd", "openapi.yaml", "http://host.docker.internal:6262"]