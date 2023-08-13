# goバージョン
FROM golang:latest

RUN apt-get install vim
# cron入れる
RUN apt-get update && apt-get install -y cron
RUN cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime
# boiler-plateディレクトリの作成
RUN mkdir /go/src/app
# ワーキングディレクトリの設定
WORKDIR /go/src/app
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /go/src/app
#cron setting
#分 時 日 月 曜日 <実行コマンド>
# 毎時30分に実行
CMD echo '30 * * * * cd /go/src/app;  go run /go/src/app/main.go >> /var/log/cron.log 2>&1' | crontab - && cron -f
