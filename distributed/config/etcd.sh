#!/bin/bash
addr="localhost:2379"
etcdctl put /config_center/ITEMSAVER_SERVICE < itemsaver.conf --endpoints=$addr
etcdctl put /config_center/CRAWLER_SERVICE < crawler.conf --endpoints=$addr
